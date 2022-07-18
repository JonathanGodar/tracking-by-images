package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/JonathanGodar/go-web-gin/models"
	"github.com/JonathanGodar/go-web-gin/server/myoto"
	"github.com/gofrs/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

type trackerService struct {
	db *sql.DB
	logger *zap.Logger
}

func (s *trackerService) Initialize(db *sql.DB, logger *zap.Logger) {
	s.db = db
	s.logger = logger
}

func (s trackerService) AddTracker(ctx context.Context, req myoto.AddTrackerRequest) (*myoto.AddTrackerResponse, error) {
	user, ok := ctx.Value(SignedInUserKey).(models.User)
	if !ok || user.ID != req.OwnerID {
		return nil, errors.New("Unauthorized")
	}

	id, _ := uuid.NewV4()
	idStr := id.String()

	tracker := &models.Tracker{
		OwnerID:  req.OwnerID,
		IsActive: req.IsActive,
		ID:       idStr,
		URL:      s.getTrackerUrl(idStr),
	}

	err := tracker.Insert(ctx, s.db, boil.Greylist())
	if err != nil {
		return nil, err
	}
	
	// trackerResp := myoto.Tracker {
	// 	ID: tracker.ID,
	// 	TimesAccessed: tracker.TimesAccessed,
	// 	OwnerID: tracker.OwnerID,
	// 	URL: tracker.URL,
	// }

	return &myoto.AddTrackerResponse{
		Tracker: *tracker,
	}, nil
}

// https://casbin.org/
func (s trackerService) UpdateTracker(ctx context.Context, update myoto.UpdateTrackerRequest) (*myoto.UpdateTrackerResponse, error) {
	tracker, err := models.FindTracker(ctx, s.db, update.ID)
	if err != nil {
		return nil, errors.New("Not found")
	}

	tracker.IsActive = update.IsActive
	_, err = tracker.Update(ctx, s.db, boil.Infer())

	if err != nil {
		return nil, err
	}

	return &myoto.UpdateTrackerResponse{
		Tracker: *tracker,
	}, nil
}

func (s trackerService) getTrackerUrl(trackerId string) string {
	return fmt.Sprintf("/tracker/?id=%v", trackerId)
}

func (s trackerService) DeleteTracker(ctx context.Context, id string) (*myoto.DeleteTrackerResponse, error) {
	user, ok := ctx.Value(SignedInUserKey).(models.User)
	if !ok {
		return nil, errors.New("Unauthorized")
	}

	rowsAffected, err := models.Trackers(qm.Where("id = ?", id), qm.And("owner_id = ?", user.ID)).DeleteAll(ctx, s.db)
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("Item not found. Make sure that the specified tracker belongs to you")
	}

	return &myoto.DeleteTrackerResponse{}, nil
}

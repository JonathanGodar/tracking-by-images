package definition

import "github.com/JonathanGodar/go-web-gin/models"

// type Tracker models.Tracker;

type TrackerService interface {
	AddTracker(AddTrackerRequest) AddTrackerResponse;
	UpdateTracker(UpdateTrackerRequest) UpdateTrackerResponse;
	DeleteTracker(ID string) DeleteTrackerResponse; 
}

type UpdateTrackerResponse struct {
	Tracker models.Tracker
}

type UpdateTrackerRequest struct {
	ID string
	IsActive bool
}

type AddTrackerRequest struct {
	OwnerID string
	IsActive bool 
}

type AddTrackerResponse struct {
	Tracker models.Tracker
}

type DeleteTrackerResponse struct {}

// type Tracker struct {
// 	ID            string
// 	TimesAccessed int
// 	OwnerID       string
// 	URL           string
// 	IsActive      bool
// }

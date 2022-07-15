// Code generated by oto; DO NOT EDIT.

package myoto

import (
	"context"
	"net/http"

	"github.com/pacedotdev/oto/otohttp"

	models "github.com/JonathanGodar/go-web-gin/models"
)

type TrackerService interface {
	AddTracker(context.Context, AddTrackerRequest) (*AddTrackerResponse, error)
	DeleteTracker(context.Context, string) (*DeleteTrackerResponse, error)
	UpdateTracker(context.Context, UpdateTrackerRequest) (*UpdateTrackerResponse, error)
}

type UserService interface {
	AddUser(context.Context, AddUserRequest) (*AddUserResponse, error)
	GetAccessToken(context.Context, GetAccessTokenRequest) (*GetAccessTokenResponse, error)
	GetMe(context.Context, interface{}) (*GetMeResponse, error)
}

type trackerServiceServer struct {
	server         *otohttp.Server
	trackerService TrackerService
}

// Register adds the TrackerService to the otohttp.Server.
func RegisterTrackerService(server *otohttp.Server, trackerService TrackerService) {
	handler := &trackerServiceServer{
		server:         server,
		trackerService: trackerService,
	}
	server.Register("TrackerService", "AddTracker", handler.handleAddTracker)
	server.Register("TrackerService", "DeleteTracker", handler.handleDeleteTracker)
	server.Register("TrackerService", "UpdateTracker", handler.handleUpdateTracker)
}

func (s *trackerServiceServer) handleAddTracker(w http.ResponseWriter, r *http.Request) {
	var request AddTrackerRequest
	if err := otohttp.Decode(r, &request); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	response, err := s.trackerService.AddTracker(r.Context(), request)
	if err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	if err := otohttp.Encode(w, r, http.StatusOK, response); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
}

func (s *trackerServiceServer) handleDeleteTracker(w http.ResponseWriter, r *http.Request) {
	var request string
	if err := otohttp.Decode(r, &request); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	response, err := s.trackerService.DeleteTracker(r.Context(), request)
	if err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	if err := otohttp.Encode(w, r, http.StatusOK, response); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
}

func (s *trackerServiceServer) handleUpdateTracker(w http.ResponseWriter, r *http.Request) {
	var request UpdateTrackerRequest
	if err := otohttp.Decode(r, &request); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	response, err := s.trackerService.UpdateTracker(r.Context(), request)
	if err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	if err := otohttp.Encode(w, r, http.StatusOK, response); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
}

type userServiceServer struct {
	server      *otohttp.Server
	userService UserService
}

// Register adds the UserService to the otohttp.Server.
func RegisterUserService(server *otohttp.Server, userService UserService) {
	handler := &userServiceServer{
		server:      server,
		userService: userService,
	}
	server.Register("UserService", "AddUser", handler.handleAddUser)
	server.Register("UserService", "GetAccessToken", handler.handleGetAccessToken)
	server.Register("UserService", "GetMe", handler.handleGetMe)
}

func (s *userServiceServer) handleAddUser(w http.ResponseWriter, r *http.Request) {
	var request AddUserRequest
	if err := otohttp.Decode(r, &request); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	response, err := s.userService.AddUser(r.Context(), request)
	if err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	if err := otohttp.Encode(w, r, http.StatusOK, response); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
}

func (s *userServiceServer) handleGetAccessToken(w http.ResponseWriter, r *http.Request) {
	var request GetAccessTokenRequest
	if err := otohttp.Decode(r, &request); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	response, err := s.userService.GetAccessToken(r.Context(), request)
	if err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	if err := otohttp.Encode(w, r, http.StatusOK, response); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
}

func (s *userServiceServer) handleGetMe(w http.ResponseWriter, r *http.Request) {
	var request interface{}
	if err := otohttp.Decode(r, &request); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	response, err := s.userService.GetMe(r.Context(), request)
	if err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	if err := otohttp.Encode(w, r, http.StatusOK, response); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
}

type AddTrackerRequest struct {
	OwnerID string `json:"ownerID"`

	IsActive bool `json:"isActive"`
}

type AddTrackerResponse struct {
	Tracker models.Tracker `json:"tracker"`

	// Error is string explaining what went wrong. Empty if everything was fine.
	Error string `json:"error,omitempty"`
}

type AddUserRequest struct {
	Email string `json:"email"`

	Password string `json:"password"`
}

type AddUserResponse struct {
	User models.User `json:"user"`

	// Error is string explaining what went wrong. Empty if everything was fine.
	Error string `json:"error,omitempty"`
}

type DeleteTrackerResponse struct {

	// Error is string explaining what went wrong. Empty if everything was fine.
	Error string `json:"error,omitempty"`
}

type GetAccessTokenRequest struct {
	Email string `json:"email"`

	Password string `json:"password"`
}

type GetAccessTokenResponse struct {
	Token string `json:"token"`

	// Error is string explaining what went wrong. Empty if everything was fine.
	Error string `json:"error,omitempty"`
}

type GetMeResponse struct {
	User models.User `json:"user"`

	// Error is string explaining what went wrong. Empty if everything was fine.
	Error string `json:"error,omitempty"`
}

type Tracker struct {
	ID string `json:"id"`

	TimesAccessed int `json:"times_accessed"`

	OwnerID string `json:"owner_id"`

	URL string `json:"url"`

	IsActive bool `json:"is_active"`
}

type UpdateTrackerRequest struct {
	ID string `json:"id"`

	IsActive bool `json:"isActive"`
}

type UpdateTrackerResponse struct {
	Tracker models.Tracker `json:"tracker"`

	// Error is string explaining what went wrong. Empty if everything was fine.
	Error string `json:"error,omitempty"`
}

type User struct {
	ID string `json:"id"`

	Email string `json:"email"`

	PasswordHash string `json:"-"`
}

type trackerL struct {
}

type trackerR struct {
	Owner *models.User `json:"Owner"`
}

type userL struct {
}

type userR struct {
	OwnerTrackers models.TrackerSlice `json:"OwnerTrackers"`
}
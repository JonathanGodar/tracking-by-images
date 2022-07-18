package definition

import "github.com/JonathanGodar/go-web-gin/models"


// type User models.User;

type UserService interface {
	AddUser(AddUserRequest) AddUserResponse
	GetAccessToken(GetAccessTokenRequest) GetAccessTokenResponse
	GetMe(interface{}) GetMeResponse
}

type GetMeResponse struct {
	User models.User
}

type AddUserRequest struct {
	Email string
	Password string
}

type AddUserResponse struct {
	User models.User
}

type GetAccessTokenRequest struct {
	Email string
	Password string
}

type GetAccessTokenResponse struct {
	Token string
}

// type User struct {
// 	ID string
// 	Email string
// }

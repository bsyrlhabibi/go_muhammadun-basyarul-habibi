package models

type UserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ListUsersResponse struct {
	Data []UserResponse `json:"data"`
}

func NewUserResponse(user *User) *UserResponse {
	if user == nil {
		return nil
	}
	return &UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}
}

func NewListUsersResponse(users []User) *ListUsersResponse {
	if users == nil {
		return nil
	}
	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, *NewUserResponse(&user))
	}
	return &ListUsersResponse{Data: userResponses}
}

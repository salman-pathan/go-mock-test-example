package domain

type AddUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AddUserResponse struct {
	UserId string `json:"user_id"`
	Err    error  `json:"err,omitempty"`
}

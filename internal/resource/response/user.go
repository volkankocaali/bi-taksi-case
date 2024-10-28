package response

type ApiResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type UserResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Token   string  `json:"token"`
	User    UserDTO `json:"user"`
}

type UserDTO struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

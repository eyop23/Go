package dto

type UserRequest struct {
	Name string `json:"name"`
	Age int32 `json:"age"`
	Password string `json:"password'`
}

type LoginRequest struct {
	Name string `json:"name"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	ID int32 `json:"id"`
	Name string `json:"name"`
	Age int32 	`json:"age"` 
}
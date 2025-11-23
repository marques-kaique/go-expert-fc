package dto

type CreateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWTRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWTResponse struct {
	AccessToken string `json:"access_token"`
}

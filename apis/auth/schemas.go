package auth

//type EmailModel struct {
//	// email in email blacklist
//	Email string `json:"email" query:"email" validate:"isValidEmail"`
//}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" minLength:"3" validate:"required,min=3"`
	Type     string `json:"type" validate:"required"`
}

type TokenResponse struct {
	Access  string `json:"access,omitempty"`
	Message string `json:"message,omitempty"`
}

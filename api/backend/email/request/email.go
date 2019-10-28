package request

type SendEmailRequest struct {
	Email string `json:"email" valid:"Required"`
}

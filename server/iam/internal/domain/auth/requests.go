package auth

import "github.com/oqamase/ozon/iam/pkg/utils/validation"

type SignInRequest struct {
	Email    string `json:"email" validate:"required,min=4,max=100"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}

func (r *SignInRequest) Validate() error {
	return validation.ValidateObjectsByTags(r)
}

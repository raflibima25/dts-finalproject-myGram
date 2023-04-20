package models

type LoginResponse struct {
	Token string `json:"token"`
}

type ResponseFailed struct {
	Message string `json:"message"`
}

type ResponseFailedUnauthorized struct {
	Message string `json:"message"`
}

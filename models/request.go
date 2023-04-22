package models

type RequestUserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RequestUserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Age      uint   `json:"age"`
}

type RequestSocialMedia struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

type RequestPhoto struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

type RequestComment struct {
	Comment string `json:"comment_message"`
}

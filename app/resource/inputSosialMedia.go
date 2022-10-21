package resource


type InputSocialMedia struct{
	Name					string		`json:"name" binding:"required" example:"Instagram"`
	SocialMediaUrl			string		`json:"social_media_url" binding:"required" example:"www.instagram.com/uzumakinaruto"`
}
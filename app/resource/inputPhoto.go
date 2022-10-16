package resource


type InputPhoto struct{
	Title			string		`json:"title" binding:"required"`
	Caption			string		`json:"caption"`
	PhotoUrl		string		`json:"photo_url" binding:"required"`
}
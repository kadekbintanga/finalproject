package resource


type InputPhoto struct{
	Title			string		`json:"title" binding:"required" example:"Selfie with team"`
	Caption			string		`json:"caption" example:"This is my greate team"`
	PhotoUrl		string		`json:"photo_url" binding:"required" example:"https://static.wikia.nocookie.net/naruto/images/5/50/Team_Kakashi.png/revision/latest?cb=20161219035928"`
}
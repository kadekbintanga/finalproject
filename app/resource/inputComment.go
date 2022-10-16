package resource


type InputComment struct{
	Message			string		`json:"message" binding:"required"`
	PhotoID			uint		`json:"photo_id" binding:"required"`
}
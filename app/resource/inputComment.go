package resource


type InputComment struct{
	Message			string		`json:"message" binding:"required"  example:"good foto!"`
	PhotoID			uint		`json:"photo_id" binding:"required"  example:"1"`
}

type UpdateComment struct{
	Message			string		`json:"message" binding:"required"  example:"amazing photo"` 
}
package resource


type InputUser struct{
	Username		string		`json:"username" binding:"required" example:"testing"`
	Email			string		`json:"email" binding:"required" validate:"email" example:"testing@gmail.com"`
	Password		string		`json:"password" binding:"required,min=6" example:"test123"`
	Age				uint		`json:"age"  binding:"required,min=9" example:"20"`
}


type LoginUser struct{
	Email			string		`json:"email" binding:"required" validate:"email" example:"testing@gmail.com"`
	Password		string		`json:"password" binding:"required,min=6" example:"test123"`
}

type UpdateUser struct{
	Email			string		`json:"email" binding:"required" validate:"email" example:"testing@gmail.com"`
	Username		string		`json:"username" binding:"required" example:"testing"`
}
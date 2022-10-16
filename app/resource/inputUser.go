package resource


type InputUser struct{
	Username		string		`json:"username" binding:"required"`
	Email			string		`json:"email" binding:"required" validate:"email"`
	Password		string		`json:"password" binding:"required,min=6"`
	Age				uint		`json:"age"  binding:"required,min=9"`
}


type LoginUser struct{
	Email			string		`json:"email" binding:"required" validate:"email"`
	Password		string		`json:"password" binding:"required,min=6"`
}

type UpdateUser struct{
	Email			string		`json:"email" binding:"required" validate:"email"`
	Username		string		`json:"username" binding:"required"`
}
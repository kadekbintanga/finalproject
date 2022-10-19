package helpers


import(
	"strings"
	"fmt"
)

func FormatValidationErrorSQL(err error) string{
	fmt.Println("error SQL")
	if strings.Contains(err.Error(),"unique constraint"){
		if strings.Contains(err.Error(),"users_username_key"){
			return "Username has been used"
		}else if strings.Contains(err.Error(),"users_email_key"){
			return "Email has been used"
		}else{
			return "Something went wrong"
		}
	}else{
		return "Something went wrong"
	}
}

func FormatValidationErrorBinding(err error) string{
	fmt.Println("error Binding")
	if strings.Contains(err.Error(),"Username"){
		return "Username is required"
	}else if strings.Contains(err.Error(),"Email"){
		return "Email is required"
	}else if strings.Contains(err.Error(),"Password"){
		if strings.Contains(err.Error(),"min"){
			return "Password has to greater than 6 character"
		}else{
			return "Password is required"
		}
	}else if strings.Contains(err.Error(),"Age"){
		if strings.Contains(err.Error(),"min"){
			return "Age has to greater than 8 years old"
		}else{
			return "Age is required"
		}
	}else if strings.Contains(err.Error(),"Title"){
		return "Title is required"
	}else if strings.Contains(err.Error(),"PhotoUrl"){
		return "PhotoUrl is required"
	}else if strings.Contains(err.Error(),"Message"){
		return "Message is required"
	}else if strings.Contains(err.Error(),"PhotoID"){
		return "PhotoID is required"
	}else if strings.Contains(err.Error(),"Name"){
		return "Name of social media is required"
	}else if strings.Contains(err.Error(),"SocialMediaUr"){
		return "Url of social media  is required"
	}else if strings.Contains(err.Error(),"username"){
		return "username must has string type"
	}else if strings.Contains(err.Error(),"email"){
		return "email must has string type"
	}else if strings.Contains(err.Error(),"password"){
		return "password must has string type"
	}else if strings.Contains(err.Error(),"age"){
		return "age must has interger type"
	}else{
		return "Something went wrong"
	}
}

func FormatValidationErrorType(err error) string{
	fmt.Println("error Type")
	if strings.Contains(err.Error(),"username"){
		return "username must has string type"
	}else if strings.Contains(err.Error(),"email"){
		return "email must has string type"
	}else if strings.Contains(err.Error(),"password"){
		return "password must has string type"
	}else if strings.Contains(err.Error(),"age"){
		return "age must has interger type"
	}else{
		return "Something went wrong"
	}
}

func FormatValidationErrorPlayground(err error) string{
	fmt.Println("error Playgrond")
	if strings.Contains(err.Error(),"email"){
		return "Invalid format email"
	}
	return "Something went wrong"
}
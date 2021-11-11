package models

import(
	jwt "github.com/dgrijalva/jwt-go"
)

type User struct{
	Username, Gmail string	
}

type Token struct{
	jwt.StandardClaims
}

type Static struct{
	Token string	
	User User
}

type UserLogin struct{
	Login, Password string	
}

type ChangePass struct{
	Login, OldPassword, NewPassword string
}


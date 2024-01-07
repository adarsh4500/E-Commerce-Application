package models

type User struct{
	Fullname string `json:"fullname" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Creds struct{
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
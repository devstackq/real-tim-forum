package models

type PageData struct{
	PageTitle string
	Categories []string
	User User
	Data interface{}
}
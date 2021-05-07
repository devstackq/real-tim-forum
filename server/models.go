package models

type User struct {
	Name  string
	Age   int
	Email string
}

type Messages struct {
	FromWho int
	ToWho   int
	Content string
}

package models

type User struct {
	Name string `bson:"nome"`
	Age  int8   `bson:"idade"`
}

package model

type Repo struct {
	Id       string `bson:"_id"`
	Name     string
	Descript string
	Attr     uint32
}

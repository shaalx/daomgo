package main

import (
	"github.com/shaalx/daomgo"
	"gopkg.in/mgo.v2/bson"
	"log"
)

var (
	// MgoDB_ = daomgo.NewMgoDB("daocloud")
	MgoDB_ = daomgo.NewMgoDB(daomgo.Conn())
	usersC = MgoDB_.GetCollection([]string{"lEyTj8hYrUIKgMfi", "users"}...)
)

type User struct {
	Id   string `bson:"id"`
	Name string `bson:"name"`
}

func TestUpsert() {
	selector := bson.M{"id": "3"}
	usersC.Insert(User{"3", "three"})
	change := User{"3", "aaasss"}
	ret := usersC.Upsert(selector, change)
	log.Println(ret)

	change2 := User{"122", "Nil selector "}
	ret = usersC.Upsert(nil, change2)
	log.Println(ret)
}

func main() {
	TestUpsert()
	selector := bson.M{"name": bson.RegEx{"a", "s"}}
	log.Println(selector)
	ret := usersC.Like(selector)
	log.Println(ret)
	log.Println(len(ret))
}

package main

import (
	"github.com/shaalx/daomgo"
	"gopkg.in/mgo.v2/bson"
	"log"

	"html/template"
	"net/http"
)

var (
	// MgoDB_ = daomgo.NewMgoDB("daocloud")
	MgoDB_ = daomgo.NewMgoDB(daomgo.Conn())
	usersC = MgoDB_.GetCollection([]string{"IEyTj8hYrUIKgMfi", "users"}...)
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
	// TestUpsert()
	// selector := bson.M{"name": bson.RegEx{"a", "s"}}
	// log.Println(selector)
	// ret := usersC.Like(selector)
	// log.Println(ret)
	// log.Println(len(ret))

	// selector = bson.M{"id": "3"}
	// ret = usersC.ISelect(nil)
	// log.Println(ret)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/", index)
	http.ListenAndServe(":80", nil)
}

func index(rw http.ResponseWriter, req *http.Request) {
	tpl, _ := template.New("index.html").ParseFiles("index.html")
	data := usersC.ISelect(nil)
	log.Println(data)
	tpl.Execute(rw, data)
}

func insert(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	cont := req.FormValue("content")
	log.Println(cont)
	slt := bson.M{"id": "12"}
	usersC.Upsert(slt, User{Id: "12", Name: cont})
	http.Redirect(rw, req, "/", 302)
}

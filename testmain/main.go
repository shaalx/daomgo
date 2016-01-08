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
	usersC = MgoDB_.GetCollection([]string{"lEyTj8hYrUIKgMfi", "OJProbs"}...)
)

type Model struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	FuncName  string `json:"func_name"`
	Content   string `json:"content"`
	ArgsType  string `json:"args_type"`
	RetsType  string `json:"rets_type"`
	TestCases string `json:"test_cases"`
}

func main() {
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/", index)
	http.ListenAndServe(":80", nil)
}

func index(rw http.ResponseWriter, req *http.Request) {
	tpl, _ := template.New("index.html").ParseFiles("index.html")
	data := usersC.ISelect(nil)
	log.Println(data)
	tpl.Execute(rw, data)
}

func edit(rw http.ResponseWriter, req *http.Request) {
	tpl, _ := template.New("edit.html").ParseFiles("edit.html")
	tpl.Execute(rw, nil)
}

func insert(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	modle := Model{}
	modle.Id = req.FormValue("id")
	modle.Title = req.FormValue("title")
	modle.Desc = req.FormValue("desc")
	modle.FuncName = req.FormValue("func_name")
	modle.Content = req.FormValue("content")
	modle.ArgsType = req.FormValue("args_type")
	modle.RetsType = req.FormValue("rets_type")
	modle.TestCases = req.FormValue("test_cases")
	log.Println(modle)
	slt := bson.M{"id": modle.Id}
	usersC.Upsert(slt, modle)
	http.Redirect(rw, req, "/", 302)
}

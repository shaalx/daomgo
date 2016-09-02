package main

import (
	"github.com/toukii/daomgo"
	"gopkg.in/mgo.v2/bson"
	"log"

	"html/template"
	"net/http"

	"fmt"
	"github.com/everfore/exc"
	. "github.com/toukii/gooj"
	"github.com/toukii/goutils"
	"os"
	"strings"
)

var (
	// MgoDB_ = daomgo.NewMgoDB("daocloud")
	MgoDB_      = daomgo.NewMgoDB(daomgo.Conn())
	usersC      = MgoDB_.GetCollection([]string{"lEyTj8hYrUIKgMfi", "OJProbs"}...)
	defaultpath string
)

func init() {
	defaultpath, _ = os.Getwd()
}

func main() {
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/test", test)
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

func parseForm(rw http.ResponseWriter, req *http.Request) *Model {
	modle := Model{}
	modle.Id = req.FormValue("id")
	modle.Title = req.FormValue("title")
	modle.Desc = req.FormValue("desc")
	modle.FuncName = req.FormValue("func_name")
	modle.Content = req.FormValue("content")
	modle.ArgsType = req.FormValue("args_type")
	modle.RetsType = req.FormValue("rets_type")
	modle.TestCases = req.FormValue("test_cases")
	if strings.EqualFold("online", req.FormValue("online")) {
		modle.Online = true
	}
	return &modle
}

func insert(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	modle := parseForm(rw, req)
	log.Println(modle)
	slt := bson.M{"id": modle.Id}
	usersC.Upsert(slt, modle)
	http.Redirect(rw, req, "/", 302)
}

func test(rw http.ResponseWriter, req *http.Request) {
	cmd := exc.NewCMD("go test -v").Cd(defaultpath)
	model := parseForm(rw, req)
	fmt.Println(model)
	path_ := strings.Split(req.RemoteAddr, ":")[1]
	err := GenerateOjModle(path_, model)
	goutils.CheckErr(err)
	ret, err := cmd.Wd().Cd(path_).Debug().Do()
	goutils.CheckErr(err)
	rw.Write(ret)
	fmt.Println(goutils.ToString(ret))
	go cmd.Reset(fmt.Sprintf("rm -rf %s", path_)).Cd(defaultpath).ExecuteAfter(5)
}

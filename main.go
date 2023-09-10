package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"d1y.io/study/conf"
	"d1y.io/study/models"
	"d1y.io/study/template"
	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
)

type IBlog struct {
	Port int
}

func newBlog(port int) *IBlog {
	return &IBlog{
		Port: port,
	}
}

func (ig *IBlog) index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emptySlice = make([]string, 0)
	var result = models.ResultVO{
		Success: true,
		Msg:     "获取成功",
		Data:    emptySlice,
	}
	data, _ := json.Marshal(result)
	w.Write(data)
}

func (ig *IBlog) indexHtml(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "html")
	w.Write([]byte(template.IndexHtmlData))
}

func (ig *IBlog) pingPong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I'm live"))
}

func (ig *IBlog) Register() {
	http.HandleFunc("/", ig.index)
	http.HandleFunc("/index", ig.indexHtml)
	http.HandleFunc("/ping", ig.pingPong)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
}

func (ig *IBlog) Run() error {
	ig.Register()
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", ig.Port), nil); err != nil {
			log.Fatal(err)
		}
	}()
	say, _ := cowsay.Say(fmt.Sprintf("server is running http://localhost:%d", ig.Port))
	fmt.Println(say)
	select {}
}

func main() {
	conf.Init()
	blog := newBlog(conf.Instance.Port)
	err := blog.Run()
	if err != nil {
		panic(err)
	}
}

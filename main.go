package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
)

//go:embed template/index.html
var indexHtmlData string

type ResultVO struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    any    `json:"data"`
}

func (ro *ResultVO) SetData(data any) {
	ro.Data = data
}

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
	var result = ResultVO{
		Success: true,
		Msg:     "获取成功",
		Data:    emptySlice,
	}
	data, _ := json.Marshal(result)
	w.Write(data)
}

func (ig *IBlog) indexHtml(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "html")
	w.Write([]byte(indexHtmlData))
}

func (ig *IBlog) pingPong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I'm live"))
}

func (ig *IBlog) Register() {
	http.HandleFunc("/", ig.index)
	http.HandleFunc("/index", ig.indexHtml)
	http.HandleFunc("/ping", ig.pingPong)
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
	blog := newBlog(8080)
	err := blog.Run()
	if err != nil {
		panic(err)
	}
}

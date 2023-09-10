package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

func (ig *IBlog) Register() {
	http.HandleFunc("/", ig.index)
}

func (ig *IBlog) Run() error {
	ig.Register()
	if err := http.ListenAndServe(fmt.Sprintf(":%d", ig.Port), nil); err != nil {
		return err
	}
	return nil
}

func main() {
	blog := newBlog(8080)
	err := blog.Run()
	if err != nil {
		panic(err)
	}
}

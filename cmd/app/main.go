package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Directory struct {
	Path string
	FS   []os.DirEntry
}

func main() {
	http.HandleFunc("/", InitialLoad)
	http.HandleFunc("/get", GetDirectory)
	http.HandleFunc("/new", NewInstance)
	http.HandleFunc("/delete", DeleteInstance)

	log.Fatal(http.ListenAndServe(":80", nil))
}

func getDirectory(path string) Directory {
	var (
		err error
		dir Directory
	)

	dir.FS, err = os.ReadDir(path)
	if err != nil {
		log.Panic(err)
	}

	dir.Path = path
	return dir
}

func InitialLoad(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	dir := getDirectory("storage")
	tmpl.Execute(w, dir)
}

func GetDirectory(w http.ResponseWriter, r *http.Request) {
	dir := getDirectory(r.Header.Get("Hx-Trigger"))
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "dir-list", dir)
}

func NewInstance(w http.ResponseWriter, r *http.Request) {
	os.Mkdir("", os.ModePerm)
}

func DeleteInstance(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)
	//os.RemoveAll()
}

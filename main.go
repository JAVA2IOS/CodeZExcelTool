package main

import (
	"net/http"
	"log"
	"codezexcel/CodeZExcelTool/service/file"
	"html/template"
)

func HTMLPageRouter(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fallthrough
	case "/index":
		tpl, gloableErr := template.ParseFiles("view/web/index.html")
		if gloableErr != nil {
			http.NotFound(w, r)
			return 
		}
		tpl.Execute(w, nil)
	default:
		http.NotFound(w, r)
	}
}

func main() {
	http.Handle("/view/web/", http.StripPrefix("/view/web/", http.FileServer(http.Dir("view/web"))))
	// index
	http.HandleFunc("/", HTMLPageRouter)
	http.HandleFunc("/index", HTMLPageRouter)
	
	// xlsx
	http.HandleFunc("/api/v0/file/upload/xlsx", file.UploadXlsxFile)


	err := http.ListenAndServe(":8080", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err, "code: ", err.Error())
    }
}
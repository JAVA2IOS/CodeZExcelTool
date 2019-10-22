package file

import (
	"net/http"
	"fmt"
	"log"
	"time"
	"codezexcel/CodeZExcelTool/config"
	"os"
	"io"
)

func UploadXlsxFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//把上传的文件存储在内存和临时文件中
	formErr := r.ParseMultipartForm(32<<20)
	if formErr != nil {
		log.Printf("解析参数错误：%v", formErr.Error())
		return 
	}
	log.Printf("参数内容:%v,\n类型：%v \n", r.MultipartForm, r.MultipartForm.Value["type"])


    //获取文件句柄，然后对文件进行存储等处理
    file, handler, err := r.FormFile("file")
    if err != nil{
        fmt.Println("form file err: ", err)
        return
    }
    defer file.Close()
    fmt.Fprintf(w, "文件头%v\n文件名称", handler.Header.Get("Content-Disposition"), r)
    //创建上传的目的文件
    f, err := os.OpenFile("./upload/" + time.Now().Format(gloableConfig.TimeForamt_yyyy_MM_dd_hh_mm_ss) + handler.Filename, os.O_WRONLY | os.O_CREATE, 0666)
    if err != nil{
        fmt.Println("open file err: ", err)
        return
    }


    defer f.Close()
    //拷贝文件
    _, err = io.Copy(f, file)
    if err != nil {
    	fmt.Fprintf(w, "文件上传失败:%v", err.Error())
    	return 
    }

    fmt.Fprintf(w, "文件上传成功")
}
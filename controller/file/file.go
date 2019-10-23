package file

import (
	"net/http"
	"fmt"
	"log"
	"time"
	"codezexcel/CodeZExcelTool/config"
	service "codezexcel/CodeZExcelTool/service/file"
	"mime/multipart"
	"codezexcel/CodeZExcelTool/ExcelManager"
	"codezexcel/CodeZExcelTool/tool/excel/service"
	"encoding/json"
	// "sync"
)

type JsonHandler struct {
	Success bool `json:"success"`
	Data string `json:"data"`
}

func UploadXlsxFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//把上传的文件存储在内存和临时文件中
	formErr := r.ParseMultipartForm(32<<20)
	if formErr != nil {
		log.Printf("解析参数错误：%v", formErr.Error())
		return 
	}

	xlsxFile, xlsxHandler, xlsxErr := r.FormFile("xlsxFile")
	if xlsxErr != nil {
		log.Printf("xlsx file open failed：%v", xlsxErr.Error())
		return 
	}
	defer xlsxFile.Close()

	txtFile, txtHandler, txtErr := r.FormFile("txtFile")

    if txtErr != nil{
        fmt.Println("text file err: ", txtErr.Error())
        return
    }
    defer txtFile.Close()


    // filePaths := make(map[string]string)

    xlsxFileName := time.Now().Format(gloableConfig.TimeForamt_yyyy_MM_dd_hh_mm_ss) + xlsxHandler.Filename
    txtFileName := time.Now().Format(gloableConfig.TimeForamt_yyyy_MM_dd_hh_mm_ss) + txtHandler.Filename

    log.Printf("文件名: %v, %v \n", xlsxFileName, txtFileName)


    xlsxPath, xlsxErr := service.SaveFileToLocalPath(xlsxFileName, xlsxFile, xlsxHandler)

	if xlsxErr != nil {
		log.Printf("xlsx文件错误: %v", xlsxErr.Error())
		return 
	}

	txtPath, txtErr := service.SaveFileToLocalPath(txtFileName, txtFile, txtHandler)
	if txtErr != nil {
		log.Printf("text文件错误: %v", txtErr.Error())
		return 
	}

	newFileName := time.Now().Format(gloableConfig.TimeForamt_yyyy_MM_dd_hh_mm_ss) + "." + excel.FileTypeXlsx

	targetPath, err := excelManager.DJsTextMatchXlsxFile(xlsxPath, txtPath, newFileName)
	if err != nil {
		log.Printf("匹配数据失败 %v\n", err.Error())
		return 
	}

	// jsonString := "success:true,data:" + targetPath;

	// jsonData := []byte{jsonString}

	handler := JsonHandler{Success:true, Data:targetPath}

	jsonString, jsonErr := json.Marshal(handler)
	if jsonErr != nil {
		log.Printf("json解析失败 %v\n", jsonErr.Error())
		return 
	}

	fmt.Fprintf(w, string(jsonString))




 //    wg := sync.WaitGroup{}
 //    wg.Add(1)

 //    go SaveFiles(xlsxFileName, xlsxFile, xlsxHandler, filePaths, "xlsx", func(){
 //    	wg.Done()
 //    })
    
	// go SaveFiles(txtFileName, txtFile, txtHandler, filePaths, "txt", func(){
	// 	wg.Done()
	// })

	// wg.Wait()

	log.Printf("结束： %v", targetPath)
}

func SaveFiles(fileName string, file multipart.File, header *multipart.FileHeader, paths chan map[string]string, key string, callback func()) {
	path, err := service.SaveFileToLocalPath(fileName, file, header)

	if err != nil {
		log.Printf("文件错误: %v", err)
		callback()
		return 
	}

	log.Printf("文件地址: %v \n", path)

	pathMap := <- paths

	pathMap[key] = path

	paths <- pathMap

	callback()
}


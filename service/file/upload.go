package file

import (
	"log"
	"os"
	"io"
	"mime/multipart"
	"errors"
)

func SaveFileToLocalPath(fileName string, file multipart.File, header *multipart.FileHeader) (string, error) {
    //创建上传的目的文件
    filePath := "./upload/" + fileName

    f, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
    if err != nil {
        log.Printf("file open failed : %v \n", err.Error())
        return "", errors.New("文件创建失败: " + err.Error())
    }
    defer f.Close()

    //拷贝文件
    _, err = io.Copy(f, file)
    if err != nil {
    	log.Printf("file copied failed : %v \n", err.Error())
    	return "", errors.New("文件数据上传失败: %v" + err.Error())
    }

    return filePath, nil
}

package excelManager

import (
	"codezexcel/CodeZExcelTool/tool/excel/service"
	"codezexcel/CodeZExcelTool/tool/yaml/controller"
	"github.com/tealeg/xlsx"
	"codezexcel/CodeZExcelTool/tool/file"
	"errors"
	"log"
)

func DJsTextMatchXlsxFile(xlsxFilePath string, textFilePath string, targetFileName string) (string, error){
	reader, _ := yamlReader.Instance()

	sheet := excel.ReadFilteredExcelFile(xlsxFilePath)


	txtStrings := fileReader.ReadTxtFile(textFilePath)
	if txtStrings == nil {
		return "", errors.New("txt文件[" + xlsxFilePath + "]读取错误")
	}

	if sheet == nil {
		return "", errors.New("xlsx文件" + textFilePath + "读取失败")
	}

	newFile := xlsx.NewFile()
	newSheet, err := newFile.AddSheet("newSheet")
	if err != nil {
		log.Panicln("新建excel文件清单失败")
		return "", errors.New("新建清单失败:" + err.Error())
	}

	for rowIndex, row := range sheet.Rows {

		if rowIndex == 0 || len(row.Cells) <= 1 {
			continue 
		}

		targetCell := row.Cells[1]
		columnValue := targetCell.String()

		if len(columnValue) == 0 {
			continue
		}

		for _, text := range txtStrings {
			if text != columnValue || len(text) == 0 {
				continue 
			}
			newRow := newSheet.AddRow()
			for _, cell := range row.Cells {
				newCell := newRow.AddCell()
				newCell.SetValue(cell.String())
				newCell.SetStyle(xlsx.NewStyle())
			}
		}
	}

	if len(newSheet.Rows) == 0 {
		return "", errors.New("匹配数据失败")
	}

	newFilePath := reader.Configure.Xlsx.SavedDirctory + targetFileName

	newErr := newFile.Save(newFilePath)

	if err != nil {
		log.Printf("创建文件失败[%v]\n", newErr.Error())
		return "", errors.New("创建文件" + newFilePath + "失败:" + err.Error())
	}


	return newFilePath, nil
}
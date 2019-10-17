package excelManager

import (
	"../tool/yaml/controller"
	"../tool/excel/service"
	"github.com/tealeg/xlsx"
	"../tool/file"
	"time"
	"errors"
	"log"
	"../config"
)

func DJsTextMatchXlsxFile() (bool, error){
	reader, _ := yamlReader.Instance()

	sheet := excel.ReadFilteredExcelFile(reader.Configure.Xlsx.AbsolutePath)


	txtStrings := fileReader.ReadTxtFile(reader.Configure.Xlsx.MatchFilePath)
	if txtStrings == nil {
		return false, errors.New("txt文件[" + reader.Configure.Xlsx.MatchFilePath + "]读取错误")
	}

	if sheet == nil {
		return false, errors.New("xlsx文件" + reader.Configure.Xlsx.AbsolutePath + "读取失败")
	}

	newFile := xlsx.NewFile()
	newSheet, err := newFile.AddSheet("newSheet")
	if err != nil {
		log.Panicln("新建excel文件清单失败")
		return false, errors.New("新建清单失败:" + err.Error())
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
		return false, errors.New("匹配数据失败")
	}

	newFileName := time.Now().Format(gloableConfig.TimeForamt_yyyy_MM_dd_hh_mm_ss) + "." + excel.FileTypeXlsx
	newFilePath := reader.Configure.Xlsx.SavedDirctory + newFileName

	newErr := newFile.Save(newFilePath)
	if err != nil {
		log.Printf("创建文件失败[%v]\n", newErr.Error())
		return false, errors.New("创建文件" + newFilePath + "失败:" + err.Error())
	}

	return true, nil
}
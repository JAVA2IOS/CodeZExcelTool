package excel

import (
	"../dao"
	"github.com/tealeg/xlsx"
)

const (
	FileTypeXlsx = "xlsx"
)

// 读取xlsx的excel文件数据
func ReadFilteredExcelFile(filePath string) *xlsx.Sheet {
	file, err := excel.ReadXlsxFile(filePath)
	if err != nil {
		return nil
	}

	if len(file.Sheets) == 0 {
		return nil
	}

	return file.Sheets[0]
}

// 比对数据，并且标红筛选后的数据
// func updateFilteredDataHighlightStyle() {
// 	sheet := ReadFilteredExcelFile(XlsxFilePath)
// 	if sheet == nil {
// 	}

// 	sheet.Row(1).Cells
// }

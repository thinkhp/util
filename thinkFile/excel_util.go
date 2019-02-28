package thinkFile

import (
	"github.com/tealeg/xlsx"
	"strconv"
	"util/think"
)

func StringSliceWriteExcel(filePath, fileName string, slice [][]string, columns []string) string {
	// 创建目录
	CreatePath(filePath)
	// 创建 EXCEL
	excel := xlsx.NewFile()
	//
	sheet, err := excel.AddSheet(strconv.Itoa(len(excel.Sheets)))
	think.IsNil(err)

	// 表头
	width := len(columns)
	row := sheet.AddRow()
	for i := 0; i < width; i++ {
		cell := row.AddCell()
		cell.SetString(columns[i])
	}

	// 主体
	for i := 0; i < len(slice); i++ {
		row := sheet.AddRow()
		//row.WriteSlice(slice[i],width)
		for j := 0; j < width; j++ {
			cell := row.AddCell()
			cell.SetString(slice[i][j])
		}
	}

	// 保存并关闭
	err = excel.Save(filePath + fileName)
	think.IsNil(err)

	return filePath + fileName
}

func SliceWriteExcel(filePath, fileName string, slice [][]interface{}, columns []string) {
	// 创建 EXCEL
	excel := xlsx.NewFile()
	//
	sheet, err := excel.AddSheet("")
	think.IsNil(err)

	// 表头
	width := len(columns)
	row := sheet.AddRow()
	for i := 0; i < width; i++ {
		cell := row.AddCell()
		cell.SetString(columns[i])
	}

	//
	for i := 0; i < len(slice); i++ {
		row := sheet.AddRow()
		for j := 0; j < width; j++ {
			cell := row.AddCell()
			cell.SetValue(slice[i][j])
		}
	}

	//
	err = excel.Save(filePath + fileName)
	think.IsNil(err)
}

//// SetInt sets a cell's value to an integer.
//func (c *Cell) SetValue(n interface{}) {
//	switch t := n.(type) {
//	case time.Time:
//		c.SetDateTime(t)
//		return
//	case int, int8, int16, int32, int64:
//		c.setNumeric(fmt.Sprintf("%d", n))
//	case float64:
//		// When formatting floats, do not use fmt.Sprintf("%v", n), this will cause numbers below 1e-4 to be printed in
//		// scientific notation. Scientific notation is not a valid way to store numbers in XML.
//		// Also not not use fmt.Sprintf("%f", n), this will cause numbers to be stored as X.XXXXXX. Which means that
//		// numbers will lose precision and numbers with fewer significant digits such as 0 will be stored as 0.000000
//		// which causes tests to fail.
//		c.setNumeric(strconv.FormatFloat(t, 'f', -1, 64))
//	case float32:
//		c.setNumeric(strconv.FormatFloat(float64(t), 'f', -1, 32))
//	case string:
//		c.SetString(t)
//	case []byte:
//		c.SetString(string(t))
//	case nil:
//		c.SetString("")
//	default:
//		c.SetString(fmt.Sprintf("%v", n))
//	}
//}

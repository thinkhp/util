package thinkFile

import (
	"util/database"
	"github.com/tealeg/xlsx"
	"util/think"
)

type ExcelTable struct {
	fileName string
	sheetIndex int
	tableName string

 	heads1 map[string]string //excel头部与table头部的对应关系
 	heads2 map[string]string

 	excelHeads []string
 	tableHeads []string
}

func (e *ExcelTable)Init(fileName string, sheetIndex int, tableName string, excelH2tableH map[string]string){
	e.fileName = fileName
	e.sheetIndex =sheetIndex
	e.tableName = tableName
	e.heads1 = excelH2tableH
	for eh, th := range e.heads1 {
		e.heads2[th] = eh
	}
}
func (e *ExcelTable)ExcelToTable(){
	// 打开表格
	excel, err := xlsx.OpenFile(e.fileName)
	think.IsNil(err)
	sheet := excel.Sheets[e.sheetIndex]
	// 获取表格表头 strings
	rowFirst := sheet.Rows[0].Cells
	strings := make([]string, 0)
	for i := 0; i < len(rowFirst); i++ {
		strings = append(strings, rowFirst[i].Value)
	}
	// 获取表格表数据 excelMaps  []map[string]string
	rowMaps := make([]map[string]string, 0)
	for i := 1; i < sheet.MaxRow; i++ {
		row := sheet.Rows[i].Cells
		rowMap := make(map[string]string)
		for j := 0; j < len(row); j++ {
			rowMap[string(strings[j])] = row[j].Value
		}
		rowMaps = append(rowMaps, rowMap)
	}
	// 重排表数据 excelData
	tableData := make([][]string, 0)
	// 对单行进行重排
	for i := 0; i < len(rowMaps); i++{
		rowT := make([]string, 0)
		//
		for j := 0; j < len(e.tableHeads); j++ {
			
			eh := e.heads2[e.tableHeads[j]]
			//
			rowT = append(rowT, rowMaps[i][eh])
		}
		tableData = append(tableData, rowT)
	}
	//
	// 拼接插入语句
	// INSERT INTO table_name (col1,col2,col3...) VALUES (v1,v2,v3...),(v1,v2,v3...),(v1,v2,v3...)
	//tx *sql.Tx, tableName string, cols []string, values [][]string
	database.InsertBatch(nil, e.tableName, e.tableHeads, tableData)

}

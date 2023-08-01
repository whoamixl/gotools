package main

import (
	"fmt"
	"github.com/whoamixl/gotools/utils/ExcelUtil"
)

func main() {
	excel := ExcelUtil.NewExcelUtil("./test.xlsx")
	excel.AddSheet("Sheet1")
	excel.Write("Sheet1", []string{"Hello", "World"})
	excel.Write("Sheet1", []string{"你好", "世界"})
	excel.Write("Sheet1", []string{"你好", "世界"})
	excel.Save()
	fmt.Println("Excel file saved as output.csv")
}

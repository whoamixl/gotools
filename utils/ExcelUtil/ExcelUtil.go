package ExcelUtil

import (
	"github.com/whoamixl/gotools/utils/FileUtil"
	"github.com/whoamixl/gotools/utils/SliceUtil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type ExcelUtil struct {
	file   *os.File
	sheets map[string][][]string
}

func NewExcelUtil(path string) *ExcelUtil {
	if !isValidFilePath(path) {
		log.Fatal("文件路径不合法")
		return nil
	}
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &ExcelUtil{
		file:   f,
		sheets: make(map[string][][]string),
	}
}
func (e *ExcelUtil) AddSheet(name string) {
	e.sheets[name] = [][]string{}
}
func (e *ExcelUtil) Write(sheet string, row []string) {
	e.sheets[sheet] = append(e.sheets[sheet], row)
}
func (e *ExcelUtil) Save() {
	defer e.file.Close()
	for sheet, rows := range e.sheets {
		_, err := e.file.WriteString(sheet + "\n")
		if err != nil {
			log.Fatal(err)
		}
		for _, row := range rows {
			_, err = e.file.WriteString(strings.Join(row, ",") + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
func isValidFilePath(filePath string) bool {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return false
	}
	fileTypes := []FileUtil.FileType{FileUtil.XLSX, FileUtil.XLS, FileUtil.CSV}
	if !SliceUtil.Contains(fileTypes, FileUtil.GetFileExt(absPath)) {
		return false
	}
	return true
}

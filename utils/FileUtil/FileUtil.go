package FileUtil

import "path/filepath"

// 获取文件后缀
func GetFileExt(filePath string) FileType {
	return fileTypeFromExt(filepath.Ext(filePath))
}

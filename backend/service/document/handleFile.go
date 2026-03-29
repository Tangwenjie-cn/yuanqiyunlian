/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package document

import (
	"errors"
	"path/filepath"
)

// 声明文档接口
type document interface {
	fileToImage(filePath string) (map[string]any, error)
}

func HandleFile(filePath string) (map[string]any, error) {
	ext := filepath.Ext(filePath)
	fileType := getFileType(ext)
	switch fileType {
	case "office", "openDocument":
		var office document = &office{}
		return office.fileToImage(filePath)
	case "pdf":
		var pdf document = &pdf{}
		return pdf.fileToImage(filePath)
	default:
		return nil, errors.New("不支持的文件类型:" + filePath)
	}
}
func getFileType(ext string) string {
	switch ext {
	case ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx":
		return "office"
	case ".pdf":
		return "pdf"
	case ".odt", ".ods", ".odp", ".odg":
		return "openDocument"
	default:
		return ""
	}
}

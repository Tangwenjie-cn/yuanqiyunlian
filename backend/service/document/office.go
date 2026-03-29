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
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// 将Office文件转换为pdf
func officeToPdf(inputPath string) (string, error) {
	// 检查输入文件是否存在
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		return "", errors.New("输入文件不存在")
	}

	// 生成输出文件路径
	dir := filepath.Dir(inputPath)
	filename := strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath))
	pdfPath := filepath.Join(dir, filename+".pdf")
	// 构建命令参数
	args := []string{
		"--headless",
		"--convert-to",
		"pdf",
		inputPath,
		"--outdir",
		dir,
	}
	// 执行转换命令
	cmd := exec.Command("soffice", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("转换失败: %v, 输出: %s", err, string(output))
	}

	return pdfPath, nil
}

type office struct{}

func (o *office) fileToImage(inputPath string) (map[string]any, error) {
	pdfPath, error := officeToPdf(inputPath)
	if error != nil {
		return nil, error
	}
	var pdf document = &pdf{}
	data, err := pdf.fileToImage(pdfPath)
	if err != nil {
		return nil, err
	}
	return data, nil

}

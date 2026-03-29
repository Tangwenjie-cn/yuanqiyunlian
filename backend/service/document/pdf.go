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
	"gin/service/set"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

// 获取PDF页数
func getPageCount(pdfPath string) (int, error) {
	// 检查文件是否存在
	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		return 0, fmt.Errorf("文件不存在: %s", pdfPath)
	}
	ctx, err := api.ReadContextFile(pdfPath)
	if err != nil {
		return 0, err
	}

	return ctx.PageCount, nil
}

// WebImageOption 网络图片转换选项
type WebImageOption struct {
	StartPage int    // 起始页码（1-based）
	EndPage   int    // 结束页码（-1表示最后一页）
	OutputDir string // 输出目录
	DPI       int    // 分辨率（网络推荐72-96，最高不超过150）
	Format    string // 图片格式
	PageCount int    // 总页数
}

type pdf struct{}

func (p pdf) fileToImage(inputPath string) (map[string]any, error) {
	config := set.GetAllSet()
	doc_start_page, _ := strconv.Atoi(config["doc_start_page"])
	doc_end_page, _ := strconv.Atoi(config["doc_end_page"])
	doc_img_dpi, _ := strconv.Atoi(config["doc_img_dpi"])
	// 获取文件名（包含扩展名）
	filename := filepath.Base(inputPath)
	// 去除扩展名，得到纯文件名
	dirName := strings.TrimSuffix(filename, filepath.Ext(filename))
	// 获取文件所在目录
	parentDir := filepath.Dir(inputPath)
	// 构建新目录的完整路径
	outputDir := filepath.Join(parentDir, dirName)
	opt := WebImageOption{
		StartPage: doc_start_page,
		EndPage:   doc_end_page,
		OutputDir: outputDir,
		DPI:       doc_img_dpi,
		Format:    "png",
	}
	pageCount, err := getPageCount(inputPath)
	if err != nil {
		return nil, err
	}
	opt.PageCount = pageCount
	images, err := PDFToWebImages(inputPath, opt)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"images": images,
		"count":  pageCount,
	}, nil

}

// PDFToWebImages 转换PDF为网络图片
func PDFToWebImages(pdfPath string, opt WebImageOption) ([]string, error) {
	// 验证基础参数
	if err := validateWebOption(pdfPath, opt); err != nil {
		return nil, err
	}

	// 创建输出目录
	if err := os.MkdirAll(opt.OutputDir, 0755); err != nil {
		return nil, fmt.Errorf("创建输出目录失败: %v", err)
	}

	// 构建输出文件名（带格式后缀）
	outputPattern := filepath.Join(opt.OutputDir, fmt.Sprintf("page_%%d.%s", opt.Format))

	// 构建命令参数
	args, err := buildConversionArgs(pdfPath, outputPattern, opt)
	if err != nil {
		return nil, err
	}
	// 执行转换命令
	cmd := exec.Command("mutool", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("转换失败: %v, 详情: %s", err, string(output))
	}

	// 生成预期的图片路径列表
	return generateWebImagePaths(opt), nil
}

// 验证网络图片转换选项
func validateWebOption(pdfPath string, opt WebImageOption) error {
	// 检查PDF文件是否存在
	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		return fmt.Errorf("PDF文件不存在: %s", pdfPath)
	}

	// 验证页码
	if opt.StartPage < 1 {
		return errors.New("起始页码必须≥1")
	}
	if opt.EndPage != -1 && opt.EndPage < opt.StartPage {
		return errors.New("结束页码必须≥起始页码")
	}

	// 验证分辨率（网络场景无需过高）
	if opt.DPI < 72 || opt.DPI > 150 {
		return errors.New("网络展示推荐DPI: 72-150（平衡清晰度和大小）")
	}

	return nil
}

// 构建转换参数
func buildConversionArgs(pdfPath, outputPattern string, opt WebImageOption) ([]string, error) {
	// 基础转换参数
	args := []string{
		"draw",
		"-r",
		strconv.Itoa(opt.DPI), // 分辨率设置
	}

	// 路径处理（Windows兼容：将反斜杠转为正斜杠）
	if runtime.GOOS == "windows" {
		outputPattern = strings.ReplaceAll(outputPattern, "\\", "/")
		pdfPath = strings.ReplaceAll(pdfPath, "\\", "/")
	}

	// 添加输出文件路径和输入PDF路径
	args = append(args,
		"-o",
		outputPattern,
		pdfPath,
	)

	// 设置页码范围
	if opt.StartPage < 1 {
		return nil, errors.New("起始页码必须大于0")
	}
	if opt.EndPage == -1 {
		args = append(args, fmt.Sprintf("%d-%d", opt.StartPage, opt.PageCount))
	} else {
		args = append(args, fmt.Sprintf("%d-%d", opt.StartPage, opt.EndPage))
	}
	fmt.Println(args)
	return args, nil
}

// 生成网络图片路径列表
func generateWebImagePaths(opt WebImageOption) []string {
	var paths []string
	if opt.EndPage == -1 {
		opt.EndPage = opt.PageCount
	}

	for i := opt.StartPage; i <= opt.EndPage; i++ {
		path := filepath.Join(
			opt.OutputDir,
			fmt.Sprintf("page_%d.%s", i, opt.Format),
		)
		paths = append(paths, path)
	}
	return paths
}

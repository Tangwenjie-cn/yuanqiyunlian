/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package util

import (
	"archive/zip"
	"bufio"
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"image/color"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 获取配置文件地址
func GetConfigPath() (string, error) {

	var configPath string
	// 定义命令行参数
	flag.StringVar(&configPath, "config", "", "Path to config file")
	// 解析命令行参数
	flag.Parse()
	if configPath == "" {

		dir, err := RootPath()
		if err != nil {
			return "", err
		}
		// 构造配置文件路径
		configPath = filepath.Join(dir, "config.toml")
	}

	// 检查文件是否存在
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return "", fmt.Errorf("配置文件不存在: %s", configPath)
	}

	return configPath, nil
}
func RootPath() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	dir := filepath.Dir(exePath)
	if runtime.GOOS == "windows" {
		if strings.Contains(dir, "\\tmp") { //判断是否在air临时目录下
			dir = filepath.Dir(dir)
		}
	}
	return dir, nil
}
func Sha256Hash(str string) string {
	// 创建一个新的SHA256哈希对象
	h := sha256.New()
	// 将字符串转换为字节切片并写入哈希对象
	h.Write([]byte(str))
	// 返回哈希值的十六进制字符串表示
	return fmt.Sprintf("%x", h.Sum(nil))
}
func Md5Hash(str string) string {
	// 创建一个新的MD5哈希对象
	h := md5.New()
	// 将字符串转换为字节切片并写入哈希对象
	h.Write([]byte(str))
	// 返回哈希值的十六进制字符串表示
	return fmt.Sprintf("%x", h.Sum(nil))

}

// 格式化时间为 "几秒前/几分钟前/几小时前/几天前/几月前/几年前"
func TimeAgo(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)
	// 按时间差从大到小判断
	switch {
	case diff < time.Minute:
		return fmt.Sprintf("%d秒前", int(diff.Seconds()))
	case diff < time.Hour:
		return fmt.Sprintf("%d分钟前", int(diff.Minutes()))
	case diff < 24*time.Hour:
		return fmt.Sprintf("%d小时前", int(diff.Hours()))
	case diff < 30*24*time.Hour:
		return fmt.Sprintf("%d天前", int(diff.Hours()/24))
	case diff < 12*30*24*time.Hour:
		return fmt.Sprintf("%d个月前", int(diff.Hours()/(30*24)))
	default:
		return fmt.Sprintf("%d年前", int(diff.Hours()/(365*24)))
	}
}

// 数字转成中文“百千万亿”表示
func FormatNumber(num float64) string {
	if num < 1000 {
		return fmt.Sprintf("%.0f", num)
	} else if num < 10000 {
		return fmt.Sprintf("%.1f千", num/1000)
	} else if num < 100000000 {
		return fmt.Sprintf("%.1f万", num/10000)
	} else {
		return fmt.Sprintf("%.1f亿", num/100000000)
	}
}

// HasIntersection 判断两个int切片是否存在交集
func HasIntersection(a, b []int) bool {
	// 创建map存储第一个切片的元素
	set := make(map[int]struct{}, len(a))
	for _, num := range a {
		set[num] = struct{}{}
	}

	// 遍历第二个切片，检查是否有元素在map中
	for _, num := range b {
		if _, exists := set[num]; exists {
			return true
		}
	}
	return false
}

// GenerateOrderNo 生成订单号：年月日+微秒
func GenerateOrderNo() string {
	now := time.Now()

	// 1. 年月日（如 20231005）
	date := now.Format("20060102")

	// 2. 微秒数（当前时间到 Unix 纪元的微秒数）
	micro := now.UnixMicro()

	// 拼接：年月日（8位）+ 微秒（13位）
	orderNo := fmt.Sprintf("%s%d", date, micro)
	return orderNo
}

// GetDayStartAndEnd 根据传入的时间，返回当天的开始时间和结束时间
// start: 当天 00:00:00
// end: 当天 23:59:59.999999999
func GetDayStartAndEnd(t time.Time) (start time.Time, end time.Time) {
	// 获取时间的年、月、日
	year, month, day := t.Date()

	// 构造当天的开始时间：00:00:00
	start = time.Date(year, month, day, 0, 0, 0, 0, t.Location())

	// 构造当天的结束时间：23:59:59.999999999
	// 注意：使用 24:00:00 减去 1 纳秒，比直接写 23:59:59 更准确（包含纳秒级）
	end = time.Date(year, month, day, 23, 59, 59, 999999999, t.Location())

	return start, end
}

// 获取当日开始结束时间
func GetTodayStartEnd() (start, end time.Time) {
	now := time.Now()
	start = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	end = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())
	return
}

// 获取本周开始结束时间
func GetWeekStartEnd() (start, end time.Time) {
	now := time.Now()
	// 计算本周的开始时间（周一）
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7 // 将周日视为一周的最后一天
	}
	start = now.AddDate(0, 0, -weekday+1).Truncate(24 * time.Hour)
	// 计算本周的结束时间（周日）
	end = start.AddDate(0, 0, 7).Add(-time.Nanosecond)
	return
}

// 获取本月开始结束时间
func GetMonthStartEnd() (start, end time.Time) {
	now := time.Now()
	// 计算本月的开始时间（1号）
	start = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()) // 计算本月的结束时间（下个月的1号减去1秒）
	end = time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, -1, now.Location())
	return
}

// 16进制转RGBA
func HexToRGBA(hex string) (color.RGBA, error) {
	var rgba color.RGBA

	if len(hex) > 0 && hex[0] == '#' {
		hex = hex[1:]
	}

	// 处理短格式，通过重复字符转换为长格式
	switch len(hex) {
	case 3:
		hex = string(hex[0]) + string(hex[0]) + string(hex[1]) + string(hex[1]) + string(hex[2]) + string(hex[2])
	case 4:
		hex = string(hex[0]) + string(hex[0]) + string(hex[1]) + string(hex[1]) + string(hex[2]) + string(hex[2]) + string(hex[3]) + string(hex[3])
	case 6:
		// 长格式 RGB，添加默认的 AA
		hex += "FF"
	case 8:
		// 长格式 RGBA，无需处理
	default:
		return rgba, fmt.Errorf("无效的十六进制颜色长度: %d", len(hex))
	}

	// 现在 hex 应该是 8 位长 (RRGGBBAA)
	val, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return rgba, err
	}

	rgba.A = uint8(val & 0xFF)
	val >>= 8
	rgba.B = uint8(val & 0xFF)
	val >>= 8
	rgba.G = uint8(val & 0xFF)
	val >>= 8
	rgba.R = uint8(val & 0xFF)

	return rgba, nil
}

// 网络请求函数
func Request(reqData map[string]any) ([]byte, error) {
	url, ok := reqData["url"].(string)
	if !ok {
		return nil, errors.New("url is not string")
	}
	method, ok := reqData["method"].(string)
	if !ok {
		return nil, errors.New("method is not string")
	}
	var body io.Reader
	data, ok := reqData["data"]
	if ok {
		var buf bytes.Buffer
		// 创建 JSON 编码器并关闭 HTML 字符转义（核心修改）
		encoder := json.NewEncoder(&buf)
		encoder.SetEscapeHTML(false) // 关闭 &/<>/'" 等字符的转义
		// 执行序列化
		if err := encoder.Encode(data); err != nil {
			return nil, err
		}
		// 移除 encoder.Encode 自动添加的换行符（关键！）
		jsonData := buf.Bytes()
		if len(jsonData) > 0 && jsonData[len(jsonData)-1] == '\n' {
			jsonData = jsonData[:len(jsonData)-1]
		}
		body = bytes.NewBuffer(jsonData)
	} else {
		body = nil
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	header, ok := reqData["header"].(map[string]string)
	if ok {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	} else {
		if data != nil {
			req.Header.Add("content-type", "application/json")
		}
	}
	timeout := 10
	_, ok = reqData["timeout"]
	if ok {
		if timeout, ok = reqData["timeout"].(int); !ok {
			return nil, errors.New("timeout is not int")
		}
	}

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http request failed with status code: %d", resp.StatusCode)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, err
}

// 获取请求的域名
func GetRequestDomain(c *gin.Context) string {
	// 首先检查 X-Forwarded-Host（反向代理）
	forwardedHost := c.GetHeader("X-Forwarded-Host")
	if forwardedHost != "" {
		// 可能有多个值，取第一个
		hosts := strings.Split(forwardedHost, ",")
		if len(hosts) > 0 {
			forwardedHost = strings.TrimSpace(hosts[0])
		}
	}

	// 确定最终的 Host
	finalHost := forwardedHost
	if finalHost == "" {
		finalHost = c.Request.Host
	}

	// 移除端口号（如果需要）
	finalHost = strings.Split(finalHost, ":")[0]

	return finalHost
}

// Unzip 解压 ZIP 文件到指定目录
// zipPath: ZIP 文件路径
// destDir: 解压目标目录
func Unzip(zipPath, destDir string) error {
	// 打开 ZIP 文件，获取文件信息
	zipFile, err := os.Open(zipPath)
	if err != nil {
		return fmt.Errorf("打开 ZIP 文件失败: %w", err)
	}
	defer zipFile.Close()

	// 获取 ZIP 文件大小（zip.NewReader 需要）
	zipFileInfo, err := zipFile.Stat()
	if err != nil {
		return fmt.Errorf("获取 ZIP 文件信息失败: %w", err)
	}

	// 创建 ZIP 读取器
	zipReader, err := zip.NewReader(zipFile, zipFileInfo.Size())
	if err != nil {
		return fmt.Errorf("创建 ZIP 读取器失败: %w", err)
	}

	// 确保目标目录存在（不存在则创建）
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return fmt.Errorf("创建目标目录失败: %w", err)
	}

	// 遍历 ZIP 中的每个文件/目录项
	for _, file := range zipReader.File {
		// 获取完整路径
		filePath := filepath.Join(destDir, file.Name)

		// 创建文件/目录
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(filePath, 0755); err != nil {
				return fmt.Errorf("创建目录失败: %w", err)
			}
		} else {
			// 确保父目录存在
			parentDir := filepath.Dir(filePath)
			if err := os.MkdirAll(parentDir, 0755); err != nil {
				return fmt.Errorf("创建目录失败: %w", err)
			}

			// 打开源文件（ZIP中的文件）
			rc, err := file.Open()
			if err != nil {
				return fmt.Errorf("打开 ZIP 文件中的文件失败: %w", err)
			}

			// 创建目标文件
			outFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
			if err != nil {
				rc.Close()
				return fmt.Errorf("创建文件失败: %w", err)
			}

			// 复制文件内容
			_, err = io.Copy(outFile, rc)
			outFile.Close()
			rc.Close()

			if err != nil {
				return fmt.Errorf("解压文件失败: %w", err)
			}
		}
	}
	return nil
}

// CopyFile 简洁版文件复制：srcPath 源文件路径，destDir 目标目录
func CopyFile(srcPath, destDir string) error {
	// 1. 打开源文件（只读模式）
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// 2. 拼接目标文件路径（保留源文件名）
	fileName := filepath.Base(srcPath)
	destPath := filepath.Join(destDir, fileName)

	// 3. 确保目标目录存在
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	// 4. 创建目标文件
	destFile, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// 5. 流式复制内容（边读边写，不占内存）
	_, err = io.Copy(destFile, srcFile)
	return err
}

// GetPublicIP 通过公网接口获取公网IP
func GetPublicIP() (string, error) {
	// 定义公网IP查询接口列表（可替换为其他可靠接口）
	urls := []string{
		"https://ifconfig.me/ip",
		"https://icanhazip.com",
		"https://api.ip.sb/ip",
	}

	// 设置HTTP客户端超时
	client := &http.Client{Timeout: 5 * time.Second}

	for _, url := range urls {
		resp, err := client.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		// 读取响应内容
		scanner := bufio.NewScanner(resp.Body)
		if scanner.Scan() {
			ipStr := strings.TrimSpace(scanner.Text())
			// 验证是否为合法IP
			if net.ParseIP(ipStr) != nil {
				return ipStr, nil
			}
		}
	}

	return "", fmt.Errorf("无法获取公网IP，所有接口请求失败")
}

/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package image

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"math"
	"os"

	"github.com/disintegration/imaging"
	"github.com/flopp/go-findfont"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
)

// getSystemDefaultFont 根据系统获取中文字体
func getSystemDefaultFont() (*sfnt.Font, error) {
	var fontPath string
	fontNames := []string{"simsun", "msyh", "simhei", "wqy", "NotoSans", "SourceHan"} // 根据系统选择字体文件路径
	for _, fontName := range fontNames {
		var err error
		fontPath, err = findfont.Find(fontName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if fontPath != "" {
			break
		}
	}
	if fontPath == "" {
		return nil, errors.New("未找到系统默认字体")
	}
	fontBytes, err := os.ReadFile(fontPath)
	if err != nil {
		return nil, err
	}
	// 解析为字体集合（处理 .ttc 情况）
	fontCollection, err := opentype.ParseCollection(fontBytes)
	if err != nil {
		return nil, err
	}

	// 4. 从集合中提取单个字体（索引 0 为第一个字体，可根据需求调整索引）
	fontIndex := 0
	ttfFont, err := fontCollection.Font(fontIndex)
	if err != nil {
		return nil, err
	}
	return ttfFont, nil
}

// 绘制水平居中且自动换行的文字
// img: 目标图片
// text: 文字内容
// face: 字体样式
// color: 文字颜色
// topY: 距离顶部的Y坐标
// lineHeight: 行高
// maxLineWidth: 每行最大宽度（像素）
func drawCenteredText(img *image.NRGBA, text string, face font.Face, color color.Color, topY, lineHeight float64, maxLineWidth int) {
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color),
		Face: face,
	}

	// 分割文字为多行（按宽度自动换行）
	lines := splitTextToLines(text, face, maxLineWidth)

	// 计算第一行的Y坐标（文字基线位置，需加上字体 ascent）
	// ascent 是字体从基线到顶部的高度
	ascent := face.Metrics().Ascent >> 6 // 转换为像素
	currentY := topY + float64(ascent)

	// 逐行绘制文字
	for _, line := range lines {
		// 计算当前行的宽度
		lineWidth := getTextWidth(line, face)
		// 水平居中：X坐标 = (图片宽度 - 行宽) / 2
		x := (img.Bounds().Dx() - lineWidth) / 2
		// 设置绘制位置
		d.Dot = fixed.Point26_6{
			X: fixed.Int26_6(x * 64),
			Y: fixed.Int26_6(currentY * 64),
		}
		// 绘制当前行
		d.DrawString(line)
		// 更新下一行的Y坐标
		currentY += lineHeight
	}
}
func drawLinesText(img *image.NRGBA, text string, face font.Face, color color.Color, startX, topY, lineHeight float64, maxLineWidth int) {
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color),
		Face: face,
	}

	// 分割文字为多行（按宽度自动换行）
	lines := splitTextToLines(text, face, maxLineWidth)

	// 计算第一行的Y坐标（文字基线位置，需加上字体 ascent）
	// ascent 是字体从基线到顶部的高度
	ascent := face.Metrics().Ascent >> 6 // 转换为像素
	currentY := topY + float64(ascent)

	// 逐行绘制文字
	for _, line := range lines {
		// 设置绘制位置
		d.Dot = fixed.Point26_6{
			X: fixed.Int26_6(startX * 64),
			Y: fixed.Int26_6(currentY * 64),
		}
		// 绘制当前行
		d.DrawString(line)
		// 更新下一行的Y坐标
		currentY += lineHeight
	}
}

// 按最大宽度分割文字为多行
func splitTextToLines(text string, face font.Face, maxWidth int) []string {
	var lines []string
	currentLine := ""
	currentWidth := 0

	for _, r := range text {
		// 测量当前字符的宽度
		charWidth := getTextWidth(string(r), face)
		// 如果加上当前字符超过最大宽度，则换行
		if currentWidth+charWidth > maxWidth && currentLine != "" {
			lines = append(lines, currentLine)
			currentLine = string(r)
			currentWidth = charWidth
		} else {
			currentLine += string(r)
			currentWidth += charWidth
		}
	}

	// 添加最后一行
	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return lines
}

// 计算文字宽度（像素）
func getTextWidth(text string, face font.Face) int {
	d := font.Drawer{Face: face}
	return d.MeasureString(text).Round()
}

// 创建圆形蒙版
func createCircleMask(diameter int) *image.NRGBA {
	radius := diameter / 2
	mask := imaging.New(diameter, diameter, color.NRGBA{0, 0, 0, 0})

	for y := range diameter {
		for x := range diameter {
			dx := x - radius
			dy := y - radius
			distance := math.Sqrt(float64(dx*dx + dy*dy))

			if distance <= float64(radius) {
				// 圆内为白色（不透明）
				mask.SetNRGBA(x, y, color.NRGBA{255, 255, 255, 255})
			}
			// 圆外保持透明
		}
	}

	return mask
}

// 将头像切成圆形
func cropToCircle(avatar image.Image, size int) *image.NRGBA {
	// 1. 调整头像大小
	resized := imaging.Resize(avatar, size, size, imaging.Lanczos)

	// 2. 创建圆形蒙版
	mask := createCircleMask(size)

	// 3. 应用蒙版
	result := imaging.New(size, size, color.NRGBA{0, 0, 0, 0})

	for y := range size {
		for x := range size {
			// 获取蒙版像素的透明度
			_, _, _, maskAlpha := mask.At(x, y).RGBA()

			if maskAlpha > 0 {
				// 如果蒙版不透明，使用头像像素
				avatarColor := resized.At(x, y)
				result.Set(x, y, avatarColor)
			}
		}
	}

	return result
}

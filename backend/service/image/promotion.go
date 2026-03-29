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
	"bytes"
	"encoding/json"
	"fmt"
	"gin/config"
	"gin/model"
	"gin/service/set"
	"gin/service/upload"
	"gin/service/wechat"
	"gin/util"
	"image"
	"image/color"
	"image/png"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/skip2/go-qrcode"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

func getThemeContent() (map[string]any, error) {
	var theme model.Theme
	if err := config.DB.Where("type=6 and status=1").First(&theme).Error; err != nil {
		return nil, err
	}
	var themeContent []map[string]any
	themeContentByte, err := theme.Content.MarshalJSON()
	if err != nil {
		return nil, err
	}
	json.Unmarshal(themeContentByte, &themeContent)
	return themeContent[0], nil
}
func getBaseImg(theme map[string]any) (image.Image, error) {
	// 网络图片的 URL
	imageURL := theme["imgUrl"].(string)
	if imageURL == "" {
		return nil, fmt.Errorf("基础图片为空")
	}
	// 1. 创建带超时的自定义客户端
	timeoutClient := &http.Client{
		Timeout: 10 * time.Second, // 总超时 5 秒
	}
	// 发送 HTTP GET 请求
	resp, err := timeoutClient.Get(imageURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // 确保响应体被关闭

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("请求基础图片失败，响应码：%d", resp.StatusCode)
	}

	// 将响应体（一个 io.Reader）解码为 image.Image
	mainImg, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	return mainImg, nil
}
func getUserAvatar(uid int64) (image.Image, string, error) {
	var user model.User
	if err := config.DB.Where("id=?", uid).First(&user).Error; err != nil {
		return nil, "", err
	}
	avatarURL := user.Avatar
	if avatarURL == "" {
		var err error
		avatarURL, err = set.GetSet("site_logo", true)
		if err != nil {
			return nil, "", err
		}
	}
	// 1. 创建带超时的自定义客户端
	timeoutClient := &http.Client{
		Timeout: 10 * time.Second, // 总超时 5 秒
	}
	resp, err := timeoutClient.Get(avatarURL)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close() // 确保响应体被关闭

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("请求头像图片失败，响应码：%d", resp.StatusCode)
	}

	// 将响应体（一个 io.Reader）解码为 image.Image
	avatarImg, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, "", err
	}
	return avatarImg, user.Nickname, nil
}
func GetUserPromotion(uid int64, platform string) (string, error) {
	themeContent, err := getThemeContent()
	if err != nil {
		return "", err
	}
	mainImg, err := getBaseImg(themeContent)
	if err != nil {
		return "", err
	}

	// 缩放主图片到指定大小 ---
	targetWidth := 300
	targetHeight := 400
	// 使用 imaging.Resize 进行缩放
	// Lanczos 算法适合缩小图片，质量较高
	resizedImg := imaging.Resize(mainImg, targetWidth, targetHeight, imaging.Lanczos)

	// 创建一个新的图像作为画布，将缩放后的图片绘制上去 ---
	dstImg := imaging.New(targetWidth, targetHeight, color.NRGBA{0, 0, 0, 0}) // 创建一个透明背景的画布
	dstImg = imaging.Paste(dstImg, resizedImg, image.Pt(0, 0))
	// 加载系统默认字体
	ttfFont, err := getSystemDefaultFont()
	if err != nil {
		return "", err
	}
	// 绘制用户信息
	avatarImg, nickName, err := getUserAvatar(uid)
	if err != nil {
		return "", err
	}
	avatarImg = imaging.Resize(avatarImg, 50, 50, imaging.Lanczos)
	avatarPos := image.Pt(50, 10)
	avatarImg1 := cropToCircle(avatarImg, 50)
	dstImg = imaging.Overlay(dstImg, avatarImg1, avatarPos, 1.0)
	// 绘制用户昵称
	face, err := opentype.NewFace(ttfFont, &opentype.FaceOptions{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return "", err
	}
	d := &font.Drawer{
		Dst:  dstImg,
		Src:  image.NewUniform(color.NRGBA{R: 0, G: 0, B: 0, A: 255}),
		Face: face,
	}
	d.Dot = fixed.P(120, 40) // 设置文字的起始位置
	d.DrawString(nickName)

	// 配置自定义字体
	fontSizeStr := themeContent["fontSize"].(string)
	fontSizeStr = strings.ReplaceAll(fontSizeStr, "px", "")
	fontSize, err := strconv.ParseFloat(fontSizeStr, 64)
	if err != nil {
		return "", err
	}
	face, err = opentype.NewFace(ttfFont, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return "", err
	}
	// 定义文字内容和样式
	text := themeContent["title"].(string)
	textColor, err := util.HexToRGBA(themeContent["color"].(string))
	if err != nil {
		return "", err
	}
	lineHeight := fontSize * 1.2            // 行高（字体大小的1.2倍）
	padding := 20                           // 左右边距20px
	maxLineWidth := targetWidth - 2*padding // 每行最大宽度

	// 绘制居中换行文字
	drawCenteredText(dstImg, text, face, textColor, 70, lineHeight, maxLineWidth)

	var qrcodeImg image.Image
	// 加载要叠加的图片
	if platform == "h5" {
		domain, err := set.GetSet("site_url", true)
		if err != nil {
			return "", err
		}
		pngImg, err := qrcode.Encode(fmt.Sprintf("%s/pages/index/index?pid=%d", domain, uid), qrcode.Medium, 200)
		if err != nil {
			return "", err
		}
		qrcodeImg, _, err = image.Decode(bytes.NewReader(pngImg))
		if err != nil {
			return "", err
		}
	}
	if platform == "weixin" {
		img := wechat.GetMiniQRCode("pages/index/index", "pid="+strconv.Itoa(int(uid)))
		qrcodeImg, _, err = image.Decode(img)
		if err != nil {
			return "", err
		}
	}
	//先缩放一下
	qrcodeImgWidth := 200
	qrcodeImgHeight := 200
	scaledQrcode := imaging.Resize(qrcodeImg, qrcodeImgWidth, qrcodeImgHeight, imaging.Lanczos)

	// 定义位置
	QrcodePos := image.Pt(50, 170)
	dstImg = imaging.Paste(dstImg, scaledQrcode, QrcodePos)

	buf := new(bytes.Buffer)
	err = png.Encode(buf, dstImg)
	if err != nil {
		return "", err
	}
	return upload.UploadFile(buf, ".png")
}
func GetGoodsPromotion(uid int64, gid int64, platform string) (string, error) {
	// defer func() {
	// 	fmt.Println("===== 完整堆栈 =====")
	// 	fmt.Println(string(debug.Stack())) // 打印堆栈信息
	// 	// if err := recover(); err != nil {
	// 	// 	fmt.Println("===== 完整堆栈 =====")
	// 	// 	fmt.Println(string(debug.Stack())) // 打印堆栈信息
	// 	// }
	// }()
	themeContent, err := getThemeContent()
	if err != nil {
		return "", err
	}
	mainImg, err := getBaseImg(themeContent)
	if err != nil {
		return "", err
	}
	// 缩放主图片到指定大小 ---
	targetWidth := 300
	targetHeight := 400
	// 使用 imaging.Resize 进行缩放
	// Lanczos 算法适合缩小图片，质量较高
	resizedImg := imaging.Resize(mainImg, targetWidth, targetHeight, imaging.Lanczos)

	// 创建一个新的图像作为画布，将缩放后的图片绘制上去 ---
	dstImg := imaging.New(targetWidth, targetHeight, color.NRGBA{0, 0, 0, 0}) // 创建一个透明背景的画布
	dstImg = imaging.Paste(dstImg, resizedImg, image.Pt(0, 0))
	// 加载系统默认字体
	ttfFont, err := getSystemDefaultFont()
	if err != nil {
		return "", err
	}
	// 绘制用户信息
	avatarImg, nickName, err := getUserAvatar(uid)
	if err != nil {
		return "", err
	}
	avatarImg = imaging.Resize(avatarImg, 50, 50, imaging.Lanczos)
	avatarPos := image.Pt(50, 10)
	avatarImg1 := cropToCircle(avatarImg, 50)
	dstImg = imaging.Overlay(dstImg, avatarImg1, avatarPos, 1.0)
	// 绘制用户昵称
	face, err := opentype.NewFace(ttfFont, &opentype.FaceOptions{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return "", err
	}
	d := &font.Drawer{
		Dst:  dstImg,
		Src:  image.NewUniform(color.NRGBA{R: 0, G: 0, B: 0, A: 255}),
		Face: face,
	}
	d.Dot = fixed.P(120, 40) // 设置文字的起始位置
	d.DrawString(nickName)
	var goods model.Goods
	if err := config.DB.Where("id = ?", gid).First(&goods).Error; err != nil {
		return "", err
	}
	// 绘制商品图片
	timeoutClient := http.Client{Timeout: 10 * time.Second}
	// 发送 HTTP GET 请求

	resp, err := timeoutClient.Get(goods.Thumb)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close() // 确保响应体被关闭

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("请求商品图片失败，响应码：%d", resp.StatusCode)
	}

	// 将响应体（一个 io.Reader）解码为 image.Image
	thumbImg, _, err := image.Decode(resp.Body)
	if err != nil {
		return "", err
	}
	thumbImg = imaging.Resize(thumbImg, 280, 200, imaging.Lanczos)
	thumbPos := image.Pt(10, 70)
	dstImg = imaging.Paste(dstImg, thumbImg, thumbPos)

	//绘制二维码图片
	var qrcodeImg image.Image
	if platform == "h5" {
		domain, err := set.GetSet("site_url", true)
		if err != nil {
			return "", err
		}
		pngImg, err := qrcode.Encode(fmt.Sprintf("%s/pages/goods/info?pid=%d&id=%d", domain, uid, gid), qrcode.Medium, 130)
		if err != nil {
			return "", err
		}
		qrcodeImg, _, err = image.Decode(bytes.NewReader(pngImg))
		if err != nil {
			return "", err
		}
	}
	if platform == "weixin" {
		img := wechat.GetMiniQRCode("pages/goods/info", "pid="+strconv.Itoa(int(uid))+"&id="+strconv.Itoa(int(gid)))
		qrcodeImg, _, err = image.Decode(img)
		if err != nil {
			return "", err
		}
	}
	//先缩放一下
	qrcodeImgWidth := 130
	qrcodeImgHeight := 110
	scaledQrcode := imaging.Resize(qrcodeImg, qrcodeImgWidth, qrcodeImgHeight, imaging.Lanczos)

	// 定义位置
	QrcodePos := image.Pt(160, 280)
	dstImg = imaging.Paste(dstImg, scaledQrcode, QrcodePos)
	// 绘制商品名称
	lineHeight := 16 * 1.2          // 行高（字体大小的1.2倍）
	padding := 10                   // 左右边距20px
	maxLineWidth := 150 - 2*padding // 每行最大宽度

	// 绘制换行文字
	drawLinesText(dstImg, goods.Title, face, color.NRGBA{R: 0, G: 0, B: 0, A: 255}, 10, 280, lineHeight, maxLineWidth)
	buf := new(bytes.Buffer)
	err = png.Encode(buf, dstImg)
	if err != nil {

		return "", err
	}
	return upload.UploadFile(buf, ".png")
}

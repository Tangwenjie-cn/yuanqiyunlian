/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package task

import (
	"fmt"
	"gin/config"
	"gin/model"
	"gin/service/document"
	"gin/service/upload"
	"os"
	"path/filepath"
	"runtime/debug"
	"strconv"
	"strings"
)

func HandleDocTask(docPath string, data model.Goods, SortIds string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("文档任务异常：", err)
			fmt.Println("===== 完整堆栈 =====")
			fmt.Println(string(debug.Stack())) // 打印堆栈信息
		}
	}()
	// 处理文件
	images, err := document.HandleFile(docPath)
	if err != nil {
		fmt.Println("处理文件失败:", err)
	}
	imgList := images["images"].([]string)
	newImgList := make([]string, 0)
	for _, img := range imgList {
		file, err := os.Open(img)
		if err != nil {
			fmt.Println("打开文件失败:", err)
			continue
		}
		newImg, err := upload.UploadFile(file, filepath.Ext(img))
		if err != nil {
			fmt.Println("上传文件失败:", err)
			continue
		}
		file.Close()
		newImgList = append(newImgList, newImg)
	}
	fmt.Println("newImgList:", newImgList)
	file, err := os.Open(docPath)
	if err != nil {
		fmt.Println("打开文档文件失败:", err)
		return
	}
	link, err := upload.UploadFile(file, filepath.Ext(docPath))
	if err != nil {
		fmt.Println("上传文档文件失败:", err)
		return
	}
	file.Close()
	data.Link = link
	data.Title = strings.TrimSuffix(filepath.Base(docPath), filepath.Ext(docPath))
	data.Thumb = newImgList[0]
	data.Pages = images["count"].(int)
	data.Type = 1
	tx := config.DB.Begin()
	if err := tx.Create(&data).Error; err != nil {
		tx.Rollback()
		fmt.Println("创建商品失败:", err)
		return
	}
	SortIdsArr := strings.Split(SortIds, ",")
	sorts := make([]model.GsCorrelation, 0)
	for _, v := range SortIdsArr {
		sid, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		sorts = append(sorts, model.GsCorrelation{
			Gid: int(data.ID),
			Sid: sid,
		})
	}
	if err := tx.Create(&sorts).Error; err != nil {
		tx.Rollback() // 回滚事务
		return
	}
	var contentStr string
	for _, v := range newImgList {
		contentStr += "<img src='" + v + "' style='max-width: 100%;max-height: 100%;' />"
	}
	content := model.GoodsContent{
		Gid:     int(data.ID),
		Content: contentStr,
	}
	if err := tx.Create(&content).Error; err != nil {
		tx.Rollback() // 回滚事务
		return
	}
	tx.Commit() // 提交事务
	dir := filepath.Dir(docPath)
	filename := strings.TrimSuffix(filepath.Base(docPath), filepath.Ext(docPath))
	if filepath.Ext(docPath) != ".pdf" {
		err = os.Remove(docPath)
		if err != nil {
			fmt.Println(err)
		}
	}
	err = os.Remove(filepath.Join(dir, filename+".pdf"))
	if err != nil {
		fmt.Println(err)
	}
	err = os.RemoveAll(filepath.Join(dir, filename))
	if err != nil {
		fmt.Println(err)
	}
}

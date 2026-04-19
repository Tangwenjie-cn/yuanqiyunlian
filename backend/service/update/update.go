package update

import (
	"encoding/json"
	"errors"
	"fmt"
	"gin/config"
	"gin/util"
	"strconv"
	"strings"
)

var giteeUrl = "https://gitee.com/api/v5/repos/chengdu-yuanqi/yuanqi-yunlian/releases/latest"

// 获取更新信息
func getUpdate() (map[string]any, error) {
	reqData := make(map[string]any, 0)
	reqData["url"] = giteeUrl
	reqData["method"] = "GET"
	body, err := util.Request(reqData)
	if err != nil {
		return nil, err
	}
	var rawResp map[string]any
	err = json.Unmarshal(body, &rawResp)
	return rawResp, err
}

// 检查更新
func CheckUpdate() (bool, map[string]any) {
	updateInfo, err := getUpdate()
	if err != nil {
		return false, nil
	}
	// 比较版本号
	res, err := CompareVersion(updateInfo["tag_name"].(string), config.Vstring)
	// 比较
	if err != nil {
		return false, nil
	}
	if res == 1 {
		return true, updateInfo
	}
	return false, nil
}

// CompareVersion
// 版本号对比函数
// v1: 待对比版本
// v2: 基准版本
// return:
//
//	1  v1 > v2
//	0  v1 == v2
//	-1 v1 < v2
//
// err: 版本号格式非法时返回
func CompareVersion(v1, v2 string) (int, error) {
	// 自动去掉 v / V 前缀
	v1 = strings.TrimPrefix(strings.ToLower(v1), "v")
	v2 = strings.TrimPrefix(strings.ToLower(v2), "v")
	// 切割版本号为数字切片
	v1Slice := strings.Split(v1, ".")
	v2Slice := strings.Split(v2, ".")

	// 转换为 int 并校验合法性
	v1Nums, err := toIntSlice(v1Slice)
	if err != nil {
		return 0, fmt.Errorf("v1 非法: %w", err)
	}
	v2Nums, err := toIntSlice(v2Slice)
	if err != nil {
		return 0, fmt.Errorf("v2 非法: %w", err)
	}

	// 取最大长度，不足补 0
	maxLen := max(len(v1Nums), len(v2Nums))
	// 补全到相同长度
	v1Nums = padSlice(v1Nums, maxLen)
	v2Nums = padSlice(v2Nums, maxLen)

	// 逐段对比
	for i := range maxLen {
		if v1Nums[i] > v2Nums[i] {
			return 1, nil
		}
		if v1Nums[i] < v2Nums[i] {
			return -1, nil
		}
	}

	// 所有段都相等
	return 0, nil
}

// toIntSlice 字符串切片转 int 切片，校验是否为合法数字
func toIntSlice(strSlice []string) ([]int, error) {
	res := make([]int, 0, len(strSlice))
	for _, s := range strSlice {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, errors.New("版本段必须是纯数字")
		}
		if num < 0 {
			return nil, errors.New("版本号不能为负数")
		}
		res = append(res, num)
	}
	return res, nil
}

// padSlice 切片补 0 到指定长度
func padSlice(slice []int, length int) []int {
	if len(slice) >= length {
		return slice
	}
	for len(slice) < length {
		slice = append(slice, 0)
	}
	return slice
}

// max 取两个整数的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

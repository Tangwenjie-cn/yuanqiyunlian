/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package model

import (
	"errors"
	"gin/config"
	"gin/service/tree"
	"strings"
	"time"
)

type Auth struct {
	ID       uint            `json:"id"`
	Name     string          `json:"name"`
	Url      string          `json:"url"`
	Pid      int             `json:"pid"`
	Path     string          `json:"path"`
	Icon     string          `json:"icon"`
	IsMenu   int             `json:"is_menu"` // 是否是菜单 1 是 0 否
	Sort     int             `json:"sort"`    // 排序
	Ctime    time.Time       `gorm:"autoCreateTime" json:"ctime"`
	Utime    time.Time       `gorm:"autoUpdateTime" json:"utime"`
	Children []tree.TreeNode `json:"children" gorm:"-"` // 子权限
}

func (a *Auth) GetId() uint                          { return a.ID }
func (a *Auth) GetPid() int                          { return a.Pid }
func (a *Auth) GetSort() int                         { return a.Sort }
func (a *Auth) SetChildren(children []tree.TreeNode) { a.Children = children }
func (a *Auth) GetChildren() []tree.TreeNode         { return a.Children }

func (a Auth) GetTree(isMenu bool, groupId int) (list []tree.TreeNode, err error) {
	// 1. 获取所有权限
	// 2. 递归获取子权限
	// 3. 返回树形结构的权限列表
	var authArr []Auth
	db := config.DB.Model(&Auth{})
	if isMenu {
		db = db.Where("is_menu = 1")
	}
	if groupId > 1 {
		var authGroup AuthGroup
		config.DB.Model(&AuthGroup{}).Where("id = ?", groupId).First(&authGroup)
		if authGroup.ID == 0 {
			return nil, errors.New("权限组不存在")
		}
		if authGroup.Auth == "" {
			return nil, errors.New("权限组没有权限")
		}
		auths := strings.Split(authGroup.Auth, ",")
		if authGroup.HalfAuth != "" { // 半选权限
			// 半选权限需要单独处理，不能直接添加到权限列表中
			auths = append(auths, strings.Split(authGroup.HalfAuth, ",")...)

		}
		db = db.Where("id in ?", auths)
	}
	db.Order("sort asc").Find(&authArr)
	treeNode := []tree.TreeNode{}

	for _, v := range authArr {
		treeNode = append(treeNode, &v)
	}
	list = tree.TreeBuilder(treeNode)
	return list, nil

}

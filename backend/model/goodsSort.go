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
	"gin/config"
	"gin/service/tree"
	"time"
)

type GoodsSort struct {
	ID       uint            `json:"id"`
	Pid      int             `json:"pid"`
	Name     string          `json:"name"`
	Thumb    string          `json:"thumb"`
	Sort     int             `json:"sort"`
	Status   int             `json:"status"`
	Ctime    time.Time       `gorm:"autoCreateTime" json:"ctime"`
	Utime    time.Time       `gorm:"autoUpdateTime" json:"utime"`
	Children []tree.TreeNode `json:"children" gorm:"-"` // 子分类
}

func (a *GoodsSort) GetId() uint                          { return a.ID }
func (a *GoodsSort) GetPid() int                          { return a.Pid }
func (a *GoodsSort) GetSort() int                         { return a.Sort }
func (a *GoodsSort) SetChildren(children []tree.TreeNode) { a.Children = children }
func (a *GoodsSort) GetChildren() []tree.TreeNode         { return a.Children }
func (g GoodsSort) GetTree() (list []tree.TreeNode, err error) {
	// 1. 获取所有分类
	// 2. 递归获取子分类
	// 3. 返回树形结构的分类列表
	var sortArr []GoodsSort
	db := config.DB.Model(&GoodsSort{})

	db.Where("status = 1").Order("sort asc").Find(&sortArr)
	treeNode := []tree.TreeNode{}
	for _, v := range sortArr {
		treeNode = append(treeNode, &v)
	}
	list = tree.TreeBuilder(treeNode)
	return list, nil
}

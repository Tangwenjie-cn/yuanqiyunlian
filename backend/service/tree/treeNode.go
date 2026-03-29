/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
package tree

import "sort"

type TreeNode interface {
	GetId() uint                     // 获取节点ID
	GetPid() int                     // 获取父节点ID
	GetSort() int                    // 获取排序值
	SetChildren(children []TreeNode) // 设置子节点
	GetChildren() []TreeNode         // 获取子节点
}

func TreeBuilder(nodes []TreeNode) []TreeNode {
	listMap := make(map[uint]TreeNode)
	for _, v := range nodes {
		listMap[v.GetId()] = v
	}
	tree := createTree(listMap, 0)
	tree = sortTree(tree)
	return tree
}

// 递归函数，传入一个权限列表和父级ID，返回树形结构的权限列表
func createTree(list map[uint]TreeNode, pid int) (tree []TreeNode) {
	//map是无序的，所以结果如果要每次一样，需要再排序
	for k, v := range list {
		if v.GetPid() == pid {
			// 删除当前节点
			delete(list, k)
			// 递归获取子节点
			v.SetChildren(createTree(list, int(v.GetId())))
			// 添加到树形结构中
			tree = append(tree, v)
		}
	}
	return tree
}

// 递归排序Auth切片及其所有子切片
func sortTree(tree []TreeNode) []TreeNode {
	// 先排序当前层级的元素
	sort.Slice(tree, func(i, j int) bool {
		return tree[i].GetSort() < tree[j].GetSort()
	})

	// 再递归排序每个元素的子元素
	for i := range tree {
		tree[i].SetChildren(sortTree(tree[i].GetChildren()))
	}

	return tree
}

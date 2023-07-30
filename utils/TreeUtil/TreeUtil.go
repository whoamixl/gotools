package TreeUtil

type TreeNode struct {
	Id       string
	Pid      string
	Atb      interface{}
	Children []TreeNode
}
type NodeList struct {
	Id  string
	Pid string
	Atb interface{}
}

func Build(list []NodeList) TreeNode {
	// 创建一个空的根节点
	root := TreeNode{}
	// 使用一个辅助函数来递归构建树
	buildTree(&root, list)
	return root
}
func buildTree(node *TreeNode, list []NodeList) {
	for _, item := range list {
		if item.Pid == node.Id {
			// 创建一个新的子节点
			child := TreeNode{
				Id:  item.Id,
				Pid: item.Pid,
				Atb: item.Atb,
			}
			// 递归构建子节点的子树
			buildTree(&child, list)
			// 将子节点添加到父节点的 Children 列表中
			node.Children = append(node.Children, child)
		}
	}
}

// 二叉搜索树的查找 插入 删除
package main

type BinarySearchTree struct {
	value int
	left  *BinarySearchTree
	right *BinarySearchTree
}

func Search(root *BinarySearchTree, x int) *BinarySearchTree {
	var result *BinarySearchTree
	if root == nil {
		return &BinarySearchTree{}
	}
	if root.value == x {
		return root
	} else if root.value < x {
		result = Search(root.right, x)
	} else {
		result = Search(root.left, x)
	}
	return result

}

func Search1(root *BinarySearchTree, x int) *BinarySearchTree {
	for root != nil {
		if root.value == x {
			return root
		} else if root.value < x {
			root = root.right
		} else {
			root = root.left
		}
	}
	return root

}

// 数据的插入
// 返回插入的那个节点的指针
func Insert(root *BinarySearchTree, value int) *BinarySearchTree {
	var result *BinarySearchTree

	// 如果二叉查收树是空的，那么就创建一个根树
	if root == nil {
		a := new(BinarySearchTree)
		a.value = value
		root = a
		return root
	}

	// 如果要创建的数据跟原有的数据一样大，不好意思我们无法帮你插入，因为这是不符合规定的，抱歉了。
	if value == root.value {
		return &BinarySearchTree{
			value: -1,
		}
	}
	// 如果数据比根节点大
	if value > root.value {
		// 如果右边是空的，直接嫁接到右边即可
		if root.right == nil {
			a := new(BinarySearchTree)
			a.value = value
			root.right = a
			return a
			//如果不是空的，那么请递归吧，直到找到是空的位置。
		} else {
			result = Insert(root.right, value)
		}
		// 这个就是数据比root小
	} else {
		// 如果left是空的，那么请嫁接上吧
		if root.left == nil {
			a := new(BinarySearchTree)
			a.value = value
			root.left = a
			return a
			// 有数据?请递归吧。
		} else {
			// 返回递归的数据，如果不返回最高的一层也拿不到数据啊。
			result = Insert(root.left, value)
		}

	}
	// 返回数据
	return result
}

// 删除 看 https://github.com/imgoogege/algo/blob/master/go/24_tree/BinarySearchTree.go 有时间再写 太费劲。

// func Delete(b *BinarySearchTree, value int) {
// 	var now *BinarySearchTree = b
// 	var parent *BinarySearchTree
// 	if now == value && now.left == nil && now.right == nil {
// 		parent
// 	}
// }

// func Delete1(b *BinarySearchTree, value int) {
// 	for b != nil {
// 		if value == b.value {
// 			fmt.Println("测试shuju ", b.value)
// 			b.value = math.MinInt64 // 将删除的数据标记为-10即可。 这样就不用真的删除了。
// 			break
// 		}
// 		if value > b.value {
// 			b = b.right
// 		} else {
// 			b = b.left
// 		}
// 	}
//
// }

func MRange(b *BinarySearchTree) []int {
	var result []int
	if b.left == nil && b.right == nil {
		return []int{b.value}
	}
	if b.left != nil {
		result = MRange(b.left)
	}
	result = append(result, b.value)
	if b.right != nil {
		result = append(result, MRange(b.right)...)
	}
	return result
}

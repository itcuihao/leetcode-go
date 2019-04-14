package structure

var BinarySearchTree = &TreeNode{
	Val: 6,
	Left: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	},
	Right: &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val: 7,
		},
		Right: &TreeNode{
			Val: 9,
		},
	},
}

var BinaryTree = &TreeNode{
	Val: 3,
	Left: &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	},
	Right: &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 8,
		},
	},
}

package collections

import (
	"errors"
)

// AVL tree object
type AVLTree struct {
	root  *treeNode
	count int
}

// AVL tree constructor
func NewAVLTree() AVLTree {
	return AVLTree{
		root:  nil,
		count: 0,
	}
}

// Method to insert cargo into AVL tree
func (tree *AVLTree) Insert(key string, cargo interface{}) (err error) {
	tree.root, err = insertAVLTreeUtil(tree.root, key, cargo)
	if err == nil {
		tree.count++
	}
	return
}

// Utility function for AVL tree insertion
func insertAVLTreeUtil(tree *treeNode, key string, cargo interface{}) (*treeNode, error) {
	var err error
	if tree == nil {
		newNode := treeNode{
			key:    key,
			cargo:  cargo,
			height: 1,
		}
		return &newNode, nil
	}
	if key < tree.key {
		tree.left, err = insertAVLTreeUtil(tree.left, key, cargo)
	} else if key > tree.key {
		tree.right, err = insertAVLTreeUtil(tree.right, key, cargo)
	} else { // Equal keys are not allowed in BST
		return tree, errors.New("element with this key already exists")
	}

	tree.height = getMax(getHeight(tree.left), getHeight(tree.right)) + 1

	diffHeight := getDiffHeight(tree)

	// Left Left Case
	if diffHeight > 1 && tree.left != nil && key < tree.left.key {
		return rightRotate(tree), err
	}

	// Right Right Case
	if diffHeight < -1 && tree.right != nil && key > tree.right.key {
		return leftRotate(tree), err
	}

	// Left Right Case
	if diffHeight > 1 && tree.left != nil && key > tree.left.key {
		tree.left = leftRotate(tree.left)
		return rightRotate(tree), err
	}

	// Right Left Case
	if diffHeight < -1 && tree.right != nil && key < tree.right.key {
		tree.right = rightRotate(tree.right)
		return leftRotate(tree), err
	}
	return tree, err
}

// A utility function to get height of the tree
func getHeight(treeNode *treeNode) int {
	if treeNode == nil {
		return 0
	} else {
		return treeNode.height
	}
}

// Get height difference of node treeNode
func getDiffHeight(treeNode *treeNode) int {
	if treeNode == nil {
		return 0
	}
	return getHeight(treeNode.left) - getHeight(treeNode.right)
}

// A utility function to right rotate subtree rooted with treeNode
func rightRotate(treeNode *treeNode) *treeNode {
	leftNode := treeNode.left
	leftRightNode := leftNode.right

	leftNode.right = treeNode
	treeNode.left = leftRightNode

	treeNode.height = getMax(getHeight(treeNode.right), getHeight(treeNode.left)) + 1
	leftNode.height = getMax(getHeight(leftNode.right), getHeight(leftNode.left)) + 1

	return leftNode
}

// A utility function to left rotate subtree rooted with treeNode
func leftRotate(treeNode *treeNode) *treeNode {
	rightNode := treeNode.right
	rightLeftNode := rightNode.left

	rightNode.left = treeNode
	treeNode.right = rightLeftNode

	treeNode.height = getMax(getHeight(treeNode.right), getHeight(treeNode.left)) + 1
	rightNode.height = getMax(getHeight(rightNode.right), getHeight(rightNode.left)) + 1

	return rightNode
}

// Method to remove cargo from an AVL tree
func (tree *AVLTree) Remove(key string) error {
	var err error
	tree.root, err = removeAVLTreeUtil(tree.root, key)
	if err == nil {
		tree.count--
	}
	return err
}

// Remove cargo from AVL tree with a given key
func removeAVLTreeUtil(tree *treeNode, key string) (*treeNode, error) {
	if tree == nil {
		return tree, errors.New("no element found with given key")
	}
	var err error
	if key < tree.key {
		tree.left, err = removeAVLTreeUtil(tree.left, key)
	} else if key > tree.key {
		tree.right, err = removeAVLTreeUtil(tree.right, key)
	} else {
		// node with only one child or no child
		if tree.left == nil {
			return tree.right, nil
		} else if tree.right == nil {
			return tree.left, nil
		} else { // node with two children
			var rightmost *treeNode
			for rightmost = tree.left; rightmost.right != nil; rightmost = rightmost.right {
			}
			tree.key = rightmost.key
			tree.cargo = rightmost.cargo
			tree.left, err = removeAVLTreeUtil(tree.left, rightmost.key)
		}
	}

	tree.height = 1 + getMax(getHeight(tree.left), getHeight(tree.right))

	diffHeight := getDiffHeight(tree)

	// Left Left Case
	if diffHeight > 1 && getDiffHeight(tree.left) >= 0 {
		return rightRotate(tree), err
	}

	// Left Right Case
	if diffHeight > 1 && getDiffHeight(tree.left) < 0 {
		tree.left = leftRotate(tree.left)
		return rightRotate(tree), err
	}

	// Right Right Case
	if diffHeight < -1 && getDiffHeight(tree.right) <= 0 {
		return leftRotate(tree), err
	}

	// Right Left Case
	if diffHeight < -1 && getDiffHeight(tree.right) > 0 {
		tree.right = rightRotate(tree.right)
		return leftRotate(tree), err
	}
	return tree, err
}

// Method returning number of elements in tree
func (tree AVLTree) Count() int {
	return tree.count
}

// Method to print AVL tree contents
func (tree AVLTree) PrintTree() {
	printTreeUtil(tree.root)
}

// Method to print AVL tree layout
// Passed argument defines how much spacing there is between tree levels
func (tree AVLTree) PrintTree2D(spacing int) {
	// Pass initial space count as 0
	printTree2DUtil(tree.root, 0, spacing)
}
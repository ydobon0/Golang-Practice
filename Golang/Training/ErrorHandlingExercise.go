package Training

import (
	"errors"
	"fmt"
	"math"
)

//------------------------------ Module 7: Practice Activities: Error Handling ------------------------------
//Activity 3
type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, errors.New("cannot Sqrt negative number")
	}
	out := math.Sqrt(x)
	return out, nil
}

func (e ErrNegativeSqrt) Error() string {
	message := fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
	return message
}

func ErrorHandlingExerciseAct3() {
	out, err := Sqrt(-2)

	if err != nil {
		message := ErrNegativeSqrt(out).Error()
		fmt.Println(message)
		return
	}
	fmt.Println(out)
}

func ErrorHandlingExercise() {
	ErrorHandlingExerciseAct3()
	ErrorHandlingExerciseAct4()
}

//Activity 4

type node struct {
	left  *node
	right *node
	data  int
}

type tree struct {
	root *node
}

func ErrorHandlingExerciseAct4() {
	n1 := node{}
	n1.data = 5

	n2 := node{}
	n2.data = 4

	n3 := node{}
	n3.data = 6

	n4 := node{}
	n4.data = 7

	n1.left = &n2
	n1.right = &n3

	n2.left = &n4

	tr := tree{}
	tr.root = &n1

	out, err := checkTree(tr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(out)
	}

	fmt.Println(searchTree(3, tr))
}

func checkTree(tr tree) (bool, error) { //make sure a binary tree is a binary search tree
	return checkNode(*tr.root)
}

func checkNode(nn node) (bool, error) { //checks to make sure a node's left and right children follow the rules of a binary search tree then recursively checks the children
	if nn.left != nil {
		if nn.left.data > nn.data {
			return false, errors.New("A node's left child is not less than the parent")
		}
		return checkNode(*nn.left)
	}
	if nn.right != nil {
		if nn.right.data < nn.data {
			return false, errors.New("A node's right child is not greater than the parent")
		}
		return checkNode(*nn.right)
	}
	return true, nil
}

func searchTree(val int, tr tree) bool { //searches a tree for a particular value
	return searchNode(val, tr.root)
}

func searchNode(val int, nn *node) bool { //checks if the current node is the requested value then checks the appropriate child recursively
	if nn.data == val {
		return true
	}
	if val < nn.data {
		if nn.left != nil {
			return searchNode(val, nn.left)
		}
	} else {
		if nn.right != nil {
			return searchNode(val, nn.right)
		}
	}
	return false
}

package main

import (
	"../utils"
	"fmt"
)

func Example_case1() {
	root := (*TreeNode)(utils.Aton("[3 [9] [20 [15] [7]]]"))
	ans := levelOrder(root)
	fmt.Print(ans)
	// Output:
	// [[3] [9 20] [15 7]]
}

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(layer *int, nodesPerLayer *[]*TreeNode, stats map[int]int) {
	if len(*nodesPerLayer) == 0 {
		return
	}

	var tmp []*TreeNode
	for _, ptr := range *nodesPerLayer {
		if ptr == nil {
			continue
		}

		stats[*layer] += ptr.Val
		tmp = append(tmp, ptr.Left)
		tmp = append(tmp, ptr.Right)
	}

	*nodesPerLayer = tmp
	*layer++
	helper(layer, nodesPerLayer, stats)

}

func maxLevelSum(root *TreeNode) int {
	var nodes []*TreeNode
	var stats = make(map[int]int, 0)
	var layer = 1
	nodes = append(nodes, root)
	helper(&layer, &nodes, stats)

	index := layer
	val := stats[layer]
	for i := layer - 1; i >= 1; i-- {
		if stats[i] >= val {
			index = i
			val = stats[i]
		}
	}
	return index
}

func main() {


}
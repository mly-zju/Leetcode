/*
 * @lc app=leetcode.cn id=733 lang=golang
 *
 * [733] 图像渲染
 *
 * https://leetcode-cn.com/problems/flood-fill/description/
 *
 * algorithms
 * Easy (52.98%)
 * Likes:    66
 * Dislikes: 0
 * Total Accepted:    11.2K
 * Total Submissions: 21.1K
 * Testcase Example:  '[[1,1,1],[1,1,0],[1,0,1]]\n1\n1\n2'
 *
 * 有一幅以二维整数数组表示的图画，每一个整数表示该图画的像素值大小，数值在 0 到 65535 之间。
 * 
 * 给你一个坐标 (sr, sc) 表示图像渲染开始的像素值（行 ，列）和一个新的颜色值 newColor，让你重新上色这幅图像。
 * 
 * 
 * 为了完成上色工作，从初始坐标开始，记录初始坐标的上下左右四个方向上像素值与初始坐标相同的相连像素点，接着再记录这四个方向上符合条件的像素点与他们对应四个方向上像素值与初始坐标相同的相连像素点，……，重复该过程。将所有有记录的像素点的颜色值改为新的颜色值。
 * 
 * 最后返回经过上色渲染后的图像。
 * 
 * 示例 1:
 * 
 * 
 * 输入: 
 * image = [[1,1,1],[1,1,0],[1,0,1]]
 * sr = 1, sc = 1, newColor = 2
 * 输出: [[2,2,2],[2,2,0],[2,0,1]]
 * 解析: 
 * 在图像的正中间，(坐标(sr,sc)=(1,1)),
 * 在路径上所有符合条件的像素点的颜色都被更改成2。
 * 注意，右下角的像素没有更改为2，
 * 因为它不是在上下左右四个方向上与初始点相连的像素点。
 * 
 * 
 * 注意:
 * 
 * 
 * image 和 image[0] 的长度在范围 [1, 50] 内。
 * 给出的初始点将满足 0 <= sr < image.length 和 0 <= sc < image[0].length。
 * image[i][j] 和 newColor 表示的颜色值在范围 [0, 65535]内。
 * 
 * 
 */

// @lc code=start
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	// floodfill算法，其实就是dfs，但是结合回溯算法框架来搞
	// 为了防止无限递归，先处理颜色相同的情况
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}

	dfs(image, sr, sc, oldColor, newColor)
	return image
}

func dfs(image [][]int, r int, c int, oldColor, newColor int) {
	// 先执行修改
	image[r][c] = newColor

	// 尝试执行探索，如果都不存在自然就结束了
	rNum, cNum := len(image), len(image[0])
	// 上
	if r - 1 >= 0 && image[r-1][c] == oldColor {
		dfs(image, r-1, c, oldColor, newColor)
	}
	// 下
	if r + 1 < rNum && image[r+1][c] == oldColor {
		dfs(image, r+1, c, oldColor, newColor)
	}
	// 左
	if c - 1 >= 0 && image[r][c-1] == oldColor {
		dfs(image, r, c-1, oldColor, newColor)
	}
	// 右
	if c + 1 < cNum && image[r][c+1] == oldColor {
		dfs(image, r, c+1, oldColor, newColor)
	}
}
// @lc code=end


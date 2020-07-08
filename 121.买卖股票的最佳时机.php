/*
 * @lc app=leetcode.cn id=121 lang=php
 *
 * [121] 买卖股票的最佳时机
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/description/
 *
 * algorithms
 * Easy (50.73%)
 * Likes:    588
 * Dislikes: 0
 * Total Accepted:    88.3K
 * Total Submissions: 172.7K
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
 * 
 * 如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。
 * 
 * 注意你不能在买入股票前卖出股票。
 * 
 * 示例 1:
 * 
 * 输入: [7,1,5,3,6,4]
 * 输出: 5
 * 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
 * ⁠    注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
 * 
 * 
 * 示例 2:
 * 
 * 输入: [7,6,4,3,1]
 * 输出: 0
 * 解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
 * 
 * 
 */
// @lc code=start
class Solution {

    /**
     * @param Integer[] $prices
     * @return Integer
     */
    function maxProfit($prices) {
        // dp思路：第i天最大收益=max{第i-1天最大收益, 第i天价格-i天之前最小价格}    
        $profits = [0];
        $minPrice = $prices[0];
        $maxProfit = 0;

        foreach ($prices as $index => $price) {
            if ($index == 0) {
                continue;
            }

            if ($profits[$index - 1] > $price - $minPrice) {
                $profits[] = $profits[$index - 1];
            } else {
                $profits[] = $price - $minPrice;
            }
            $maxProfit = $profits[$index] > $maxProfit ? $profits[$index] : $maxProfit;
            // 重新计算最小价格
            $minPrice = $minPrice > $price ? $price : $minPrice;
        }

        return $maxProfit;
    }
}
// @lc code=end


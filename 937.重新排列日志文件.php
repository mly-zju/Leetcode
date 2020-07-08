<?php
class Solution {

    /**
     * @param String[] $logs
     * @return String[]
     */
    function reorderLogFiles($logs) {
        $charLogs = [];
        $numLogs = [];    
        // 遍历logs, 分别筛选字母和数字log
        foreach ($logs as $v) {
            // 判断是否是数字log
            if ($this->isNumber($v)) {
                $numLogs[] = $v;
            } else {
                $charLogs[] = $v;
            }
        }

        // 对字母log做排序
        usort($charLogs, function ($a, $b) {
            $aIndex = strpos($a, " ") + 1;
            $bIndex = strpos($b, " ") + 1;
            $aSub = substr($a, $aIndex);
            $bSub = substr($b, $bIndex);
            // 如果不相等，就可以直接返回了
            if (strcmp($aSub, $bSub) != 0) {
                return strcmp($aSub, $bSub);
            } else {
                $aFlag = substr($a, 0, $aIndex);
                $bFlag = substr($b, 0, $bIndex);
                return strcmp($aFlag, $bFlag);
            }
        });

        // 合并返回
        return array_merge($charLogs, $numLogs);
    }

    /**
     * 判断是否是数字log
     * 
     * @param string $log
     * @return bool
     */
    function isNumber($log) {
        $index = strpos($log, " ") + 1;
        // 由于只包含数字或者只包含字母，只需要看第一位是不是数字
        if ($log[$index] >= '0' && $log[$index] <= '9') {
            return true;
        }
        return false;
    }
}

$do = new Solution();
$logs = ["a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"];
$newLogs = $do->reorderLogFiles($logs);
var_dump($newLogs);
<?php
class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function balancedStringSplit($s) {
        $num = 1;
        $startChar = $s[0];
        $len = strlen($s);
        $sum = 0;
        for ($i = 1; $i < $len; $i++) {
            if ($s[$i] == $startChar) {
                $num++;
            } else {
                $num--;
            }
            if ($num == 0) {
                $sum++;
            }
        } 

        return $sum;
    }
}

$do = new Solution();
$s = "RLRRLLRLRL";
$sum = $do->balancedStringSplit($s);
var_dump($sum);
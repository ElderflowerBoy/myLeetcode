<?php

class Solution
{

    /**
     * @param Integer $num
     *
     * @return Integer
     */
    function largestInteger($num)
    {
        $even = [];
        $odd = [];
        foreach (str_split(strval($num)) as $key => $value) {
            if ($value % 2 === 0) {
                $even[$key] = $value;
            } else {
                $odd[$key] = $value;
            }
        }

        $eKeys = array_keys($even);
        $oKeys = array_keys($odd);
        rsort($even);
        rsort($odd);
        $arr = array_combine($eKeys, $even) + array_combine($oKeys, $odd);
        ksort($arr);
        return intval(implode($arr));
    }
}
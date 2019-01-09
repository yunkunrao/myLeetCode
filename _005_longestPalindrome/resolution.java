package _005_longestPalindrome;

import java.util.HashMap;
import java.util.Map.Entry;

class Solution {
    public static String longestPalindrome(String s) {
        if (s.length() <= 1)
            return s;

        String longest = "";

        int mid = s.length() / 2;
        int r = 1;
        while (Math.abs(r) <= mid + 1) {
            String l1 = longestPalindrome(s, mid);
            String l2 = longestPalindrome(s, mid - r);

            String tmp = l1.length() > l2.length() ? l1 : l2;

            if (tmp.length() > longest.length()) {
                longest = tmp;
            }

            r *= -1;
            if (r > 0)
                ++r;
        }

        return longest;
    }

    public static String longestPalindrome(String s, int mid) {

        int halfSpan = 0;
        int beginIndex = 0;
        int endIndex = 1;
        int l = 0;

        // 中心对称
        while ((mid - halfSpan - 1 >= 0) && (mid + halfSpan + 1 < s.length())) {
            if (s.charAt(mid - halfSpan - 1) != s.charAt(mid + halfSpan + 1))
                break;
            halfSpan++;
        }
        if (halfSpan != 0) {
            beginIndex = mid - halfSpan;
            endIndex = mid + halfSpan + 1;
            l = endIndex - beginIndex;
        }

        // 左右对称
        // reset halfSpan
        halfSpan = 0;
        while ((mid - halfSpan >= 0) && (mid + halfSpan + 1 < s.length())) {
            if (s.charAt(mid - halfSpan) != s.charAt(mid + halfSpan + 1)) {
                break;
            }
            halfSpan++;
        }
        if (halfSpan != 0) {
            if (l < 2 * halfSpan) {
                beginIndex = mid - halfSpan + 1;
                endIndex = mid + halfSpan + 1;
            }
        }
        return s.substring(beginIndex, endIndex);
    }

    public static void main(String[] args) {
        boolean succeed = true;

        HashMap<String, String> map = new HashMap<String, String>();
        map.put("babad", "bab");
        map.put("cbbd", "bb");
        map.put("ac", "a");
        map.put("aaaa", "aaaa");
        map.put("abadd", "aba");
        map.put("bb", "bb");

        for (Entry<String, String> entry : map.entrySet()) {
            String str = entry.getKey();
            String expected = entry.getValue();
            String actual = longestPalindrome(str);
            if (!expected.equals(actual)) {
                succeed = false;
                System.out.println("[" + str + "]" + " got " + actual + " ,expected " + expected);
                break;
            }
        }
        System.out.println(succeed ? "Nice!" : "Fucking fucked!");
        return;
    }
}
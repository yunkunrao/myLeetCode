package _003_lengthOfLongestSubstring;

import java.util.HashMap;
import java.util.Map.Entry;
import java.util.HashSet;

class Solution {
    public static int lengthOfLongestSubstring(String s) {
        int max = 0;

        HashSet<Character> set = new HashSet<Character>();

        int start = 0;
        int end = -1;
        int index = 0;

        for (int i = 0; i < s.length(); i++) {
            if (set.contains(s.charAt(i))) {
                max = Math.max(max, end - start + 1);

                index = s.substring(start, end + 1).indexOf(s.charAt(i));
                start = start + index + 1;
                if (start == s.length() - 1)
                    break;
                end = i;
            } else {
                set.add(s.charAt(i));
                ++end;
            }
            if (end == s.length() - 1) {
                max = Math.max(max, end - start + 1);
            }
        }

        return max;
    }

    public static void main(String[] args) {

        boolean succeed = true;
        HashMap<String, Integer> map = new HashMap<String, Integer>();

        map.put("abcabcbb", 3);
        map.put("bbbbb", 1);
        map.put("pwwkew", 3);
        map.put("", 0);
        map.put(" ", 1);

        for (Entry<String, Integer> entry : map.entrySet()) {
            String str = entry.getKey();
            int expected = entry.getValue();
            int actual = lengthOfLongestSubstring(str);
            if (expected != actual) {
                succeed = false;
                System.out.println(str + " ,got " + actual + " ,expected " + expected);
                break;
            }
        }
        System.out.println(succeed ? "Nice!" : "Fucking fucked!");
    }
}
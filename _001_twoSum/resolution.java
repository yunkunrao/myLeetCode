package _001_twoSum;

import java.util.Arrays;

public class resolution {
    public static int[] twoSum(int[] nums, int target) {
        int index[] = { -1, -1 };

        int[] nums2 = copyArray(nums);
        Arrays.sort(nums2);
        int start = 0;
        int end = nums2.length - 1;
        while (start < end) {
            if (nums2[start] + nums2[end] < target) {
                start++;
            } else if (nums2[start] + nums2[end] > target) {
                end--;
            } else {
                break;
            }
        }

        for (int i = 0; i < nums.length; i++) {
            if (nums[i] == nums2[start])
                index[0] = i;
            if (nums[nums.length - 1 - i] == nums2[end])
                index[1] = nums.length - 1 - i;
        }

        if (index[0] > index[1])
            swap(index, 0, 1);

        return index;
    }

    public static int[] copyArray(int[] arr) {
        if (arr == null) {
            return null;
        }
        int[] res = new int[arr.length];
        for (int i = 0; i < arr.length; i++) {
            res[i] = arr[i];
        }
        return res;
    }

    public static void swap(int[] arr, int i, int j) {
        int tmp = arr[i];
        arr[i] = arr[j];
        arr[j] = tmp;
    }

    public static void main(String[] args) {

        int testTime = 500000;
        int maxSize = 100;
        int maxValue = 100;
        boolean succeed = true;
        int target = 6;

        for (int i = 0; i < testTime; i++) {
            int[] testcase = Test.g(target, maxSize, maxValue);
            int[] res = twoSum(testcase, target);
            if ((testcase[res[0]] + testcase[res[1]]) != target) {
                succeed = false;
                System.out.println("[error index] " + res[0] + " and " + res[1]);
                break;
            }
        }

        System.out.println(succeed ? "Nice!" : "Fucking fucked!");

    }

}

package _004_findMedianSortedArrays;

// m    n	   length			 Mid
// 偶数 偶数   yes				 (leftMax + rightMin)/2
// 奇数 偶数   left + 1 = right  rightMin
// 偶数 奇数   left + 1 = right	 rightMin
// 奇数 奇数   yes			     (leftMax + rightMin)/2

class Solution {

    public static double findMedianSortedArrays(int[] A, int[] B) {
        int m = A.length;
        int n = B.length;

        // to make sure m < n
        if (m > n) {
            int[] tmpArr = B;
            B = A;
            A = tmpArr;

            int tmp = m;
            m = n;
            n = tmp;
        }

        int leftMax = 0;
        int rightMin = 0;
        for (int i = 0; i <= m; i++) {
            int j = (m + n) / 2 - i;

            if (i >= 1) {
                if (j > 0) {
                    leftMax = Math.max(A[i - 1], B[j - 1]);
                } else { // j = 0
                    leftMax = A[i - 1];
                }
            } else {
                if (j >= 1) {
                    leftMax = B[j - 1];
                } else { // i = 0 j=0
                    leftMax = Integer.MIN_VALUE;
                }
            }
            if (i < m) {
                if (j < n) {
                    rightMin = Math.min(A[i], B[j]);
                } else { // j = n
                    rightMin = A[i];
                }

            } else {
                rightMin = B[j];
            }
            if (leftMax <= rightMin) {
                int leftLength = i + j;
                int rightLength = m + n - leftLength;
                if (leftLength == rightLength) {
                    break;
                } else {
                    leftMax = rightMin;
                    break;
                }
            } else {
                continue;
            }
        }

        return (leftMax + rightMin) / 2.0;
    }

    public static void main(String[] args) {

        boolean succeed = true;

        int[] nums1 = { 1 };
        int[] nums2 = { 1 };

        double expected = 1.0;
        double actual = findMedianSortedArrays(nums1, nums2);
        if (expected != actual) {
            succeed = false;
            System.out.println("got " + actual + " ,expected " + expected);

        }
        System.out.println(succeed ? "Nice!" : "Fucking fucked!");
        return;
    }

}
package demo;

public class BubbleSort {
    // 冒泡排序：
    // 从左到右不断交换相邻逆序的元素，在一轮的循环之后，可以让未排序的最大元素上浮到右侧。
    // 优化 - 在一轮循环中，如果没有发生交换，就说明数组已经是有序的，此时可以直接退出。
    public static void sort(int[] arr) {
        if (arr == null || arr.length < 2) {
            return;
        }
        boolean hasSorted = false;
        for (int e = arr.length - 1; e > 0 && !hasSorted; e--) {
            hasSorted = true;
            for (int i = 0; i < e; i++) {
                if (arr[i] > arr[i + 1]) {
                    hasSorted = false;
                    swap(arr, i, i + 1);
                }
            }
        }
    }

    // 异或运算的本质是无进位加法
    public static void swap(int[] arr, int i, int j) {
        arr[i] = arr[i] ^ arr[j];
        arr[j] = arr[i] ^ arr[j];
        arr[i] = arr[i] ^ arr[j];
    }

}

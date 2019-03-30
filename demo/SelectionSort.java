package demo;

public class SelectionSort {
    // 选择排序：
    // 选择出数组中的最小元素，将它与数组的第一个元素交换位置。再从剩下的元素中
    // 选择出最小的元素，将它与数组的第二个元素交换位置。不断进行这样的操作，直
    // 到将整个数组排序。
    public static void sort(int[] arr) {
        if (arr == null || arr.length < 2) {
            return;
        }
        for (int i = 0; i < arr.length - 1; i++) {
            int minIndex = i;
            for (int j = i + 1; j < arr.length; j++) {
                minIndex = arr[j] < arr[minIndex] ? j : minIndex;
            }
            utils.swap(arr, i, minIndex);
        }
    }
}

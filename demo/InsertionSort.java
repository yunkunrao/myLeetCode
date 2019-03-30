package demo;

public class InsertionSort {

    // 插入排序：
    // 每次都将当前元素插入到左侧已经排序的数组中，使得插入之后左侧数组依然有序。
    public void sort(int[] arr) {
        if (arr == null || arr.length < 2) {
            return;
        }
        for (int i = 1; i < arr.length; i++) {
            for (int j = i; j > 0 && arr[j] < arr[j - 1]; j--) {
                utils.swap(arr, j, j - 1);
            }
        }
    }

}

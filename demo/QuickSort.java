package demo;

public class QuickSort {

    // 快速排序：
    // 快速排序通过一个切分元素将数组分为两个子数组，左子数组小于等于切分元素，
    // 右子数组大于等于切分元素，将这两个子数组排序也就将整个数组排序了。
    public static void quickSort(int[] arr) {
        if (arr == null || arr.length < 2) {
            return;
        }
        quickSort(arr, 0, arr.length - 1);
    }

    public static void quickSort(int[] arr, int l, int r) {
        if (l < r) {
            utils.swap(arr, l + (int) (Math.random() * (r - l + 1)), r);
            int[] p = partition(arr, l, r);
            quickSort(arr, l, p[0] - 1);
            quickSort(arr, p[1] + 1, r);
        }
    }

    public static int[] partition(int[] arr, int l, int r) {
        int less = l - 1;
        int more = r;
        while (l < more) {
            if (arr[l] < arr[r]) {
                utils.swap(arr, ++less, l++);
            } else if (arr[l] > arr[r]) {
                utils.swap(arr, --more, l);
            } else {
                l++;
            }
        }
        utils.swap(arr, more, r);
        return new int[] { less + 1, more };
    }

}

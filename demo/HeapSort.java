package demo;

public class HeapSort {

    // 堆排序：
    // 堆中某个节点的值总是大于等于其子节点的值，并且堆是一颗完全二叉树。
    public static void heapSort(int[] arr) {
        if (arr == null || arr.length < 2) {
            return;
        }
        for (int i = 0; i < arr.length; i++) {
            heapInsert(arr, i);
        }
        int size = arr.length;
        utils.swap(arr, 0, --size);
        while (size > 0) {
            // 每次拿到堆中最大元素并放入数组末尾
            heapify(arr, 0, size);
            utils.swap(arr, 0, --size);
        }
    }
    // 大根堆
    public static void heapInsert(int[] arr, int index) {
        //
        while (arr[index] > arr[(index - 1) / 2]) {
            utils.swap(arr, index, (index - 1) /2);
            index = (index - 1)/2 ;
        }
    }

    public static void heapify(int[] arr, int index, int size) {
        int left = index * 2 + 1;
        while (left < size) {
            int largest = left + 1 < size && arr[left + 1] > arr[left] ? left + 1 : left;
            largest = arr[largest] > arr[index] ? largest : index;
            if (largest == index) {
                break;
            }
            utils.swap(arr, largest, index);
            index = largest;
            left = index * 2 + 1;
        }
    }
}

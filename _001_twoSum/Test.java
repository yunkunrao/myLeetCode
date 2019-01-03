package _001_twoSum;

// test case related
public class Test {

    public static int[] generateRandomArray(int maxSize, int maxValue) {
        int[] arr = new int[(int) ((maxSize + 1) * Math.random())];
        for (int i = 0; i < arr.length; i++) {
            arr[i] = (int) ((maxValue + 1) * Math.random()) - (int) (maxValue * Math.random());
        }
        return arr;
    }

    public static void printArray(int[] arr) {
        if (arr == null) {
            return;
        }
        for (int i = 0; i < arr.length; i++) {
            System.out.print(arr[i] + " ");
        }
        System.out.println();
    }

    public static void swap(int[] arr, int i, int j) {
        int tmp = arr[i];
        arr[i] = arr[j];
        arr[j] = tmp;
    }

    public static int getRandomComplement(int[] arr, int index, int target, int maxValue) {
        int ran = (int) ((maxValue + 1) * Math.random()) - (int) (maxValue * Math.random());

        int i = 0;
        while (i < index) {
            if (arr[i] + ran == target) {
                i = 0;
                ran = (int) ((maxValue + 1) * Math.random()) - (int) (maxValue * Math.random());
                continue;
            }
            i++;
        }

        return ran;
    }

    public static int[] g(int target, int maxSize, int maxValue) {

        int[] arr = new int[(int) (2 + (maxSize - 1) * Math.random())];
        arr[0] = (int) ((target + 1) * Math.random());
        arr[1] = target - arr[0];

        for (int i = 2; i < arr.length; i++) {
            arr[i] = getRandomComplement(arr, i, target, maxValue);
        }

        for (int k = 0; k < arr.length; k++) {
            int ranIndex = (int) ((k + 1) * Math.random());
            swap(arr, k, ranIndex);
        }

        return arr;
    }

}
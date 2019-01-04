package _002_addTwoNumbers;

public class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
        val = x;
    }

    public static ListNode generateListNode(int[] arr) {
        int len = arr.length;
        int i = 0;
        ListNode dump = new ListNode(0);
        ListNode cur = dump;
        while (i < len) {
            cur.next = new ListNode(arr[i]);
            i++;
            cur = cur.next;
        }

        return dump.next;
    }

}

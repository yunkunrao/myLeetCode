package _002_addTwoNumbers;

public class resolution {

    public static ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        return addTwoNumbers(l1, l2, 0);
    }

    public static ListNode addTwoNumbers(ListNode l1, ListNode l2, int carry) {
        if (l1 == null && l2 == null) {
            return carry == 0 ? null : new ListNode(carry);
        }

        if (l1 == null) {
            if (carry == 0) {
                return copyListNode(l2);
            } else {
                l1 = new ListNode(carry);
                // reset carry
                carry = 0;
            }
        }
        if (l2 == null) {
            if (carry == 0) {
                return copyListNode(l1);
            } else {
                l2 = new ListNode(carry);
                // reset carry
                carry = 0;
            }
        }
        int curSum = l1.val + l2.val + carry;
        ListNode head = new ListNode(curSum % 10);
        head.next = addTwoNumbers(l1.next, l2.next, curSum / 10);

        return head;
    }

    public static ListNode copyListNode(ListNode header) {
        if (header == null)
            return null;

        ListNode h = new ListNode(header.val);
        ListNode ret = h;
        header = header.next;
        while (header != null) {
            h.next = new ListNode(header.val);
            h = h.next;
            header = header.next;
        }

        return ret;
    }

    public static int getListSum(ListNode head) {
        if (head == null)
            return 0;

        int sum = head.val + 10 * getListSum(head.next);

        return sum;
    }

    public static void printListNode(ListNode head) {
        while (head != null) {
            System.out.print("" + head.val + "->");
            head = head.next;
        }
        System.out.println("null");
    }

    public static void main(String[] args) {

        int[] list1 = { 9, 8 };
        ListNode head1 = ListNode.generateListNode(list1);

        int[] list2 = { 1 };
        ListNode head2 = ListNode.generateListNode(list2);

        boolean succeed = true;
        // System.out.println(Integer.MAX_VALUE);

        if ((getListSum(head1) + getListSum(head2)) != getListSum(addTwoNumbers(head1, head2))) {
            succeed = false;
            printListNode(head1);
            printListNode(head2);
            printListNode(addTwoNumbers(head1, head2));
        }

        System.out.println(succeed ? "Nice!" : "Fucking fucked!");

    }
}
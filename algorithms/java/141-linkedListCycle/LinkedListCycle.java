public class LinkedListCycle {
    public static void main(String args[]) {
        Solution s = new Solution();

        if (true) {
            ListNode n1 = new ListNode(3);
            ListNode n2 = new ListNode(2);
            ListNode n3 = new ListNode(0);
            ListNode n4 = new ListNode(-4);

            n1.next = n2;
            n2.next = n3;
            n3.next = n4;
            n4.next = n2;

            System.out.println(s.hasCycle(n1));
        }

        if (true) {
            ListNode n1 = new ListNode(1);
            ListNode n2 = new ListNode(2);

            n1.next = n2;
            n2.next = n1;

            System.out.println(s.hasCycle(n1));
        }

        if (true) {
            ListNode n1 = new ListNode(1);

            System.out.println(s.hasCycle(n1));
        }
    }
}

/**
 * Definition for singly-linked list.
 * class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */

class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
        val = x;
        next = null;
    }
}

class Solution {
    public boolean hasCycle(ListNode head) {
        if (head == null) {
            return false;
        }

        ListNode slow = head;
        ListNode fast = head.next;

        while (fast != null) {
            if (fast == slow) {
                return true;
            }

            fast = fast.next;
            slow = slow.next;

            if (fast != null) {
                fast = fast.next;
            }
        }

        return false;
    }
}
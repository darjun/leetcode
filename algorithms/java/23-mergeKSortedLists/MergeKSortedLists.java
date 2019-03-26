import java.util.Comparator;
import java.util.PriorityQueue;
import java.util.Queue;

public class MergeKSortedLists {
    public static void main(String args[]) {
        Solution s = new Solution();

        if (true) {
            ListNode n1 = new ListNode(1);
            ListNode n2 = new ListNode(4);
            ListNode n3 = new ListNode(5);
            n1.next = n2;
            n2.next = n3;

            ListNode n4 = new ListNode(1);
            ListNode n5 = new ListNode(3);
            ListNode n6 = new ListNode(4);
            n4.next = n5;
            n5.next = n6;

            ListNode n7 = new ListNode(2);
            ListNode n8 = new ListNode(6);
            n7.next = n8;

            ListNode[] lists = new ListNode[] { n1, n4, n7 };
            printList(s.mergeKLists(lists));
        }
    }

    public static void printList(ListNode head) {
        if (head == null) {
            return;
        }

        while (head.next != null) {
            System.out.print(head.val);
            System.out.print(" -> ");
            head = head.next;
        }

        System.out.println(head.val);
    }
}

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */

class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
        val = x;
    }
}

class Solution {
    private static Comparator<ListNode> listNodeComparator = new Comparator<ListNode>() {
        @Override
        public int compare(ListNode l1, ListNode l2) {
            return l1.val - l2.val;
        }
    };

    public ListNode mergeKLists(ListNode[] lists) {
        if (lists.length == 0) {
            return null;
        }

        Queue<ListNode> q = new PriorityQueue<>(lists.length, listNodeComparator);
        ListNode head = null, tail = null;
        for (int i = 0; i < lists.length; i++) {
            if (lists[i] != null) {
                q.add(lists[i]);
            }
        }

        while (q.size() > 0) {
            ListNode smallestNode = q.poll();
            ListNode newNode = new ListNode(smallestNode.val);

            if (head == null) {
                head = tail = newNode;
            } else {
                tail.next = newNode;
                tail = newNode;
            }

            if (smallestNode.next != null) {
                q.add(smallestNode.next);
            }
        }

        return head;
    }
}
public class SwapNodesInPairs {
    public static void main(String[] args) {
        {
            ListNode n1 = new ListNode(1);
            ListNode n2 = new ListNode(2);
            ListNode n3 = new ListNode(3);
            ListNode n4 = new ListNode(4);
            ListNode n5 = new ListNode(5);
            ListNode n6 = new ListNode(6);
            n1.next = n2;
            n2.next = n3;
            n3.next = n4;
            n4.next = n5;
            n5.next = n6;

            System.out.println(new Solution().swapPairs(n1));
        }
    }
}


class Solution {
    public ListNode swapPairs(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }

        ListNode prev = null;
        ListNode cur = head;
        ListNode next = head.next;
        ListNode newHead = next;

        while (next != null) {
            cur.next = next.next;
            if (prev != null) {
                prev.next = next;
            }
            next.next = cur;

            // next loop
            prev = cur;
            cur = cur.next;
            if (cur != null) {
                next = cur.next;
            } else {
                next = null;
            }
        }

        return newHead;
    }
}

class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
        val = x;
    }
    
    @Override
    public String toString() {
        String ret = String.valueOf(val);
        ListNode cur = this;
        while (cur.next != null) {
            ret += " -> " + String.valueOf(cur.next.val);
            cur = cur.next;
        }
        return ret;
    }
}
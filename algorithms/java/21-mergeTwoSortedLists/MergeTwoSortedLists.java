public class MergeTwoSortedLists {
    public static void main(String[] args) {
        {
            ListNode n1 = new ListNode(1);
            ListNode n2 = new ListNode(2);
            ListNode n4 = new ListNode(4);
            n1.next = n2;
            n2.next = n4;

            ListNode m1 = new ListNode(1);
            ListNode m3 = new ListNode(3);
            ListNode m4 = new ListNode(4);
            m1.next = m3;
            m3.next = m4;

            System.out.println(new Solution().mergeTwoLists(n1, m1));
        }
    }
}

class Solution {
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        ListNode ret = null;
        ListNode cur = null;
        while (l1 != null || l2 != null) {
            int val1 = l1 != null ? l1.val : Integer.MAX_VALUE;
            int val2 = l2 != null ? l2.val : Integer.MAX_VALUE;

            ListNode n;
            if (val1 < val2) {
                n = new ListNode(val1);
                l1 = l1.next;
            } else {
                n = new ListNode(val2);
                l2 = l2.next;
            }
            if (ret == null) {
                ret = cur = n;
            } else {
                cur.next = n;
                cur = n;
            }
        }
        return ret;
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
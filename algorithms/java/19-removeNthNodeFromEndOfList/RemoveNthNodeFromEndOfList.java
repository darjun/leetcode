public class RemoveNthNodeFromEndOfList {
    public static void main(String[] args) {
        {
            ListNode n1 = new ListNode(1);
            ListNode n2 = new ListNode(2);
            ListNode n3 = new ListNode(3);
            ListNode n4 = new ListNode(4);
            ListNode n5 = new ListNode(5);
            n1.next = n2;
            n2.next = n3;
            n3.next = n4;
            n4.next = n5;

            System.out.println(new Solution().removeNthFromEnd(n1, 2));
        }
    }
}

class Solution {
    public ListNode removeNthFromEnd(ListNode head, int n) {
        // invalid
        if (n == 0 || head == null) {
            return head;
        }

        ListNode n1 = head;
        ListNode n2 = head;
        int i = n;
        // make the distance between n1 and n2 is n
        while (i-- > 0 && n2 != null) {
            n2 = n2.next;
        }
        
        // move n2 to the node past end
        ListNode prev = null;
        while (n2 != null) {
            prev = n1;
            n1 = n1.next;
            n2 = n2.next;
        }

        if (n1 == head) {
            return head.next;
        }

        prev.next = n1.next; 
        return head;
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
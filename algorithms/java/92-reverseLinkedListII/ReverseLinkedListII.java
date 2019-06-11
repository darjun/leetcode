public class ReverseLinkedListII {
  public static void main(String[] args) {
//    {
//      ListNode n1 = new ListNode(1);
//      ListNode n2 = new ListNode(2);
//      ListNode n3 = new ListNode(3);
//      ListNode n4 = new ListNode(4);
//      ListNode n5 = new ListNode(5);
//
//      n1.next = n2;
//      n2.next = n3;
//      n3.next = n4;
//      n4.next = n5;
//
//      System.out.println(new Solution().reverseBetween(n1, 2, 4));
//    }
//
//    {
//      ListNode n1 = new ListNode(3);
//      ListNode n2 = new ListNode(5);
//
//      n1.next = n2;
//
//      System.out.println(new Solution().reverseBetween(n1, 1, 2));
//    }

    {
      ListNode n1 = new ListNode(1);
      ListNode n2 = new ListNode(2);
      ListNode n3 = new ListNode(3);

      n1.next = n2;
      n2.next = n3;

      System.out.println(new Solution().reverseBetween(n1, 1, 3));
    }
  }

}

class Solution {
  public ListNode reverseBetween(ListNode head, int m, int n) {
    if (m >= n || head == null || head.next == null) {
      return head;
    }

    ListNode prev = null;
    ListNode cur = head;
    int i = m;
    while (cur != null && --i > 0) {
      prev = cur;
      cur = cur.next;
    }

    // tail before reverse position
    ListNode lastTail = prev;

    // cur is the first node to reverse
    prev = cur;
    cur = cur.next;
    if (cur == null) {
      return head;
    }
    
    ListNode reverseHead = prev;
    ListNode next = cur.next;
    int count = n - m;
    prev.next = null;
    do {
      cur.next = prev;
      prev = cur;
      cur = next;
      if (next != null) {
        next = next.next;
      }
      --count;
    } while (count > 0 && cur != null);

    if (m != 1) {
      // link first part and the head of reverse part
      lastTail.next = prev;
    }

    // link the tail of reverse head and the left part
    reverseHead.next = cur;

    return m == 1 ? prev : head;
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
public class RemoveDuplicatesFromSortedListII {
  public static void main(String[] args) {
//     {
//       ListNode n11 = new ListNode(1);
//       ListNode n12 = new ListNode(1);
//       ListNode n2 = new ListNode(2);
//
//       n11.next = n12;
//       n12.next = n2;
//
//       System.out.println(new Solution().deleteDuplicates(n11));
//     }
//
//    {
//      ListNode n11 = new ListNode(1);
//      ListNode n12 = new ListNode(1);
//      ListNode n13 = new ListNode(1);
//      ListNode n2 = new ListNode(2);
//      ListNode n3 = new ListNode(3);
//
//      n11.next = n12;
//      n12.next = n13;
//      n13.next = n2;
//      n2.next = n3;
//
//      System.out.println(new Solution().deleteDuplicates(n11));
//    }

    {
      ListNode n0 = new ListNode(0);
      ListNode n1 = new ListNode(1);
      ListNode n21 = new ListNode(2);
      ListNode n22 = new ListNode(2);
      ListNode n3 = new ListNode(3);
      ListNode n4 = new ListNode(4);

      n0.next = n1;
      n1.next = n21;
      n21.next = n22;
      n22.next = n3;
      n3.next = n4;

      System.out.println(new Solution().deleteDuplicates(n0));
    }
  }
}

class Solution {
  public ListNode deleteDuplicates(ListNode head) {
    if (head == null || head.next == null) {
      return head;
    }

    ListNode newHead = null;
    ListNode newTail = null;

    ListNode cur = head;

    while (cur != null) {
      ListNode prev = cur;
      cur = cur.next;

      while (cur != null && prev.val != cur.val) {
        if (newHead == null) {
          newHead = newTail = prev;
        } else {
          newTail.next = prev;
          newTail = prev;
        }

        prev = cur;
        cur = cur.next;
      }

      if (cur == null) {
        if (newHead == null) {
          newHead = newTail = prev;
        } else {
          newTail.next = prev;
          newTail = prev;
        }
      }

      if (cur != null) {
        int curVal = prev.val;
        do {
          cur = cur.next;
        } while (cur != null && cur.val == curVal);
      }
    }

    if (newTail != null) {
      newTail.next = null;
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

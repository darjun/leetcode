public class RemoveDuplicatesFromSortedListI {
  public static void main(String[] args) {
    // {
    //   ListNode n1 = new ListNode(1);
    //   ListNode n2 = new ListNode(2);
    //   ListNode n31 = new ListNode(3);
    //   ListNode n32 = new ListNode(3);
    //   ListNode n41 = new ListNode(4);
    //   ListNode n42 = new ListNode(4);
    //   ListNode n5 = new ListNode(5);

    //   n1.next = n2;
    //   n2.next = n31;
    //   n31.next = n32;
    //   n32.next = n41;
    //   n41.next = n42;
    //   n42.next = n5;

    //   System.out.println(new Solution().deleteDuplicates(n1));
    // }

    // {
    //   ListNode n11 = new ListNode(1);
    //   ListNode n12 = new ListNode(1);
    //   ListNode n2 = new ListNode(2);

    //   n11.next = n12;
    //   n12.next = n2;

    //   System.out.println(new Solution().deleteDuplicates(n11));
    // }
    
    {
      ListNode n11 = new ListNode(1);
      ListNode n12 = new ListNode(1);
      ListNode n2 = new ListNode(2);
      ListNode n31 = new ListNode(3);
      ListNode n32 = new ListNode(3);
      
      n11.next = n12;
      n12.next = n2;
      n2.next = n31;
      n31.next = n32;

      System.out.println(new Solution().deleteDuplicates(n11));
    }
  }
}

class Solution {
  public ListNode deleteDuplicates(ListNode head) {
    if (head == null || head.next == null) {
      return head;
    }

    ListNode newHead = head;
    ListNode newTail = head;

    int curVal = head.val;
    ListNode cur = head.next;
    while (cur != null) {
      if (cur.val != curVal) {
        newTail.next = cur;
        newTail = cur;

        curVal = cur.val;
      }

      cur = cur.next;
    }

    newTail.next = null;
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

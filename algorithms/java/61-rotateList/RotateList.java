public class RotateList {
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

      System.out.println(new Solution().rotateRight(n1, 1));
    }

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

      System.out.println(new Solution().rotateRight(n1, 2));
    }

    for (int k = 1; k <= 4; k++) {
      ListNode n1 = new ListNode(0);
      ListNode n2 = new ListNode(1);
      ListNode n3 = new ListNode(2);
      n1.next = n2;
      n2.next = n3;

      System.out.println(new Solution().rotateRight(n1, k));
    }
  }
}

class Solution {
  public ListNode rotateRight(ListNode head, int k) {
    if (head == null || head.next == null) {
      return head;
    }

    ListNode tailNode = head;
    int size = 1;
    while (tailNode.next != null) {
      tailNode = tailNode.next;
      ++size;
    }

    k %= size;
    if (k == 0) {
      return head;
    }

    ListNode splitNode = head;
    int i = size - k - 1;
    while (i > 0) {
      splitNode = splitNode.next;
      --i;
    }

    ListNode newHead = splitNode.next;
    splitNode.next = null;
    tailNode.next = head;

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
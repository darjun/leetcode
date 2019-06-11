public class PartitionList {
  public static void main(String[] args) {
    {
      ListNode n1 = new ListNode(1);
      ListNode n2 = new ListNode(4);
      ListNode n3 = new ListNode(3);
      ListNode n4 = new ListNode(2);
      ListNode n5 = new ListNode(5);
      ListNode n6 = new ListNode(2);

      n1.next = n2;
      n2.next = n3;
      n3.next = n4;
      n4.next = n5;
      n5.next = n6;

      System.out.println(new Solution().partition(n1, 3));
    }
  }

}

class Solution {
  public ListNode partition(ListNode head, int x) {
    if (head == null || head.next == null) {
      return head;
    }

    ListNode left = null;
    ListNode leftTail = null;
    ListNode right = null;
    ListNode rightTail = null;

    ListNode cur = head;

    while (cur != null) {
      if (cur.val < x) {
        if (left == null) {
          left = leftTail = cur;
        } else {
          leftTail.next = cur;
          leftTail = cur;
        }
      } else {
        if (right == null) {
          right = rightTail = cur;
        } else {
          rightTail.next = cur;
          rightTail = cur;
        }
      }
      cur = cur.next;
    }

    if (leftTail != null) {
      leftTail.next = right;
    }
    if (rightTail != null) {
      rightTail.next = null;
    }

    return left != null ? left : right;
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
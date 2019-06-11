public class ReverseNodesInKGroup {
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

      System.out.println(new Solution().reverseKGroup(n1, 2));
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

      System.out.println(new Solution().reverseKGroup(n1, 3));
    }
  }
}

class Solution {
  public ListNode reverseKGroup(ListNode head, int k) {
    if (k == 1 || head == null || head.next == null) {
      return head;
    }

    ListNode newHead = null;
    ListNode prev = null;
    ListNode curTail = null;
    ListNode lastTail = null;
    ListNode curNode = head;
    while (true) {
      ListNode sentiel = curNode;
      int i = k;
      while (i > 0 && sentiel != null) {
        sentiel = sentiel.next;
        --i;
      }

      // there is less than k nodes left.
      if (i != 0) {
        if (lastTail != null) {
          lastTail.next = curNode;
        }
        break;
      }

      ListNode next = curNode.next;
      curTail = curNode;
      while (curNode != sentiel) {
        curNode.next = prev;

        prev = curNode;
        curNode = next;
        if (next != null) {
          next = next.next;
        }
      }

      if (newHead == null) {
        newHead = prev;
      } else {
        lastTail.next = prev;
      }

      lastTail = curTail;
    }

    return newHead != null ? newHead : head;
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
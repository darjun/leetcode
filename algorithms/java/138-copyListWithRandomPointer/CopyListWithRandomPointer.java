import java.util.HashMap;

public class CopyListWithRandomPointer {
  public static void main(String[] args) {
    {
      Node n1 = new Node();
      n1.val = 1;
      Node n2 = new Node();
      n2.val = 2;
      
      n1.random = n2;
      n1.next = n2;
      n2.random = n2;
      System.out.println(new Solution().copyRandomList(n1));
    }
  }
}

/*
// Definition for a Node.
class Node {
    public int val;
    public Node next;
    public Node random;

    public Node() {}

    public Node(int _val,Node _next,Node _random) {
        val = _val;
        next = _next;
        random = _random;
    }
};
*/
class Solution {
  public Node copyRandomList(Node head) {
    if (head == null) {
      return null;
    }

    Node newHead = null;
    Node newTail = null;
    Node current = head;
    HashMap<Node, Node> nodeMap = new HashMap<>();
    while (current != null) {
      Node newNode = new Node(current.val, null, null);
      if (newHead == null) {
        newHead = newTail = newNode;
      } else {
        newTail.next = newNode;
        newTail = newNode;
      }

      nodeMap.put(current, newNode);
      current = current.next;
    }

    current = head;
    Node newCurrent = newHead;
    while (current != null) {
      if (current.random != null) {
        newCurrent.random = nodeMap.get(current.random);
      }

      current = current.next;
      newCurrent = newCurrent.next;
    }

    return newHead;
  }
}

class Node {
  public int val;
  public Node next;
  public Node random;

  public Node() {}

  public Node(int _val, Node _next, Node _random) {
    val = _val;
    next = _next;
    random = _random;
  }

  @Override
  public String toString() {
    String ret = String.valueOf(val);
    if (random != null) {
      ret += " random " + String.valueOf(random.val);
    }
    Node cur = this.next;
    while (cur != null) {
      ret += " -> " + String.valueOf(cur.val);
      if (cur.random != null) {
        ret += " random " + String.valueOf(cur.random.val);
      }
      cur = cur.next;
    }
    return ret;
  }
}
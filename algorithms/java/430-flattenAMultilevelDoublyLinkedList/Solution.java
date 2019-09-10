class Solution {
  public Node flatten(Node head) {
    _flatten(head);
    return head;
  }
  
  public Node _flatten(Node head) {
    Node tail = null;
    for (Node p = head; p != null; p = p.next) {
      if (p.child != null) {
        Node t = _flatten(p.child);

        t.next = p.next;
        if (p.next != null) {
          p.next.prev = t;
        }
        p.child.prev = p;
        p.next = p.child;
        p.child = null;
        p = t;
      }

      tail = p;
    }

    return tail;
  }
}
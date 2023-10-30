/**
 * @param {ListNode} head
 * @param {number} left
 * @param {number} right
 * @return {ListNode}
 */
var reverseBetween = function (head, left, right) {
  if (!head) {
    return null;
  }
  const root = new ListNode(null);
  root.next = head;

  let curr = root,
    pos = 0,
    leftNode = head,
    rightNode = head;
  while (pos <= right) {
    if (pos == left - 1) {
      leftNode = curr;
    }
    if (pos == right) {
      rightNode = curr;
      break;
    }
    curr = curr.next;
    pos++;
  }
  let prev = rightNode.next;
  curr = leftNode.next;
  for (let i = left; i <= right; i++) {
    const next = curr.next;
    curr.next = prev;
    prev = curr;
    curr = next;
  }
  leftNode.next = prev;
  return root.next;
};

class ListNode {
  /** @type {number} */
  val;

  /** @type {ListNode} */
  next;

  /**
   * @param {number} val
   * @param {ListNode} next
   */
  constructor(val, next) {
    this.val = val === undefined ? 0 : val;
    this.next = next === undefined ? null : next;
  }
}

/**
 * @param {ListNode} head
 */
var printList = (head) => {
  /** @type {number[]} */
  let s = "";
  let node = head;
  while (node) {
    s += node.val + " -> ";
    node = node.next;
  }
  if (s.length > 0) {
    s += "null";
  } else {
    s = "null";
  }
  console.log(s);
};

{
  const head = new ListNode(1);
  head.next = new ListNode(2);
  head.next.next = new ListNode(3);
  head.next.next.next = new ListNode(4);
  head.next.next.next.next = new ListNode(5);
  printList(head);
  printList(reverseBetween(head, 2, 4));
}

{
  const head = new ListNode(1);
  head.next = new ListNode(2);
  head.next.next = new ListNode(3);
  head.next.next.next = new ListNode(4);
  head.next.next.next.next = new ListNode(5);
  printList(head);
  printList(reverseBetween(head, 1, 5));
}

{
  printList(null);
  printList(reverseBetween(null, 0, 0));
}

{
  const head = new ListNode(1);
  printList(head);
  printList(reverseBetween(head, 1, 1));
}

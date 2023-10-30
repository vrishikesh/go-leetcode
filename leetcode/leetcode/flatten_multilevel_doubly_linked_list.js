class DLLNode {
  /** @type {number} */
  val;

  /** @type {DLLNode} */
  prev;

  /** @type {DLLNode} */
  next;

  /** @type {DLLNode} */
  child;

  /**
   * @param {number} val
   * @param {DLLNode} prev
   * @param {DLLNode} next
   * @param {DLLNode} child
   */
  constructor(val, prev, next, child) {
    this.val = val;
    this.prev = prev == null ? null : prev;
    this.next = next == null ? null : next;
    this.child = child == null ? null : child;
  }
}

/**
 * @param {DLLNode} head
 * @return {DLLNode}
 */
var flatten = function (head) {
  recurse(head);
  return head;
};

/**
 * @param {DLLNode} head
 * @return {DLLNode}
 */
var recurse = function (head) {
  let curr = head;
  let tail = curr;
  while (curr) {
    if (curr.child) {
      const next = curr.next;
      const childTail = recurse(curr.child);
      const child = curr.child;
      curr.child = null;
      curr.next = child;
      child.prev = curr;
      if (next) {
        childTail.next = next;
        next.prev = childTail;
      }
    }
    tail = curr;
    curr = curr.next;
  }
  return tail;
};

/**
 * @param {DLLNode} head
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
  const one = new DLLNode(1);
  const two = new DLLNode(2);
  const three = new DLLNode(3);
  const four = new DLLNode(4);
  const five = new DLLNode(5);
  const six = new DLLNode(6);
  const seven = new DLLNode(7);
  const eight = new DLLNode(8);
  const nine = new DLLNode(9);
  const ten = new DLLNode(10);
  const eleven = new DLLNode(11);
  const twelve = new DLLNode(12);

  one.next = two;
  two.prev = one;
  two.next = three;
  three.prev = two;
  three.next = four;
  three.child = seven;
  four.prev = three;
  four.next = five;
  five.prev = four;
  five.next = six;
  six.prev = five;

  seven.next = eight;
  eight.prev = seven;
  eight.next = nine;
  eight.child = eleven;
  nine.prev = eight;
  nine.next = ten;
  ten.prev = nine;

  eleven.next = twelve;
  twelve.prev = eleven;

  printList(flatten(one)); // 1 2 3 7 8 11 12 9 10 4 5 6
}

{
  const one = new DLLNode(1);
  const two = new DLLNode(2);
  const three = new DLLNode(3);

  one.next = two;
  one.child = three;
  two.prev = one;

  printList(flatten(one)); // 1 3 2
}

{
  printList(flatten(null)); // null
}

{
  const one = new DLLNode(1);
  const two = new DLLNode(2);
  const three = new DLLNode(3);

  one.child = two;
  two.child = three;

  printList(flatten(one)); // 1 2 3
}

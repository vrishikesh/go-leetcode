/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var reverseList = function (head) {
  let node = head;
  while (node && node.next) {
    const nn = node.next,
      nnn = nn.next;
    node.next = nnn;
    nn.next = head;
    head = nn;
  }
  return head;
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
  const ll = [];
  let node = head;
  while (node) {
    ll.push(node.val);
    node = node.next;
  }
  console.log(ll.join(" -> "));
};

const head = new ListNode(1);
head.next = new ListNode(2);
head.next.next = new ListNode(3);
head.next.next.next = new ListNode(4);
head.next.next.next.next = new ListNode(5);
printList(head);
printList(reverseList(head));

printList(null);
printList(reverseList(null));

head.next = null;
printList(head);
printList(reverseList(head));

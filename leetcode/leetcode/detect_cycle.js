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
    this.val = val;
    this.next = next == null ? null : next;
  }
}

/**
 * Storing Nodes
 * @param {ListNode} head
 * @return {ListNode}
 */
// var detectCycle = function (head) {
//   const m = new Map();
//   let curr = head;
//   while (curr) {
//     if (m.has(curr)) {
//       return curr;
//     }
//     m.set(curr, curr);
//     curr = curr.next;
//   }
//   return null;
// };

/**
 * Floyd's Algorithm
 * @param {ListNode} head
 * @return {ListNode}
 */
var detectCycle = function (head) {
  let slow = head,
    fast = head;
  while (fast) {
    fast = fast.next;
    slow = slow.next;
    if (fast && fast.next) {
      fast = fast.next;
    } else return null;
    if (fast == slow) break;
  }

  let curr = head;
  while (curr != fast) {
    curr = curr.next;
    fast = fast.next;
  }

  return curr;
};

{
  const three = new ListNode(3);
  const two = new ListNode(2);
  const zero = new ListNode(0);
  const mFour = new ListNode(-4);

  three.next = two;
  two.next = zero;
  zero.next = mFour;
  mFour.next = two;

  console.log(detectCycle(three)); // 1
}

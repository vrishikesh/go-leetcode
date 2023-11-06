/**
 * @param {number} val
 * @param {TreeNode} left
 * @param {TreeNode} right
 */
function TreeNode(val, left, right) {
  this.val = val === undefined ? 0 : val;
  this.left = left === undefined ? null : left;
  this.right = right === undefined ? null : right;
}

/**
 * @param {TreeNode} root
 * @return {number[][]}
 */
var levelOrder = function (root) {
  if (root == null) {
    return [];
  }

  /** @type {[TreeNode, level][]} */
  const q = [];
  let curr = root;
  q.push([curr, 0]);

  /** @type {number[][]} */
  const ans = [];

  while (q.length) {
    const nl = q.shift();
    if (nl && nl[0]) {
      const node = nl[0];
      const level = nl[1];
      ans[level] = ans[level] || [];
      ans[level].push(node.val);
      q.push([node.left, level + 1]);
      q.push([node.right, level + 1]);
    }
  }

  return ans;
};

{
  const three = new TreeNode(3);
  const nine = new TreeNode(9);
  const twenty = new TreeNode(20);
  const fifteen = new TreeNode(15);
  const seven = new TreeNode(7);

  three.left = nine;
  three.right = twenty;
  twenty.left = fifteen;
  twenty.right = seven;

  console.log(levelOrder(three));
}

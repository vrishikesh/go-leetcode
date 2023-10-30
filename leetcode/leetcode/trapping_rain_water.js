/**
 * @param {number[]} height
 * @return {number}
 */
var trap = function (height) {
  console.log(height);
  let totalWater = 0,
    p1 = 0,
    p2 = height.length - 1,
    maxLeft = 0,
    maxRight = 0;
  while (p1 < p2) {
    if (height[p1] <= height[p2]) {
      if (height[p1] > maxLeft) {
        maxLeft = height[p1];
      } else {
        totalWater += Math.max(maxLeft - height[p1], 0);
      }
      p1++;
    } else {
      if (height[p2] > maxRight) {
        maxRight = height[p2];
      } else {
        totalWater += Math.max(maxRight - height[p2], 0);
      }
      p2--;
    }
  }
  return totalWater;
};

console.log(trap([0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]));
console.log(trap([0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2]));

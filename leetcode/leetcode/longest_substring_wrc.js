/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function (s) {
  if (s.length < 2) {
    return s.length;
  }

  let longest = 0,
    left = 0,
    right = 0;
  /** @type {Map<string, number>} */
  let seen = new Map();
  while (right < s.length) {
    const currenctChar = s[right];
    const seenChar = seen.get(currenctChar);
    if (seenChar >= left) {
      left = seenChar + 1;
    }
    seen.set(currenctChar, right);
    right++;
    longest = Math.max(longest, right - left);
  }
  return longest;
};

// console.log(lengthOfLongestSubstring("abcbdaac")); // 4
// console.log(lengthOfLongestSubstring("abcabcbb")); // 3
// console.log(lengthOfLongestSubstring("abba")); // 2
console.log(lengthOfLongestSubstring("au")); // 2

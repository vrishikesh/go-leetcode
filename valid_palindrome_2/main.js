/**
 * @param {string} s
 * @return {boolean}
 */
var validPalindrome = function (s) {
  for (let i = 0; i < s.length / 2; i++) {
    if (s[i] !== s[s.length - i - 1]) {
      const secondIndex = s.length - i - 1;
      return (
        compare(s.slice(i, secondIndex)) ||
        compare(s.slice(i + 1, secondIndex + 1))
      );
    }
  }
  return true;
};

/**
 * @param {string} s
 * @return {boolean}
 */
var compare = function (s) {
  for (let i = 0; i < s.length / 2; i++) {
    if (s[i] !== s[s.length - i - 1]) {
      return false;
    }
  }
  return true;
};

console.log(validPalindrome("aba")); // true
console.log(validPalindrome("abca")); // true
console.log(validPalindrome("abc")); // false
console.log(validPalindrome("abccdba")); // true

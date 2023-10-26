/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var backspaceCompare = function (s, t) {
  let p1 = s.length - 1,
    p2 = t.length - 1;
  while (p1 >= 0 || p2 >= 0) {
    console.log({ p1, p2 });
    if (s[p1] === "#" || t[p2] === "#") {
      if (s[p1] === "#") {
        let backCount = 2;
        while (backCount > 0) {
          p1--;
          backCount--;
          if (s[p1] === "#") {
            backCount += 2;
          }
        }
        console.log({ p1, p2 });
      }
      if (t[p2] === "#") {
        let backCount = 2;
        while (backCount > 0) {
          p2--;
          backCount--;
          if (t[p2] === "#") {
            backCount += 2;
          }
        }
        console.log({ p1, p2 });
      }
      continue;
    }
    if (s[p1] !== t[p2]) {
      return false;
    }
    p1--;
    p2--;
  }
  console.log({ p1, p2 });
  return true;
};

console.log(backspaceCompare("ab#c", "ad#c")); //  true
console.log(backspaceCompare("ab#c", "ad##c")); // false
console.log(backspaceCompare("ab#c", "ad####c")); // false
console.log(backspaceCompare("xywrrmp", "xywrrmu#p")); // true
console.log(backspaceCompare("ab##", "c#d#")); // true
console.log(backspaceCompare("a##c", "#a#c")); // true

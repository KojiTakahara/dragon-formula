var f = angular.module('stringFilter', []);
f.filter('abbreviate', function () {
  return function (text, len, end) {
    if (len === undefined) { // デフォルトは10文字
      len = 10;
    }
    if (end === undefined) {
      end = "…";
    }
    if(text !== undefined) {
      if(text.length > len) {
        return text.substring(0, len - 1) + end;
      } else {
        return text;
      }
    }
  };
});
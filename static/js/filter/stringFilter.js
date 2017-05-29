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

f.filter("statusToJapanese", function() {
    return function(status) {
        switch(status) {
        case "REVIEW":
            return "レビュー待ち";
        case "FINALCHECK":
            return "最終確認";
        case "APPROVED":
            return "承認済み";
        case "TURNDOWN":
            return "取り下げ";
        case "REJECTED":
            return "却下";
        default:
            return "";
        }
    };
});
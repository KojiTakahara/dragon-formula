var dir = angular.module('cardTextDirective', [
]);

dir.directive('replacecardtext', function() {
  return {
    restrict: 'E',
    replace: true,
    scope: {
      text: '@'
    },
    template: '<div compile="text | addModalLink"></div>',
    controller: ['$scope', '$sce', '$mdDialog', function($scope, $sce, $mdDialog) {
      $scope.showModal = function(cardName, ev) {
        var text = "";
        // TODO データから名前に一致するテキストを取得する仕組みを作る
        if (cardName == "ガンリキ・インディゴ・カイザー") {
            text = "相手がクリーチャーを召喚した時または呪文を唱えた時、このターン、相手のクリーチャーは、攻撃またはブロックできない。";
        } else {
            text = "スピードアタッカー";
        }
        var alert = $mdDialog.alert().content(text).ok('閉じる');
        $mdDialog.show(alert);
      };
    }]
  };
});

dir.filter('addModalLink', [function() {
  return function(text) {
    var result = text.replace( /(《)((?!《》).)*?(》)/g , function(all) {
      var name = all.substr(1);
      name = name.substr(0, name.length - 1);
      return "<a ng-click=\"showModal('" + name + "')\">" + all + "</a>";
    });
    return result;
  }
}]);

dir.directive('compile', ['$compile', function ($compile) {
  return function(scope, element, attrs) {
    scope.$watch(
      function(scope) {
        return scope.$eval(attrs.compile);
      },
      function(value) {
        element.html(value);
        $compile(element.contents())(scope);
      }
    );
  };
}]);
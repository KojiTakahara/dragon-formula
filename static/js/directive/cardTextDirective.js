var dir = angular.module("cardTextDirective", [
]);

dir.directive("replacecardtext", function() {
  return {
    restrict: "E",
    replace: true,
    scope: {
      text: "@",
      annotations: "@"
    },
    template: "<div compile=\"text | addModalLink\"></div>",
    controller: ["$scope", "$sce", "$mdDialog", "questionAnnotationService", function($scope, $sce, $mdDialog, questionAnnotationService) {
      $scope.showModal = function(cardName, ev) {
        var annotations = JSON.parse($scope.annotations);
        var text = "";
        if (!angular.isUndefined(annotations)) {
          for (var i = 0; i < annotations.length; i++) {
            if (cardName === annotations[i].CardName) {
              text = annotations[i].Annotation;
            }
          }
        }
        var alert = $mdDialog.alert().content(text).ok("閉じる");
        $mdDialog.show(alert);
      };
    }]
  };
});

dir.filter("addModalLink", [function() {
  return function(text) {
    var result = text.replace( /(《)((?!《》).)*?(》)/g , function(all) {
      var name = all.substr(1);
      name = name.substr(0, name.length - 1);
      return "<a ng-click=\"showModal('" + name + "')\">" + all + "</a>";
    });
    return result;
  }
}]);

dir.directive("compile", ["$compile", function ($compile) {
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
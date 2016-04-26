var dir = angular.module('rubricDirective', [
]);

dir.directive('rubricButton', function() {
  return {
    restrict: 'E',
    replace: true,
    scope: {
      text: '@'
    },
    template: '<span><md-button ng-click="openRubric($event, text)" class="md-raised">解説</md-button></span>',
    controller: ['$scope', '$sce', '$mdDialog', function($scope, $sce, $mdDialog) {
      $scope.openRubric = function(ev, text) {
        var alert = $mdDialog.alert().content(text).ok('閉じる');
        $mdDialog.show(alert);
      };
    }]
  };
});
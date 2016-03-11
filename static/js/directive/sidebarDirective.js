var dir = angular.module('sidebarDirective', [
  'userService'
]);

dir.directive('sidebar', function() {
  return {
    restrict: 'E',
    replace: true,
    templateUrl: '/view/common/sidebar.html',
    controller: ['$rootScope', '$scope', '$window', '$location', 'userService', function($rootScope, $scope, $window, $location, userService) {
      $scope.login = function() {
        $window.location.href = '/api/twitter/login';
      };
      $rootScope.user = {};
      /**
       * セッションからログインユーザの情報をとる
       */
      userService.getLoginUser().then(function(data) {
        $rootScope.user = data; // 成功
      }, function(e) {
        $rootScope.user = undefined;
      });
    }]
  };
});

dir.directive('copyright', function() {
  return {
    restrict: 'E',
    replace: true,
    scope: {
      name: '@'
    },
    template: '<small>Copyright &copy; {{year}} {{name}} All Rights Reserved.</small>',
    link: ['$scope', function($scope) {
      $scope.year = new Date().getFullYear();
    }]
  };
});
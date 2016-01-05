"use strict";

var app = angular.module('indexCtrl', [
  'userService'
]);
app.controller('indexCtrl', ['$scope', '$http', '$sce', '$window', '$mdDialog', 'userService',
function($scope, $http, $sce, $window, $mdDialog, userService) {
  $scope.user = {};
  userService.getLoginUser().then(function(data) {
    $scope.user = data; // 成功
  }, function(e) {
    $scope.user = undefined;
  });

  $scope.movePage = function(ev, nextPageUrl) {
    console.log(nextPageUrl);
    if ($scope.user) {
      $window.location.href = nextPageUrl;
    } else {
      var confirm = $mdDialog.confirm()
          .title('Twitterを使ってログインします。よろしいですか？')
          .targetEvent(ev)
          .ok('OK')
          .cancel('キャンセル');
      $mdDialog.show(confirm).then(function() {
        $window.location.href = '/api/twitter/login';
      }, function() { // cancel
        // none
      });
    }
  };

}]);
"use strict";

var app = angular.module("indexCtrl", [
  "userService"
]);

app.controller("indexCtrl", ["$rootScope", "$scope", "$http", "$sce", "$window", "$mdDialog", "userService",
function($rootScope, $scope, $http, $sce, $window, $mdDialog, userService) {

  /**
   * 次画面への遷移. 次画面URLを受け取って, ログインしていればOK.
   */
  $scope.movePage = function(ev, nextPageUrl) {
    if (true) {
    // if ($rootScope.user) {
      $window.location.href = nextPageUrl;
    } else {
      var confirm = $mdDialog.confirm()
          .title("Twitterを使ってログインします。よろしいですか？")
          .targetEvent(ev)
          .ok("OK")
          .cancel("キャンセル");
      $mdDialog.show(confirm).then(function() {
        $window.location.href = "/api/twitter/login";
      });
    }
  };

}]);

app.directive("maincard", function() {
  return {
    restrict: "E",
    replace: true,
    scope: {
      title: "@",
      description: "@",
      link: "@",
      image: "@"
    },
    templateUrl: "/view/common/top_card.html",
    controller: "indexCtrl"
  };
});
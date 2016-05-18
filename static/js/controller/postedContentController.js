"use strict";

var app = angular.module("postedContentCtrl", [
  "categoryService",
  "questionService",
  "userService"
]);
app.controller("postedContentCtrl", ["$scope", "$rootScope", "$http", "$sce", "$window", "$mdDialog", "categoryService", "questionService", "userService",
function($scope, $rootScope, $http, $sce, $window, $mdDialog, categoryService, questionService, userService) {
  $scope.processed = false;
  $scope.qFilter = {};
  $scope.sortReverse = false;

  questionService.search(null, null, null, $rootScope.user.Key).then(function(data) {
    $scope.postedQuestions = data; 
  }, function(e) {
    console.log(e);
  });
  
  /**
   * 取り下げダイアログを表示する
   */
  $scope.showConfirm = function(ev, question) {
    $scope.processed = true;
    var confirm = $mdDialog.confirm()
          .title("投稿を取り下げてもよろしいですか？")
          .targetEvent(ev)
          .ok("OK")
          .cancel("キャンセル");
    $mdDialog.show(confirm).then(function() {
      submission(ev, question);      
    }, function() { // cancel
      $scope.processed = false;
    });
  };
  
  var submission = function(ev, question) {
    question.Status = "TURNDOWN";
    questionService.update(question).then(function(data) {
      var confirm = $mdDialog.confirm()
          .title("投稿を取り下げました。")
          .targetEvent(ev)
          .ok("OK")
      $mdDialog.show(confirm);
	  });
  };
  
  // 問題の確認
  $scope.showContentModal = function($event, question) {
    $mdDialog.show({
      clickOutsideToClose: true,
      parent: angular.element(document.body),
      targetEvent: $event,
      templateUrl: "/view/common/questionModal.html",
      locals: {
        question: question
      },
      controller: function ModalController($scope, $mdDialog, question) {
        $scope.q = question;
        $scope.closeDialog = function() {
          $mdDialog.hide();
        };
      }
    });
  };

}]);
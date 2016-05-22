/** カードルールコントローラー */
"use strict";

var app = angular.module("cardRulesCtrl", []);
app.controller("cardRulesCtrl", ["$scope", "$rootScope", "$http", "$sce", "$window", "$mdDialog", "questionService", "questionAnnotationService", "userService", "userAnswerService",
function($scope, $rootScope, $http, $sce, $window, $mdDialog, questionService, questionAnnotationService, userService, userAnswerService) {
  $scope.user = {};
  $scope.processed = false;
  $scope.showAnswerResult = false;
  $scope.rightAnswer = 0;
  $scope.questions = [];
  var questionCount = 15;
  var ruleCategory = "rule_3";

  /**
   * 初期処理
   */
  var init = function() {
    questionService.search(ruleCategory, "APPROVED", questionCount).then(function(data) {
	    $scope.questions = data;
      setTimeout(function() {
        $(".carousel").slick({infinite: false, dots: false, arrows: true});
        $(".slick-prev").css("display", "none");
        $(".slick-next").css("display", "none");
      }, 0);
      $scope.processed = false;
    }, function(e) {
      $scope.processed = true;
    });
  };
  init();

  /**
   * 回答送信ダイアログを表示する
   */
  $scope.showConfirm = function(ev) {
    $scope.processed = true;
    var quantity = questionService.getQuantityResponses($scope.questions);
    var confirm = $mdDialog.confirm()
          .title("送信してもよろしいですか？")
          .textContent("回答数 " + quantity + "/" + $scope.questions.length)
          .targetEvent(ev)
          .ok("OK")
          .cancel("キャンセル");
    $mdDialog.show(confirm).then(function() {
      postUserAnswer();
    }, function() {
      $scope.processed = false;
    });
  };

  /**
   * ユーザの解答データを送信する
   */
  var postUserAnswer = function() {
    var answer = userAnswerService.convertData($scope.questions, ruleCategory, $rootScope.user, questionService);
    $scope.rightAnswer = answer.rightAnswer;
    userAnswerService.create(answer).then(function(data) {
      $scope.showAnswerResult = true;
	  });
  };

  $scope.moveTopPage = function() {
    $window.location.href = "/";
  };

  $scope.prev = function() {
    $(".slick-prev").click();
  };

  $scope.next = function() {
    $(".slick-next").click();
  };
}]);
"use strict";

var app = angular.module('competitionCtrl', []);
app.controller('competitionCtrl', ['$scope', '$http', '$sce', '$window', '$mdDialog', 'questionService', 'userService', 'userAnswerService',
function($scope, $http, $sce, $window, $mdDialog, questionService, userService, userAnswerService) {
  $scope.user = {};
  $scope.processed = false;
  $scope.showAnswerResult = false;
  $scope.rightAnswer = 0;
  $scope.questions = [];
  var questionCount = 15;

  var init = function() {
    questionService.search("rule_2", null, questionCount).then(function(data) {
			$scope.questions = data;
      setTimeout(function() {
        $(".carousel").slick({
          infinite: false,
          dots: false,
          arrows: true,
        });
        $('.slick-prev').css("display", "none");
        $('.slick-next').css("display", "none");
      }, 0);
      $scope.processed = false;
		}, function(e) {
      $scope.processed = true;
      console.log(e);
    });
    userService.getLoginUser().then(function(data) {
      $scope.user = data; // 成功
    }, function(e) {
      $scope.user = undefined;
    });

	};
	init();

	/**
   * 回答送信ダイアログを表示する
   */
	$scope.showConfirm = function(ev) {
    $scope.processed = true;
    var confirm = $mdDialog.confirm()
          .title('送信してもよろしいですか？')
          .textContent('回答数 ' + getQuantityResponses() + '/' + $scope.questions.length)
          .targetEvent(ev)
          .ok('OK')
          .cancel('キャンセル');
    $mdDialog.show(confirm).then(function() {
      createUserAnswer();
    }, function() { // cancel
      $scope.processed = false;
    });
  };

  var createUserAnswer = function() {
    var answer = {
      userKey: $scope.user.screen_name,
      categoryKey: "rule_2",
    }, wrongAnswer = 0;
    for (var i = 1; i <= $scope.questions.length; i++) {
      answer['question' + i] = $scope.questions[i - 1].Key;
      answer['category' + i] = $scope.questions[i - 1].SmallCategoryKey;
      var result = getTrueFalse($scope.questions[i -1]);
      answer['corrected' + i] = result;
      result ? $scope.rightAnswer++ : wrongAnswer++;
    }
    answer.rightAnswer = $scope.rightAnswer;
    answer.wrongAnswer = wrongAnswer;
    userAnswerService.create(answer).then(function(data) {
      $scope.showAnswerResult = true;
		});
  };

  var getTrueFalse = function(question) {
    switch (question.selected - 0) {
      case question.Choice1.Key:
        return question.Choice1.TrueFalse;
      case question.Choice2.Key:
        return question.Choice2.TrueFalse;
      case question.Choice3.Key:
        return question.Choice3.TrueFalse;
    }
    return false;
  };

  /**
   * 回答数のカウント
   */
  var getQuantityResponses = function() {
    var result = 0;
    for (var i = 0; i < $scope.questions.length; i++) {
      if ($scope.questions[i].selected) {
        result++;
      }
    }
    return result;
  };

  $scope.moveTopPage = function() {
    $window.location.href = "/";
  };
  $scope.prev = function() {
    $('.slick-prev').click();
  };
  $scope.next = function() {
    $('.slick-next').click();
  };

}]);
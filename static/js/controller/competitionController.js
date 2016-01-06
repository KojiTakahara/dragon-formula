"use strict";

var app = angular.module('competitionCtrl', []);
app.controller('competitionCtrl', ['$scope', '$http', '$sce', '$window', '$mdDialog', 'questionService', 'userService', 'userAnswerService',
function($scope, $http, $sce, $window, $mdDialog, questionService, userService, userAnswerService) {
  $scope.user = {};
  $scope.processed = false;
  
  var init = function() {
    	questionService.search("rule_2", null, 15).then(function(data) {
			$scope.questions = data;
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
          .textContent('回答数 1/' + $scope.questions.length)
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
      rightAnswer: 1, // TODO
      wrongAnswer: 1  // TODO
    };
    for (var i = 1; i <= $scope.questions.length; i++) {
      answer['question' + i] = $scope.questions[i - 1].Key;
      answer['category' + i] = $scope.questions[i - 1].SmallCategoryKey;
      answer['corrected' + i] = ""; // TODO
    }
    userAnswerService.create(answer).then(function(data) {
		});
  };

}]);
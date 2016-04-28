"use strict";

var app = angular.module('submissionCtrl', [
  'questionService',
  'userService'
]);
app.controller('submissionCtrl', ['$scope', '$http', '$sce', '$window', '$mdDialog', 'categoryService', 'questionService', 'userService',
function($scope, $http, $sce, $window, $mdDialog, categoryService, questionService, userService) {
  $scope.processed = false;
  
  var init = function() {
    $scope.categories = [
        "カードルール",
        "総合ゲームルール",
        "競技イベント運営ルール"
    ];
    $scope.trueFalse = ["○", "×"];
    
      
    categoryService.search(null, "rule_1").then(function(data) {
		$scope.rule1Categories = data;
	});
    categoryService.search(null, "rule_2").then(function(data) {
		$scope.rule2Categories = data;
	});
  };	
  init();

  /**
   * 送信ダイアログを表示する
   */
	$scope.showConfirm = function(ev) {
    $scope.processed = true;
    var confirm = $mdDialog.confirm()
          .title('送信してもよろしいですか？')
          .targetEvent(ev)
          .ok('OK')
          .cancel('キャンセル');
    $mdDialog.show(confirm).then(function() {
      submission(ev);      
    }, function() { // cancel
      $scope.processed = false;
    });
  };
  
  var submission = function(ev) {
    switch ($scope.question.categoryName) {
      case "カードルール":
        $scope.question.category = "rule_3";
        $scope.question.correctAnswer = $scope.question.trueFalse;
        if ($scope.question.trueFalse === "○") {
          $scope.question.wrongAnswer1 = "×";
        } else if ($scope.question.trueFalse === "×") {
          $scope.question.wrongAnswer1 = "○";
        }
        break;
      case "競技イベント運営ルール":
        $scope.question.category = "rule_2";
        break;
      case "総合ゲームルール":
        $scope.question.category = "rule_1";
        break;
    }

    var question = {
      content: $scope.question.content,
      choice1Content: $scope.question.correctAnswer,
      choice2Content: $scope.question.wrongAnswer1,
      choice3Content: $scope.question.wrongAnswer2,
      choice1Bool: true,
      choice2Bool: false,
      choice3Bool: false,
      largeCategoryKey: null,
      mediumCategoryKey: null,
      smallCategoryKey: $scope.question.category,
      rubric: $scope.question.rubric,
      userKey: "test" //$scope.user.screen_name
    };
    
    if ($scope.question.category === "rule_3" && $scope.question.correctAnswer === "×") {
      question.choice1Content = $scope.question.wrongAnswer1;
      question.choice2Content = $scope.question.correctAnswer;
      question.choice1Bool = false;
      question.choice2Bool = true;
    }
    
    questionService.create(question).then(function(data) {
      var confirm = $mdDialog.confirm()
          .title('送信完了しました。')
          .targetEvent(ev)
          .ok('OK')
      $mdDialog.show(confirm).then(function() {
        $window.location.href = '/';
      });
		});
  };

}]);
"use strict";

var app = angular.module('submissionCtrl', [
  'questionService',
  'userService'
]);
app.controller('submissionCtrl', ['$scope', '$http', '$sce', '$window', '$mdDialog', 'categoryService', 'questionService', 'userService',
function($scope, $http, $sce, $window, $mdDialog, categoryService, questionService, userService) {
  $scope.processed = false;
  //$scope.user = {};
  
  var init = function() {
    categoryService.search(null, "rule_1").then(function(data) {
			$scope.rule1Categories = data;
		});
    categoryService.search(null, "rule_2").then(function(data) {
			$scope.rule2Categories = data;
		});
    // userService.getLoginUser().then(function(data) {
    //   $scope.user = data; // 成功
    // }, function(e) {
    //   $scope.user = undefined;
    // });
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
      userKey: $scope.user.screen_name
    };
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
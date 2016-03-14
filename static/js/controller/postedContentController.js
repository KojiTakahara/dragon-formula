"use strict";

var app = angular.module('postedContentCtrl', [
  'categoryService',
  'questionService',
  'userService'
]);
app.controller('postedContentCtrl', ['$scope', '$http', '$sce', '$window', '$mdDialog', 'categoryService', 'questionService', 'userService',
function($scope, $http, $sce, $window, $mdDialog, categoryService, questionService, userService) {
  $scope.processed = false;
  $scope.user = {};
  
  
  $scope.qFilter = {};
  $scope.sortReverse = false;
  $scope.searchFish = "";

  questionService.search(null, null, null).then(function(data) {
    $scope.postedQuestions = data; 
  }, function(e) {
    console.log(e);
  });
  
  var init = function() {
  };
  init();

  /**
   * 取り下げダイアログを表示する
   */
  $scope.showConfirm = function(ev) {
    $scope.processed = true;
    var confirm = $mdDialog.confirm()
          .title("投稿を取り下げてもよろしいですか？")
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
    //   content: $scope.question.content,
    //   choice1Content: $scope.question.correctAnswer,
    //   choice2Content: $scope.question.wrongAnswer1,
    //   choice3Content: $scope.question.wrongAnswer2,
    //   choice1Bool: true,
    //   choice2Bool: false,
    //   choice3Bool: false,
    //   largeCategoryKey: null,
    //   mediumCategoryKey: null,
    //   smallCategoryKey: $scope.question.category,
    //   userKey: $scope.user.screen_name
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
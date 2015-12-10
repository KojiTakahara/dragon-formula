"use strict";

var app = angular.module('competitionCtrl', []);
app.controller('competitionCtrl', ['$scope', '$http', '$sce', '$window', '$mdDialog', 'questionService',
function($scope, $http, $sce, $window, $mdDialog, questionService) {
	var init = function() {
    	questionService.search().then(function(data) {
			$scope.questions = data;
		});
	};
	init();

	/**
   * 回答送信ダイアログを表示する
   */
	$scope.showConfirm = function(ev) {
    var confirm = $mdDialog.confirm()
          .title('送信してもよろしいですか？')
          .textContent('回答数 1/' + $scope.questions.length)
          .targetEvent(ev)
          .ok('OK')
          .cancel('キャンセル');
    $mdDialog.show(confirm).then(function() {
      $scope.status = 'You decided to get rid of your debt.';
    }, function() { // cancel
      $scope.status = 'You decided to keep your debt.';
    });
  };

}]);
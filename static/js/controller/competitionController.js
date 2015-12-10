"use strict";

var app = angular.module('competitionCtrl', []);
app.controller('competitionCtrl', ['$scope', '$http', '$sce', '$window', '$mdDialog', 'questionService',
function($scope, $http, $sce, $window, $mdDialog, questionService) {
	$scope.questions = [];
	for (var i = 0; i < 25; i++) {
		$scope.questions.push(i + 1);
	}
	var init = function() {
    	questionService.search().then(function(data) {
			$scope.list = data;
		});
	};
	init();
	
	/**
   * 回答送信ダイアログを表示する 
   */
	$scope.showConfirm = function(ev) {
    var confirm = $mdDialog.confirm()
          .title('送信してもよろしいですか？')
          .textContent('回答数 1/25')
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
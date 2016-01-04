"use strict";

var app = angular.module('submissionCtrl', []);
app.controller('submissionCtrl', ['$scope', '$http', '$sce', '$window', '$mdDialog', 'categoryService',
function($scope, $http, $sce, $window, $mdDialog, categoryService) {
	var init = function() {
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
    var confirm = $mdDialog.confirm()
          .title('送信してもよろしいですか？')
          //.textContent('問題文:' + $scope.question.content)
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
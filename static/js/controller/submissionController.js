"use strict";

var app = angular.module('submissionCtrl', []);
app.controller('submissionCtrl', ['$scope', '$http', '$sce', '$window', '$mdDialog', 'categoryService',
function($scope, $http, $sce, $window, $mdDialog, categoryService) {
	var init = function() {
    categoryService.search().then(function(data) {
			$scope.categories = data;
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
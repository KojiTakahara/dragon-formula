"use strict";

var app = angular.module('competitionCtrl', []);
app.controller('competitionCtrl', ['$scope', '$http', '$sce', '$window', 'questionService', function($scope, $http, $sce, $window, questionService) {
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

}]);
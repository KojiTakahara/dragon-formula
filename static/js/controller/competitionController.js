"use strict";

var app = angular.module('competitionCtrl', []);
app.controller('competitionCtrl', ['$scope', '$http', '$sce', '$window', 'questionService', function($scope, $http, $sce, $window, questionService) {

	var init = function() {
    	questionService.search().then(function(data) {
			$scope.list = data;
		});
	};
	init();

}]);
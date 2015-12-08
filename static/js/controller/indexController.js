"use strict";

var app = angular.module('indexCtrl', []);
app.controller('indexCtrl', ['$scope', '$http', '$sce', '$window', function($scope, $http, $sce, $window) {
  $scope.isStarted = false;
  $scope.questions = [];
  for (var i = 0; i < 25; i++) {
    $scope.questions.push(i + 1);
  }
  $scope.start = function() {
    $scope.isStarted = true;
  };
  $scope.login = function() {
    $window.location.href = '/api/twitter/login';
  };
}]);
"use strict";

var app = angular.module('app', [
  'categoryService',
  'questionService',
  'userAnswerService',
  'userService',
  'stringFilter',
  'ui.router',
  'ngMaterial',
]);

app.controller('indexCtrl', ['$scope', '$http', '$sce', '$window', '$mdDialog', 'questionService',
function($scope, $http, $sce, $window, $mdDialog, questionService) {

  $scope.qFilter = {};
  $scope.sortReverse = false;
  $scope.searchFish = "";

  questionService.search(null, null, null).then(function(data) {
    $scope.questions = data; 
  }, function(e) {
    console.log(e);
  });
  
  

}]);

app.config(['$httpProvider', '$locationProvider', '$stateProvider', '$urlRouterProvider',
function($httpProvider, $locationProvider, $stateProvider, $urlRouterProvider) {
  $httpProvider.defaults.headers.common = {'X-Requested-With': 'XMLHttpRequest'};
  $httpProvider.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded; charset=utf-8';
  $httpProvider.defaults.transformRequest = function(data) {
    if (data === undefined) {
      return data;
    }
    return $.param(data);
  }
  $locationProvider.html5Mode({
    enabled: true,
    requireBase: false
  });
  $urlRouterProvider.otherwise("");
}]);
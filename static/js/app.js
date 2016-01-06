"use strict";

var app = angular.module('app', [
  'indexCtrl',
  'competitionCtrl',
  'submissionCtrl',
  'categoryService',
  'questionService',
  'userAnswerService',
  'userService',
  'sidebarDirective',
  'ui.router',
  'slick',
  'ngMaterial',
  'angular-loading-bar'
]);

var closeSideMenu = function() {
  angular.element(".mdl-layout__obfuscator").click();
}

app.config(['$httpProvider', '$locationProvider', '$stateProvider', '$urlRouterProvider', 'cfpLoadingBarProvider',
function($httpProvider, $locationProvider, $stateProvider, $urlRouterProvider, cfpLoadingBarProvider) {
  cfpLoadingBarProvider.includeBar = false;
  cfpLoadingBarProvider.includeSpinner = false;
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
  $stateProvider.state('top', {
    url: '/',
    views: {
      mainContent: {
        templateUrl: '/view/top.html',
        controller: 'indexCtrl'
      }
    }
  });
  $stateProvider.state('competition', {
    url: '/competition',
    views: {
      mainContent: {
        templateUrl: '/view/competition.html',
        controller: 'competitionCtrl'
      }
    }
  });
  $stateProvider.state('submission', {
    url: '/submission',
    views: {
      mainContent: {
        templateUrl: '/view/submission.html',
        controller: 'submissionCtrl'
      }
    }
  });
}]);
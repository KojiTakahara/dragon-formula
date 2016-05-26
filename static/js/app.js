"use strict";

var app = angular.module("app", [
  "indexCtrl",
  "allCtrl",
  "allContentCtrl",
  "cardRulesCtrl",
  "competitionCtrl",
  "comprehensiveCtrl",
  "postedContentCtrl",
  "submissionCtrl",
  "categoryService",
  "questionService",
  "questionAnnotationService",
  "userAnswerService",
  "userService",
  "cardTextDirective",
  "rubricDirective",
  "sidebarDirective",
  "stringFilter",
  "ui.router",
  "ngMaterial",
]);

var closeSideMenu = function() {
  angular.element(".mdl-layout__obfuscator").click();
}

app.config(["$httpProvider", "$locationProvider", "$stateProvider", "$urlRouterProvider",
function($httpProvider, $locationProvider, $stateProvider, $urlRouterProvider) {
  $httpProvider.defaults.headers.common = {"X-Requested-With": "XMLHttpRequest"};
  $httpProvider.defaults.headers.post["Content-Type"] = "application/x-www-form-urlencoded; charset=utf-8";
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
  $stateProvider.state("top", {
    url: "/",
    views: {
      mainContent: {
        templateUrl: "/view/top.html",
        controller: "indexCtrl"
      }
    }
  });
  $stateProvider.state("cardRules", {
    url: "/cardrules",
    views: {
      mainContent: {
        templateUrl: "/view/answer_page.html",
        controller: "cardRulesCtrl"
      }
    }
  });
  $stateProvider.state("competition", {
    url: "/competition",
    views: {
      mainContent: {
        templateUrl: "/view/answer_page.html",
        controller: "competitionCtrl"
      }
    }
  });
  $stateProvider.state("comprehensive", {
    url: "/comprehensive",
    views: {
      mainContent: {
        templateUrl: "/view/answer_page.html",
        controller: "comprehensiveCtrl"
      }
    }
  });
  $stateProvider.state("all", {
    url: "/all",
    views: {
      mainContent: {
        templateUrl: "/view/answer_page.html",
        controller: "allCtrl"
      }
    }
  });
  $stateProvider.state("submission", {
    url: "/submission",
    views: {
      mainContent: {
        templateUrl: "/view/submission.html",
        controller: "submissionCtrl"
      }
    }
  });
  $stateProvider.state("posted_content", {
    url: "/posted_content",
    views: {
      mainContent: {
        templateUrl: "/view/posted_content.html",
        controller: "postedContentCtrl"
      }
    }
  });
  $stateProvider.state("all_content", {
    url: "/all_content",
    views: {
      mainContent: {
        templateUrl: "/view/all_content.html",
        controller: "allContentCtrl"
      }
    }
  });
}]);
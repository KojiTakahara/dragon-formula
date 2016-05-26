"use strict";

var app = angular.module("allContentCtrl", [
  "questionService",
]);
app.controller("allContentCtrl", ["$scope", "$http", "$sce", "$window", "$mdDialog", "questionService", "categoryService", "userService",
function($scope, $http, $sce, $window, $mdDialog, questionService, categoryService, userService) {
  $scope.qFilter = {};
  $scope.sortReverse = false;
  var categories = [];
  
  var init = function() {
    questionService.search(null, "APPROVED", null, null).then(function(data) {
      $scope.questions = data;
    }, function(e) {
      console.log(e);
    });
    categoryService.search().then(function(data) {
      categories = data;
    });
    categoryService.search(null, "").then(function(data) {
      $scope.largeCategories = data;
      $scope.largeCategories.push({Name:"すべて"});
    });
  };
  init();
  
  $scope.categoryJapanese = function(key) {
    return categoryService.getCategoryJapanese(categories, key);
  }

  // 問題の確認
  $scope.showContentModal = function($event, question) {
    $mdDialog.show({
      clickOutsideToClose: true,
      parent: angular.element(document.body),
      targetEvent: $event,
      templateUrl: "/view/common/questionModal.html",
      locals: {
        question: question
      },
      controller: function ModalController($scope, $mdDialog, question, userService) {
        if (!angular.isUndefined(question.UserKey)) {
          userService.getTwitterUser(question.UserKey).then(function(data) {
            $scope.userInfo = data;
          });
        }
        $scope.q = question;
        $scope.closeDialog = function() {
          $mdDialog.hide();
        };
      }
    });
  };
}]);
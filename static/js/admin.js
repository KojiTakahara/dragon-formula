"use strict";

var app = angular.module("app", [
  "categoryService",
  "questionAnnotationService",
  "questionService",
  "userAnswerService",
  "userService",
  "stringFilter",
  "ui.router",
  "ngMaterial",
]);

app.controller("indexCtrl", ["$scope", "$http", "$sce", "$window", "$mdDialog", "$mdToast", "questionService", "questionAnnotationService",
function($scope, $http, $sce, $window, $mdDialog, $mdToast, questionService, questionAnnotationService) {

  $scope.qFilter = {};
  $scope.sortReverse = false;
  $scope.searchFish = "";

  questionService.search(null, null, null).then(function(data) {
    $scope.questions = data; 
  }, function(e) {
    $mdToast.showSimple(e);
    console.log(e);
  });
  
  var reDialog = function(ev, question) {
    $scope.showDialog(ev, question);
  }
  
  $scope.showDialog = function($event, question) {
    $mdDialog.show({
      clickOutsideToClose: true,
      parent: angular.element(document.body),
      targetEvent: $event,
      templateUrl: "/admin/questionModal.html",
      locals: {
        question: question
      },
      controller: function ModalController($scope, $mdDialog, question) {
        console.log(question);
        $scope.annotations = [];
        $scope.question = question;
        
        questionAnnotationService.search(question.Key).then(function(data) {
          console.log(data);
          $scope.annotations = data; 
        }, function(e) {
          $mdToast.showSimple(e);
          console.log(e);
        });
        
        
        
        
        $scope.update = function(ev, question) {
          var confirm = $mdDialog.confirm()
                .title("変更してもよろしいですか？")
                .targetEvent(ev)
                .ok("OK")
                .cancel("キャンセル");
          $mdDialog.show(confirm).then(function() {
            questionService.update(question).then(function(data) {
              $mdToast.showSimple("成功しました");
            }, function(e) {
              $mdToast.showSimple("失敗しました");
            });
          }, function() {
            reDialog(ev, question);
          });
        };
        $scope.closeDialog = function() {
          $mdDialog.hide();
        };
        // Annotationの追加
        $scope.addAnnotation = function() {
          $scope.annotations.push({
            Key: undefined,
            QuestionKeyId: question.Key,
            CardName: "",
            Annotation: ""
          });
        };
        // 削除
        $scope.deleteAnnotation = function(i) {
          var annotation = $scope.annotations[i];
          if (angular.isUndefined(annotation.Key)) {
            $scope.annotations.splice(i, 1); 
          } else {
            questionAnnotationService.delete(annotation).then(function() {
              alert("success");
              $scope.annotations.splice(i, 1);
            });
          }
        };
        // 登録・更新
        $scope.submitAnnotation = function(i) {
          var annotation = $scope.annotations[i];
          if (angular.isUndefined(annotation.Key)) {
            questionAnnotationService.create(annotation).then(function() {
              alert("success");
            });
          } else {
            questionAnnotationService.update(annotation).then(function() {
              alert("success");
            });
          }
        };
      }
    });
  }

}]);

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
}]);
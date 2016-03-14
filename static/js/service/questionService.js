var service = angular.module('questionService', []);
service.factory('questionService', ['$http', '$q', function($http, $q) {
  var service = {};

  /**
   * 検索
   */
  service.search = function(largeCategory, status, limit, userKey) {
    var deferred = $q.defer();
    $http({
      method: 'GET',
      url: '/api/question',
      params: {
        largeCategoryKey: largeCategory,
        status:status,
        limit: limit,
        userKey: userKey
      },
      cache: true
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };
  
  /**
   * 登録
   */
  service.create = function(question) {
    var deferred = $q.defer();
    $http({
      method: 'POST',
      url: '/api/question',
      params: question,
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };
  
  
  service.getTrueFalse = function(question) {
    switch (question.selected - 0) {
      case question.Choice1.Key:
        return question.Choice1.TrueFalse;
      case question.Choice2.Key:
        return question.Choice2.TrueFalse;
      case question.Choice3.Key:
        return question.Choice3.TrueFalse;
    }
    return false;
  };
  
  /**
   * 回答数のカウント
   */
  service.getQuantityResponses = function(questions) {
    var result = 0;
    for (var i = 0; i < questions.length; i++) {
      if (questions[i].selected) {
        result++;
      }
    }
    return result;
  };

  return service;
}]);
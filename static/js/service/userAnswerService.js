var service = angular.module('userAnswerService', []);
service.factory('userAnswerService', ['$http', '$q', function($http, $q) {
  var service = {};

  /**
   * 登録
   */
  service.create = function(userAnswer) {
    var deferred = $q.defer();
    $http({
      method: 'POST',
      url: '/api/userAnswer',
      params: userAnswer,
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };

  return service;
}]);
var service = angular.module('questionService', []);
service.factory('questionService', ['$http', '$q', function($http, $q) {
  var service = {};

  service.search = function(largeCategory, status, limit) {
    var deferred = $q.defer();
    $http({
      method: 'GET',
      url: '/api/question',
      params: {
        largeCategoryKey: largeCategory,
        status:status,
        limit: limit
      },
      cache: true
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };

  return service;
}]);
var service = angular.module('categoryService', []);
service.factory('categoryService', ['$http', '$q', function($http, $q) {
  var service = {};

  service.search = function(type, parent) {
    var deferred = $q.defer();
    $http({
      method: 'GET',
      url: '/api/category',
      params: {
        type: type,
        parentKey: parent
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
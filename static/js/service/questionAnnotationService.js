var service = angular.module('questionAnnotationService', []);
service.factory('questionAnnotationService', ['$http', '$q', function($http, $q) {
  var service = {};

  /**
   * 検索
   */
  service.search = function(questionKeyId) {
    var deferred = $q.defer();
    $http({
      method: 'GET',
      url: '/api/questionAnnotation',
      params: {
        questionKeyId: questionKeyId
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
  service.create = function(qa) {
    var deferred = $q.defer();
    $http({
      method: 'POST',
      url: '/api/questionAnnotation',
      params: qa,
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };
  
  /**
   * 更新
   */
  service.update = function(qa) {
    var deferred = $q.defer();
    $http({
      method: 'PUT',
      url: '/api/questionAnnotation',
      params: qa,
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };
  
  /**
   * 削除
   */
  service.delete = function(qa) {
    var deferred = $q.defer();
    $http({
      method: 'DELETE',
      url: '/api/questionAnnotation',
      params: qa,
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };
  
  return service;
}]);
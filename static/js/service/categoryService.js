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
  
  /** カテゴリの日本語名を返す */
  service.getCategoryJapanese = function(categories, key) {
    for (var i = 0; i < categories.length; i++) {
      var category = categories[i];
      if (category.Key === key) {
        return category.Name;
      }
    }
    return "";
  };
  
  return service;
}]);
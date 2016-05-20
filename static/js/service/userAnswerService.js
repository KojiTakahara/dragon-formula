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
  
  service.convertData = function(questions, category, userName, questionService) {
    var userAnswer = {
      userKey: userName,
      categoryKey: category,
    }, rightAnswer = 0, wrongAnswer = 0;
    for (var i = 1; i <= questions.length; i++) {
      userAnswer['question' + i] = questions[i - 1].Key;
      userAnswer['category' + i] = questions[i - 1].SmallCategoryKey;
      var result = questionService.getTrueFalse(questions[i -1]);
      userAnswer['corrected' + i] = result;
      result ? rightAnswer++ : wrongAnswer++;
    }
    userAnswer.rightAnswer = rightAnswer;
    userAnswer.wrongAnswer = wrongAnswer;
    return userAnswer;
  };

  return service;
}]);
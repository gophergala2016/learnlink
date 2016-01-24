angular
  .module('LEARNLINK_LOGIN', ['ngRoute'])
  .controller('LoginCtrl', function ($scope) {
    $scope.user = {
      username: undefined,
      password: undefined,
      email: undefined
    };
    $scope.submit = function (user) {
      console.log(user);
    };
  });
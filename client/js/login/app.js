angular
  .module('LEARNLINK_LOGIN', ['ngRoute', 'ngCookies'])
  .service('LoginService', function ($http) {

    return {
      login: function (user_info) {
        return $http.post('/login', user_info);
      },
      signup: function (user_info) {
        return $http.post('/signup', user_info);
      }
    };

  })
  .controller('LoginCtrl', function ($scope, $cookies, LoginService) {

    $scope.loginType = 'login';

    $scope.user = {
      username: undefined,
      password: undefined,
      email: undefined
    };

    $scope.changeLoginType = function (login_type) {
      $scope.loginType = login_type;
    };

    $scope.submit = function (user, login_type) {

      LoginService[login_type](user)
        .then(function (res) {
          $cookies.put('token', res.token);
          window.user = res.user;
        }, function (err) {
          alert('Something occurred!', err.toString());
        });

    };

  });
var dependencies = [
  'ngMaterial',
  'ngRoute',
  'LEARNLINK_LOGIN'
];

angular
  .module('LEARNLINK', dependencies)
  .config(function ($routeProvider, $locationProvider, $mdThemingProvider) {

    $mdThemingProvider.theme('default')
      .primaryPalette('green')
      .accentPalette('orange')

    $routeProvider
      .when('/', {
        templateUrl: '/client/templates/index.html',
        controller: 'MainCtrl'
      })
      .when('/login', {
        templateUrl: '/client/templates/login/index.html',
        controller: 'LoginCtrl'
      });

    $locationProvider.html5Mode(true);

  })
  .controller('MainCtrl', function ($scope) {
    $scope.foo = 'bar';
  });
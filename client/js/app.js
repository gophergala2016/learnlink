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
        controller: 'CourseCtrl'
      })
      .when('/login', {
        templateUrl: '/client/templates/login/index.html',
        controller: 'LoginCtrl'
      });

    $locationProvider.html5Mode(true);

  })
  .controller('CourseCtrl', function ($scope) {
    $scope.hard_coded_courses = [
      {
      id: 4,
      name: "Algorithms & Data Structures",
      email: "plumbus@fleeb.com",
      url: "www.courseofmine.com",
      priority: 1,
      checkoff: 2,
      checkoff_time: "2016-01-23T05:18:02.706584Z",
      note: "This is my favorite course",
      created_at: "2016-01-24T05:18:02.706584Z",
      updated_at: "2016-01-24T05:18:02.706584Z"
      },
      {
      id: 5,
      name: "How to Make Kimchi Fried Rice",
      email: "plumbus@fleeb.com",
      url: "www.courseofmine.com/gero",
      priority: 3,
      checkoff: 2,
      checkoff_time: "2016-03-25T03:08:49.516044Z",
      note: "This is my favorite OFO FOFOFOFOOFFF FFFFFcourse",
      created_at: "2016-01-24T05:18:02.706584Z",
      updated_at: "2016-01-24T11:06:49.150208Z"
      }
    ];
  });
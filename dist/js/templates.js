angular.module('LEARNLINK').run(['$templateCache', function($templateCache) {
  'use strict';

  $templateCache.put('/client/index.html',
    "<head>\n" +
    "  <link rel=\"stylesheet\" href=\"http://ajax.googleapis.com/ajax/libs/angular_material/1.0.0/angular-material.min.css\">\n" +
    "  <base href=\"/\">\n" +
    "</head>\n" +
    "<body ng-app=\"LEARNLINK\">\n" +
    "  <ng-view>\n" +
    "  </ng-view>\n" +
    "  <script src=\"http://ajax.googleapis.com/ajax/libs/angularjs/1.4.8/angular.min.js\"></script>\n" +
    "  <script src=\"http://ajax.googleapis.com/ajax/libs/angularjs/1.4.8/angular-animate.min.js\"></script>\n" +
    "  <script src=\"http://ajax.googleapis.com/ajax/libs/angularjs/1.4.8/angular-aria.min.js\"></script>\n" +
    "  <script src=\"http://ajax.googleapis.com/ajax/libs/angularjs/1.4.8/angular-messages.min.js\"></script>\n" +
    "  <script src=\"http://ajax.googleapis.com/ajax/libs/angular_material/1.0.0/angular-material.min.js\"></script>\n" +
    "  <script src=\"http://ajax.googleapis.com/ajax/libs/angularjs/1.4.8/angular-route.js\"></script>\n" +
    "  <script type=\"text/javascript\" src=\"/js/base_module.js\"></script>\n" +
    "  <script type=\"text/javascript\" src=\"/js/login_module.js\"></script>\n" +
    "    <script type=\"text/javascript\" src=\"/js/templates.js\"></script>\n" +
    "</body>"
  );


  $templateCache.put('/client/templates/index.html',
    "<h1>YOU DID IT</h1>\n" +
    "<a href=\"/login\">Click</a>"
  );


  $templateCache.put('/client/templates/login/index.html',
    "<h1>Welcome to Login</h1>"
  );

}]);

angular.module('LEARNLINK').run(['$templateCache', function($templateCache) {
  'use strict';

  $templateCache.put('/client/index.html',
    "<head>\n" +
    "  <link rel=\"stylesheet\" href=\"http://ajax.googleapis.com/ajax/libs/angular_material/1.0.0/angular-material.min.css\">\n" +
    "  <link rel=\"stylesheet\" href=\"/css/index.css\">\n" +
    "  <base href=\"/\">\n" +
    "</head>\n" +
    "<body ng-app=\"LEARNLINK\">\n" +
    "  <md-toolbar md-scroll-shrink>\n" +
    "    <div class=\"md-toolbar-tools\">\n" +
    "      <h3>\n" +
    "        <span>LearnLink</span>\n" +
    "      </h3>\n" +
    "    </div>\n" +
    "  </md-toolbar>\n" +
    "  <div ng-view class=\"center-layout\">\n" +
    "  </div>\n" +
    "  <script src=\"http://ajax.googleapis.com/ajax/libs/angularjs/1.4.8/angular.min.js\"></script>\n" +
    "  <script src=\"http://ajax.googleapis.com/ajax/libs/angularjs/1.4.8/angular-animate.min.js\"></script>\n" +
    "  <script src=\"http://ajax.googleapis.com/ajax/libs/angularjs/1.4.8/angular-aria.min.js\"></script>\n" +
    "  <script src=\"http://ajax.googleapis.com/ajax/libs/angularjs/1.4.8/angular-messages.min.js\"></script>\n" +
    "  <script src=\"http://ajax.googleapis.com/ajax/libs/angularjs/1.4.8/angular-cookies.min.js\"></script>\n" +
    "  <script src=\"http://ajax.googleapis.com/ajax/libs/angular_material/1.0.0/angular-material.min.js\"></script>\n" +
    "  <script src=\"http://ajax.googleapis.com/ajax/libs/angularjs/1.4.8/angular-route.js\"></script>\n" +
    "  <script type=\"text/javascript\" src=\"/js/base_module.js\"></script>\n" +
    "  <script type=\"text/javascript\" src=\"/js/login_module.js\"></script>\n" +
    "  <script type=\"text/javascript\" src=\"/js/templates.js\"></script>\n" +
    "</body>"
  );


  $templateCache.put('/client/templates/index.html',
    "<div layout=\"column\" flex>\n" +
    "  <!-- add a course -->\n" +
    "  <!-- courses -->\n" +
    "  <md-card ng-repeat=\"course in hard_coded_courses\">\n" +
    "    <!-- <img ng-src=\"{{imagePath}}\" class=\"md-card-image\" alt=\"Washed Out\"> -->\n" +
    "    <md-card-title>\n" +
    "      <md-card-title-text>\n" +
    "        <span class=\"md-headline\">{{ course.name }}</span>\n" +
    "      </md-card-title-text>\n" +
    "    </md-card-title>\n" +
    "    <md-card-content>\n" +
    "      <p>{{ course.note }}</p>\n" +
    "    </md-card-content>\n" +
    "    <md-card-actions layout=\"row\" layout-align=\"end center\">\n" +
    "      <md-button>Checkoff</md-button>\n" +
    "    </md-card-actions>\n" +
    "  </md-card>\n" +
    "</div>"
  );


  $templateCache.put('/client/templates/login/index.html',
    "<div layout=\"column\">\n" +
    "  <div flex>\n" +
    "    <md-content layout-padding>\n" +
    "\n" +
    "        <form name=\"userForm\" ng-submit=\"submit(user, loginType)\">\n" +
    "          <div layout-gt-sm=\"row\">\n" +
    "            <md-input-container class=\"md-block\" flex-gt-sm>\n" +
    "              <label>Username</label>\n" +
    "              <input type=\"text\" ng-model=\"user.username\">\n" +
    "            </md-input-container>\n" +
    "          </div>\n" +
    "          <div layout-gt-sm=\"row\">\n" +
    "            <md-input-container class=\"md-block\" flex-gt-sm>\n" +
    "              <label>Password</label>\n" +
    "              <input type=\"password\" ng-model=\"user.password\">\n" +
    "            </md-input-container>\n" +
    "          </div>\n" +
    "          <div layout-gt-sm=\"row\">\n" +
    "            <md-input-container class=\"md-block\" flex-gt-sm>\n" +
    "              <label>Email</label>\n" +
    "              <input type=\"email\" ng-model=\"user.email\">\n" +
    "            </md-input-container>\n" +
    "          </div>\n" +
    "          <md-button type=\"submit\" class=\"md-raised md-primary\">{{ loginType.toUpperCase() }}</md-button>\n" +
    "        </form>\n" +
    "\n" +
    "        <a href=\"#\" ng-if=\"loginType === 'login'\" ng-click=\"changeLoginType('signup')\">Not a registered user? Sign up!</a>\n" +
    "        <a href=\"#\" ng-if=\"loginType === 'signup'\" ng-click=\"changeLoginType('login')\">Already a user? Log in!</a>\n" +
    "    </md-content>\n" +
    "  </div>\n" +
    "</div>"
  );

}]);

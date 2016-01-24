module.exports = function(grunt) {

  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    concat: {
      options: {
        separator: '; ',
      },
      css: {
        src: ['client/index.css'],
        dest: 'dist/css/index.css'
      },
      base_module: {
        src: ['client/js/app.js'],
        dest: 'dist/js/base_module.js'
      },
      login_module: {
        src: ['client/js/login/app.js', 'client/js/login/**/*.js'],
        dest: 'dist/js/login_module.js'
      }
    },
    ngtemplates: {
      LEARNLINK: {
        src: 'client/**/*.html',
        dest: 'dist/js/templates.js',
        options: {
          prefix: '/'
        }
      }
    },
    watch: {
      files: ['client/**/*'],
      tasks: ['default']
    }
  });

  grunt.loadNpmTasks('grunt-contrib-concat');
  grunt.loadNpmTasks('grunt-angular-templates');
  grunt.loadNpmTasks('grunt-contrib-watch');

  // Default task(s).
  grunt.registerTask('default', ['concat', 'ngtemplates']);

};
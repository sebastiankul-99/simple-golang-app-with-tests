pipeline {
    agent {
        dockerfile {
            filename 'Dockerfile-build'
            dir 'build'
            label 'docker_app_build_image'
            args '-v in-vol:/build  -v out-vol:/output'
        }
        dockerfile {
            filename 'Dockerfile-test'
            dir 'build'
            label 'docker_app_build_test'
            args '-v in-vol:/build  -v out-vol:/output'
        }
    }
   
    stages {
        stage('Build') {
            agent {
                docker { 
                    image 'docker_app_build_image:latest'
                    }
                }
            steps {
                 
                    sh 'rm -r ../../build/*'
                    sh 'rm -r ../../output/*'
                    sh 'cp -r !(simple-golang-app-with-tests)  ../../build/'
                    sh 'cp -r  . ../../output/'
                }
                       
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}
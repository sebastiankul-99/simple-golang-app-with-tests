pipeline {
    agent none
    stages {
        stage('CreateBuildImage') {
            agent any
            steps {
                   
                    sh "docker build --file Dockerfile-build --tag docker_app_build_image:latest ."
                    sh "docker images "
                    
                 }
        }
        stage('Build') {
            agent {
                docker {
                    image'docker_app_build_image:latest'
                    args '-v in-vol:/build  -v out-vol:/output '
                    reuseNode false
                    }
                }
            steps {
                sh 'ls'
                sh 'ls ../'
                sh 'ls ../../'
                sh 'ls ../../workspace'
                sh 'go version'
             //   sh 'rm -r ../../build/*'
              //  sh 'rm -r ../../output/*'
                sh 'pwd'
                sh 'cd '
                sh 'ls'
                
               // sh 'cp -r !(simple-golang-app-with-tests)  ../../build/'
               // sh 'cp -r  . ../../output/'
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
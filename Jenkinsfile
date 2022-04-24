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
                    args '-v in-vol:/build  -v out-vol:/output -w /app/simple-golang-app-with-tests '
                    reuseNode true
                    }
                }
            steps {
                sh 'ls'
                sh 'ls ../'
                sh 'ls ../../'
                sh 'ls ~/'
                sh 'ls /app/simple-golang-app-with-tests'
                sh 'go version'

                dir(path: '../../../../app/simple-golang-app-with-tests') {
               

                sh 'pwd'
                 }
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
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
                    args '-v in-vol:/build  -v out-vol:/output  --user root'
                    reuseNode false
                    }
                }
            steps {
                sh 'rm -rf /build/*'
                sh 'rm -rf /output/*'
                sh 'pwd'
                sh 'cd '
                sh 'ls'
                
               // sh 'cp -r /app/simple-golang-app-with-tests/!(simple-golang-app-with-tests)  /build/'
                sh 'cp -r . /build/'
                sh 'cp -r  /app/simple-golang-app-with-tests /output/'
                sh 'ls /build'
                sh 'ls /output'
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
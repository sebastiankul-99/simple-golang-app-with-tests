pipeline {
    agent any

    stages {
        stage('CreateBuildImage') {
            steps {
                   
                    
                    sh "docker images "
                    sh "docker volume ls"
                    sh 'ls '
                    sh 'ls ../'
                    
                 }
        }
        stage('Build') {
            
                
            steps {
                sh "docker build --file Dockerfile-build --tag docker_app_build_image:latest ."
                sh "docker run -v in-vol:/build -v out-vol:/output --name --rm -it docker_app_build_image:latest"
                sh 'ls'
                sh 'ls ../'
                sh 'ls ../../'
                sh 'rm -r ../../build/*'
                sh 'rm -r ../../output/*'
                sh 'cp -r !(simple-golang-app-with-tests)  ../../build/'
                sh 'cp -r  . ../../output/'
            }*/
           
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
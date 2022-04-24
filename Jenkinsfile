pipeline {
    agent any

    stages {
        stage('Build') {
            
            steps {
                sh "docker build --file Dockerfile-build --tag docker_app_build_image:latest ."
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
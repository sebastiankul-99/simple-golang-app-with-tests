pipeline {
    agent any

    stages {
        stage('Build') {
            agent {
                dockerfile {
                    filename 'Dockerfile-build'
                }
            }
            steps {
                sh "docker images"
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
pipeline {
    agent any

    stages {
        stage('CreateBuildImage') {
            steps {
                   
                    sh "docker build --file Dockerfile-build --tag docker_app_build_image:latest ."
                    sh "docker images "
                    
                 }
        }
        stage('Build') {
            agent {
                docker {
                    image 'node'
                    // Run the container on the node specified at the top-level of the Pipeline, in the same workspace, rather than on a new node entirely:
                    reuseNode true
                }
            }
            steps {
                sh 'node --version'
                sh 'ls '
            }
            /*agent {
                docker {
                    image'docker_app_build_image:latest'
                    args '-v in-vol:/build  -v out-vol:/output '
                    reuseNode true
                    }
                }
            steps {
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
pipeline {
    agent any

    def dockerfile_build = 'Dockerfile-build'
    def dockerfile_test = 'Dockerfile-test'
    def buildImage = docker.build("docker_app_build_image:latest", "-f ${dockerfile_build} .") 
    def testImage = docker.build("docker_app_build_test:latest", "-f ${dockerfile_test} .") 
    stages {
        stage('Build') {
            agent {
                docker { 
                    image:'docker_app_build_image:latest'
                    args '-v in-vol:/build  -v out-vol:/output --rm'
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
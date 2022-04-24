pipeline {
    agent any

   // def dockerfile_build = 'Dockerfile-build'
   // def dockerfile_test = 'Dockerfile-test'
   // def buildImage = docker.build("docker_app_build_image:latest", "-f ${dockerfile_build} .") 
   // def testImage = docker.build("docker_app_build_test:latest", "-f ${dockerfile_test} .") 
    stages {
        stage('Build') {
            agent {
                docker { 
                
                    }
                }
            steps {
                sh "docker build --file Dockerfile-build --tag docker_app_build_image:latest ."
                docker.image('docker_app_build_image:latest').withRun('-v in-vol:/build  -v out-vol:/output --rm') { c ->
                    
                    sh 'rm -r ../../build/*'
                    sh 'rm -r ../../output/*'
                    
                    sh 'cp -r !(simple-golang-app-with-tests)  ../../build/'
                     sh 'cp -r  . ../../output/'
                }
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
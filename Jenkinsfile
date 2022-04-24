pipeline {
    agent none

    /*agent {
        dockerfile {
            filename 'Dockerfile-build'
            dir 'build'
            //label 'docker_app_build_image'
            args '-v in-vol:/build  -v out-vol:/output'
        }
        
    }*/
    /*
    agent {

        dockerfile {
            filename 'Dockerfile-test'
            dir 'build'
            label 'docker_app_build_test'
            args '-v in-vol:/build  -v out-vol:/output'
        }
    }*/

  /* script{
        docker.build("docker_app_build_image:latest", "-f Dockerfile-build .") 
        docker.build("docker_app_build_test:latest", "-f Dockerfile-test .") 
   }*/
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
                    }
                    // Run the container on the node specified at the top-level of the Pipeline, in the same workspace, rather than on a new node entirely:
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
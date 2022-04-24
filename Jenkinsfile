pipeline {
    agent none
    stages {
        stage('Build') {
            agent any
            steps {
                   
                    sh "docker build --file Dockerfile-build --tag docker_app_build_image:latest ."
                    sh "docker images "
                    scripts {
                        
                           build_container = docker.image('docker_app_build_image:latest').withRun('-v in-vol:/build  -v out-vol:/output --rm -it ') { c ->
                                //docker.image('docker_app_build_image:latest').inside{
                                    sh 'rm -rf /build/*'
                                    sh 'rm -rf /output/*'
                                // sh 'cp -r /app/simple-golang-app-with-tests/!(simple-golang-app-with-tests)  /build/'
                                    sh 'cp -r . /build/'
                                    sh 'cp -r  /app/simple-golang-app-with-tests /output/'
                                    sh 'ls /build'
                                    sh 'ls /output'
                                    sh 'ls /output/simple-golang-app-with-tests'
                                //}
                           }
                        }
                    
                 }
        }
        /*
        stage('CopyBuildToVolumes') {
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
               // sh 'cp -r /app/simple-golang-app-with-tests/!(simple-golang-app-with-tests)  /build/'
                sh 'cp -r . /build/'
                sh 'cp -r  /app/simple-golang-app-with-tests /output/'
                sh 'ls /build'
                sh 'ls /output'
                sh 'ls /output/simple-golang-app-with-tests'
            }
           
        }*/
        stage('BuildTest') {
            agent any
            steps {
                   
                    sh "docker build --file Dockerfile-test --tag docker_app_build_test:latest ."
                    sh "docker images "
                    
                 }
        }
        stage('Test') {
             agent {
                docker {
                    image'docker_app_build_test:latest'
                    args '-v in-vol:/build  -v out-vol:/output  --user root'
                    reuseNode false
                    }
                }
            steps {
                sh 'cd /output/simple-golang-app-with-tests && go test '   
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}
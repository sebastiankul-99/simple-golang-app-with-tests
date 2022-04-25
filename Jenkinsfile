pipeline {
    agent none
    stages {
        stage('Logging') {
             agent any
             steps {
                sh 'mkdir -p logs'
                sh 'cd logs && mkdir -p test.log'
                sh 'cd logs and touch test.log.0000_first'
                sh 'echo "these are logs collected by fluentd " >logs/test.log.first'
                sh 'docker run -d  --name fluentd --user root -v /var/lib/docker/containers:/fluentd/log/containers -v `pwd`/fluent.conf:/fluentd/etc/fluent.conf -v `pwd`/logs:/output --log-driver local fluent/fluentd:v1.14.6-debian-1.0'
               // sh ' docker run --rm --name iperf-server --network devops-net  -p 5201:5201 -d  networkstatic/iperf3 -s '
               // sh 'docker run  --rm --name  iperf-client --network devops-net    networkstatic/iperf3 -c iperf-server'
              // sh 'sleep 25s'
             }
        }
       
        stage('Build') {
            agent any
            steps {
                    
                    sh "docker build --file Dockerfile-build --tag docker_app_build_image:latest ."
                    sh "docker images "
                    /*scripts {
                         
                         def  build_container = docker.image('docker_app_build_image:latest', "--rm build_container").withRun('--name build_container -v in-vol:/build  -v out-vol:/output  --user root') { 
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
                        }*/
                    
            }     

        }
        
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
                sh 'cp -r /app/simple-golang-app-with-tests/*  /build/'
                sh 'rm /build/simple-golang-app-with-tests' 
                sh 'cp -r  /app/simple-golang-app-with-tests/* /output/'
                sh 'ls /build'
                sh 'ls /output'
                sh 'echo "these are building container logs" >&2'
                //sh 'sleep 60s'
                
            }
           
        }
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
                sh 'cd /output && go test ' 
                 sh 'echo "these are building container logs" >&2'  
                //sh 'sleep 60s'
                
            }
        }
        stage('clean-up') {
            agent any
            steps {

                sh 'docker ps -a'
                sh 'ls logs'
              //  sh 'ls /var/lib/docker/containers'
                echo 'Deploying....'
                //sh 'docker stop iperf-server'
                
                sh 'docker stop fluentd'
                sh 'docker rm fluentd'
               
                script {
                    docker.image('docker_app_build_test').withRun('--user root') { c->
                    sh 'ls'
                    sh 'rm output.log'
                    sh 'cat logs/test.log.* > output.log'
                
                    sh 'rm -rf logs'
                    
                    }
                }    
                archiveArtifacts artifacts: 'output.log', fingerprint: true        
            }
        }
        stage('Deploy') {
             agent {
                docker {
                    image'docker_app_build_test:latest'
                    args '-v in-vol:/build  -v out-vol:/output  --user root'
                    reuseNode false
                    }
                }
            steps {
                sh 'ls /output'
                sh 'cd /output && go test ' 
                
                //sh 'sleep 60s'
                
            }
        }
    }
}
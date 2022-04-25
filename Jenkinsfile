pipeline {
    agent none
    stages {
        stage('Logging') {
             agent any
             steps {
          
                sh 'docker run -d  --name fluentd --user root -v /var/lib/docker/containers:/fluentd/log/containers -v `pwd`/fluent.conf:/fluentd/etc/fluent.conf -v `pwd`/logs:/output --log-driver local fluent/fluentd:v1.11-debian'
            //    sh ' docker run --rm --name iperf-server --network devops-net  -p 5201:5201 -d  networkstatic/iperf3 -s '
              //  sh 'docker run  --rm --name  iperf-client --network devops-net    networkstatic/iperf3 -c iperf-server'
               sh 'sleep 1s'
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
                sh 'cp -r . /build/'
                sh 'cp -r  /app/simple-golang-app-with-tests /output/'
                sh 'ls /build'
                sh 'ls /output'
                sh 'ls /output/simple-golang-app-with-tests'
                sh 'sleep 1s'
                
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
                sh 'cd /output/simple-golang-app-with-tests && go test '   
                sh 'sleep 1s'
            }
        }
        stage('Deploy') {
            agent any
            steps {

                sh 'docker ps -a'
                sh 'ls logs'
              //  sh 'ls /var/lib/docker/containers'
                echo 'Deploying....'
               // sh 'docker stop iperf-server'
                
                sh 'docker stop fluentd'
                sh 'docker rm fluentd'
                sh ' cd logs && cat test.log.* > output.log'
                 scripts {
    
                        docker.image('docker_app_build_test').withRun('--name clean_up --rm --user root') { c ->
                            /* Wait until mysql service is up */
                            sh 'ls'
                        }
          
                       
                    }

                
                
            }
        }
    }
}
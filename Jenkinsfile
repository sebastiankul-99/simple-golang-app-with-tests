pipeline {
    agent any
    
    parameters{
        string(name: 'VERSION', defaultValue: '1.0.0', description: 'release version')
        booleanParam(name: 'RELEASE', defaultValue: false, description: 'should promote to release version')
    }

    stages {
        stage('Logging') {
            
             steps {
                sh 'mkdir -p logs'
                sh 'cd logs && mkdir -p test.log'
                sh 'cd logs and touch test.log.zz_first'
                sh 'echo "these are logs collected by fluentd " >logs/test.log.first'
                sh 'docker run -d  --rm --name fluentd --user root -v /var/lib/docker/containers:/fluentd/log/containers -v `pwd`/fluent.conf:/fluentd/etc/fluent.conf -v `pwd`/logs:/output --log-driver local fluent/fluentd:v1.14.6-debian-1.0'
                sh 'docker run --rm -d --name red redis '
                
                script {    
                    env.GIT_COMMIT_REV = sh(returnStdout: true, script: "git log -n 1 --pretty=format:'%h'").trim()
                }
                sh 'docker stop red'
              
             }
             
        }
       
        stage('Build') {

            steps {
                    
                    sh "docker build --file Dockerfile-build --tag docker_app_build_image:latest ."
                    sh "docker images "
            }     

        }
        
        stage('CopyBuildToVolumes') {
            agent {
                docker {
                    image'docker_app_build_image:latest'
                    args '-v in-vol:/build  -v out-vol:/output  --user root'
                    reuseNode true
                    }
                }
            steps {
                sh 'rm -rf /build/*'
                sh 'rm -rf /output/*'
                sh 'cp -r /app/simple-golang-app-with-tests/*  /build/'
                sh 'rm /build/simple-golang-app-with-tests' 
                sh 'cp -r  /app/simple-golang-app-with-tests/*.go /output/'   
                sh 'cp -r  /app/simple-golang-app-with-tests/go.* /output/' 
                sh 'cp -r  /app/simple-golang-app-with-tests/simple-golang-app-with-tests /output/' 
            }
           
        }
        stage('BuildTest') {
        
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
                    reuseNode true
                    }
                }
            steps {
                sh 'cd /output && go test ' 
     
            }
        }
        
        stage('Deploy') {
             agent {
                docker {
                    image'docker_app_build_test:latest'
                    args '-v in-vol:/build  -v out-vol:/output  --user root'
                    reuseNode true
                    }
                }
            steps {
                sh 'ls /output'
                sh 'cd /output && ./simple-golang-app-with-tests ' 
            }
        }
         stage('Prepublish') {
             agent {
                docker {
                    image'docker_app_build_test:latest'
                    args '-v in-vol:/build  -v out-vol:/output  --user root'
                    reuseNode true
                    }
                }
            steps {
                sh 'rm -rf publish_app'
                sh 'rm -rf checksum.txt'
                sh 'mkdir  publish_app'
                sh 'rm -f simple_go_app*.tar.gz'
                sh 'cp /output/*.go ./publish_app/' 
                sh 'cp /output/go.* ./publish_app/'
                sh 'cp ./Instruction.md ./publish_app'     
            }
        }
         stage('Publish unofficial version') {
             when{
                 environment name: 'RELEASE', value: 'false'
             }
            steps {

                sh 'tar -zcvf simple_go_app_${GIT_COMMIT_REV}.tar.gz ./publish_app'
                sh 'cat simple_go_app_${GIT_COMMIT_REV}.tar.gz | sha512sum > checksum.txt'
                archiveArtifacts artifacts: 'simple_go_app*.tar.gz', fingerprint: true   
                archiveArtifacts artifacts: 'checksum.txt', fingerprint: true      
            }
        }
         stage('Publish release version') {
             when{
                 environment name: 'RELEASE', value: 'true'
             }
            steps {

                sh 'tar -zcvf simple_go_app_v_${VERSION}.tar.gz ./publish_app'
                sh 'cat simple_go_app_${VERSION}.tar.gz | sha512sum > checksum.txt'
                archiveArtifacts artifacts: 'simple_go_app*.tar.gz', fingerprint: true   
                archiveArtifacts artifacts: 'checksum.txt', fingerprint: true 
            }
        }
        stage('Clean-up') {
            steps {
                sh 'docker stop fluentd'
                script {
                    docker.image('docker_app_build_test').withRun('--user root') { c->
                    sh 'rm -rf containers*.log'
                    sh 'cat logs/test.log.* > containers_${GIT_COMMIT_REV}.log'
                    sh 'rm -rf logs'
                    }
                }    
                archiveArtifacts artifacts: ' containers*.log', fingerprint: true        
            }
        }
    }
    post{
        failure{
          
                sh 'docker stop fluentd'
            
        }
        always{
           
                sh 'docker rmi docker_app_build_test'   
                sh 'docker rmi docker_app_build_image'
            
        }
    }
}
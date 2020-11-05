pipeline {
    agent any

    stages {
        stage('unit test') {
            steps {
                sh label: '', script: '''cd helloworld
                go mod download
                go test -v 2>&1 | go-junit-report > report.xml'''
            }
            post {
                always {
                    junit 'helloworld/report.xml'
                }
            }
        }
        stage('build and start application') {
            steps {
                sh label: '', script: 'docker-compose up -d --build --force-recreate'
            }
        }
        stage('run api test') {
            steps {
                sh label: '', script: '''cd test/api
                robot greeting.robot'''
            }
            post {
                always {
                    robot outputPath: 'test/api', unstableThreshold: 100.0
                    sh label: '', script: 'docker-compose down'
                }
            }
        }
    }
}
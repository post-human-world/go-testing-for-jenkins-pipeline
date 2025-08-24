pipeline {
    agent {
        docker {
            image 'golang:1.25-trixie'
            reuseNode false
            // run as root
            args '-u 0'
        }
    }
    options {
        // Timeout counter starts AFTER agent is allocated
        timeout(time: 600, unit: 'SECONDS')
    }
    environment {
        // Test project
        Github_Project = 'https://github.com/post-human-world/go-testings-for-jenkins-pipeline.git'
    }
    stages {
        stage('Git') {
            steps {
                git Github_Project
            }
        }
        stage('Test') {
            // Unit testing and Integration testing
            steps {
                sh '''
                    #!/bin/bash
                    mkdir coverage

                    # coverage profiling
                    go build -cover -o app ./main.go
                    chmod a+x ./app
                    GOCOVERDIR=coverage ./app

                    # coverage
                    go test -v -coverprofile coverage/cover.out ./...
                    go tool cover -html coverage/cover.out -o coverage/cover.html
                '''

                script {
                    recordCoverage(
                        id: 'GoCoverage', name: 'Golang Coverage',
                        // sourceCodeRetention: 'EVERY_BUILD',
                        qualityGates: [
                            [threshold: 60.0, metric: 'LINE', baseline: 'PROJECT', unstable: true],
                            [threshold: 60.0, metric: 'BRANCH', baseline: 'PROJECT', unstable: true]
                        ],
                        ignoreParsingErrors: true,
                        sourceCodeEncoding: 'UTF-8',
                        tools: [[parser: 'GO_COV', pattern: '**/coverage/']]
                    )
                }
            }
        }
    }
    post {
        cleanup {
            // Clear data
            sh '''
                #!/bin/bash
                ls -l
                rm -rf ./*
            '''
        }
    }
}

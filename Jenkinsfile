pipeline {
    agent any

    environment {
        REGISTRY = "ghcr.io"
        IMAGE_NAME = "ase-cx/url-shortener-backend"
        TAG = sh(script: "git describe --tags", returnStdout: true).trim()
    }

    stages {
        stage('Build docker image') {
            steps {
                script {
                    docker.build("${env.REGISTRY}/${env.IMAGE_NAME}:${env.TAG}")
                }
            }
        }

        stage('Push docker image') {
            steps {
                script {
                    docker.withRegistry("https://${env.REGISTRY}", 'github-credentials') {
                        docker.image("${env.REGISTRY}/${env.IMAGE_NAME}:${env.TAG}").push()
                    }
                }
            }
        }
    }
}
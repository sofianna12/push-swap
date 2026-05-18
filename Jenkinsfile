pipeline {
    agent any

    stages {

        stage('Clone') {
            steps {
                git branch: 'main',
                    url: 'https://github.com/sofianna12/push-swap'
            }
        }

        stage('Build Image') {
            steps {
                sh 'docker build -t asofia32/push-swap:latest .'
            }
        }

        stage('Push to Docker Hub') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: 'dockerhub',
                    usernameVariable: 'USER',
                    passwordVariable: 'PASS'
                )]) {
                    sh 'echo $PASS | docker login -u $USER --password-stdin'
                    sh 'docker push asofia32/push-swap:latest'
                }
            }
        }

    }

    post {
        success { echo 'Pipeline ολοκληρώθηκε επιτυχώς!' }
        failure { echo 'Κάτι πήγε στραβά.' }
    }
}
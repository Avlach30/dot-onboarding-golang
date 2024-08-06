def secrets = [
  [
    path: 'kv/codespace-x/prod',
    engineVersion: 2,
    secretValues: [],
  ],
]
def configuration = [
    vaultUrl: 'https://vault.ubed.dev',
    vaultCredentialId: 'vaultapprole',
    engineVersion: 2,
]

pipeline {
    agent any
    environment{
        DIGITALOCEAN_REGISTRY_CREDS = credentials('DigitaloceanRegistry')
        GCS_ACCOUNT = credentials('GCS_ACCOUNT')
        SVR_JENKINS_PASS = credentials('SvrExavPass')
    }
    stages {
         stage('Clone Repo Dev') {
             when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/dev' }
                }
            }
            steps {
                checkout scm
                sh '''#!/bin/bash
                addgroup jenkins docker
                docker ps
                rm -rf assets
                mkdir assets
                chmod 760 assets
                cp $GCS_ACCOUNT ./assets/gcs_account.json
                '''
            }
        }
        stage('Clone Repo Master') {
             when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/master' }
                }
            }
            steps {
                checkout scm
                sh '''#!/bin/bash
                addgroup jenkins docker
                docker ps
                rm -rf assets
                mkdir assets
                chmod 760 assets
                '''
            }
        }
        stage('Download ENV Dev') {
            when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/dev' }
                }
            }
            steps {
                withVault([configuration: configuration, vaultSecrets: secrets]) {
                    sh '''
                    docker exec vault sh -c 'export VAULT_ADDR=http://127.0.0.1:8200;rm -rf env.json;vault kv get -format=json kv/sirqu-be/dev > env.json;exit'
                    rm -rf .env
                    docker cp vault:env.json env.json
                    cat env.json | jq -r '.data.data | to_entries[] | join("=")' > .env
                    '''
                }
            }
        }
        stage('Download ENV Prod') {
            when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/master' }
                }
            }
            steps {
                withVault([configuration: configuration, vaultSecrets: secrets]) {
                    sh '''
                    docker exec vault sh -c 'export VAULT_ADDR=http://127.0.0.1:8200;rm -rf env.json;vault kv get -format=json kv/codespace-x/prod > env.json;exit'
                    rm -rf .env
                    docker cp vault:env.json env.json
                    cat env.json | jq -r '.data.data | to_entries[] | join("=")' > .env
                    '''
                }
            }
        }
        stage('Build Image Dev') {
            when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/dev' }
                }
            }
            steps {
		         sh '''#!/bin/bash
                 docker build -t sirqu-be:1 .
                 '''
            }
        }
         stage('Build Image Prod') {
            when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/master' }
                }
            }
            steps {
		         sh '''#!/bin/bash
                 docker build -t codespace-x:2 .
                 '''
            }
        }
        stage('DOCR Login') {
            steps {
                sh 'echo $DIGITALOCEAN_REGISTRY_CREDS_PSW | docker login registry.digitalocean.com -u $DIGITALOCEAN_REGISTRY_CREDS_USR --password-stdin'
            }
         }
        stage('DOCR Push Dev') {
            when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/dev' }
                }
            }
            steps {
                sh 'docker tag sirqu-be:1 registry.digitalocean.com/sirqu-container-registry/sirqu-be:1'
                sh 'docker push registry.digitalocean.com/sirqu-container-registry/sirqu-be:1'
            }
         }
         stage('DOCR Push Prod') {
            when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/master' }
                }
            }
            steps {
                sh 'docker tag codespace-x:2 registry.digitalocean.com/sirqu-container-registry/codespace-x:2'
                sh 'docker push registry.digitalocean.com/sirqu-container-registry/codespace-x:2'
            }
         }
        stage('Deploy Dev') {
            when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/dev' }
                }
            }
            steps {
                build job: "Sirqu-Deploy-Dev", wait: true
            }
        }
        stage('Deploy Prod') {
            when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/master' }
                }
            }
            steps {
                sh '''
                ssh admin@194.233.70.62 'bash /var/www/codespace-x/codespace-x-deploy.sh'
                ''', wait: true
            }
        }
        stage('Send Discord Notif Dev') {
            when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/dev' }
                }
            }
            environment {
                DISCORD_WEBHOOK_URL = credentials('webhook_discord')
            }
            steps {
                discordSend description: "New SIRQU BE DEV pipeline triggered for $env.GIT_BRANCH", footer: 'SIRQU BE DEV Pipeline result', link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: env.DISCORD_WEBHOOK_URL
            }
        }
        stage('Send Discord Notif Prod') {
            when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/master' }
                }
            }
            environment {
                DISCORD_WEBHOOK_URL = credentials('webhook_discord')
            }
            steps {
                discordSend description: "New CODESPACE X PROD pipeline triggered for $env.GIT_BRANCH", footer: 'CODESPACE X PROD Pipeline result', link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: env.DISCORD_WEBHOOK_URL
            }
        }
   }
    post {
		always {
			sh 'docker logout'
		}
	 }
    }
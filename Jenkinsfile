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
        SVR_JENKINS_PASS = credentials('SvrExavPass')
    }
    stages {
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
                '''
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
        stage('Deploy Prod') {
            when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/master' }
                }
            }
            steps {
                build job: "codespace-x-deploy", wait: true
            }
        }
        stage('Send Discord Notif Prod') {
            when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/master' }
                }
            }
            environment {
                DISCORD_WEBHOOK_URL = credentials('webhook_discord_codespace')
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
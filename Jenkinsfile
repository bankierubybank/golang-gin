// Declarative //
pipeline {
    // Use Linux Docker host as default agent for this pipeline
    agent {
        label 'slave-linux'
    }
    environment {
        // SCM: Git URL
        APP_REPOSITORY = 'https://github.com/bankierubybank/golang-gin.git'
        BRANCH_NAME = 'main'

        // Sonatype Nexus RM Docker Registry Configuration
        DOCKER_REPOSITORY_HOST = credentials('DOCKER_HOSTED_NEXUS_HOST')
        DEV_DOCKER_REPOSITORY_HOST = credentials('DOCKER_DEV_NEXUS_HOST')
        DOCKER_USER = credentials('NEXUS_JENKINS_USER')
        DOCKER_PASS = credentials('NEXUS_JENKINS_PASS')
        IMAGE_NAME = 'golang-gin'
        NEXUS_AUTH = credentials('NEXUS_AUTH')
    }
    stages {
        stage('CI') {
            parallel {
                stage('BUILD') {
                    stages {
                        stage('SONATYPE: SCA') {
                            steps {
                                // Git clone
                                sh 'git clone --branch ${BRANCH_NAME} ${APP_REPOSITORY} ${IMAGE_NAME}'
                                dir(env.IMAGE_NAME) {
                                    // Scan dependencies using NexusIQ, required Nexus Platform plugin on Jenkins
                                    nexusPolicyEvaluation advancedProperties: '',
                                        enableDebugLogging: false,
                                        failBuildOnNetworkError: false,
                                        iqApplication: selectedApplication('golang-gin'),
                                        iqInstanceId: 'nexusiq.devops.demo',
                                        iqScanPatterns: [[scanPattern: "**/go.sum"]],
                                        iqStage: 'build',
                                        jobCredentialsId: ''
                                }
                            }
                        }
                        stage('JENKINS: BUILD IMAGE') {
                            steps {
                                // Build image from source code with secrets injection for custom PyPI host, required Buildkit to be enabled on built host
                                sh "docker image build --no-cache --progress=plain --tag ${IMAGE_NAME}:$BUILD_NUMBER ${IMAGE_NAME}"
                            }
                        }
                        stage('AQUA: IMAGE SCAN') {
                            steps {
                                sh 'docker login -u ${DOCKER_USER} -p ${DOCKER_PASS} aqua.nexus.devops.demo'
                                // Scan image using Aqua, required Aqua Security Scanner plugin on Jenkins
                                aqua containerRuntime: 'docker',
                                    customFlags: '',
                                    hideBase: false,
                                    hostedImage: '',
                                    localImage: "golang-gin:$BUILD_NUMBER",
                                    locationType: 'local',
                                    notCompliesCmd: '',
                                    onDisallowed: 'fail',
                                    policies: '',
                                    register: false,
                                    registry: '',
                                    scannerPath: '',
                                    showNegligible: true,
                                    tarFilePath: ''
                            }
                        }
                    }
                }
                stage('SAST') {
                    stages {
                        stage('SAST') {
                            steps {
                                // Todo
                                echo 'SAST'
                            }
                        }
                    }
                }
            }
        }
        stage('JENKINS: PUSH IMAGE') {
            //agent {label 'slave-linux'}
            steps {
                // Tag built image for pushing
                sh "docker image tag ${IMAGE_NAME}:$BUILD_NUMBER ${DEV_DOCKER_REPOSITORY_HOST}/${IMAGE_NAME}:$BUILD_NUMBER"
                // Login to Sonatype Nexus
                sh 'docker login -u ${DOCKER_USER} -p ${DOCKER_PASS} ${DEV_DOCKER_REPOSITORY_HOST}'
                // Push image to Sonatype Nexus
                sh "docker image push ${DEV_DOCKER_REPOSITORY_HOST}/${IMAGE_NAME}:$BUILD_NUMBER"
            }
        }
    }
    post { 
        // Always do below tasks when pipeline ends
        always {
            // Remove workspace
            cleanWs()
            // Remove built application images with any tags
            // for i in $(docker image ls | grep ${IMAGE_NAME}:$BUILD_NUMBER | awk '{print $1":"$2}'); do docker rmi $i; done
            sh "for i in \$(docker image ls | grep ${IMAGE_NAME} | gerp $BUILD_NUMBER | awk '{print \$1\":\"\$2}'); do docker rmi \$i; done"
            // Remove usused data
            sh 'docker system prune --force'
        }
    }
}
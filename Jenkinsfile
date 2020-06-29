
node ('master'){
    def dockerFile = 'api.Dockerfile'

    def appName = "mitra-integrasi"
    def nexusPluginsRepoUrl = 'https://nexus.tapera.go.id/repository/maven-central/'
    def imageTag = 'latest'

    def namespace = 'mitra-integrasi'
    def memLimit = '512Mi'
    def cpuLimit = '500m'
    def imagePullSecret = 'nexus-dev-repo'
    def serviceTypeGKE = 'LoadBalancer'
    def serviceTypeALI = 'NodePort'
    def gateway = 'istio-gateway'
    def host = 'hello.sotech.info'
    def gkeKubeConfig = '/jenkinsdev01/jenkins-home/kubeconfig/config-gke-dev'
    def aliKubeConfig = '/jenkinsdev01/jenkins-home/kubeconfig/config-ali-dev'
    def jmeterTestFileGCP = 'petclinic-gke.jmx'
    def jmeterTestFileALI = 'petclinic-ali.jmx'
    def jmeterNumberThreads = '20'
    def jmeterRampUp = '3'
    def jmeterLoopCount = '10'
    def jmeterErrorRateThresholdPercent = '99'
    def jmeterGitRepo = 'https://bitbucket.tapera.go.id/scm/sam/jmeter.git'
    def katalonGitRepo = 'https://bitbucket.tapera.go.id/scm/sam/katalon.git'
    def katalonBranch = 'master'
    def katalonProjectName = 'TestPlease.prj'
    def katalonTestSuiteName = 'TestSuiteTapera'
    def jiraIssueKey
    def jiraUrl = 'https://jira.tapera.go.id'
    def nexusUsername = 'nexusUsername'
    def nexusPassword = 'nexusPassword'
    def branch = 'master'

    def nexusGoCentral = 'nexus.tapera.go.id/repository/go-central'

    def nexusDockerDevRepoGCP = '10.172.24.50:8082'
    def nexusDockerDevRepoALI = '10.172.24.50:8082'

    // ganti git url tapera
    def gitUrl = 'https://bitbucket.tapera.go.id/scm/int/integrasi.git'

    environment {
        GO111MODULE = 'on'
    }

    stage('Checkout'){
        echo 'Checking out SCM'
        checkout scm: [$class: 'GitSCM', userRemoteConfigs: [[credentialsId: 'ci-cd', url: "${gitUrl}"]], branches: [[name: "${branch}"]]]
    }

    stage('Unit Test') {
        echo 'unit test'
        //withCredentials([usernamePassword(credentialsId: 'ci-cd', passwordVariable: nexusPassword, usernameVariable: nexusUsername)]) {
            sh """
                go version
                #export GOPROXY=https://${nexusUsername}:${nexusPassword}@${nexusGoCentral}
                cd api
                rm -rf report.xml
                go mod tidy

                go get golang.org/x/lint/golint
                golint ./controller/...
                
                #go get google.golang.org/grpc
                #go get github.com/golang/protobuf/{proto,protoc-gen-go}
                
                go mod download
                go get github.com/jstemmer/go-junit-report
                go clean -testcache
            """
            
            try {
                sh returnStdout: true, script: '''
                    cd api
                    CGO_ENABLED=0 go test ./controller/... -cover -v 2>&1 | go-junit-report -set-exit-code > ./report.xml; echo $?
                '''
            }
            finally{
                if (fileExists('./report.xml')) { 
                    junit '**/api/report.xml'
                }
            }     
        //}
    }

    stage("SonarQube Analysis"){
        //scan
		withSonarQubeEnv(credentialsId: 'sonarqube-token', installationName: 'sonarqube') {
            // kalau tidak bisa pakai bat
			sh """
                sonar-scanner -Dsonar.projectKey=${appName} -Dsonar.sources=api/controller -Dsonar.qualitygate.wait=true
            """
		}
		//quality gate
		timeout(time: 1, unit: 'HOURS') {
              def qg = waitForQualityGate()
              if (qg.status != 'OK') {
                  error "Pipeline aborted due to quality gate failure: ${qg.status}"
              }
        }
    }

    stage('Build App'){
        echo 'build app'
        //withCredentials([usernamePassword(credentialsId: 'ci-cd', passwordVariable: nexusPassword, usernameVariable: nexusUsername)]) {
            sh """
                #export GOPROXY=https://${nexusUsername}:${nexusPassword}@${nexusGoCentral}
                cd api
                rm -rf bin 
                mkdir -p bin
                CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin
            """
        //}
    }

    stage("Security Check"){
        echo 'security check'
    }

    stage('Build Image') {
        echo 'build image'
        sh """ 
            docker build -t ${appName}:${imageTag}  -f ${dockerFile} --build-arg APP_NAME=${appName} --build-arg APP_VER=${imageTag}  .
        """
    }

    stage('Publish to GKE'){
        echo 'publish to GKE'
        def nexusUrl = nexusDockerDevRepoGCP.replace("http://","")
        sh """
            docker login -u=${nexusUsername} -p=${nexusPassword} ${nexusDockerDevRepoGCP}
            docker push ${nexusUrl}/${appName}:${imageTag} 
        """
    }

    stage('Publish to ALI'){
        echo 'publish to li'
        def nexusUrl = nexusDockerDevRepoGCP.replace("http://","")
        sh """
            docker login -u=${nexusUsername} -p=${nexusPassword} ${nexusDockerDevRepoGCP}
            docker push ${nexusUrl}/${appName}:${imageTag} 
        """
    }
        
    stage('Deploy to GKE'){
        echo 'deploy to GKE'
        
        sh """
		cat kubernetes/gcp-deployment-template.yaml | sed 's/{APP_NAME}/${appName}/g'  \
		| sed 's/{NEXUS_REPO}/${nexusDockerDevRepoGCP}/g' | sed 's/{IMG_TAG}/${imageTag}/g' \
		| sed 's/{MEMORY_LIMIT}/${memLimit}/g' | sed 's/{CPU_LIMIT}/${cpuLimit}/g' \
		| sed 's/{IMG_PULL_SECRET}/${imagePullSecret}/g' |  sed 's/{SERVICE_TYPE}/${serviceTypeGKE}/g' \
		| kubectl --kubeconfig='${gkeKubeConfig}' apply -n '${namespace}' -f -
		
		kubectl --kubeconfig='${gkeKubeConfig}' rollout status deployment/'${appName}' -n '${namespace}'
		"""
    }

    stage('Deploy to ALI'){
        echo 'deploy to ALI'

        /*
        sh """
		cat kubernetes/ali-deployment-template.yaml | sed 's/{APP_NAME}/${appName}/g'  \
		| sed 's/{NEXUS_REPO}/${nexusDockerDevRepoGCP}/g' | sed 's/{IMG_TAG}/${imageTag}/g' \
		| sed 's/{MEMORY_LIMIT}/${memLimit}/g' | sed 's/{CPU_LIMIT}/${cpuLimit}/g' \
		| sed 's/{IMG_PULL_SECRET}/${imagePullSecret}/g' |  sed 's/{SERVICE_TYPE}/${serviceTypeALI}/g' \
		| kubectl --kubeconfig='${aliKubeConfig}' apply -n '${namespace}' -f -
		
		cat kubernetes/ali-virtual-service-template.yaml | sed 's/{APP_NAME}/${appName}/g'  \
		| sed 's/{GATEWAY}/${gateway}/g' | sed 's/{HOST}/${host}/g' \
		| sed 's/{NAMESPACE}/${namespace}/g' \
		| kubectl --kubeconfig='${aliKubeConfig}' apply -f -
		
		kubectl --kubeconfig='${aliKubeConfig}' rollout status deployment/'${appName}' -n '${namespace}'
		"""
        */
    }

    stage("Delete Image"){
        echo "delete image"
        sh """
            docker rmi -f ${appName}:${imageTag}
        """       
    }
    
    stage ("Regression Test") {
        /*
        echo "Build"
		node ('kre-centos') {
			sleep 60
			cleanWs deleteDirs: true
			//git credentialsId: 'portal', branch: 'development', url: 'https://bitbucket.tapera.go.id/scm/PRT/portal-peserta-katalon-script.git'
            checkout scm: [$class: 'GitSCM', userRemoteConfigs: [[credentialsId: 'ci-cd', url: "${gitUrlKatalon}"]], branches: [[name: "${branch}"]]]
			
			withCredentials([string(credentialsId: 'katalon-api-key', variable: 'secret')]) {
				echo "workspace : ${workspace}"
				sh """
				pwd
				ls
				
				/katalon01/katalon-studio-engine/katalonc -noSplash -runMode=console \
				-projectPath='${workspace}/${katalonProjectName}' -retry=0 \
				-testSuitePath='Test Suites/${katalonTestSuiteName}' -executionProfile='default' \
				-browserType="Chrome (headless)" -apiKey='${secret}'
				"""
			}
		}*/
	}
        
    stage ("Jira Update Status") {
        echo "jira update status"
    }
}


#!groovy

@Library('jenkinslib@test') _

// func from shareibrary
def mytools = new org.devops.tools()

// env 环境变量信息 GROUP_NAME PROJECT_NAME VISIBILITY DESC
mytools.PrintMessage("组：${GROUP_NAME}        权限：${VISIBILITY} \n\r 项目：${PROJECT_NAME}        描述：${DESC}", "skyblue")


// 此次构建基础信息输出
currentBuild.description = "组：${GROUP_NAME}        权限：${VISIBILITY} \n\r项目：${PROJECT_NAME}"


// 启用 podTemplate
def label = "slave-${UUID.randomUUID().toString()}"

podTemplate(
    label: label,
    containers: [
        containerTemplate(name: 'main', image: 'harbor.evescn.com/devops/gitlab-project:v1.0', command: 'cat', ttyEnabled: true),
    ], 
    serviceAccount: 'devops',
) {
    node(label) { 
        try {
            stage('创建 GitLab 项目') {
                container('main') {
                    mytools.PrintMessage("创建 GitLab 项目", "green")
                    sh '''
                    cd /app/
                    /app/main --groupname ${GROUP_NAME} --projectname ${PROJECT_NAME} --visibility ${VISIBILITY} --desc "${DESC}"
                    '''
                }
            }

        } catch (Exception e) {
            println(e)
        } finally {
            
        }
    }
}

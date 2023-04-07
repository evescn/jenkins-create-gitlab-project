package org.devops

//格式化输出
def PrintMessage(value, color='skyblue'){
    colors = ['red'         : "\033[1;31m ${value} \033[0m",
              'skyblue'     : "\033[1;36m ${value} \033[0m",
              'blue'        : "\033[1;34m ######## ${value} ######## \033[0m",
              'green'       : "\033[1;32m ================> ${value} <================ \033[0m" ]
              
    ansiColor('xterm') {
        println(colors[color])
    }
}

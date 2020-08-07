#!/bin/bash

CURRENT_PATH=$(cd `dirname $0` && pwd)

BIN_FILE=BIN_FILE_CONST

ENV=ENV_CONST

BIN_PATH=${CURRENT_PATH}/bin/$BIN_FILE

SUPERVISORCTL_CMD="supervisorctl -c /usr/local/etc/supervisord.ini"

if [ ! -d "/data/logs/app/topics" ]; then
  mkdir -p /data/logs/app/topics
fi

function create_supervisor_file() {
cat > /data/etc/supervisord/${BIN_FILE}.ini <<EOF
[program:${BIN_FILE}]
directory=${CURRENT_PATH}
command=${BIN_PATH} --env=${ENV}
autostart=true
autorestart=true
startsecs=1
startretries=3
user=work
redirect_stderr=true
stdout_logfile=/data/logs/app/${BIN_FILE}/console.out
EOF
}

 #如果当前qtt 没有权限 使用sudo权限
function stop() {
	sudo supervisorctl stop ${BIN_FILE}

	if [ $? -ne 0 ]; then
		echo "stop ${BIN_FILE} fail"
		exit 1
	fi
}

function reload() {
    sudo supervisorctl status ${BIN_FILE} |grep -q RUNNING
    if [ $? -ne 0 ]; then
        echo "${BIN_FILE} is not runing"
        exit 1
    fi
    pid=$(sudo supervisorctl status ${BIN_FILE}|awk '{print $4}'|awk -F, '{print $1}')
	create_supervisor_file && (sudo supervisorctl update) && kill -USR2 $pid
    if [ $? -ne 0 ]; then
        echo "reload ${BIN_FILE} fail"
        exit 1
    fi

}

function startOrReload() {
    sudo supervisorctl status ${BIN_FILE} |grep -q RUNNING

    #进程没有运行则start，运行则reload
    if [ $? -ne 0 ]; then
        start
    else
        reload
    fi
}


function start() {
	sudo supervisorctl status ${BIN_FILE} |grep -q RUNNING

	if [ $? -eq 0 ]; then
		echo "${BIN_FILE} is already runing"
		exit 1
	fi

	create_supervisor_file && (sudo supervisorctl update) && (sudo supervisorctl start ${BIN_FILE})
	if [ $? -ne 0 ]; then
		echo "start ${BIN_FILE} fail"
		exit 1
	fi
}

if [ $# -lt 1 ]; then
	echo "usage: $0 [start|stop|reload|restart|startOrReload]"
	exit 1
else
	if [ "$1" == 'stop' ] ; then
		stop
		echo "${BIN_FILE} stop"
	elif [ "$1" == 'start' ] ; then
		start
		echo "${BIN_FILE} start"
	elif [ "$1" == 'reload' ] ; then
        reload
		echo "${BIN_FILE} reload"
	elif [ "$1" == 'startOrReload' ] ; then
        startOrReload
		echo "${BIN_FILE} startOrReload"
	elif [ "$1" == 'restart' ] ; then
		stop
		start
		echo "${BIN_FILE} restart"
	else
		echo "usage: $0 [start|stop|reload|restart|startOrReload]"
		exit 1
	fi
fi

exit 0

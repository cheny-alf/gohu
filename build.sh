#!/bin/bash

export PROJECT_NAME=$1

export THREAD=$2

docker_names=('oauth-api' 'oauth-rpc-token-enhancer' 'oauth-rpc-token-store' \
'user-api' 'user-rpc-crud' 'user-rpc-info' 'user-rpc-vip' 'notification-api' \
'notification-rpc-crud' 'notification-rpc-info' 'mq-asynq-scheduler' 'mq-asynq-processor' \
'mq-nsq-consumer')

function docker_build() {
  if [ "$1" -ef "" ]; then
    return 0
  fi

  array=$(echo "$1" | tr '-' '\n')
  path='./app/service'
  for var in $array
  do
    path="${path}""/""${var}"
  done
  docker build -t "$PROJECT_NAME""_""$1" "${path}"
  return 1
}

[ -e ./build/fd1 ] || mkfifo ./build/fd1
exec 3<>./build/fd1
rm -rf ./build/df1

for ((i=1;i<=THREAD;i++))
do
  echo >&5
done

cd /www/site/"$PROJECT_NAME" || exit

remain_build=${#docker_names[@]}

echo "start building images, remain: ""${remain_build}"

for docker_name in ${docker_names[*]}
do
  read -r -u5
{
  docker_build "${docker_name}"
  remain_build=${remain_build}-1
  echo "build ""${docker_name}"" complete, remain: ""${remain_build}"
  echo >&5
} &
done

wait
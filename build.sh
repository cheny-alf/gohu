#!/bin/bash

docker_names=('oauth-api' 'oauth-crud-rpc')

function docker_build() {
  if [ "$1" -ef "" ]; then
    return 0
  fi

  array=$(echo "$1" | tr '-' '\n')
  path='./service'
  for var in $array
  do
    path="${path}""/""${var}"
  done

  docker build -t "$PROJECT_NAME""_""$1" path
  return 1
}

cd /www/site/"$PROJECT_NAME" || exit

echo "docker_images: ""${#docker_names[@]}"

for docker_name in ${docker_names[*]}
do
{
  docker_build "${docker_name}"
} &
done

wait

echo "complete"
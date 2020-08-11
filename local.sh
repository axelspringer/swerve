#!/bin/bash

NAME_NGINX=swerve-nginx
NAME_SWERVE=swerve-app
NAME_DYNAMODB=dynamodb
NAME_PEBBLE=pebble
VOLUME_DYNAMODB=swerve-dynamodb
NAME_AWS_CLI=swerve-aws-cli
export TABLE_USERS=Users


function docker_stop() {
    docker ps | grep -q "$1" &&  docker stop "$1" && echo -e "\tstopped"
}

function start_pebble() {
  docker_stop $NAME_PEBBLE
  docker run --rm --network swerve \
    -e "PEBBLE_VA_ALWAYS_VALID=1" \
    -e "PEBBLE_VA_NOSLEEP=1" \
    --name $NAME_PEBBLE \
    -p 14000:14000 \
    -p 15000:15000 \
    -d letsencrypt/pebble
}

function create_volume() {
    docker volume create $VOLUME_DYNAMODB
    echo "docker volume $VOLUME_DYNAMODB created"
}

function wait_for_dynamodb_is_up {
  running=0
  while [ $running -eq 0 ]
  do
    echo "wait for  dynamodb is up"
    if [[ "$(nc -z localhost 8000 &> /dev/null; echo $?)" == "0" ]]; then
      echo "local dynamodb is up and running"
      running=1
    else
      sleep 1
    fi
  done
}

function start_dynamodb {
  if [[ "$(nc -z localhost 8000 &> /dev/null; echo $?)" == "0" ]]; then
    echo "[INFO] dynamodb already running"
    return
  fi
  docker_stop $NAME_DYNAMODB
  docker run --rm --network swerve \
    --mount source=$VOLUME_DYNAMODB,target=/data \
    --name $NAME_DYNAMODB \
    -p 8000:8000 \
    -d dwmkerr/dynamodb -sharedDb -dbPath /data
}


function init_dynamodb {
  echo "start aws cli"
  docker_stop $NAME_AWS_CLI
  docker run -ti --rm --network swerve \
    --mount type=bind,source="$(pwd)"/scripts,target=/tmp/scripts \
    -e AWS_ACCESS_KEY_ID=0 \
    -e AWS_SECRET_ACCESS_KEY=0 \
    -e AWS_DEFAULT_REGION=eu-west-1 \
    -e TABLE_REDIRECTS=Redirects \
    -e TABLE_CERTCACHE=CertCache \
    -e TABLE_USERS=Users \
    garland/aws-cli-docker \
    /bin/sh /tmp/scripts/init_dynamo.sh
}

function start_nginx {
  echo "start nginx"
  docker_stop $NAME_NGINX
  docker run --rm --network swerve \
    -e NGINX_HOST=demo-target \
    -e NGINX_PORT=80 \
    -p 8090:80 \
    --name $NAME_NGINX \
    -d nginx:alpine
}


function build_swerve() {
  echo "docker build -t $NAME_SWERVE ."
  docker build -t $NAME_SWERVE .
}

function run_swerve() {
  echo "start red-swerve"
  docker_stop $NAME_SWERVE
  docker run --rm --network swerve \
   -e SWERVE_HTTP=:8080 \
   -e SWERVE_HTTPS=:8081 \
   -e SWERVE_API=:8082 \
   -e SWERVE_DB_REGION=eu-west-1 \
   -e SWERVE_DB_ENDPOINT=http://dynamodb:8000 \
   -e SWERVE_API_UI_URL=http://not-in-use.local \
   -p 8080:8080 \
   -p 8081:8081 \
   -p 8082:8082 \
   --name $NAME_SWERVE \
   -d $NAME_SWERVE
}

###
if [[ "$1" == "aws-cli" ]]; then
 docker run --rm --network swerve \
   -e AWS_ACCESS_KEY_ID=0 \
   -e AWS_SECRET_ACCESS_KEY=0 \
   -e AWS_DEFAULT_REGION=eu-west-1 \
   -ti garland/aws-cli-docker /bin/sh
   exit 0
fi

if [[ "$1" == "stop" ]]; then
  exitcode=0
  docker_stop $NAME_DYNAMODB || exitcode=1
  docker_stop $NAME_SWERVE || exitcode=1
  docker_stop $NAME_NGINX || exitcode=1
  exit $exitcode
fi

if [[ "$1" == "build" ]]; then
  build_swerve
fi

if [[ "$1" == "dyno" ]]; then
 init_dynamodb
fi


if [[ "$1" == "dep" ]]; then
  start_dynamodb
  start_pebble
fi

if [[ "$1" == "run" ]]; then
  echo "check if docker volume exists"
  test $(docker volume ls -q -f "name=$VOLUME_DYNAMODB") = "$VOLUME_DYNAMODB" || init_dynamodb
  echo "start dynamodb"
  start_dynamodb
  echo "init dynamodb tables"
  init_dynamodb
  echo "wait for running dynamodb"
  wait_for_dynamodb_is_up
  build_swerve
  echo "run docker swerve app"
  run_swerve
fi




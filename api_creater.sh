#!/bin/bash

check_input() {
    if [ -z "$1" ]; then
        echo "Input error"
        exit 1
    fi
}

check_number() {
    regex='[^0-9]+'
    if [[ $1 =~ $regex ]] ;
    then
        echo "input error"
        exit 1
    fi
}

if [ $# -ne 1 ]; then
  echo "Please set api dir"
  exit 1
fi

if [ -d ${1} ]; then
    read -p "Press add api name: " API_NAME
    check_input $API_NAME
    
    read -p "Press add api ip address: " API_IP
    check_input $API_IP
    check_number $API_IP

    read -p "Press add db ip address: " DB_IP
    check_input $DB_IP
    check_number $DB_IP
fi

cp -a base_api/ ${API_NAME}_api/

sed -i -e "s/\${default_name}/${API_NAME}/g" ./${API_NAME}_api/docker-compose.yml
sed -i -e "s/\${default_name}/${API_NAME}/g" ./${API_NAME}_api/postgres/init/create_db.sql
sed -i -e "s/\${default_name}/${API_NAME}/g" ./${API_NAME}_api/setting/dbconfig.yml

sed -i -e "s/\${default_db_ip}/${DB_IP}/g" ./${API_NAME}_api/docker-compose.yml
sed -i -e "s/\${default_db_ip}/${DB_IP}/g" ./${API_NAME}_api/setting/dbconfig.yml

sed -i -e "s/\${default_api_ip}/${API_IP}/g" ./${API_NAME}_api/docker-compose.yml

mv ./${API_NAME}_api $1
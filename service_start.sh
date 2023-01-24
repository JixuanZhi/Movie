docker-compose -f deployment.yml down

docker-compose -f deployment.yml rm -f

while getopts "k" arg
do
   case $arg in
     k)
       exit 0 
       ;;
     ?)
       echo "Unknow args $arg... exit..."
       exit 1
       ;;
   esac
done

ENV_FILE=".env"
rm -f $ENV_FILE

CUR_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
HOST_IP=`ifconfig | grep 'inet ' | grep -Fv 127.0.0.1 | awk '{print $2}'`

function encode_env {
  echo "$1=${!1}" >> $ENV_FILE
}

encode_env "CUR_DIR"
encode_env "HOST_IP"

docker-compose -f deployment.yml up --build -d

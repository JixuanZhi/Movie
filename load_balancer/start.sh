docker build -f ./Dockerfile -t nginx:live . 
HOST_IP=`ifconfig | grep 'inet ' | grep -Fv 127.0.0.1 | awk '{print $2}'`
docker run -it -p 5002:80 -d -v "$(pwd)":/load_balancer nginx:live bash init.sh $HOST_IP 5001

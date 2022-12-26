ID=$(docker ps -qf publish=5001)
if [ -z $ID ]
then
 echo ""
else 
 docker stop $ID
fi

while getopts "b:" arg
do
   case $arg in
     b)
       IS_BUILD=true
       TAG_NAME=$OPTARG
       ;;
     ?)
       echo "Unknow args $arg... exit..."
       exit 1
       ;;
   esac
done

if [ "$IS_BUILD" = true ] ; then
        docker tag go:live go:$TAG_NAME
	docker build -f ./Dockerfile -t go:live .
fi

docker run --rm -p 5001:80 -it -d -v "$(pwd)":/go_source go:live go run server.go
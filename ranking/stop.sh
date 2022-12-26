ID=$(docker ps -qf status=running)
for id in $ID
do
 docker stop $id
done

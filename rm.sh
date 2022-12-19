IDS=$(docker ps -q -a)
for ID in $IDS
do
  docker rm  $ID
done

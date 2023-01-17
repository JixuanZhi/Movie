curl -X PUT "localhost:5003/customer/_bulk?pretty" -H 'Content-Type: application/json' -d'
{ "create": { } }
{ "movie_name": "The Shawshank Redemption","genre":["Drama"],"director": "Frank Darabont", "rating": "5"}
{ "create": { } }
{ "movie_name": "The Godfather","genre":["Drama","Crime"],"director": "Francis Ford Coppola", "rating": "4.9"}
{ "create": { } }
{"movie_name": "The Dark Knight","genre":["Action","Crime","Drama","Thriller"],"director": "Christopher Nolan", "rating": "4.8"}
{ "create": { } }
{ "movie_name": "12 Angry Men","genre":["Drama","Crime"],"director": "Sidney Lumet", "rating": "4.7"}
'


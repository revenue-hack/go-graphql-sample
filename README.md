# go-graphql-sample


## exec
createUser
```
curl -H 'Content-Type: application/json' -X POST -d 'mutation{createUser(userName:"mituba", description: "des", photoURL: "photo", email: "email"){userId, userName, description, photoURL, email}}' http://localhost:8080/graphql 
```

user
```
curl -H 'Content-Type: application/json' -X POST -d '{user(id: "111"){userId, userName}}' http://localhost:8080/graphql
```

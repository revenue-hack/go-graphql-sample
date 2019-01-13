# go-graphql-sample

## install
install docker for mac
https://docs.docker.com/docker-for-mac/install/

go version is more than 1.11

## exec
createUser
```
curl -H 'Content-Type: application/json' -X POST -d 'mutation{createUser(userName:"mituba", description: "des", photoURL: "photo", email: "email"){userId, userName, description, photoURL, email}}' http://localhost:8080/graphql 
```

user(need to change uuid key)
```
curl -H 'Content-Type: application/json' -X POST -d '{user(id: "uuid key"){userId, userName}}' http://localhost:8080/graphql
```

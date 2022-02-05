# API Gateway

1. Copy the contents of `.env-example` to `.env`
2. Run `docker-compose up -d --build` to start the project locally
4. Visit the page: http://localhost:9000

## Database migrations
#### Install golang-migrations
```
brew install golang-migrate
```
from: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

### Run migrations
```
 make migrate
```

## Build
```
 make
```

## Run
```
 make run
```

## API requests 
### Import post
```
curl -X "POST" "localhost:9000/v1/post/import" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{
    "quantity":50
}'
```
### Create post
```
curl -X "POST" "localhost:9000/v1/post" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{
    "userID": 1,
    "title": "Example title",
    "body": "Example body"
}'
```
### Get post by id
```
curl "localhost:9000/v1/post/1" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```
### Update post
```
curl -X "PUT" "localhost:9000/v1/post" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{
    "userID": 1,
    "title": "Example title",
    "body": "Example body"
}'
```
### Delete post by id
```
curl -X "DELETE" "localhost:9000/v1/post/1" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

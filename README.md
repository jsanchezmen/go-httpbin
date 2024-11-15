# GO Httpbin

This is a reppo that will contains an httpbin server the possible endpoints are
- /hello
- /{code}

Default port to listen is 8080 if you want to change it set up the ENV variable `SERVER_PORT`

## /hello

This will return a json response with the name of the server stored on the ENV variable `SERVER_NAME`

expected response:
``` json
{
  "code": 200,
  "message": "Hello from [ server1 ] =D"
}

```
## /{code}

This will return a json response with the code provided

``` json
// /200
{
  "code": 200,
  "message": "Status OK"
}

```

## Create a Tag/Release

```
git tag -l
git tag -a v0.0.1-beta -m "v0.0.1-beta"
git push --tags
# Force push
git push -f --tags
```

## Building Docker image

```
docker build -t sjulian/go-httpbin:local .
docker push sjulian/go-httpbin:local
docker run -d -p 8080:8080 --name httpbin sjulian/go-httpbin:local
```
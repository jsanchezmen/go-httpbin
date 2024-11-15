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

```
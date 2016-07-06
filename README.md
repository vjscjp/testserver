# testserver

A simple Load Test application.

### Default: Index Page

StatusCode    : Expected Status Code in response (if left blank api will return actual response code)

Delay         : Amount of Delay required from server side for getting reponse.

Hits          : No of Hits requierd

### Api: /hits
Request URL:http://localhost:9090/hits

Request Method:POST

Payload
```
{"URL":"","StatusCode":200,"Delay":100,"Hits":1}
``` 

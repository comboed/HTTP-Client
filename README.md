A lightweight client for sending reusable HTTP/HTTPS requests

## Features

**Thread Safe**: Safe to use in multi-threaded applications.

**Connection Pooling**: Efficient management of connections for improved performance.

**Simple Request Building**: Easy to create and customize requests.

## Creating a client
**Arguments:**

Max Connections

Read Buffer Size

TLS Config
```go
var client *Client = createClient(100, 0, nil)
```

## Creating a request
```go
var request *Request = createRequest()
```

## Usage
```go
	// Not SSL
	request.SetURI("http://example.org/")
	var body = client.Do(request)
	fmt.Println(string(body))

	// SSL
	request.SetURI("https://example.org/")
	body = client.Do(request)
	fmt.Println(string(body))
```

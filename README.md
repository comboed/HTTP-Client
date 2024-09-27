A lightweight client for sending reusable HTTP/HTTPS requests

## Features

**Thread Safe**: Safe to use in multi-threaded applications.

**Connection Pooling**: Efficient management of connections for improved performance.

**Simple Request Building**: Easy to create and customize requests.

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

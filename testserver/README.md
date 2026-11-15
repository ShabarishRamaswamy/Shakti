# Local HTTP test server

From the Shakti repo root:

```sh
go run ./testserver
```

Connect to `127.0.0.1:6011` over plain TCP (no TLS). For example, send:

```http
GET / HTTP/1.0

```

Use CRLF (`\r\n`) line endings, including the empty line after the request.
The server uses Go’s standard `net/http` package. With the request above, it
returns the following response (header order may differ), then closes the connection:

```text
HTTP/1.0 200 OK\r\n
Content-Length: 22\r\n
Content-Type: text/html\r\n
Connection: close\r\n
\r\n
<h1>Hello, World!</h1>
```

The body has no trailing newline. Every path returns the same response.
HEAD requests return only the headers. There is no encryption, compression,
chunked encoding, or persistent connection. The server binds only to localhost.

To inspect the response from another terminal:

```sh
curl --http1.0 -i http://127.0.0.1:6011/
```

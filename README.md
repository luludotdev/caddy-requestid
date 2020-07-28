# Caddy RequestID
> Caddy v2 Module that adds a unique ID to response headers.


## Caddyfile Syntax
```
request_id [<matcher>] [<header>] {
  header <text>
}
```

### Example
```
{
  order request_id before respond
}

localhost {
  request_id * x-request-id
  respond * "hello world" 200
}
```

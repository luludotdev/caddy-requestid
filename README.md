# Caddy Request ID
> Caddy v2 Module that sets unique request ID placeholders.

## Usage
### Caddyfile
```
request_id [<length>] {
  [<key> <length>]
  ...
}
```
* **length** - length of ID to generate, defaults to 21
* **key, length** - additional keys to generate independent IDs for

If you wish to use the directive in a top level block, you must explicitly define the order.
```
{
  order request_id before header
}
```

### With JSON Config
```json5
{
  "handler": "request_id",
  "length": 21, // optional
  "additional": { // optional
    "header": 21,
  }
}
```

### Placeholders
The top level request ID will be set in the `{http.request_id}` placeholder. Any additional IDs will be set in the `{http.request_id.<key>}` placeholder.

## Example
The following example Caddyfile a different request ID for response bodies and headers.
```
{
  order request_id before header
}

localhost {
  request_id {
    body 10
    header 21
  }

  header * x-request-id "{http.request_id.header}"
  respond * "{http.request_id.body}" 200
}
```

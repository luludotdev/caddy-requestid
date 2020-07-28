# Caddy RequestID
> Caddy v2 Module that sets a unique request ID placeholder.

## Usage
### With a Caddyfile
Usage with a Caddyfile is fairly straightforward. Simply add the `request_id` directive to a handler block, and the `{http.request_id}` template will be set.

If you wish to use the directive in a top level block, you must explicitly define the order.
```
{
  order request_id before header
}
```

### With JSON Config
In the JSON Config, all you need to do is add the `request_id` hander to your `handle[]` chain. The same template as documented above will be set. Note that you must set this before any handlers that you want to use the template in.
```json
{
  "handler": "request_id"
}
```

### Example
The following example Caddyfile sets the `x-request-id` header for all responses.
```
{
  order request_id before header
}

localhost {
  request_id

  header * x-request-id "{http.request_id}"
  respond * "{http.request_id}" 200
}
```

# Caddy RequestID
> Caddy v2 Module that adds a unique ID to response headers.

## Caddyfile Syntax
The header and template value used can be set using the Caddyfile. The tempate string can use [placeholders](https://caddyserver.com/docs/conventions#placeholders), and an additional placeholder `{uid}` is defined for this module. The header defaults to `x-request-id` and the template string defaults to `{uid}`.

```
request_id [<matcher>] [<header>] {
  header <text>
  template <text>
}
```

### Example
```
# Required to use in top-level blocks
{
  order request_id before handle
}

localhost {
  request_id /api/* {
    header x-ray-id
    template "{system.hostname}-{uid}"
  }

  respond * "hello world" 200
}
```

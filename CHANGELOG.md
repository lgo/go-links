# Changelog

## 0.2
Added Redis compatibility, allowing persistence between application starts.
- Use by setting `GOLINKS_STORE=redis`
- Configure Redis via. `REDIS_URL=redis://h:<PASSWORD>@<HOST>:<PORT>`

## 0.1
Initial release, working but limited functionality.
- Supports ephemeral dictionary backend, losing all short links and metrics on startup
- No way to fetch recorded

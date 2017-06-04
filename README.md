# go-links

A small and simple URL shortener.

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

## Usage

For information on building, see the section at the end.

### Running

```
ADMIN_KEY=<secret> PORT=80 GOLINK_ENV=prod ./golinks
```

- `ADMIN_KEY` is used for authenticating the administrative endpoints and is required.
- `PORT` is optional, defaulting to `8080`.
- `GOLINK_ENV` defaults to `dev`. `prod` runs on `0.0.0.0`, `dev` runs on `127.0.0.1`
- `GOLINK_LOGLEVEL` defaults to `info`, and has the options `debug`, `info`, `warn`.

### Adding short links

```bash
curl localhost/admin/api/links/htn \
  -X POST -H "Content-type: application/json" \
  -H "Golink-Auth: <ADMIN_KEY>" \
  -d '{"url": "https://hackthenorth.com"}'

curl localhost/htn
> <a href="https://hackthenorth.com">See Other</a>.
# A 304 Redirect (temporary redirect)
```

To see what a short link is set to, along with the metrics
```bash
curl localhost/admin/api/links/htn \
  -H "Golink-Auth: <ADMIN_KEY>" \
> {"url":"https://hackthenorth.com","metrics":0}
```

### Building

If you're looking to just get started, there are pre-built binaries under Github Releases.

This library was built and only tested with `go1.8.3`. In order to use it, it also assumes you have the following packages in your `GOPATH`, along with this repo[0].
```
go get "github.com/sirupsen/logrus"
go get "github.com/gorilla/mux"
```

[0]: If the latest dependancy versions don't work, let me know and I'll vendor them.

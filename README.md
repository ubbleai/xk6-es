# xk6-es
Yet another (working) elasticsearch output for k6.
It relies on https://github.com/olivere/elastic/v7 library & uses bulking for K6 output sample ingestion.

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git
- [xk6](https://github.com/grafana/xk6)

1. Build with `xk6`:

```bash
xk6 build --with github.com/ubbleai/xk6-es
```

This will result in a `k6` binary in the current directory.

2. Run with the just build `k6:

```bash
./k6 run -o xk6-es <script.js>
```

## Configuration
| ENV | Default value | Description |
|-----|---------------|-------------|
| `K6_ES_PUSH_INTERVAL`  | "1s"                    | K6 samples flush interval |
| `K6_ES_ADDRESS`        | "http://127.0.0.1:9200" | Elasticsearch endpoint |
| `K6_ES_USERNAME`       | ""                      | Elasticsearch username |
| `K6_ES_PASSWORD`       | ""                      | Elasticsearch password |
| `K6_ES_INDEX`          | ""                      | Index used to store samples |
| `K6_ES_ENABLE_SNIFFER` | "false"                 | Enable Elastic endpoints discovery mecanism |
| `K6_ES_MAX_BULKSIZE`   | "2048"                  | Maximum bulk size before forcing flush |

## Document template

* `name`, `type` = `string`
* `@timestamp` = `int64` _(epoch milli)_
* `Value` = `float64`
* `Tags` = `map[string]string`

### output example:

```
{
  "name": "http_req_waiting",
  "type": "trend",
  "@timestamp": 1654968030083,
  "Tags": {
    "expected_response": "true",
    "function": "foo",
    "group": "::bar",
    "method": "POST",
    "name": "https://foo.bar/foo",
    "proto": "HTTP/2.0",
    "scenario": "default",
    "status": "200",
    "tls_version": "tls1.3",
    "url": "https://foo.bar/foo"
  },
  "Value": 970.557
}
```

### Unsupported features / TODO

* Unauthenticated write
* Embedded index templating pushed by the plugin (auto creation)
* Dashboard examples
* log verbosity / error catching edge cases

Even if this project doesn't really require daily support, we're using it in real life !
_PR & issues are open to anyone._
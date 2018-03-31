# saclogsend

A wrapper around Go's syslog package for transmitting logs to a remote syslog server over UDP.

Used by the Go microservices of [mailsac.com](https://mailsac.com/docs/api).

## Usage

```go
package main

import (
  "flag"

  "github.com/mailsac/saclogsend"
)

var syslogNet = flag.String("syslog", "127.0.0.1:514", "net location of a syslog server")

var slog *saclogsend.Slog

func main() {
  flag.Parse()
  var err error
  slog, err = saclogsend.New("inbound", *syslogNet)
  if err != nil {
    slog.Err("failed connecting to syslog, %s", err.Error())
  } else {
    slog.Info("syslog connected:" + *syslogNet)
  }
}
```



## License

MIT

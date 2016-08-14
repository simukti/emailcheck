## EMAILCHECK
A small and handy Go utility to check whether email address is disposable.

### DETAILS
Disposable email domains is ported from [FGRibreau/mailchecker](https://github.com/FGRibreau/mailchecker/blob/master/list.json).

### INSTALL
`go get -u -v github.com/simukti/emailcheck`

### UPDATE DOMAINS
To update domain list, just simply run :
```
cd $GOPATH/src/github.com/simukti/emailcheck

cd ./generator && go run main.go && cd ../ && go fmt .
```
### USAGE SAMPLE
```go
package main

import (
    "github.com/simukti/emailcheck"
)

func main() {
    isDisposable := emailcheck.IsDisposableEmail("kadalkesit@mailinator.com")
    if isDisposable {
        // detected as disposable email
        // do something here
    }
}
```

### TEST
```
$ go test -v ./...
=== RUN   TestIsDisposable
--- PASS: TestIsDisposable (0.00s)
=== RUN   TestIsRegularEmail
--- PASS: TestIsRegularEmail (0.00s)
PASS
ok  	github.com/simukti/emailcheck	0.009s
testing: warning: no tests to run
PASS
ok  	github.com/simukti/emailcheck/generator	0.008s
```

### LICENSE
This project is released under the MIT licence.
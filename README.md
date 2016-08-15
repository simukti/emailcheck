## EMAILCHECK
A small and handy Golang utility to check whether email address is disposable email.

### DETAILS
Disposable email domain list is ported from [FGRibreau/mailchecker](https://github.com/FGRibreau/mailchecker/blob/master/list.json), with duplication removal and transform all domain name to lowercase.

### INSTALL
`go get -u -v github.com/simukti/emailcheck`

### UPDATE DOMAINS
To update domain list, just simply run :
```
cd $GOPATH/src/github.com/simukti/emailcheck

cd ./generator && go run main.go && cd ../ && go fmt .
```

### TEST
`$ go test -v ./...`

### HOW TO USE
```go
package main

import (
    "github.com/simukti/emailcheck"
)

func main() {
    emailAddress := "kadalkesit@mailinator.com"
    
    // validate email format here :
    // ....
    
    if isDisposable := emailcheck.IsDisposableEmail(emailAddress); isDisposable {
        // detected as disposable email
        // do something here    
    }
}
```

### LICENSE
This project is released under the MIT licence.
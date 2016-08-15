// Copyright (c) 2016 - Sarjono Mukti Aji <me@simukti.net>
// Unless otherwise noted, this source code license is MIT-License

package emailcheck

import "strings"

// IsDisposableEmail will return true if email's domain is in disposableDomains list
func IsDisposableEmail(email string) bool {
	// I assume this is a valid email address from your validation process
	// you can use github.com/asaskevich/govalidator
	// govalidator.IsEmail(emailAddress)
	s := strings.Split(email, "@")
	if len(s) != 2 {
		// dont mark it as disposable if it was not a valid email address
		return false
	}

	domain := strings.ToLower(strings.TrimSpace(s[1]))

	return disposableDomains[domain]
}

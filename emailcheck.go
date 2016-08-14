// Copyright (c) 2016 - Sarjono Mukti Aji <me@simukti.net>
// Unless otherwise noted, this source code license is MIT-License

package emailcheck

import "strings"

func IsDisposableEmail(email string) bool {
	// I assume this is a valid email address from your validation process
	// you can use github.com/asaskevich/govalidator
	// govalidator.IsEmail(emailAddress)
	s := strings.Split(email, "@")
	if len(s) != 2 {
		return false
	}

	domain := strings.ToLower(strings.TrimSpace(s[1]))

	return disposableDomain[domain]
}

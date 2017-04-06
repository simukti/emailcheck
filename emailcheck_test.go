// Copyright (c) 2016 - Sarjono Mukti Aji <me@simukti.net>
// Unless otherwise noted, this source code license is MIT-License

package emailcheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDisposable(t *testing.T) {
	disposableEmails := []string{"prapto@mailinator.com", "bohay@letthemeatspam.com", "kadalkesit@magicbox.ro"}
	for _, email := range disposableEmails {
		isDisposable := IsDisposableEmail(email)
		assert.True(t, isDisposable)
	}
}

func TestIsRegularEmail(t *testing.T) {
	regularEmails := []string{"prapto@google.com", "staff@github.com", "kadalkesit@fb.com", "postmaster@simukti.net"}
	for _, email := range regularEmails {
		isDisposable := IsDisposableEmail(email)
		assert.False(t, isDisposable)
	}
}

func TestWithInvalidEmailAddressWhichShouldBeFalse(t *testing.T) {
	invalidEmails := []string{"prapto", "sodiq", "holil"}
	for _, email := range invalidEmails {
		isDisposable := IsDisposableEmail(email)
		assert.False(t, isDisposable)
	}
}

// Copyright (c) 2016 - Sarjono Mukti Aji <me@simukti.net>
// Unless otherwise noted, this source code license is MIT-License

package emailcheck

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsDisposable(t *testing.T) {
	disposableEmails := []string{"prapto@mailinator.com", "bohay@letthemeatspam.com", "kadalkesit@magicbox.ro"}
	for _, email := range disposableEmails {
		isDisposable := IsDisposableEmail(email)
		assert.True(t, isDisposable)
	}
}

func TestIsRegularEmail(t *testing.T) {
	regularEmails := []string{"prapto@google.com", "staff@github.com", "kadalkesit@fb.com"}
	for _, email := range regularEmails {
		isDisposable := IsDisposableEmail(email)
		assert.False(t, isDisposable)
	}
}

/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package cli

import (
	"os"
	"testing"

	"github.com/akamensky/argparse"
	"github.com/stretchr/testify/assert"
)

func newParser() *argparse.Parser {
	return argparse.NewParser("test", "Test parser")
}

func TestParseList(t *testing.T) {
	t.Run("test short flag", func(t *testing.T) {
		parser := newParser()
		list := ParseList(parser)

		assert.NotNil(t, list)
		assert.False(t, *list)

		os.Args = []string{"sentinel", "-l" /*+ string(sentinel.List)*/}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.True(t, *list)
	})

	t.Run("test long flag", func(t *testing.T) {
		parser := newParser()
		list := ParseList(parser)

		assert.NotNil(t, list)
		assert.False(t, *list)

		os.Args = []string{"sentinel", "--list" /*+ string(sentinel.ListExp)*/}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.True(t, *list)
	})
}

func TestParseDeleteSecret(t *testing.T) {
	t.Run("test short flag", func(t *testing.T) {
		parser := newParser()
		remove := ParseDeleteSecret(parser)

		assert.NotNil(t, remove)
		assert.False(t, *remove)

		os.Args = []string{"sentinel", "-d" /*+ string(sentinel.Remove)*/}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.True(t, *remove)
	})

	t.Run("test long flag", func(t *testing.T) {
		parser := newParser()
		remove := ParseDeleteSecret(parser)

		assert.NotNil(t, remove)
		assert.False(t, *remove)

		os.Args = []string{"sentinel", "--delete" /*+ string(sentinel.RemoveExp)*/}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.True(t, *remove)
	})
}

func TestParseAppendSecret(t *testing.T) {
	t.Run("test short flag", func(t *testing.T) {
		parser := newParser()
		join := ParseAppendSecret(parser)

		assert.NotNil(t, join)
		assert.False(t, *join)

		os.Args = []string{"sentinel", "-a" /*+ string(sentinel.Join)*/}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.True(t, *join)
	})

	t.Run("test long flag", func(t *testing.T) {
		parser := newParser()
		join := ParseAppendSecret(parser)

		assert.NotNil(t, join)
		assert.False(t, *join)

		os.Args = []string{"sentinel", "--append" /*+ string(sentinel.JoinExp)*/}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.True(t, *join)
	})
}

func TestParseNamespaces(t *testing.T) {
	t.Run("test short flag", func(t *testing.T) {
		parser := newParser()
		namespaces := ParseNamespaces(parser)

		assert.NotNil(t, namespaces)
		assert.Equal(t, []string{}, *namespaces)

		os.Args = []string{"sentinel", "-n" /*+ string(sentinel.Namespace)*/, "ns1", "-n", "ns2"}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.Equal(t, []string{"ns1", "ns2"}, *namespaces)
	})

	t.Run("test long flag", func(t *testing.T) {
		parser := newParser()
		namespaces := ParseNamespaces(parser)

		assert.NotNil(t, namespaces)
		assert.Equal(t, []string{}, *namespaces)

		os.Args = []string{"sentinel", "--namespace" /*+ string(sentinel.NamespaceExp)*/, "ns1", "-n", "ns2"}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.Equal(t, []string{"ns1", "ns2"}, *namespaces)
	})
}

func TestParseInputKeys(t *testing.T) {
	t.Run("test short flag", func(t *testing.T) {
		parser := newParser()
		inputKeys := ParseInputKeys(parser)

		assert.NotNil(t, inputKeys)
		assert.Equal(t, "", *inputKeys)

		os.Args = []string{"sentinel", "-i" /*+ string(sentinel.Keys)*/, "AGE-SECRET-KEY-1RZU...\nage1...\na6...ceec"}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.Equal(t, "AGE-SECRET-KEY-1RZU...\nage1...\na6...ceec", *inputKeys)
	})

	t.Run("test long flag", func(t *testing.T) {
		parser := newParser()
		inputKeys := ParseInputKeys(parser)

		assert.NotNil(t, inputKeys)
		assert.Equal(t, "", *inputKeys)

		os.Args = []string{"sentinel", "--input-keys" /*+ string(sentinel.KeysExp)*/, "AGE-SECRET-KEY-1RZU...\nage1...\na6...ceec"}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.Equal(t, "AGE-SECRET-KEY-1RZU...\nage1...\na6...ceec", *inputKeys)
	})
}

func TestParseEncrypt(t *testing.T) {
	t.Run("test short flag", func(t *testing.T) {
		parser := newParser()
		encrypt := ParseEncrypt(parser)

		assert.NotNil(t, encrypt)
		assert.False(t, *encrypt)

		os.Args = []string{"sentinel", "-e" /*+ string(sentinel.Encrypt)*/}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.True(t, *encrypt)
	})

	t.Run("test long flag", func(t *testing.T) {
		parser := newParser()
		encrypt := ParseEncrypt(parser)

		assert.NotNil(t, encrypt)
		assert.False(t, *encrypt)

		os.Args = []string{"sentinel", "--encrypt" /*+ string(sentinel.EncryptExp)*/}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.True(t, *encrypt)
	})
}

func TestParseExpires(t *testing.T) {
	t.Run("test default", func(t *testing.T) {
		parser := newParser()
		expires := ParseExpires(parser)

		assert.NotNil(t, expires)
		assert.Equal(t, "", *expires)

		os.Args = []string{"sentinel"}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.Equal(t, "never", *expires)
	})

	t.Run("test short flag", func(t *testing.T) {
		parser := newParser()
		expires := ParseExpires(parser)

		assert.NotNil(t, expires)
		assert.Equal(t, "", *expires)

		os.Args = []string{"sentinel", "-E" /*+ string(sentinel.Expires)*/, "2023-12-31"}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.Equal(t, "2023-12-31", *expires)
	})

	t.Run("test long flag", func(t *testing.T) {
		parser := newParser()
		expires := ParseExpires(parser)

		assert.NotNil(t, expires)
		assert.Equal(t, "", *expires)

		os.Args = []string{"sentinel", "--exp" /*+ string(sentinel.ExpiresExp)*/, "2023-12-31"}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.Equal(t, "2023-12-31", *expires)
	})
}

func TestParseNotBefore(t *testing.T) {

	t.Run("test default", func(t *testing.T) {
		parser := newParser()
		notBefore := ParseNotBefore(parser)

		assert.NotNil(t, notBefore)
		assert.Equal(t, "", *notBefore)

		os.Args = []string{"sentinel"}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.Equal(t, "now", *notBefore)
	})

	t.Run("test short flag", func(t *testing.T) {
		parser := newParser()
		notBefore := ParseNotBefore(parser)

		assert.NotNil(t, notBefore)
		assert.Equal(t, "", *notBefore)

		os.Args = []string{"sentinel", "-N" /*+ string(sentinel.NotBefore)*/, "2023-12-31"}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.Equal(t, "2023-12-31", *notBefore)
	})

	t.Run("test long flag", func(t *testing.T) {
		parser := newParser()
		notBefore := ParseNotBefore(parser)

		assert.NotNil(t, notBefore)
		assert.Equal(t, "", *notBefore)

		os.Args = []string{"sentinel", "--nbf" /*+ string(sentinel.NotBeforeExp)*/, "2023-12-31"}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.Equal(t, "2023-12-31", *notBefore)
	})
}

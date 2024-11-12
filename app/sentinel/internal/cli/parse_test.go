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

func TestParseWorkload(t *testing.T) {
	t.Run("test short flag", func(t *testing.T) {
		parser := newParser()
		workload := ParseWorkload(parser)

		assert.NotNil(t, workload)
		assert.Equal(t, []string{}, *workload)

		os.Args = []string{"sentinel", "-w", "short workload"}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.Equal(t, []string{"short workload"}, *workload)
	})

	t.Run("test long flag", func(t *testing.T) {
		parser := newParser()
		workload := ParseWorkload(parser)

		assert.NotNil(t, workload)
		assert.Equal(t, []string{}, *workload)

		os.Args = []string{"sentinel", "--workload", "long workload"}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.Equal(t, []string{"long workload"}, *workload)
	})
}

func TestParseSecret(t *testing.T) {
	t.Run("test short flag", func(t *testing.T) {
		parser := newParser()
		secret := ParseSecret(parser)

		assert.NotNil(t, secret)
		assert.Equal(t, "", *secret)

		os.Args = []string{"sentinel", "-s", "short secret option"}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.Equal(t, "short secret option", *secret)
	})

	t.Run("test long flag", func(t *testing.T) {
		parser := newParser()
		secret := ParseSecret(parser)

		assert.NotNil(t, secret)
		assert.Equal(t, "", *secret)

		os.Args = []string{"sentinel", "--secret", "long secret option"}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.Equal(t, "long secret option", *secret)
	})
}

func TestParseTemplate(t *testing.T) {
	t.Run("test short flag", func(t *testing.T) {
		parser := newParser()
		template := ParseTemplate(parser)

		assert.NotNil(t, template)
		assert.Equal(t, "", *template)

		os.Args = []string{"sentinel", "-t", "short template"}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.Equal(t, "short template", *template)
	})

	t.Run("test long flag", func(t *testing.T) {
		parser := newParser()
		template := ParseTemplate(parser)

		assert.NotNil(t, template)
		assert.Equal(t, "", *template)

		os.Args = []string{"sentinel", "--template", "long template"}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.Equal(t, "long template", *template)
	})
}

func TestParseFormat(t *testing.T) {
	t.Run("test short flag", func(t *testing.T) {
		parser := newParser()
		format := ParseFormat(parser)

		assert.NotNil(t, format)
		assert.Equal(t, "", *format)

		os.Args = []string{"sentinel", "-f", "short format"}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.Equal(t, "short format", *format)
	})

	t.Run("test long flag", func(t *testing.T) {
		parser := newParser()
		format := ParseFormat(parser)

		assert.NotNil(t, format)
		assert.Equal(t, "", *format)

		os.Args = []string{"sentinel", "--format", "long format"}

		err := parser.Parse(os.Args)
		assert.NoError(t, err)
		assert.Equal(t, "long format", *format)
	})
}

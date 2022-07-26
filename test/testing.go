package test

import "testing"

const (
	ok     = "\u2713" // √
	failed = "\u2717" // ×
)

func Failed(t *testing.T, descFmt string, params ...any) {
	fmt := descFmt + "  [%s]"
	if params != nil {
		params = append(params, failed)
		t.Errorf(fmt, params...)
	} else {
		t.Errorf(fmt, failed)
	}
}

func Pass(t *testing.T, descFmt string, params ...any) {
	fmt := descFmt + "  [%s]"
	if params != nil {
		params = append(params, ok)
		t.Logf(fmt, params...)
	} else {
		t.Logf(fmt, ok)
	}
}

func Log(t *testing.T, pass bool, descFmt string, params ...any) {
	if pass {
		Pass(t, descFmt, params...)
	} else {
		Failed(t, descFmt, params...)
	}
}

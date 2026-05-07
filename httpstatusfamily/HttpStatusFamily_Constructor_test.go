package httpstatusfamily

import "testing"

func TestHttpStatusFamily_New_KnownNames(t *testing.T) {
	for _, name := range []string{"Informational", "Successful", "Redirection", "ClientError", "ServerError"} {
		v, err := New(name)
		if err != nil {
			t.Fatalf("New(%q) unexpected error: %v", name, err)
		}
		if v.IsInvalid() {
			t.Fatalf("New(%q) returned Invalid", name)
		}
		if v.Name() != name {
			t.Fatalf("Round-trip: New(%q).Name() = %q, want %q", name, v.Name(), name)
		}
	}
}

func TestHttpStatusFamily_New_Unknown(t *testing.T) {
	v, err := New("Bogus")
	if err == nil {
		t.Fatal("New(\"Bogus\") expected error, got nil")
	}
	if !v.IsInvalid() {
		t.Fatalf("New(\"Bogus\") expected Invalid, got %v", v)
	}
}

func TestHttpStatusFamily_Constructors(t *testing.T) {
	if Min() != Informational {
		t.Fatalf("Min() = %v, want Informational", Min())
	}
	if Max() != ServerError {
		t.Fatalf("Max() = %v, want ServerError", Max())
	}
}

func TestHttpStatusFamily_FromStatusCode(t *testing.T) {
	cases := []struct {
		code int
		want Variant
	}{
		{100, Informational}, {199, Informational},
		{200, Successful}, {201, Successful}, {299, Successful},
		{301, Redirection}, {399, Redirection},
		{400, ClientError}, {404, ClientError}, {499, ClientError},
		{500, ServerError}, {503, ServerError}, {599, ServerError},
		{0, Invalid}, {99, Invalid}, {600, Invalid}, {-1, Invalid},
	}
	for _, c := range cases {
		if got := FromStatusCode(c.code); got != c.want {
			t.Fatalf("FromStatusCode(%d) = %v, want %v", c.code, got, c.want)
		}
	}
}

func TestHttpStatusFamily_Predicates(t *testing.T) {
	if !ServerError.IsError() || !ClientError.IsError() {
		t.Fatal("ClientError/ServerError must be errors")
	}
	if Successful.IsError() || Redirection.IsError() || Informational.IsError() {
		t.Fatal("Non-4xx/5xx families must not be errors")
	}
	if !ServerError.IsRetryable() {
		t.Fatal("ServerError must be retryable")
	}
	if ClientError.IsRetryable() || Successful.IsRetryable() {
		t.Fatal("Non-5xx families must not be retryable")
	}
}

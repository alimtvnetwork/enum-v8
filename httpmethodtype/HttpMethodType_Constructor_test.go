package httpmethodtype

import "testing"

func TestHttpMethodType_New_KnownNames(t *testing.T) {
	for _, name := range []string{"Get", "Post", "Put", "Patch", "Delete", "Head", "Options"} {
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

func TestHttpMethodType_New_Unknown(t *testing.T) {
	v, err := New("Bogus")
	if err == nil {
		t.Fatal("New(\"Bogus\") expected error, got nil")
	}
	if !v.IsInvalid() {
		t.Fatalf("New(\"Bogus\") expected Invalid, got %v", v)
	}
}

func TestHttpMethodType_Constructors(t *testing.T) {
	if Min() != Get {
		t.Fatalf("Min() = %v, want Get", Min())
	}
	if Max() != Options {
		t.Fatalf("Max() = %v, want Options", Max())
	}
}

func TestHttpMethodType_Predicates(t *testing.T) {
	if !Get.IsSafe() || !Head.IsSafe() || !Options.IsSafe() {
		t.Fatal("Get/Head/Options should be safe")
	}
	if Post.IsSafe() || Put.IsSafe() || Delete.IsSafe() || Patch.IsSafe() {
		t.Fatal("Post/Put/Delete/Patch should not be safe")
	}
	if !Get.IsIdempotent() || !Put.IsIdempotent() || !Delete.IsIdempotent() {
		t.Fatal("Get/Put/Delete should be idempotent")
	}
	if Post.IsIdempotent() || Patch.IsIdempotent() {
		t.Fatal("Post/Patch should not be idempotent")
	}
	if !Post.IsBodyAllowed() || !Put.IsBodyAllowed() || !Patch.IsBodyAllowed() {
		t.Fatal("Post/Put/Patch should allow body")
	}
	if Get.IsBodyAllowed() || Head.IsBodyAllowed() {
		t.Fatal("Get/Head should not allow body")
	}
}

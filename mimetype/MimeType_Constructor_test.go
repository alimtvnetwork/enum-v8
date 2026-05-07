package mimetype

import "testing"

func TestMimeType_New_KnownNames(t *testing.T) {
	for _, name := range []string{"Application", "Audio", "Font", "Image", "Message", "Model", "Multipart", "Text", "Video"} {
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

func TestMimeType_New_Unknown(t *testing.T) {
	v, err := New("Bogus")
	if err == nil {
		t.Fatal("New(\"Bogus\") expected error, got nil")
	}
	if !v.IsInvalid() {
		t.Fatalf("New(\"Bogus\") expected Invalid, got %v", v)
	}
}

func TestMimeType_Constructors(t *testing.T) {
	if Min() != Application {
		t.Fatalf("Min() = %v, want Application", Min())
	}
	if Max() != Video {
		t.Fatalf("Max() = %v, want Video", Max())
	}
}

func TestMimeType_FromContentType(t *testing.T) {
	cases := []struct {
		in   string
		want Variant
	}{
		{"text/html", Text},
		{"text/html; charset=utf-8", Text},
		{"  Application/json  ", Application},
		{"image/png", Image},
		{"audio/mpeg", Audio},
		{"video/mp4", Video},
		{"font/woff2", Font},
		{"message/rfc822", Message},
		{"model/gltf+json", Model},
		{"multipart/form-data; boundary=x", Multipart},
		{"", Invalid},
		{"bogus/whatever", Invalid},
	}
	for _, c := range cases {
		if got := FromContentType(c.in); got != c.want {
			t.Fatalf("FromContentType(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestMimeType_Predicates(t *testing.T) {
	if !Audio.IsMedia() || !Image.IsMedia() || !Video.IsMedia() {
		t.Fatal("Audio/Image/Video must be media")
	}
	if Text.IsMedia() || Application.IsMedia() {
		t.Fatal("Text/Application must not be media")
	}
	if !Text.IsTextual() {
		t.Fatal("Text must be textual")
	}
	if !Application.IsBinary() || !Image.IsBinary() {
		t.Fatal("Application/Image must be binary")
	}
	if Text.IsBinary() {
		t.Fatal("Text must not be binary")
	}
}

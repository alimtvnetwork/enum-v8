package promptclitype

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core-v9/issetter"
)

// AL2-06 Batch F coverage suite for promptclitype.

func TestPromptCliType_NewAndAccessors(t *testing.T) {
	for _, name := range []string{"Accept", "Reject", "Later", "Review"} {
		v, err := New(name)
		if err != nil {
			t.Errorf("New(%q): %v", name, err)
		}
		if v.Name() != name {
			t.Errorf("Name mismatch: %q != %q", v.Name(), name)
		}
		_ = NewMust(name)
	}
	// alternative names go through newOtherWays
	if v, err := New("yes"); err != nil || v != Accept {
		t.Errorf("New(yes) = %v err=%v", v, err)
	}
	if v, err := New("no"); err != nil || v != Reject {
		t.Errorf("New(no) = %v err=%v", v, err)
	}
	if v, err := New("ask"); err != nil || v != Later {
		t.Errorf("New(ask) = %v err=%v", v, err)
	}
	if _, err := New("__bogus__"); err == nil {
		t.Error("New bogus should fail")
	}
	if Min() != Invalid || Max() != Review {
		t.Errorf("Min/Max wrong: %v %v", Min(), Max())
	}
	if RangesInvalidErr() == nil {
		t.Error("RangesInvalidErr nil")
	}

	if NewUsingBool(true) != Accept || NewUsingBool(false) != Reject {
		t.Error("NewUsingBool wrong")
	}
	if NewUsingAndBooleans(true, true, true) != Accept ||
		NewUsingAndBooleans(true, false) != Reject {
		t.Error("NewUsingAndBooleans wrong")
	}
	if v, err := NewUsingSetter(issetter.True); err != nil || v != Accept {
		t.Errorf("NewUsingSetter(True) = %v err=%v", v, err)
	}
	if v, err := NewUsingSetter(issetter.False); err != nil || v != Reject {
		t.Errorf("NewUsingSetter(False) = %v err=%v", v, err)
	}

	v := Accept
	if !v.IsAccept() || v.IsReject() || v.IsLater() {
		t.Error("predicates wrong")
	}
	if !Reject.IsReject() || !Later.IsLater() {
		t.Error("predicates wrong (2)")
	}
	if v.OnOffLowercaseName() == "" || v.TrueFalseName() == "" || v.ToNumberString() == "" {
		t.Error("string accessors empty")
	}
	if !Is("Accept", Accept) || Is("__bogus__", Accept) {
		t.Error("Is wrong")
	}
	if len(v.AllNameValues()) == 0 || len(v.IntegerEnumRanges()) == 0 {
		t.Error("ranges empty")
	}
	data, err := json.Marshal(v)
	if err != nil || len(data) == 0 {
		t.Errorf("Marshal: %v", err)
	}
	var got Variant
	if err := json.Unmarshal(data, &got); err != nil {
		t.Errorf("Unmarshal: %v", err)
	}
	_ = v.OnlySupportedErr("Accept")
	_ = v.OnlySupportedMsgErr("ctx", "Accept")
	_ = v.ToIsSetter()
}

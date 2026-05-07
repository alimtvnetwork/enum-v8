package rootcmdnames

import (
	"encoding/json"
	"testing"
)

// AL2-06 Batch F coverage suite for cmdenumtypes/rootcmdnames.

func TestRootCmdNames_NewAndAccessors(t *testing.T) {
	for _, name := range []string{"Help", "Compress", "Cron", "Decompress", "Dns", "Download", "Ftp", "Ssh", "Ssl", "User"} {
		v, err := New(name)
		if err != nil {
			t.Errorf("New(%q): %v", name, err)
		}
		if v.Name() != name {
			t.Errorf("Name mismatch: %q != %q", v.Name(), name)
		}
		_ = NewMust(name)
	}
	if _, err := New("__bogus__"); err == nil {
		t.Error("New bogus should fail")
	}
	v := Help
	if len(v.AllNameValues()) == 0 || len(v.IntegerEnumRanges()) == 0 {
		t.Error("ranges empty")
	}
	if v.MinValueString() == "" || v.MaxValueString() == "" {
		t.Error("string accessors empty")
	}
	if min, max := v.MinMaxAny(); min == nil || max == nil {
		t.Error("MinMaxAny nil")
	}
	data, err := json.Marshal(v)
	if err != nil || len(data) == 0 {
		t.Errorf("Marshal: %v", err)
	}
	var got Variant
	if err := json.Unmarshal(data, &got); err != nil {
		t.Errorf("Unmarshal: %v", err)
	}
	_ = v.OnlySupportedErr("Help")
	_ = v.OnlySupportedMsgErr("ctx", "Help")
}

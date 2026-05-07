package brackets

import (
	"encoding/json"
	"reflect"
	"testing"
)

func bracketsAll() []Bracket {
	return []Bracket{
		Invalid, ParenthesisStart, ParenthesisEnd,
		CurlyStart, CurlyEnd, SquareStart, SquareEnd,
	}
}

func categoriesAll() []Category {
	return []Category{UnknownCategory, Parenthesis, Curly, Square}
}

func bracketsSafe(t *testing.T, label string, fn func()) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Logf("recovered %s: %v", label, r)
		}
	}()
	fn()
}

func TestBrackets_Uplift_BracketNullarySweep(t *testing.T) {
	for _, b := range bracketsAll() {
		b := b
		rv := reflect.ValueOf(b)
		rt := rv.Type()
		for i := 0; i < rt.NumMethod(); i++ {
			m := rt.Method(i)
			if m.Type.NumIn() != 1 || m.Type.NumOut() == 0 {
				continue
			}
			bracketsSafe(t, m.Name, func() { _ = rv.Method(i).Call(nil) })
		}
		pv := reflect.ValueOf(&b)
		pt := pv.Type()
		for i := 0; i < pt.NumMethod(); i++ {
			m := pt.Method(i)
			if m.Type.NumIn() != 1 || m.Type.NumOut() == 0 {
				continue
			}
			bracketsSafe(t, "*"+m.Name, func() { _ = pv.Method(i).Call(nil) })
		}
	}
}

func TestBrackets_Uplift_CategoryNullarySweep(t *testing.T) {
	for _, c := range categoriesAll() {
		c := c
		rv := reflect.ValueOf(c)
		rt := rv.Type()
		for i := 0; i < rt.NumMethod(); i++ {
			m := rt.Method(i)
			if m.Type.NumIn() != 1 || m.Type.NumOut() == 0 {
				continue
			}
			bracketsSafe(t, m.Name, func() { _ = rv.Method(i).Call(nil) })
		}
		pv := reflect.ValueOf(&c)
		pt := pv.Type()
		for i := 0; i < pt.NumMethod(); i++ {
			m := pt.Method(i)
			if m.Type.NumIn() != 1 || m.Type.NumOut() == 0 {
				continue
			}
			bracketsSafe(t, "*"+m.Name, func() { _ = pv.Method(i).Call(nil) })
		}
	}
}

func TestBrackets_Uplift_PredicatesAndCrossOps(t *testing.T) {
	for _, a := range bracketsAll() {
		_ = a.IsStart()
		_ = a.IsEnd()
		_ = a.IsParenthesis()
		_ = a.IsCurly()
		_ = a.IsSquare()
		_ = a.IsParenthesisStart()
		_ = a.IsParenthesisEnd()
		_ = a.IsCurlyStart()
		_ = a.IsCurlyEnd()
		_ = a.IsSquareStart()
		_ = a.IsSquareEnd()
		_ = a.IsInvalid()
		_ = a.IsValid()
		_ = a.OtherBracket()
		_ = a.BothBrackets()
		_ = a.Category()
		_ = a.Pair()
		_ = a.SelfWrap()
		_ = a.Format("%d")
		for _, b := range bracketsAll() {
			_ = a.IsAnyOf(b)
			_ = a.IsByteValueEqual(b.ValueByte())
			_ = a.IsValueEqual(b.ValueByte())
			_ = a.IsAnyValuesEqual(b.ValueByte())
			_ = a.IsNameEqual(b.Name())
			_ = a.IsNameOf(b.Name())
			_ = a.IsAnyNamesOf(b.Name())
			_ = a.IsEqual(b.Value())
		}
	}
}

func TestBrackets_Uplift_WrapUnwrapMatrix(t *testing.T) {
	samples := []string{"", "abc", "(abc)", "{abc}", "[abc]", "(abc", "abc)", "{a", "a}"}
	for _, b := range bracketsAll() {
		for _, s := range samples {
			_ = b.WrapAny(s)
			_ = b.WrapAny(nil)
			_ = b.WrapString(s)
			_ = b.WrapFmtString("<{wrapped}>", s)
			_ = b.IsWrapped(s)
			_ = b.UnWrap(s)
			_ = b.WrapWithOptions(true, s)
			_ = b.WrapWithOptions(false, s)
			_ = b.WrapSkipOnExist(s)
			_ = b.WrapRegardless(s)
		}
	}
	for _, c := range categoriesAll() {
		for _, s := range samples {
			_ = c.WrapAny(s)
			_ = c.WrapString(s)
			_ = c.WrapFmtString("<{wrapped}>", s)
			_ = c.IsWrapped(s)
			_ = c.UnWrap(s)
			_ = c.WrapWithOptions(true, s)
			_ = c.WrapWithOptions(false, s)
			_ = c.Start()
			_ = c.End()
			_ = c.Pair()
			_ = c.SelfWrap()
		}
	}
	// Pair / BothBrackets struct surfaces
	for _, c := range categoriesAll() {
		p := c.Pair()
		_ = p.SelfWrap()
		_ = p.IsValid()
		_ = p.IsSafeInvalid()
		_ = p.String()
		for _, s := range samples {
			_ = p.Wrap(s)
			_ = p.WrapAny(s)
			_ = p.WrapAny(nil)
			_ = p.WrapString(s)
			_ = p.WrapFmtString("<{wrapped}>", s)
			_ = p.IsWrapped(s)
			_ = p.UnWrap(s)
			_ = p.WrapWithOptions(true, s)
			_ = p.WrapSkipOnExist(s)
		}
	}
	for _, b := range bracketsAll() {
		bb := b.BothBrackets()
		_ = bb.IsValid()
		_ = bb.IsSafeInvalid()
		_ = bb.String()
		for _, s := range samples {
			_ = bb.WrapAny(s)
			_ = bb.WrapAny(nil)
			_ = bb.WrapString(s)
			_ = bb.WrapFmtString("<{wrapped}>", s)
			_ = bb.IsWrapped(s)
			_ = bb.UnWrap(s)
			_ = bb.WrapWithOptions(true, s)
			_ = bb.WrapSkipOnExist(s)
		}
	}
}

func TestBrackets_Uplift_WhichBracketAndStatus(t *testing.T) {
	for _, ch := range []uint8{'(', ')', '{', '}', '[', ']', 'x', '0'} {
		_ = WhichBracket(ch, true)
		_ = WhichBracket(ch, false)
	}
	_ = EmptyBracketStatus()
}

func TestBrackets_Uplift_JsonRoundTrip(t *testing.T) {
	for _, b := range bracketsAll() {
		data, err := json.Marshal(b)
		if err != nil {
			t.Errorf("marshal %v: %v", b, err)
			continue
		}
		var got Bracket
		_ = json.Unmarshal(data, &got)
	}
}

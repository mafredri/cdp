package cdptype

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestAccessibilityAXValueType_Marshal(t *testing.T) {
	var v AccessibilityAXValueType

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Boolean.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Tristate.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test BooleanOrUndefined.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

	// Test IDRef.
	v = 4
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 4 {
		t.Errorf("Unmarshal(4): v == %d, want 4", v)
	}

	// Test IdrefList.
	v = 5
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 5 {
		t.Errorf("Unmarshal(5): v == %d, want 5", v)
	}

	// Test Integer.
	v = 6
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 6 {
		t.Errorf("Unmarshal(6): v == %d, want 6", v)
	}

	// Test Node.
	v = 7
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 7 {
		t.Errorf("Unmarshal(7): v == %d, want 7", v)
	}

	// Test NodeList.
	v = 8
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 8 {
		t.Errorf("Unmarshal(8): v == %d, want 8", v)
	}

	// Test Number.
	v = 9
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 9 {
		t.Errorf("Unmarshal(9): v == %d, want 9", v)
	}

	// Test String.
	v = 10
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 10 {
		t.Errorf("Unmarshal(10): v == %d, want 10", v)
	}

	// Test ComputedString.
	v = 11
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 11 {
		t.Errorf("Unmarshal(11): v == %d, want 11", v)
	}

	// Test Token.
	v = 12
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 12 {
		t.Errorf("Unmarshal(12): v == %d, want 12", v)
	}

	// Test TokenList.
	v = 13
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 13 {
		t.Errorf("Unmarshal(13): v == %d, want 13", v)
	}

	// Test DOMRelation.
	v = 14
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 14 {
		t.Errorf("Unmarshal(14): v == %d, want 14", v)
	}

	// Test Role.
	v = 15
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 15 {
		t.Errorf("Unmarshal(15): v == %d, want 15", v)
	}

	// Test InternalRole.
	v = 16
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 16 {
		t.Errorf("Unmarshal(16): v == %d, want 16", v)
	}

	// Test ValueUndefined.
	v = 17
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 17 {
		t.Errorf("Unmarshal(17): v == %d, want 17", v)
	}

}

func TestAccessibilityAXValueSourceType_Marshal(t *testing.T) {
	var v AccessibilityAXValueSourceType

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Attribute.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Implicit.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test Style.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

	// Test Contents.
	v = 4
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 4 {
		t.Errorf("Unmarshal(4): v == %d, want 4", v)
	}

	// Test Placeholder.
	v = 5
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 5 {
		t.Errorf("Unmarshal(5): v == %d, want 5", v)
	}

	// Test RelatedElement.
	v = 6
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 6 {
		t.Errorf("Unmarshal(6): v == %d, want 6", v)
	}

}

func TestAccessibilityAXValueNativeSourceType_Marshal(t *testing.T) {
	var v AccessibilityAXValueNativeSourceType

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Figcaption.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Label.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test Labelfor.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

	// Test Labelwrapped.
	v = 4
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 4 {
		t.Errorf("Unmarshal(4): v == %d, want 4", v)
	}

	// Test Legend.
	v = 5
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 5 {
		t.Errorf("Unmarshal(5): v == %d, want 5", v)
	}

	// Test Tablecaption.
	v = 6
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 6 {
		t.Errorf("Unmarshal(6): v == %d, want 6", v)
	}

	// Test Title.
	v = 7
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 7 {
		t.Errorf("Unmarshal(7): v == %d, want 7", v)
	}

	// Test Other.
	v = 8
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 8 {
		t.Errorf("Unmarshal(8): v == %d, want 8", v)
	}

}

func TestAccessibilityAXGlobalStates_Marshal(t *testing.T) {
	var v AccessibilityAXGlobalStates

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Disabled.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Hidden.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test HiddenRoot.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

	// Test Invalid.
	v = 4
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 4 {
		t.Errorf("Unmarshal(4): v == %d, want 4", v)
	}

	// Test Keyshortcuts.
	v = 5
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 5 {
		t.Errorf("Unmarshal(5): v == %d, want 5", v)
	}

	// Test Roledescription.
	v = 6
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 6 {
		t.Errorf("Unmarshal(6): v == %d, want 6", v)
	}

}

func TestAccessibilityAXLiveRegionAttributes_Marshal(t *testing.T) {
	var v AccessibilityAXLiveRegionAttributes

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Live.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Atomic.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test Relevant.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

	// Test Busy.
	v = 4
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 4 {
		t.Errorf("Unmarshal(4): v == %d, want 4", v)
	}

	// Test Root.
	v = 5
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 5 {
		t.Errorf("Unmarshal(5): v == %d, want 5", v)
	}

}

func TestAccessibilityAXWidgetAttributes_Marshal(t *testing.T) {
	var v AccessibilityAXWidgetAttributes

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Autocomplete.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Haspopup.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test Level.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

	// Test Multiselectable.
	v = 4
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 4 {
		t.Errorf("Unmarshal(4): v == %d, want 4", v)
	}

	// Test Orientation.
	v = 5
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 5 {
		t.Errorf("Unmarshal(5): v == %d, want 5", v)
	}

	// Test Multiline.
	v = 6
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 6 {
		t.Errorf("Unmarshal(6): v == %d, want 6", v)
	}

	// Test Readonly.
	v = 7
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 7 {
		t.Errorf("Unmarshal(7): v == %d, want 7", v)
	}

	// Test Required.
	v = 8
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 8 {
		t.Errorf("Unmarshal(8): v == %d, want 8", v)
	}

	// Test Valuemin.
	v = 9
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 9 {
		t.Errorf("Unmarshal(9): v == %d, want 9", v)
	}

	// Test Valuemax.
	v = 10
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 10 {
		t.Errorf("Unmarshal(10): v == %d, want 10", v)
	}

	// Test Valuetext.
	v = 11
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 11 {
		t.Errorf("Unmarshal(11): v == %d, want 11", v)
	}

}

func TestAccessibilityAXWidgetStates_Marshal(t *testing.T) {
	var v AccessibilityAXWidgetStates

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Checked.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Expanded.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test Modal.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

	// Test Pressed.
	v = 4
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 4 {
		t.Errorf("Unmarshal(4): v == %d, want 4", v)
	}

	// Test Selected.
	v = 5
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 5 {
		t.Errorf("Unmarshal(5): v == %d, want 5", v)
	}

}

func TestAccessibilityAXRelationshipAttributes_Marshal(t *testing.T) {
	var v AccessibilityAXRelationshipAttributes

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Activedescendant.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Controls.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test Describedby.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

	// Test Details.
	v = 4
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 4 {
		t.Errorf("Unmarshal(4): v == %d, want 4", v)
	}

	// Test Errormessage.
	v = 5
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 5 {
		t.Errorf("Unmarshal(5): v == %d, want 5", v)
	}

	// Test Flowto.
	v = 6
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 6 {
		t.Errorf("Unmarshal(6): v == %d, want 6", v)
	}

	// Test Labelledby.
	v = 7
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 7 {
		t.Errorf("Unmarshal(7): v == %d, want 7", v)
	}

	// Test Owns.
	v = 8
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 8 {
		t.Errorf("Unmarshal(8): v == %d, want 8", v)
	}

}

func TestBrowserWindowState_Marshal(t *testing.T) {
	var v BrowserWindowState

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Normal.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Minimized.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test Maximized.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

	// Test Fullscreen.
	v = 4
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 4 {
		t.Errorf("Unmarshal(4): v == %d, want 4", v)
	}

}

func TestCSSStyleSheetOrigin_Marshal(t *testing.T) {
	var v CSSStyleSheetOrigin

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Injected.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test UserAgent.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test Inspector.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

	// Test Regular.
	v = 4
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 4 {
		t.Errorf("Unmarshal(4): v == %d, want 4", v)
	}

}

func TestDOMPseudoType_Marshal(t *testing.T) {
	var v DOMPseudoType

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test FirstLine.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test FirstLetter.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test Before.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

	// Test After.
	v = 4
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 4 {
		t.Errorf("Unmarshal(4): v == %d, want 4", v)
	}

	// Test Backdrop.
	v = 5
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 5 {
		t.Errorf("Unmarshal(5): v == %d, want 5", v)
	}

	// Test Selection.
	v = 6
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 6 {
		t.Errorf("Unmarshal(6): v == %d, want 6", v)
	}

	// Test FirstLineInherited.
	v = 7
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 7 {
		t.Errorf("Unmarshal(7): v == %d, want 7", v)
	}

	// Test Scrollbar.
	v = 8
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 8 {
		t.Errorf("Unmarshal(8): v == %d, want 8", v)
	}

	// Test ScrollbarThumb.
	v = 9
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 9 {
		t.Errorf("Unmarshal(9): v == %d, want 9", v)
	}

	// Test ScrollbarButton.
	v = 10
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 10 {
		t.Errorf("Unmarshal(10): v == %d, want 10", v)
	}

	// Test ScrollbarTrack.
	v = 11
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 11 {
		t.Errorf("Unmarshal(11): v == %d, want 11", v)
	}

	// Test ScrollbarTrackPiece.
	v = 12
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 12 {
		t.Errorf("Unmarshal(12): v == %d, want 12", v)
	}

	// Test ScrollbarCorner.
	v = 13
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 13 {
		t.Errorf("Unmarshal(13): v == %d, want 13", v)
	}

	// Test Resizer.
	v = 14
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 14 {
		t.Errorf("Unmarshal(14): v == %d, want 14", v)
	}

	// Test InputListButton.
	v = 15
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 15 {
		t.Errorf("Unmarshal(15): v == %d, want 15", v)
	}

}

func TestDOMShadowRootType_Marshal(t *testing.T) {
	var v DOMShadowRootType

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test UserAgent.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Open.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test Closed.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

}

func TestDOMDebuggerDOMBreakpointType_Marshal(t *testing.T) {
	var v DOMDebuggerDOMBreakpointType

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test SubtreeModified.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test AttributeModified.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test NodeRemoved.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

}

func TestEmulationVirtualTimePolicy_Marshal(t *testing.T) {
	var v EmulationVirtualTimePolicy

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Advance.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Pause.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test PauseIfNetworkFetchesPending.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

}

func TestInputGestureSourceType_Marshal(t *testing.T) {
	var v InputGestureSourceType

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Default.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Touch.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test Mouse.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

}

func TestMemoryPressureLevel_Marshal(t *testing.T) {
	var v MemoryPressureLevel

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Moderate.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Critical.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

}

func TestNetworkErrorReason_Marshal(t *testing.T) {
	var v NetworkErrorReason

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Failed.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Aborted.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test TimedOut.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

	// Test AccessDenied.
	v = 4
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 4 {
		t.Errorf("Unmarshal(4): v == %d, want 4", v)
	}

	// Test ConnectionClosed.
	v = 5
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 5 {
		t.Errorf("Unmarshal(5): v == %d, want 5", v)
	}

	// Test ConnectionReset.
	v = 6
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 6 {
		t.Errorf("Unmarshal(6): v == %d, want 6", v)
	}

	// Test ConnectionRefused.
	v = 7
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 7 {
		t.Errorf("Unmarshal(7): v == %d, want 7", v)
	}

	// Test ConnectionAborted.
	v = 8
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 8 {
		t.Errorf("Unmarshal(8): v == %d, want 8", v)
	}

	// Test ConnectionFailed.
	v = 9
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 9 {
		t.Errorf("Unmarshal(9): v == %d, want 9", v)
	}

	// Test NameNotResolved.
	v = 10
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 10 {
		t.Errorf("Unmarshal(10): v == %d, want 10", v)
	}

	// Test InternetDisconnected.
	v = 11
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 11 {
		t.Errorf("Unmarshal(11): v == %d, want 11", v)
	}

	// Test AddressUnreachable.
	v = 12
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 12 {
		t.Errorf("Unmarshal(12): v == %d, want 12", v)
	}

}

func TestNetworkTimestamp_Marshal(t *testing.T) {
	var v NetworkTimestamp

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test non-empty.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "1" {
		t.Errorf("Marshal() got %s, want 1", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
}

func TestNetworkHeaders_Marshal(t *testing.T) {
	var v NetworkHeaders

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test non-empty.
	v = []byte("\"test\"")
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if !bytes.Equal(v, b) {
		t.Errorf("Marshal() got %s, want %s", b, v)
	}
	v = nil
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if !bytes.Equal(v, b) {
		t.Errorf("Unmarshal() got %s, want %s", b, v)
	}
}

func TestNetworkConnectionType_Marshal(t *testing.T) {
	var v NetworkConnectionType

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test None.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Cellular2g.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test Cellular3g.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

	// Test Cellular4g.
	v = 4
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 4 {
		t.Errorf("Unmarshal(4): v == %d, want 4", v)
	}

	// Test Bluetooth.
	v = 5
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 5 {
		t.Errorf("Unmarshal(5): v == %d, want 5", v)
	}

	// Test Ethernet.
	v = 6
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 6 {
		t.Errorf("Unmarshal(6): v == %d, want 6", v)
	}

	// Test Wifi.
	v = 7
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 7 {
		t.Errorf("Unmarshal(7): v == %d, want 7", v)
	}

	// Test Wimax.
	v = 8
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 8 {
		t.Errorf("Unmarshal(8): v == %d, want 8", v)
	}

	// Test Other.
	v = 9
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 9 {
		t.Errorf("Unmarshal(9): v == %d, want 9", v)
	}

}

func TestNetworkCookieSameSite_Marshal(t *testing.T) {
	var v NetworkCookieSameSite

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Strict.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Lax.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

}

func TestNetworkResourcePriority_Marshal(t *testing.T) {
	var v NetworkResourcePriority

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test VeryLow.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Low.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test Medium.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

	// Test High.
	v = 4
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 4 {
		t.Errorf("Unmarshal(4): v == %d, want 4", v)
	}

	// Test VeryHigh.
	v = 5
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 5 {
		t.Errorf("Unmarshal(5): v == %d, want 5", v)
	}

}

func TestNetworkBlockedReason_Marshal(t *testing.T) {
	var v NetworkBlockedReason

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Csp.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test MixedContent.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test Origin.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

	// Test Inspector.
	v = 4
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 4 {
		t.Errorf("Unmarshal(4): v == %d, want 4", v)
	}

	// Test SubresourceFilter.
	v = 5
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 5 {
		t.Errorf("Unmarshal(5): v == %d, want 5", v)
	}

	// Test Other.
	v = 6
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 6 {
		t.Errorf("Unmarshal(6): v == %d, want 6", v)
	}

}

func TestOverlayInspectMode_Marshal(t *testing.T) {
	var v OverlayInspectMode

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test SearchForNode.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test SearchForUAShadowDOM.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test None.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

}

func TestPageResourceType_Marshal(t *testing.T) {
	var v PageResourceType

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Document.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Stylesheet.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test Image.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

	// Test Media.
	v = 4
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 4 {
		t.Errorf("Unmarshal(4): v == %d, want 4", v)
	}

	// Test Font.
	v = 5
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 5 {
		t.Errorf("Unmarshal(5): v == %d, want 5", v)
	}

	// Test Script.
	v = 6
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 6 {
		t.Errorf("Unmarshal(6): v == %d, want 6", v)
	}

	// Test TextTrack.
	v = 7
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 7 {
		t.Errorf("Unmarshal(7): v == %d, want 7", v)
	}

	// Test XHR.
	v = 8
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 8 {
		t.Errorf("Unmarshal(8): v == %d, want 8", v)
	}

	// Test Fetch.
	v = 9
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 9 {
		t.Errorf("Unmarshal(9): v == %d, want 9", v)
	}

	// Test EventSource.
	v = 10
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 10 {
		t.Errorf("Unmarshal(10): v == %d, want 10", v)
	}

	// Test WebSocket.
	v = 11
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 11 {
		t.Errorf("Unmarshal(11): v == %d, want 11", v)
	}

	// Test Manifest.
	v = 12
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 12 {
		t.Errorf("Unmarshal(12): v == %d, want 12", v)
	}

	// Test Other.
	v = 13
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 13 {
		t.Errorf("Unmarshal(13): v == %d, want 13", v)
	}

}

func TestPageTransitionType_Marshal(t *testing.T) {
	var v PageTransitionType

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Link.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Typed.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test AutoBookmark.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

	// Test AutoSubframe.
	v = 4
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 4 {
		t.Errorf("Unmarshal(4): v == %d, want 4", v)
	}

	// Test ManualSubframe.
	v = 5
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 5 {
		t.Errorf("Unmarshal(5): v == %d, want 5", v)
	}

	// Test Generated.
	v = 6
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 6 {
		t.Errorf("Unmarshal(6): v == %d, want 6", v)
	}

	// Test AutoToplevel.
	v = 7
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 7 {
		t.Errorf("Unmarshal(7): v == %d, want 7", v)
	}

	// Test FormSubmit.
	v = 8
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 8 {
		t.Errorf("Unmarshal(8): v == %d, want 8", v)
	}

	// Test Reload.
	v = 9
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 9 {
		t.Errorf("Unmarshal(9): v == %d, want 9", v)
	}

	// Test Keyword.
	v = 10
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 10 {
		t.Errorf("Unmarshal(10): v == %d, want 10", v)
	}

	// Test KeywordGenerated.
	v = 11
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 11 {
		t.Errorf("Unmarshal(11): v == %d, want 11", v)
	}

	// Test Other.
	v = 12
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 12 {
		t.Errorf("Unmarshal(12): v == %d, want 12", v)
	}

}

func TestPageDialogType_Marshal(t *testing.T) {
	var v PageDialogType

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Alert.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Confirm.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test Prompt.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

	// Test Beforeunload.
	v = 4
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 4 {
		t.Errorf("Unmarshal(4): v == %d, want 4", v)
	}

}

func TestPageNavigationResponse_Marshal(t *testing.T) {
	var v PageNavigationResponse

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Proceed.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Cancel.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test CancelAndIgnore.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

}

func TestRuntimeUnserializableValue_Marshal(t *testing.T) {
	var v RuntimeUnserializableValue

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Infinity.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test NaN.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test NegativeInfinity.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

	// Test Negative0.
	v = 4
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 4 {
		t.Errorf("Unmarshal(4): v == %d, want 4", v)
	}

}

func TestRuntimeTimestamp_Marshal(t *testing.T) {
	var v RuntimeTimestamp

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test non-empty.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "1" {
		t.Errorf("Marshal() got %s, want 1", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
}

func TestSecurityState_Marshal(t *testing.T) {
	var v SecurityState

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Unknown.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Neutral.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test Insecure.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

	// Test Warning.
	v = 4
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 4 {
		t.Errorf("Unmarshal(4): v == %d, want 4", v)
	}

	// Test Secure.
	v = 5
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 5 {
		t.Errorf("Unmarshal(5): v == %d, want 5", v)
	}

	// Test Info.
	v = 6
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 6 {
		t.Errorf("Unmarshal(6): v == %d, want 6", v)
	}

}

func TestSecurityCertificateErrorAction_Marshal(t *testing.T) {
	var v SecurityCertificateErrorAction

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Continue.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Cancel.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

}

func TestServiceWorkerVersionRunningStatus_Marshal(t *testing.T) {
	var v ServiceWorkerVersionRunningStatus

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Stopped.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Starting.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test Running.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

	// Test Stopping.
	v = 4
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 4 {
		t.Errorf("Unmarshal(4): v == %d, want 4", v)
	}

}

func TestServiceWorkerVersionStatus_Marshal(t *testing.T) {
	var v ServiceWorkerVersionStatus

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test New.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Installing.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test Installed.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

	// Test Activating.
	v = 4
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 4 {
		t.Errorf("Unmarshal(4): v == %d, want 4", v)
	}

	// Test Activated.
	v = 5
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 5 {
		t.Errorf("Unmarshal(5): v == %d, want 5", v)
	}

	// Test Redundant.
	v = 6
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 6 {
		t.Errorf("Unmarshal(6): v == %d, want 6", v)
	}

}

func TestStorageType_Marshal(t *testing.T) {
	var v StorageType

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test bad input.
	v = 9001
	_, err = json.Marshal(&v)
	if err == nil {
		t.Error("Marshal(9001) got no error, want error")
	}
	err = json.Unmarshal([]byte("9001"), &v)
	if err == nil {
		t.Error("Unmarshal(9001) got no error, want error")
	}

	// Test Appcache.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 1 {
		t.Errorf("Unmarshal(1): v == %d, want 1", v)
	}

	// Test Cookies.
	v = 2
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 2 {
		t.Errorf("Unmarshal(2): v == %d, want 2", v)
	}

	// Test FileSystems.
	v = 3
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 3 {
		t.Errorf("Unmarshal(3): v == %d, want 3", v)
	}

	// Test Indexeddb.
	v = 4
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 4 {
		t.Errorf("Unmarshal(4): v == %d, want 4", v)
	}

	// Test LocalStorage.
	v = 5
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 5 {
		t.Errorf("Unmarshal(5): v == %d, want 5", v)
	}

	// Test ShaderCache.
	v = 6
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 6 {
		t.Errorf("Unmarshal(6): v == %d, want 6", v)
	}

	// Test Websql.
	v = 7
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 7 {
		t.Errorf("Unmarshal(7): v == %d, want 7", v)
	}

	// Test ServiceWorkers.
	v = 8
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 8 {
		t.Errorf("Unmarshal(8): v == %d, want 8", v)
	}

	// Test CacheStorage.
	v = 9
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 9 {
		t.Errorf("Unmarshal(9): v == %d, want 9", v)
	}

	// Test All.
	v = 10
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 10 {
		t.Errorf("Unmarshal(10): v == %d, want 10", v)
	}

	// Test Other.
	v = 11
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if strings.Contains(v.String(), string(b)) {
		t.Errorf("Marshal() got %s, want ~~ %s", b, v.String())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if v != 11 {
		t.Errorf("Unmarshal(11): v == %d, want 11", v)
	}

}

func TestTracingMemoryDumpConfig_Marshal(t *testing.T) {
	var v TracingMemoryDumpConfig

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test non-empty.
	v = []byte("\"test\"")
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if !bytes.Equal(v, b) {
		t.Errorf("Marshal() got %s, want %s", b, v)
	}
	v = nil
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
	if !bytes.Equal(v, b) {
		t.Errorf("Unmarshal() got %s, want %s", b, v)
	}
}

func TestTimestamp_Marshal(t *testing.T) {
	var v Timestamp

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}

	// Test non-empty.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %v, want no error", err)
	}
	if string(b) != "1" {
		t.Errorf("Marshal() got %s, want 1", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %v, want no error", err)
	}
}

package entities

import (
	"encoding/json"
	"testing"
)

func TestErrorSerialization(t *testing.T) {
	e := Error{Code: 404, Message: "Not Found"}
	bytes, err := json.Marshal(e)
	if err != nil {
		t.Fatalf("Failed to marshal error: %v", err)
	}

	expected := `{"code":404,"message":"Not Found"}`
	if string(bytes) != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, string(bytes))
	}

	var de Error
	err = json.Unmarshal(bytes, &de)
	if err != nil {
		t.Fatalf("Failed to unmarshal error: %v", err)
	}

	if de != e {
		t.Errorf("Expected %v but got %v", e, de)
	}
}

func TestMessageSerialization(t *testing.T) {
	m := Message{Code: 200, Message: "OK"}
	bytes, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("Failed to marshal message: %v", err)
	}

	expected := `{"code":200,"message":"OK"}`
	if string(bytes) != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, string(bytes))
	}

	var dm Message
	err = json.Unmarshal(bytes, &dm)
	if err != nil {
		t.Fatalf("Failed to unmarshal message: %v", err)
	}

	if dm != m {
		t.Errorf("Expected %v but got %v", m, dm)
	}
}

func TestResultInfoSerialization(t *testing.T) {
	ri := ResultInfo{Count: 5, Page: 2, PerPage: 25, TotalCount: 100}
	bytes, err := json.Marshal(ri)
	if err != nil {
		t.Fatalf("Failed to marshal result info: %v", err)
	}

	expected := `{"count":5,"page":2,"per_page":25,"total_count":100}`
	if string(bytes) != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, string(bytes))
	}

	var dri ResultInfo
	err = json.Unmarshal(bytes, &dri)
	if err != nil {
		t.Fatalf("Failed to unmarshal result info: %v", err)
	}

	if dri != ri {
		t.Errorf("Expected %v but got %v", ri, dri)
	}
}

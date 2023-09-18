package log

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"testing"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
		DisableQuote:     true,
	})
}

func TestSetOutput(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	Info("test")

	expected := "level=info msg=test\n"
	if buf.String() != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, buf.String())
	}
}

func TestSetLevel(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)

	SetLevel(LevelDebug)
	Trace("Should not appear")
	if buf.String() != "" {
		t.Errorf("Expected empty string but got '%s'", buf.String())
	}

	SetLevel(LevelTrace)
	Trace("Should appear")

	expected := "level=trace msg=Should appear\n"
	if buf.String() != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, buf.String())
	}
}

func TestParseLevel(t *testing.T) {
	tests := []struct {
		input  string
		output Level
		err    bool
	}{
		{"t", LevelTrace, false},
		{"trace", LevelTrace, false},
		{"TRACE", LevelTrace, false},
		{"Trace", LevelTrace, false},
		{"d", LevelDebug, false},
		{"debug", LevelDebug, false},
		{"DEBUG", LevelDebug, false},
		{"Debug", LevelDebug, false},
		{"i", LevelInfo, false},
		{"info", LevelInfo, false},
		{"INFO", LevelInfo, false},
		{"Info", LevelInfo, false},
		{"w", LevelWarn, false},
		{"warn", LevelWarn, false},
		{"warning", LevelWarn, false},
		{"WARN", LevelWarn, false},
		{"WARNING", LevelWarn, false},
		{"Warn", LevelWarn, false},
		{"Warning", LevelWarn, false},
		{"e", LevelError, false},
		{"error", LevelError, false},
		{"ERROR", LevelError, false},
		{"Error", LevelError, false},
		{"f", LevelFatal, false},
		{"fatal", LevelFatal, false},
		{"FATAL", LevelFatal, false},
		{"Fatal", LevelFatal, false},
		{"p", LevelPanic, false},
		{"panic", LevelPanic, false},
		{"PANIC", LevelPanic, false},
		{"Panic", LevelPanic, false},
		{"unknown", 0, true},
		{"z", 0, true},
	}

	for _, test := range tests {
		result, err := ParseLevel(test.input)
		if err != nil && !test.err {
			t.Errorf("expected no error for input %s but got %v", test.input, err)
		}
		if err == nil && test.err {
			t.Errorf("expected an error for input %s but got none", test.input)
		}
		if result != test.output {
			t.Errorf("for input %s, expected %v but got %v", test.input, test.output, result)
		}
	}
}

func TestTrace(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)

	Trace("Trace test")
	expected := "level=trace msg=Trace test\n"
	if buf.String() != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, buf.String())
	}
}

func TestTraceWithFields(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)

	TraceWithFields("Trace with fields", FieldsMap{"key": "value"})
	expected := "level=trace msg=Trace with fields key=value\n"
	if buf.String() != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, buf.String())
	}
}

func TestDebug(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)

	Debug("Debug test")
	expected := "level=debug msg=Debug test\n"
	if buf.String() != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, buf.String())
	}
}

func TestInfo(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)

	Info("Info test")
	expected := "level=info msg=Info test\n"
	if buf.String() != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, buf.String())
	}
}

func TestWarn(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)

	Warn("Warn test")
	expected := "level=warning msg=Warn test\n"
	if buf.String() != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, buf.String())
	}
}

func TestError(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)

	Error("Error test")
	expected := "level=error msg=Error test\n"
	if buf.String() != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, buf.String())
	}
}

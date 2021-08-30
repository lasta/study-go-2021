package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

// cf. Go で標準出力に書き込まれた文字列のテストをする方法
// https://zenn.dev/glassonion1/articles/8ac939208bd455
func captureStdout(t *testing.T, block func()) string {
	t.Helper()

	stdout := os.Stdout
	defer func() {
		os.Stdout = stdout
	}()

	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	block()

	err := writer.Close()
	if err != nil {
		t.Fatalf("writer is nil unexpectedly")
	}
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(reader); err != nil {
		t.Fatalf("failed to read buf: %v", err)
	}

	return strings.TrimRight(buf.String(), "\n")
}

func TestEx01Main(t *testing.T) {
	os.Args = []string{"first", "second", "third"}

	actual := captureStdout(t, main)
	expected := "first second third"

	if actual != expected {
		t.Fatalf("expected: %s, but actual: %s", expected, actual)
	}
}

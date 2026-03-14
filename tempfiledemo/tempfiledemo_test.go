package tempfiledemo

import (
	"os"
	"strings"
	"testing"
)

func assertFileContent(t *testing.T, filename, want string) {
	t.Helper()
	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Cannot read file %s: %v", filename, err)
	}
	if string(data) != want {
		t.Errorf("Get %q, want %q", string(data), want)
	}
}

func assertFileNotExist(t *testing.T, filename string) {
	t.Helper()
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		t.Errorf("File %s is still exist", filename)
	}
}

func assertFileLines(t *testing.T, filename string, want []string) {
	t.Helper()
	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Cannot read file %s: %v", filename, err)
	}

	lines := strings.Split(string(data), "\n")

	if len(lines) != len(want) {
		t.Errorf("Get %d lines, want %d lines", len(lines), len(want))
	}

	for i, v := range want {
		if lines[i] != v {
			t.Errorf("Line %d: got %q, but want %q", i, lines[i], v)
		}
	}
}

func TestWriteLinesToTemp(t *testing.T) {

	t.Run("File is exist and contains strings", func(t *testing.T) {
		content := []string{"hello"}
		filename, err := WriteLinesToTemp("testfile", content)
		if err != nil {
			t.Fatalf("Cannot create tmp file: %v", err)
		}
		t.Logf("File %s was created", filename)

		t.Cleanup(func() {
			os.Remove(filename) // We can ignore error here
			t.Log("Cleanup test env")
		})

		assertFileContent(t, filename, "hello")

		if err := os.Remove(filename); err != nil {
			t.Fatalf("Cannot delete tmp file %q: %v", filename, err) // Cannot remove file
		}

		assertFileNotExist(t, filename)
		t.Logf("File %s was removed", filename)
	})

	t.Run("Check line separator and order", func(t *testing.T) {
		//
		content := []string{"hello", "hallo", "bonjour", "hola", "ciao"}
		filename, err := WriteLinesToTemp("testfile", content)
		if err != nil {
			t.Fatalf("Cannot create tmp file: %v", err)
		}
		t.Logf("File %s was created", filename)

		t.Cleanup(func() {
			os.Remove(filename) // We can ignore error here
			t.Log("Cleanup test env")
		})

		assertFileLines(t, filename, content)
	})

	t.Run("Check empty list", func(t *testing.T) {
		content := []string{}
		filename, err := WriteLinesToTemp("testfile", content)
		if err != nil {
			t.Fatalf("Cannot create tmp file: %v", err)
		}
		t.Logf("File %s was created", filename)

		t.Cleanup(func() {
			os.Remove(filename) // We can ignore error here
			t.Log("Cleanup test env")
		})

		assertFileContent(t, filename, "")
	})

}

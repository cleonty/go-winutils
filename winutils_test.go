// winutils_test.go
package winutils

import (
	"fmt"
	"os"
	"testing"
)

func TestRemoveFileOnReboot(t *testing.T) {
	filename := "file_to_remove.txt"
	file, err := os.Create(filename)
	if err != nil {
		t.Error(err)
	}
	file.WriteString("Hello world!\n")
	file.Close()
	result := RemoveFileOnReboot(filename)
	if result != 0 {
		t.Errorf("Expected 0 got %d\n", result)
	}
	fmt.Println(result)
}

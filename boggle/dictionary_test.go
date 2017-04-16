package boggle

import (
	"strings"
	"testing"
)

func TestDictImportWhiteSpace(t *testing.T) {
	input := `cat
						dog
						bird
						rabbit`
	reader := strings.NewReader(input)
	var dict Dictionary
	err := dict.ImportFromReader(reader)
	if err != nil {
		t.Errorf("Failed received during import: %s", err)
	}
}

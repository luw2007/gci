package section

import (
	"testing"

	"github.com/luw2007/gci/pkg/specificity"
)

func TestStandardPackageSpecificity(t *testing.T) {
	testCases := []specificityTestData{
		{"context", Standard{}, specificity.StandardMatch{}},
		{"contexts", Standard{}, specificity.MisMatch{}},
		{"crypto", Standard{}, specificity.StandardMatch{}},
		{"crypto1", Standard{}, specificity.MisMatch{}},
		{"crypto/ae", Standard{}, specificity.MisMatch{}},
		{"crypto/aes", Standard{}, specificity.StandardMatch{}},
		{"crypto/aes2", Standard{}, specificity.MisMatch{}},
	}
	testSpecificity(t, testCases)
}

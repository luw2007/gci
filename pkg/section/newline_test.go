package section

import (
	"testing"

	"github.com/luw2007/gci/pkg/specificity"
)

func TestNewLineSpecificity(t *testing.T) {
	testCases := []specificityTestData{
		{`""`, NewLine{}, specificity.MisMatch{}},
		{`"x"`, NewLine{}, specificity.MisMatch{}},
		{`"\n"`, NewLine{}, specificity.MisMatch{}},
	}
	testSpecificity(t, testCases)
}

// func TestNewLineToString(t *testing.T) {
// 	testSectionToString(t, NewLine{})
// }

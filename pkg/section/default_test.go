package section

import (
	"testing"

	"github.com/luw2007/gci/pkg/specificity"
)

func TestDefaultSpecificity(t *testing.T) {
	testCases := []specificityTestData{
		{`""`, Default{}, specificity.Default{}},
		{`"x"`, Default{}, specificity.Default{}},
	}
	testSpecificity(t, testCases)
}

// func TestDefaultSectionParsing(t *testing.T) {
// 	testCases := []sectionTestData{
// 		{"def", Default{}, nil},
// 		{"defAult(invalid)", nil, SectionTypeDoesNotAcceptParametersError},
// 	}
// 	testSectionParser(t, testCases)
// }

// func TestDefaultSectionToString(t *testing.T) {
// 	testSectionToString(t, Default{})
// }

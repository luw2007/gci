package section

import (
	"fmt"
	"strings"

	"github.com/luw2007/gci/pkg/parse"
	"github.com/luw2007/gci/pkg/specificity"
)

type Custom struct {
	Prefix string
}

const CustomType = "custom"

func (c Custom) MatchSpecificity(spec *parse.GciImports) specificity.MatchSpecificity {
	for _, prefix := range strings.Split(c.Prefix, ";") {
		if strings.HasPrefix(spec.Path, prefix) {
			return specificity.Match{Length: len(prefix)}
		}
	}

	return specificity.MisMatch{}
}

func (c Custom) String() string {
	return fmt.Sprintf("prefix(%s)", c.Prefix)
}

func (c Custom) Type() string {
	return CustomType
}

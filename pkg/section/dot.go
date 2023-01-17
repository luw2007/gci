package section

import (
	"github.com/luw2007/gci/pkg/parse"
	"github.com/luw2007/gci/pkg/specificity"
)

type Dot struct{}

const DotType = "dot"

func (d Dot) MatchSpecificity(spec *parse.GciImports) specificity.MatchSpecificity {
	if spec.Name == "." {
		return specificity.NameMatch{}
	}
	return specificity.MisMatch{}
}

func (d Dot) String() string {
	return DotType
}

func (d Dot) Type() string {
	return DotType
}

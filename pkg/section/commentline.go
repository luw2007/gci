package section

import (
	"fmt"

	"github.com/luw2007/gci/pkg/parse"
	"github.com/luw2007/gci/pkg/specificity"
)

type CommentLine struct {
	Comment string
}

func (c CommentLine) MatchSpecificity(spec *parse.GciImports) specificity.MatchSpecificity {
	return specificity.MisMatch{}
}

func (c CommentLine) String() string {
	return fmt.Sprintf("commentline(%s)", c.Comment)
}

func (c CommentLine) Type() string {
	return "commentline"
}

package format

import (
	"fmt"

	"github.com/luw2007/gci/pkg/config"
	"github.com/luw2007/gci/pkg/log"
	"github.com/luw2007/gci/pkg/parse"
	"github.com/luw2007/gci/pkg/section"
	"github.com/luw2007/gci/pkg/specificity"
)

type Block struct {
	Start, End int
}

type resultMap map[string][]*Block

func Format(data []*parse.GciImports, cfg *config.Config) (resultMap, error) {
	result := make(resultMap, len(cfg.Sections))
	for _, d := range data {
		// determine match specificity for every available section
		var bestSection section.Section
		var bestSectionSpecificity specificity.MatchSpecificity = specificity.MisMatch{}
		for _, section := range cfg.Sections {
			sectionSpecificity := section.MatchSpecificity(d)
			if sectionSpecificity.IsMoreSpecific(specificity.MisMatch{}) && sectionSpecificity.Equal(bestSectionSpecificity) {
				// specificity is identical
				// return nil, section.EqualSpecificityMatchError{}
				return nil, nil
			}
			if sectionSpecificity.IsMoreSpecific(bestSectionSpecificity) {
				// better match found
				bestSectionSpecificity = sectionSpecificity
				bestSection = section
			}
		}
		if bestSection == nil {
			return nil, section.NoMatchingSectionForImportError{Imports: d}
		}
		log.L().Debug(fmt.Sprintf("Matched import %v to section %s", d, bestSection))
		result[bestSection.String()] = append(result[bestSection.String()], &Block{d.Start, d.End})
	}

	return result, nil
}

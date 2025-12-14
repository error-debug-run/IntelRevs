package parser

func NewGenericParser() Parser {
	return &genericParser{}
}

type genericParser struct{}

func (p *genericParser) parse(html string) ([]string, error) {

	return []string{}, nil
}

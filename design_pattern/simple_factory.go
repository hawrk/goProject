package design_pattern

type IruleConfigParser interface {
	Parse([]byte)
}

type JsonRuleConfigParser struct {
}

func (j JsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type YamlRuleConfigParser struct {
}

func (y YamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

func NewIRuleConfigParser(config string) IruleConfigParser {
	switch config {
	case "json":
		return JsonRuleConfigParser{}
	case "ymal":
		return YamlRuleConfigParser{}
	}
	return nil
}

package rule

type Element struct {
	Name string
	Hint string
}

type Rule struct {
	Name     string
	Elements []Element
}

func NewRule(name string) *Rule {
	return &Rule{Name: name}
}

func (r *Rule) AddElement(name string, hint string) {
	r.Elements = append(r.Elements, Element{Name: name, Hint: hint})
}

func (r *Rule) RemoveElement(name string) {
	// ...
}

func (r *Rule) UpdateElement(name string, hint string) {
	// ...
}

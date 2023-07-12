package rule

import (
	"fmt"
	"github.com/letusgogo/nopass/config"
	"github.com/letusgogo/nopass/log"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
)

type Element struct {
	Name string
	Hint string
	Sort int
}

type Rule struct {
	Name            string
	ElementWithVals []*ElementWithVal
}

func NewRuleFromConfig(ruleName string, c *config.Config) (*Rule, error) {
	ruleConfig := c.Rules[ruleName]
	if ruleConfig == nil {
		return nil, fmt.Errorf("can not find rule: %s", ruleName)
	}

	rule := &Rule{
		Name:            ruleName,
		ElementWithVals: []*ElementWithVal{},
	}

	sortSet := make(map[string]struct{})
	for _, elemConfig := range ruleConfig {
		if _, exists := sortSet[elemConfig.Name]; exists {
			return nil, fmt.Errorf("duplicate element in rule '%s': '%s'", ruleName, elemConfig.Name)
		}
		sortSet[elemConfig.Name] = struct{}{}

		rule.ElementWithVals = append(rule.ElementWithVals, &ElementWithVal{
			Element: Element{
				Name: elemConfig.Name,
				Hint: elemConfig.Hint,
				Sort: 0,
			},
			Value: "",
		})
	}

	//sort.Sort(ElementWithVals(rule.ElementWithVals))

	return rule, nil
}

type ElementWithVal struct {
	Element
	Value string
}

//type ElementWithVals []*ElementWithVal
//
//func (a ElementWithVals) Len() int           { return len(a) }
//func (a ElementWithVals) Less(i, j int) bool { return a[i].Sort < a[j].Sort }
//func (a ElementWithVals) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// FullElements full all elements with args
func (r *Rule) FullElements() error {
	for _, elemVal := range r.ElementWithVals {
		log.Hintf("please input %s: (%s) and then input Enter.", elemVal.Name, elemVal.Hint)
		val, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			log.Fatalf("\nerror reading element val: %v\n", err)
		}
		elemVal.Value = string(val)
		log.Hint("")
	}

	return nil
}

// Display display all elements with values
func (r *Rule) Display() {
	log.DrawParagraph("please verify generator info", log.DebugLevel, func() {
		for _, elem := range r.ElementWithVals {
			log.Infof("     %s: %s\n", elem.Name, elem.Value)
		}
	})
}

func (r *Rule) Compose() string {
	result := ""
	for _, elem := range r.ElementWithVals {
		result += fmt.Sprintf("%s=%s&", elem.Name, elem.Value)
	}

	strings.TrimSuffix(result, "&")
	return result
}

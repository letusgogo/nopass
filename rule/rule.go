package rule

import (
	"fmt"
	"github.com/letusgogo/nopass/config"
	"github.com/letusgogo/nopass/log"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"sort"
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

	sortSet := make(map[int]bool)

	for elemName, elemConfig := range ruleConfig {
		if _, exists := sortSet[elemConfig.Sort]; exists {
			return nil, fmt.Errorf("duplicate sort value in rule '%s': %d", ruleName, elemConfig.Sort)
		}
		sortSet[elemConfig.Sort] = true

		rule.ElementWithVals = append(rule.ElementWithVals, &ElementWithVal{
			Element: Element{
				Name: elemName,
				Hint: elemConfig.Hint,
				Sort: elemConfig.Sort,
			},
			Value: "",
		})
	}

	sort.Sort(ElementWithVals(rule.ElementWithVals))

	return rule, nil
}

type ElementWithVal struct {
	Element
	Value string
}

type ElementWithVals []*ElementWithVal

func (a ElementWithVals) Len() int           { return len(a) }
func (a ElementWithVals) Less(i, j int) bool { return a[i].Sort < a[j].Sort }
func (a ElementWithVals) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// FullElements full all elements with args
func (r *Rule) FullElements() error {
	for _, elemVal := range r.ElementWithVals {
		log.Hintf("please input %s: (%s) and then input Enter.", elemVal.Name, elemVal.Hint)
		val, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			log.Fatalf("error reading element val: %v\n", err)
		}
		elemVal.Value = string(val)
		log.Hint("")
	}

	return nil
}

// Display display all elements with values
func (r *Rule) Display() {
	log.DrawPhase("please verify generator info", log.DebugLevel, func() {
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

package conditionable

import (
	"strconv"
)

type Conditionable interface {
	PassRequirement(*Condition) bool
}

func GetConditionable(c *Condition, ci Conditionable) bool {
	return ci.PassRequirement(c)
}

type Condition struct {
	key			string
	value		string
	isMandatory bool
}

func (c *Condition) Key() string {
	return c.key
}

func (c *Condition) Value() string {
	return c.value
}

func (c *Condition) ValueInt() int {
	intval, e := strconv.Atoi(c.value)
	if (e != nil) {
		panic("Condition can't convert to int in ValueInt")
	}
	return intval
}

func (c *Condition) IsMandatory() bool {
	return c.isMandatory
}
/*type SupermarketCashier struct {
	Conds	    map[string]Condition
}

func (s *SupermarketCashier) PassRequirement(c *Condition) bool {
	validates := false
	switch c.Key {
	case "MaxItems":
		sval, e1 := strconv.Atoi(s.Conds[c.Key].Value)
		cval, e2 := strconv.Atoi(c.Value)
		if e1 != nil || e2 != nil {				
			validates = false
		} else {
			validates = sval >= cval
		}
	case "IsActive":
		validates = s.Conds[c.Key].Value == "1"
	}
	return validates
}

func main() {
	super := new(SupermarketCashier)
	
	var cond1 Condition
	cond1.Key = "IsActive"
	cond1.Value = "1"
	cond1.IsMandatory = true
	
	var cond2 Condition
	cond2.Key = "MaxItems"
	cond2.Value = "10"
	cond2.IsMandatory = true
	
	var condClient Condition
	condClient.Key = "MaxItems"
	condClient.Value = "10"
	condClient.IsMandatory = true
	
	super.Conds = make(map[string]Condition)
	super.Conds[cond1.Key] = cond1
	super.Conds[cond2.Key] = cond2
	
	if GetConditionable(&condClient, super) {
		fmt.Println("client can access to super")
	} else {
		fmt.Println("client can't access to super")
	}
}*/
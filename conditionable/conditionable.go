package main

import (
	"fmt"
	"strconv"
)

type Conditionable interface {
	// Validate(c1, c2 Condition) 			bool
	Conditions()		 	   		map[string]Condition
	HasCondition(condKey string)	bool
	ConditionValidates(Condition) 	bool
}

type Condition struct {
	Key			string
	Value		string
	IsMandatory bool
}

type SupermarketCashier struct {
	MaxItems 	int
	IsActive 	int
	Conds	    map[string]Condition
}

func (s *SupermarketCashier) Conditions() map[string]Condition {
	return s.Conds
}

func (s *SupermarketCashier) HasCondition(key string) bool {
	if s.Conds[key].Key == key {
		return true
	}
	return false
}

func (s *SupermarketCashier) ConditionValidates(c Condition) bool {
	validates := false
	if !c.IsMandatory {
		return true
	} else if  s.HasCondition(c.Key) {
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
	}

	return validates
}

func main() {
	super := new(SupermarketCashier)
	super.MaxItems = 5
	super.IsActive = 1
	
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
	condClient.Value = "8"
	condClient.IsMandatory = true
	
	super.Conds = make(map[string]Condition)
	super.Conds[cond1.Key] = cond1
	super.Conds[cond2.Key] = cond2
	
	fmt.Println("Super conditions: ", super.Conditions())
	fmt.Println("Client condition: ", condClient)
	
	fmt.Println("Super has cond1: ", super.HasCondition(cond1.Key))
	fmt.Println("Super has cond2: ", super.HasCondition(cond2.Key))
	fmt.Println("Super has condClient: ", super.HasCondition(condClient.Key))
	
	fmt.Println("Super validates cond1: ", super.ConditionValidates(cond1))
	fmt.Println("Super has condClient: ", super.ConditionValidates(condClient))
}
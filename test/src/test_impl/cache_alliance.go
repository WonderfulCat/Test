package test_impl

import (
	"bytes"
	"test/src/test_interface"
)

type CacheAlliance struct {
	AllianceMap map[string]test_interface.AllianceI
}

func (c *CacheAlliance) Build() {
	c.AllianceMap = make(map[string]test_interface.AllianceI)

}

func (c *CacheAlliance) AllianceGetByName(name string) (test_interface.AllianceI, bool) {
	if ret, ok := c.AllianceMap[name]; ok {
		return ret, true
	}

	return nil, false
}

func (c *CacheAlliance) AllianceAdd(alliance test_interface.AllianceI) bool {
	if _, ok := c.AllianceGetByName(alliance.GetName()); ok {
		return false
	}

	c.AllianceMap[alliance.GetName()] = alliance
	return true
}

func (c *CacheAlliance) AllianceRemove(name string) {
	delete(c.AllianceMap, name)
}

func (c *CacheAlliance) AllianceList() string {
	var ret bytes.Buffer

	for k := range c.AllianceMap {
		ret.WriteString(k + "\n")
	}
	return ret.String()
}

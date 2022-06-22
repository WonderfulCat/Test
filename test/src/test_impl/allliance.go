package test_impl

import (
	"bytes"
	"errors"
	"test/src/test_constant"
)

type AllianceInfo struct {
	Name    string
	Members map[string]*Member
	*AllianceStoreInfo
}

func (c *AllianceInfo) Build(name string) {
	c.Name = name
	c.AllianceStoreInfo = InitAllianceStoreInfo()
	c.Members = make(map[string]*Member, test_constant.ALLIANCE_MAX_MEMBERS)

}

func (c *AllianceInfo) GetName() string {
	return c.Name
}

func (c *AllianceInfo) GetMemberList() string {
	var ret bytes.Buffer

	for k := range c.Members {
		ret.WriteString(k + "\n")
	}
	return ret.String()
}

func (c *AllianceInfo) CreateAlliance(aname string) bool {
	return true
}

func (c *AllianceInfo) JoinAlliance(cname string, permission int32) (bool, error) {
	if len(c.Members) >= test_constant.ALLIANCE_MAX_MEMBERS {
		return false, errors.New(test_constant.RES_ERR_MSG_5)
	}

	if _, ok := c.Members[cname]; ok {
		return false, errors.New(test_constant.RES_ERR_MSG_10)
	}

	c.Members[cname] = AddMember(cname, permission)
	return true, nil
}

func (c *AllianceInfo) DismissAlliance(cname string) (bool, error) {
	if !c.CheckPermission(cname, test_constant.MEMBER_PERMISSION_ADMIN) {
		return false, errors.New(test_constant.RES_ERR_MSG_15)
	}
	return true, nil
}

func (c *AllianceInfo) CheckPermission(cname string, permission int32) bool {
	if m, ok := c.Members[cname]; ok {
		return m.Permission&permission == permission
	}
	return false
}

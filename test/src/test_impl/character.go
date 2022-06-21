package test_impl

type CharacterInfo struct {
	Name         string
	Pswd         string
	AllianceName string
}

func (c *CharacterInfo) Build(name, pswd string) {
	c.Name = name
	c.Pswd = pswd

}

func (c *CharacterInfo) GetName() string {
	return c.Name
}

func (c *CharacterInfo) GetPswd() string {
	return c.Pswd
}

func (c *CharacterInfo) GetAllianceName() string {
	return c.AllianceName
}

func (c *CharacterInfo) SetAllianceName(name string) {
	c.AllianceName = name
}

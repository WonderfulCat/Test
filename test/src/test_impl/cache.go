package test_impl

type Cache struct {
	*CacheAlliance
	*CacheCharacter
}

func (c *Cache) Build() {
	c.CacheAlliance = &CacheAlliance{}
	c.CacheCharacter = &CacheCharacter{}

	c.CacheAlliance.Build()
	c.CacheCharacter.Build()

}

package test_interface

type CacheI interface {
	Build()
	CacheCharacterI
	CacheAlliance
}

type CacheCharacterI interface {
	CharacterGetByNamePswd(name, pswd string) (CharacterI, int32)
	CharacterAdd(character CharacterI) bool
	CharacterGetByName(name string) CharacterI
}

type CacheAlliance interface {
	AllianceGetByName(name string) (AllianceI, bool)
	AllianceAdd(alliance AllianceI) bool
	AllianceRemove(name string)
	AllianceList() string
}

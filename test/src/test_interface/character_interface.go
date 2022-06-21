package test_interface

type CharacterI interface {
	Build(name, pswd string)
	GetName() string
	GetPswd() string
	GetAllianceName() string
	SetAllianceName(name string)
}

package test_impl

type Member struct {
	Name       string
	Permission int32
}

func AddMember(name string, permission int32) *Member {
	return &Member{Name: name, Permission: permission}
}

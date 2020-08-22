package flag

type flag struct{}

func (f *flag) Parse() {}

func New() *flag {
	return new(flag)
}

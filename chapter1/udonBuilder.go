package main

type Portion int

const (
	Regular Portion = iota + 1
	Small
	Large
)

type fluentOpt struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

type Udon struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdon(p Portion) *fluentOpt {
	return &fluentOpt{
		men:      p,
		aburaage: false,
		ebiten:   0,
	}
}

func (f *fluentOpt) Aburaage() *fluentOpt {
	f.aburaage = true
	return f
}

func (f *fluentOpt) Ebiten(n uint) *fluentOpt {
	f.ebiten = n
	return f
}

func (f *fluentOpt) Order() *Udon {
	return &Udon{
		men:      f.men,
		aburaage: f.aburaage,
		ebiten:   f.ebiten,
	}
}

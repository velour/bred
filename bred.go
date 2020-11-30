package bred

import "net/http"

type Bred struct {
	addr string
	tplDir string
}

func New(addr, tplDir string) *Bred {
	return &Bred{addr, tplDir }
}

func (b *Bred) Start() {
	fs := http.FileServer(http.Dir(b.tplDir))
	http.Handle("/", http.StripPrefix("/", fs))
	http.ListenAndServe(b.addr, nil)
}

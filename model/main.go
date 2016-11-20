package model

type gopher struct{
	name string
	age int
}


// all gophers respond to jump method
func (g gopher)Jump()string{
	if g.age < 65 {
		return g.name  + " yükseğe sıçrıyabiliyor"
	}else{
		return g.name + " zıplayıver çekirge"
	}
}

type horse struct {
	name string
	weight int
}

func (h horse)Jump() string{
	if h.weight > 200 {
		return "çok ağır bu at bea"
	}else{
		return "sıska kalmış la bu at arpa verin şuna"
	}
}


func GetList() []jumper{
	ali := &gopher{name: "Ali", age: 20}
	veli := &gopher{name: "Veli", age: 80}
	barbaro := &horse{ name: "yebeni", weight: 300}
	list := []jumper{ali,veli,barbaro}
	return list
}

type jumper interface {
	Jump() string //jump methodunu kim uyguluyorsa get liste girer kardeşim
}

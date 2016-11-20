package main

import "fmt"


type gopher struct {
	name string
	age int
	isAdult bool
}


func (g gopher) jump() string{
	if g.age < 65 {
		return g.name  + " can jump high"
	}else{
		return g.name + " can still jump"
	}
}


func validateAge(g *gopher){
	g.isAdult = g.age >= 21
}


func main(){
	
	gopher1 := &gopher{ name: "oğan", age : 38}
	validateAge(gopher1)
	fmt.Println(gopher1)
	gopher2 := &gopher{name:"ali", age :90}
	validateAge(gopher2)
	fmt.Println(gopher2)
	fmt.Println(gopher2.jump())	
}


func highJump(name string){
	fmt.Println(name, "can jump high")
}


func lowJump(name string){
	fmt.Println(name, "can't jump")
}
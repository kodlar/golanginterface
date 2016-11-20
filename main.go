package main

import ( 
		"fmt"
		"github.com/kodlar/datastruct/model"
        )



/* interfacesiz method
func main(){
	gopherList := getList()
	for _,gopher := range gopherList{
		fmt.Println(gopher.jump())
	}
	
}
*/

func main(){
	jumperList := model.GetList() //getList()
	for _,jumper := range jumperList{
		fmt.Println(jumper.Jump())
	}
	
}




/* interfacesiz gopher getlist
func getList() []*gopher{
	ali := &gopher{name: "Ali", age: 20}
	veli := &gopher{name: "Veli", age: 80}
	
	list := []*gopher{ali,veli}
	return list
}
*/

package main
import "fmt"
func main(){
	mymap:=make(map[string]int)
	mymap["apple"]=1
	mymap["orange"]=2
	mymap["bananna"]=3
	mymap["cake"]=4
	for key,value:=range mymap{
		fmt.Printf("key",key,"value",value)
	}
	delete(mymap,"apple")
	fmt.Println("after",mymap)
	value,exists:=mymap["cake"]
	if exists{
		fmt.Printf("it is exists",value)
	}else{
		fmt.Printf("not exist")
	}


}
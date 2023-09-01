package main
import "fmt"
func triplet(n int)(int,int,int){
	for a:=1;a<=n/3;a++{
		for b:=a;b<=(n-a)/2;b++{
			c:=n-a-b
			if a*a+b*b==c*c{
				return a,b,c
			}
		}
	}
	return -1,-1,-1
}
func main(){
	a,b,c:=triplet(150)
	fmt.Println("values are %d,%d,%d",a,b,c)
}
        

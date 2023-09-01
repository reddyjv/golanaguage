package main
import( "fmt"
        "math"
)
func largestprime(n int)(int){
	largest:=1
	if n%2==0{
		largest=2
		n/=2
	}
	for i:=3;i<=int(math.Sqrt(float64(n)));i=+2{
		if n%i==0{
			largest=i
			n/=i

		}
		if n>2{
			largest=n
		}
		
	}
	return largest
}
func main(){
	number:=14
	largest:=largestprime(number)
	fmt.Printf("the largest prime factor is %d is %d",number,largest)
}
//a go program to find Largest prime factor
package main
import ("fmt")
func main(){
	var n int
	
	fmt.Println("Enter the number")
	fmt.Scanf("%d", &n)
	l:= -1
	for i := 2; i <= n; {
		if n%i==0{
			n/=i
			l=i
		}else{
			i++
		}
	}
	fmt.Println("The largest prime factor",l)
}
//a go program to find Largest palindrome product
package main
import ("fmt")
func isp(n int)bool{
s:=fmt.Sprintf("%d",n)
return s==reverse(s)

}

func reverse(s string)string{
	c:=[]rune(s)
	for i,j:=0,len(c)-1;i<j;i,j=i+1,j-1{
		c[i],c[j]=c[j],c[i]
	}
	return string(c)

}

func main(){
	l:=0
	max:=0
	min:=0
	fmt.Println("Enter the minimum value")
	fmt.Scanln(&min)
	fmt.Println("Enter the maximum value")
	fmt.Scanln(&max)
	for i:=max;i>min;i--{
		for j:=max;j>min;j--{
			p:=i*j
			if isp(p) && p>l{
				l=p
			}
		}
	}

	fmt.Println("the largest palindrome is",l)
}
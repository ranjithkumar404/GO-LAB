
// a go program to illustrate Maps
package main
import("fmt")
func main(){
	data:=make(map[string]string)
	for{
		var k string
	fmt.Println("Enter a key or 0 to quit")
	fmt.Scanln(&k)
	if k=="0"{
		break
	}
	var v string
	fmt.Println("Enter the value")
	fmt.Scanln(&v)
	data[k]=v
	}
	for{
		var n string
		fmt.Println("Enter the key to search or 0 to quit")
		fmt.Scanln(&n)
		if n=="0"{
			break
		}
		val,f:=data[n]
		if f {
			fmt.Println("The value for the key is",val)
		}else{
			fmt.Println("Value not found")
		}
	}


}
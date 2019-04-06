
package golangexamples

import (
	"fmt"
	"github.com/ehteshamz/greetings
)

func ConcatSlice(SliceToConcat []byte) string{

		var a string
		for i:=0; i<len(SliceToConcat); i++{
			a += string([]byte{SliceToConcat[i]})
			a +="-"
		}
return a
}

func Encrypt(a string,b int) string{

	var c string
	for i:=0;i<len(a);i++{
		j :=((int(a[i])+3))
		if(j>90){
			j=j-26
		}
		c=c+string(j)
	}	
	
	return c

}






func EZGreetings(name string) string {
	return (greetings.PrintGreetings(name))

}

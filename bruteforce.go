//findElement method given array and k element
package main


import (
	"fmt"
)


func findElement(arr [10]int, k int,n int) bool {
	var i int
	for i = 0; i < n; i++ {

		if arr[i] == k {
			return true
		}
	}
	return false
}


func main() {
var n int
fmt.Scan(&n)
 var arr [10] int  
   var i int  
   for i = 0; i < n; i++ {  
      
      fmt.Scan(&arr[i])
   }  
  var find int
   fmt.Scan(&find)
	var check bool = findElement(arr, find, n)

	fmt.Println(check)

	

}

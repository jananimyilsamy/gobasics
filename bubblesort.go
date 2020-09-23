package main

import (
    "fmt"
)


func bubbleSort(input [10] int) {
    
    n := 10
    
    swapped := true
    
    for swapped {
        
        swapped = false
        
        for i := 1; i < n; i++ {
            
            if input[i-1] > input[i] {
                
                input[i], input[i-1] = input[i-1], input[i]
                
                swapped = true
            }
        }
    }
    
    fmt.Println(input)
}

func main() {
    var x [10] int  
   var i int  
   for i = 0; i < 10; i++ {  
      
      fmt.Scan(&x[i])
   }  
    fmt.Println("bubble sort")
    bubbleSort(x)
}

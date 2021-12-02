/* Alta3 Research | RZFeeser
   Channels - Create a basic channel   
   
   the value of mychannel is a memory address (provided it is created by make(). Keep in mind that channels are nothing but pointers. That’s why we can pass them to goroutines, and we can easily put the data or read the data. 
   
   */

package main
  
import "fmt"
  
func main() {
  
    // Creating a channel
    // Using var keyword - initializes with 'nil' - cannot transport data with us
    var mychannel chan int
    fmt.Println("Value of the channel: ", mychannel)
    fmt.Printf("Type of the channel: %T ", mychannel)
  
    // Creating a channel using make() function
    mychannel1 := make(chan int)
    fmt.Println("\nValue of the channel1: ", mychannel1)
    fmt.Printf("Type of the channel1: %T ", mychannel1)
}


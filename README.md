# go

1. Variables and Constants
Variables in Go are declared using the var keyword, and constants are declared using the const keyword.
Syntax:

go
Copy code
var x int = 10
var y = 20 // Type inferred
const pi = 3.14
You can also declare multiple variables at once:
go
Copy code
var x, y, z int
2. Data Types
Go supports basic data types like int, float64, string, bool, and more complex types such as struct, array, slice, and map.
Example:

go
Copy code
var name string = "John"
var age int = 30
var isActive bool = true
3. Arrays and Slices
Array: Fixed-size collection of elements.
Slice: More flexible and common than arrays; dynamic size.
Syntax:

go
Copy code
var arr [3]int = [3]int{1, 2, 3} // Array
slice := []int{1, 2, 3}          // Slice
Adding elements to a slice:
go
Copy code
slice = append(slice, 4)
4. Maps
Maps are key-value pairs and similar to dictionaries in Python.
Syntax:

go
Copy code
var bookings = make(map[string]string)
bookings["fname"] = "John"
bookings["lname"] = "Doe"
5. Functions
Functions are defined with the func keyword.
Syntax:

go
Copy code
func greeting() {
    fmt.Println("Hello!")
}
Functions can take multiple return values:
go
Copy code
func add(a int, b int) (int, int) {
    return a + b, a - b
}
Named return values:
go
Copy code
func add(a int, b int) (sum int, diff int) {
    sum = a + b
    diff = a - b
    return
}
6. Structs
Structs are used to define complex data types. They group multiple variables together.
Syntax:

go
Copy code
type verification struct {
    ageVerification uint
    residentOfIndia string
}

var v verification
v.ageVerification = 25
v.residentOfIndia = "Yes"
7. Control Flow: If, Else, Switch
If-Else is used for conditional statements.
Syntax:

go
Copy code
if age > 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}
Switch provides multiple condition checks.
Syntax:

go
Copy code
switch age {
case 18:
    fmt.Println("18 years old")
case 25:
    fmt.Println("25 years old")
default:
    fmt.Println("Unknown age")
}
8. Loops: For Loop
Go has only the for loop, which is versatile.
Syntax:

go
Copy code
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
Infinite loop:
go
Copy code
for {
    fmt.Println("This is an infinite loop!")
}
9. Goroutines and Concurrency
Goroutines are lightweight threads used for concurrent execution. You can create them using the go keyword.
Syntax:

go
Copy code
go functionName()  // Starts a goroutine
WaitGroup is used to synchronize goroutines, ensuring all goroutines finish execution.
Syntax:

go
Copy code
var wg sync.WaitGroup
wg.Add(1)          // Increment counter
go func() {
    defer wg.Done() // Decrement counter when goroutine finishes
}()
wg.Wait()           // Wait for all goroutines to complete
10. Interfaces
An interface defines behavior (methods) without implementation. Any type that has the methods specified by the interface is considered to implement that interface.
Syntax:

go
Copy code
type Speaker interface {
    Speak() string
}

type Person struct {
    Name string
}

func (p Person) Speak() string {
    return "Hello, " + p.Name
}

var s Speaker = Person{Name: "John"}
fmt.Println(s.Speak()) // Output: Hello, John
11. Error Handling
Go uses explicit error handling via the error type. Functions that might return an error usually return a value and an error.
Syntax:

go
Copy code
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

result, err := divide(10, 0)
if err != nil {
    fmt.Println("Error:", err)
}
12. Packages and Imports
Go supports modular code with packages, which are imported using the import keyword.
Syntax:

go
Copy code
import "fmt"
import (
    "math"
    "strings"
)
Custom package import (e.g., helper in your project):
go
Copy code
import "path/to/helper"
13. Random Numbers
You can generate random numbers using the math/rand package.
Syntax:

go
Copy code
import "math/rand"
randomNum := rand.Intn(50) // Generates a random number between 0 and 49
14. String Manipulation
Strings can be manipulated using functions from the strings package.
Example:

go
Copy code
import "strings"
contains := strings.Contains("hello", "e") // Returns true
15. Type Conversion
In Go, you can convert between types explicitly.
Syntax:

go
Copy code
var num int = 42
var str string = strconv.Itoa(num) // int to string
16. fmt Package (Formatted I/O)
The fmt package is used for formatted input and output operations.
Syntax:

go
Copy code
fmt.Println("Hello World!")
fmt.Printf("Hello, %v\n", "Go") // %v for general formatting
fmt.Sprintf("Hello, %v", "Go") // Return formatted string

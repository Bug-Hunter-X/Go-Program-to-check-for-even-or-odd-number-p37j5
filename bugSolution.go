func main() {
        fmt.Println("Enter a number")
        var a int
        _, err := fmt.Scanln(&a)
        if err != nil {
                fmt.Println("Invalid input. Please enter an integer.")
                return
        }

        if a%2 == 0 {
                fmt.Println("Even")
        } else {
                fmt.Println("Odd")
        }
} 
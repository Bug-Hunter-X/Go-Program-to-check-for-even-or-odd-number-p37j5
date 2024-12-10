# Go Program to Check Even or Odd Number

This program demonstrates a common error in Go programs that handle user input: failure to account for invalid input.  The original code attempts to read an integer from the user, but if the user enters a non-integer value, the program will panic.  The solution provides robust error handling to gracefully manage incorrect input.

## Bug
The bug lies in the use of `fmt.Scanln`. If the user inputs something that cannot be parsed as an integer (e.g., "abc", "1.5", or just leaves it blank), the program will crash with a runtime error.

## Solution
The solution utilizes error handling using `fmt.Scanln`'s return value. The code now checks for errors to ensure valid input before proceeding. It will print a suitable message if non-numeric input is entered.

package main

import ("fmt" "os")

func main() {
  fmt.Println("Hello Docker Actions")
  
  // Access Inputs as environment vars
  firstGreeting := os.Getenv("INPUT_FIRSTGREETING")
  secondGreeting := os.Getenv("INPUT_SECONGREETING")
  thirdGreeting := os.Getenv("INPUT_THIRDGREETING")
  
  // Use those inputs in the action logic
  fmt.Println("Hello " + firstGreeting)
  fmt.Println("Hello " + secondGreeting)
  
  //someimes inputs are not "required" and we can build around that
  if thirdGreeting != "" {
    fmt.Println("Hello " + thirdGreeting)
  }
}

# Exercise 2

This exercise will reinforce our understanding of method sets:
- Create a type person struct
- Attach a method speak to type person using a pointer receiver
  - *person
- Create a type human interface
  - To implicitly implement the interface, a human must have the speak method
- Create func “saySomething”
  - Have it take in a human as a parameter
  - Have it call the speak method
- Show the following in your code
  - You CAN pass a value of type *person into saySomething
  - You CANNOT pass a value of type person into saySomething
- Here is a hint if you need some help
  - https://play.golang.org/p/FAwcQbNtMG

Receivers       Values
-----------------------------------------------
(t T)           T and *T
(t *T)          *T

Solution: [https://github.com/GoesToEleven/go-programming](https://github.com/GoesToEleven/go-programming)


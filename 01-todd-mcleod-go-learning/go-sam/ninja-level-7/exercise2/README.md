# Exercise 2

- Create a person struct
- Create a func called “changeMe” with a *person as a parameter
  - Change the value stored at the *person address
    - important
      - to dereference a struct, use (*value).field
        - p1.first
        - (*p1).first
      - “As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method), x.f is shorthand for (*x).f.”
        - https://golang.org/ref/spec#Selectors
- create a value of type person
  - print out the value
- in func main
  - call “changeMe”
- in func main
  - print out the value

Solution: [https://play.golang.org/p/AehM662HKS](https://play.golang.org/p/AehM662HKS)

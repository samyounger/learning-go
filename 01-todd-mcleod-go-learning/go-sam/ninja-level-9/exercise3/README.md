# Exercise 3

- Using goroutines, create an incrementer program
  - Have a variable to hold the incrementer value
  - Launch a bunch of goroutines
    - Each goroutine should
      - Read the incrementer value
        - Store it in a new variable
      - Yield the processor with runtime.Gosched()
      - Increment the new variable
      - Write the value in the new variable back to the incrementer variable
- Use waitgroups to wait for all of your goroutines to finish
- The above will create a race condition.
- Prove that it is a race condition by using the -race flag
- If you need help, here is a hint: [https://play.golang.org/p/FYGoflKQej](https://play.golang.org/p/FYGoflKQej)

Solution: [https://github.com/GoesToEleven/go-programming](https://github.com/GoesToEleven/go-programming)

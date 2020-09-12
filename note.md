## data types
string
bool
int
    - signed ints
        - int, int8, int16, int32, int64
    - unsigned ints
        - uint, uint8, uint16, uint32, uint64
byte - alias for uint8
rune - alias for uint32
fload32, float64
complex64, complex128
## packages
- package is nothing but directory and any exported memeber of a package should start with uppercase

## arrays and slice

- both are pretty much the same thing but only slice doesnt want you specify the size in other word
it's dynamically setting it's size
    - len() - used to get the length of the size
    - we can slice the array eg.fruits[1:3]
    - we can iterate through array items using **range** 
    for index, item range items {
        fmt.Println(index, item)
    }
## maps
- they are data structures usually used to store key-value pairs
- nameEmail := make(map[string]string)

## Pointers
- as the name suggests it points but what? Ans: memory location of variable

```
a := 5
b := &a // b is now has memory location of variable a
```
to read and write using pointers we use * (astriks). note the variable should be pointer type (varibale having &)
fmt.Println(*b) // which gives 5
*b = 15 // set value of a to 15
fmt.Println(a) // will be 15

## Structs
is same as a class

## Variadic functions
**Variadic functions** can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.

## Closures
Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it.
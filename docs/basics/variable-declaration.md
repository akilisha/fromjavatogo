---
sidebar_position: 4
---

# Variable Declaration

:::info[JAVA]
Variables
:::

Java requires explicitly declaring a variable together with its before it can used.

```Java
int x;
float y = 10f;
```

When declaring and assigning a variable in the same line, Java can use type inference to bind the declared variable to a type;

```Java
var x = 10;
var y = 10f;
```

When declaring a _const_, java uses the keyword final

```java
final var c = 'a';
```

The keyword _final_ can be used in several different contexts:

1. variable - the primitive value assigned cannot be changed (for reference variables, the memory address assigned value cannot be changed)
2. method - the method cannot be overriden in a sub-class
3. class - the class cannot be extended


:::info[GO]
Variables
:::

In Go, variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.

_var_ declares one of more variables without initializing (or assgining) a value. If the declared variable is not explicitly initialized, 
Go will initialize it with its _zeroth_value_. 

Go will infer the type of initialized variables if it is not done explicitly.

The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case. This syntax is only available inside functions.
Outside a function, every statement begins with a keyword (var, func, and so on). 

Go has various value types:

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32 which represents a Unicode code point

float32 float64

complex64 complex128
```

_const_ declares a constant value, similar to a _final_ variable in java.

A _const_ statement can appear anywhere a var statement can. _const_ expressions perform arithmetic with arbitrary precision.

A numeric constant has no type until it’s given one, such as by an explicit conversion.



```go
package main

import "fmt"

func main() {

    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d bool = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "apple"
    fmt.Println(f)

    const n = 500000000

    const d = 3e20 / n
    fmt.Println(d)

    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))

}

initial
1 2
true
0
apple
6e+11
600000000000
-0.28470407323754404
```

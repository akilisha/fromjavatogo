---
sidebar_position: 6
---

# If/Else Statement

:::info[JAVA]
Variables
:::

Java's _if_ statement has remained very basic, and does't merit much discussion.

```java
public class HelloWorld {

    public static void main(String[] args) {

        if (7 % 2 == 0) {
            System.out.println("7 is even");
        } else {
            System.out.println("7 is odd");
        }

        if (12 % 4 == 0) {
            System.out.println("12 is divisible by 4");
        }

        if (8 % 2 == 0 || 7 % 2 == 0) {
            System.out.println("either 8 or 7 are even");
        }

        int num = 9;
        if (num < 0) {
            System.out.printf("%d is negative\n", num);
        } else if (num < 10) {
            System.out.printf("%d has 1 digit\n", num);
        } else {
            System.out.printf("%d has multiple digits\n", num);
        }
    }
}

7 is odd
12 is divisible by 4
either 8 or 7 are even
9 has 1 digit

```

Java has a _ternary operator_ which is short-form of a basic _if-else_ statement (this is not included in Go)

```java
var num = 10
var volume = num < 5 ? "ok tone" : "tone too high"

System.out.println(volume);

"tone too high"
```

:::info[GO]
Variables
:::

Unlike java, the _if_ statement is **not** srrounded by _parenthesis ()_. 

The _else_ **must** start on the same line as the closing _curly brace }_ and it's body is also **not** sorrounded by _parenthesis ()_ either.

In other words, you don’t need parentheses around conditions in Go, but that the braces are required.

There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.

```go
package main

import "fmt"

func main() {

    // Here’s a basic example.
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    // You can have an if statement without an else.
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    // Logical operators like && and || are often useful in conditions.
    if 8%2 == 0 || 7%2 == 0 {
        fmt.Println("either 8 or 7 are even")
    }

    // A statement can precede conditionals; any variables declared in this statement are available in the current and all subsequent branches.
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}
```

Like _for_, the _if_ statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the if. 

```go
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

```

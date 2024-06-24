---
sidebar_position: 7
---

# Switch Statement

:::info[JAVA]
Variables
:::

Unlike if-then and if-then-else statements, the switch statement can have a number of possible execution paths. A switch works with the _byte_, _short_, _char_, and _int_ primitive data types. It also works with enumerated types, the String class, and a few special classes that wrap certain primitive types: _Character_, _Byte_, _Short_, and _Integer_.

Another point of interest is the break statement. Each break statement terminates the enclosing switch statement. Control flow continues with the first statement following the switch block. The break statements are necessary because without them, statements in switch blocks fall through: All statements after the matching case label are executed in sequence, regardless of the expression of subsequent case labels, until a break statement is encountered.

```java
public class HelloWorld {

    // calculates the number of days in a particular month:
    public static void main(String[] args) {

        int month = 2;
        int year = 2000;
        int numDays = 0;

        switch (month) {
            case 1: case 3: case 5:
            case 7: case 8: case 10:
            case 12:
                numDays = 31;
                break;
            case 4: case 6: case 9:
            case 11:
                numDays = 30;
                break;
            case 2:
                if (((year % 4 == 0) &&
                        !(year % 100 == 0))
                        || (year % 400 == 0))
                    numDays = 29;
                else
                    numDays = 28;
                break;
            default:
                System.out.println("Invalid month.");
                break;
        }
        System.out.printf("Number of Days = %d\n", numDays);

    }
}

Number of Days = 29

```

## Switch expression

Since _java 13_, the _switch expression_ was intruduced, which is now able to return a value from the switch. Notice that the new syntax uses the -> operator instead of the colon we’re used to with switch statements. Also, there’s no break keyword: The switch expression doesn’t fall through cases.

```java
public class HelloWorld {

    enum Months {JANUARY, FEBRUARY, MARCH, APRIL, MAY, JUNE, JULY, AUGUST, SEPTEMBER, OCTOBER, NOVEMBER, DECEMBER}

    public static void main(String[] args) {

        var month = Months.APRIL;
        
        var result = switch(month) {
            case JANUARY, JUNE, JULY -> 3;
            case FEBRUARY, SEPTEMBER, OCTOBER, NOVEMBER, DECEMBER -> 1;
            case MARCH, MAY, APRIL, AUGUST -> 2;
            default -> 0; 
        };

        System.out.println(result);
    }
}
```

## yield keyword

Going a bit further, there’s a possibility to obtain fine-grain control over what’s happening on the right side of the expression by using code blocks, and then _yielding_ instead of _returning_ from the block.

```java
var result = switch (month) {
    case JANUARY, JUNE, JULY -> 3;
    case FEBRUARY, SEPTEMBER, OCTOBER, NOVEMBER, DECEMBER -> 1;
    case MARCH, MAY, APRIL, AUGUST -> {
        int monthLength = month.toString().length();
        yield monthLength * 4;
    }
    default -> 0;
};
```

:::info[GO]
Variables
:::

Unlike java, the _switch_ statement does not require a _break_ after every case. Execution will always exit the switch wnenever a mtching case statement is executed.

A swicth case also be used without evaluating an expression, in which case it becomes an alternative way of implementing an _if/else_ statement



```go
package main

import (
    "fmt"
    "time"
)

func main() {

    // Here’s a basic switch.
    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    // You can use commas to separate multiple expressions in the same case statement. We use the optional default case in this example as well.
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

    // switch without an expression is an alternate way to express if/else logic. Here we also show how the case expressions can be non-constants.
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    // A type switch compares types instead of values. You can use this to discover the type of an interface value. In this example, the variable t will have the type corresponding to its clause.
    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}
```

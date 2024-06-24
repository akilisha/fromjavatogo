---
sidebar_position: 5
---

# Looping Constructs

:::info[JAVA]
Variables
:::

Java uses seberal method tyo achieve looping. Declare an array to demostrate

```java 
int[] numbers = {1,2,3,4,5}
```

The output will be the same in all cases

```
num[0]=1
num[1]=2
num[2]=3
num[3]=4
num[4]=5
```

## For loop

```java
for(int i = 0; i < numbers.length; i++){
    System.out.printf("num[%d]=%d\n", i, numbers[i]);
}
```

_For loop_ has a different construct which handles the loop boundary variable implicity;

```java
int l = 0;
for(int num : numbers){
    System.out.printf("num[%d]=%d\n", l++, num);
}

List<Integer> list = List.of(1,2,3,4,5);
for(int num : list){
    System.out.printf("num[%d]=%d\n", l++, num);
}

```

_For loop_ has a special construct also that can be used with any iterator value;

```java
List<Integer> list = List.of(1,2,3,4,5);
for(Iterator<Integer> iter = list.iterator(); iter.hasNext(); ){
    System.out.printf("num[%d]=%d\n", l++, iter.next());
}
```

## While loop

```java
int j = 0;
while (j < numbers.length){
    System.out.printf("num[%d]=%d\n", j, numbers[j]);
    j++; //missing to increment will produce an infinite loop
}
```

## Do While loop

```java
int k = 0;
do {
    System.out.printf("num[%d]=%d\n", k, numbers[k]);
    k++; //missing to increment will produce an infinite loop
} while (k < numbers.length);
```

:::info[GO]
Variables
:::

For all the huffing and puffing in Java, _for_ is Go’s only looping construct. Here are some basic types of for loops.

```go
package main

import "fmt"

func main() {

    // The most basic type, with a single condition.
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    // A classic initial/condition/after for loop.
    for j := 0; j < 3; j++ {
        fmt.Println(j)
    }

    // Another way of accomplishing the basic “do this N times” iteration is range over an integer.
    for i := range 3 {
        fmt.Println("range", i)
    }

    // for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
    for {
        fmt.Println("loop")
        break
    }

    // You can also continue to the next iteration of the loop.
    for n := range 6 {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}

1
2
3
0
1
2
range 0
range 1
range 2
loop
1
3
5
```


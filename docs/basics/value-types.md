---
sidebar_position: 3
---

# Values Types

:::info[JAVA]
Java primitives
:::

One of the original sins that Java is working really hard to rectify, is to do with primitives and their boxed counterparts.
They families of data types use different amounts of memory to represent the same value, and this can be quite inconvenient 
in some application domains. However, these data types have remained consistently the same throughout the existence of Java. 

```java
package some.folder;

public class Main {

  public static void main(String...args){
    // String
    var someString = "go" + "lang";
    System.out.printf("string = %s, type=%s\n", someString, ((Object)someString).getClass().getSimpleName());

    // byte/Byte
    var someBytes = "java".getBytes();
    System.out.printf("bytes = %s, type=%s\n", someBytes, ((Object)someBytes).getClass().getSimpleName());

    // char/Character
    var someCharacter = 'A';
    System.out.printf("char = %C, type=%s\n", someCharacter, ((Object)someCharacter).getClass().getSimpleName());

    // int/Integer
    var someShort = Short.valueOf("1") + Short.valueOf("1");
    System.out.printf("1+1 = %d, type=%s\n", someShort, ((Object)someShort).getClass().getSimpleName());

    // int/Integer
    var someInteger = 1+1;
    System.out.printf("1+1 = %d, type=%s\n", someInteger, ((Object)someInteger).getClass().getSimpleName());

    // long/Long
    var someLong = 1l+1l;
    System.out.printf("1+1 = %d, type=%s\n", someLong, ((Object)someLong).getClass().getSimpleName());

    // float/Float
    var someFloat = 7.0f/3.0f;
    System.out.printf("7.0/3.0 = %f, type=%s\n", someFloat, ((Object)someFloat).getClass().getSimpleName());

    // double/Double
    var someDouble = 7.0d/3.0d;
    System.out.printf("7.0/3.0 = %f type=%s\n", someDouble, ((Object)someDouble).getClass().getSimpleName());

    // boolean/Boolean
    var someBoolean = true;
    System.out.printf("boolean = %b, type=%s\n", someBoolean, ((Object)someBoolean).getClass().getSimpleName());
  }
}
```

Build and run the java code

```bash
javac -cp . some/folder/Main.java

java -cp . some.folder.Main

string = golang, type=String
bytes = [B@65ab7765, type=byte[]
char = A, type=Character
1+1 = 2, type=Integer
1+1 = 2, type=Integer
1+1 = 2, type=Long
7.0/3.0 = 2.333333, type=Float
7.0/3.0 = 2.333333 type=Double
boolean = true, type=Boolean
```

:::info[GO]
Value types
:::

Go has various value types including strings, integers, floats, booleans.

```go
package main

import "fmt"

func main() {

    fmt.Println("go" + "lang")

    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}
```

Run the Go code

```bash
go run values.go
golang
1+1 = 2
7.0/3.0 = 2.3333333333333335
false
true
false
```
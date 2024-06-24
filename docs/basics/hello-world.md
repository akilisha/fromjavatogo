---
sidebar_position: 2
---

# Create Hello World

:::info[JAVA]
Hello world
:::

There is so much ceremony involved in writing a _Hello world_ program, it almost feels like I came out of the gym.

Create a project directory

```bash
mkdir projectjava
cd projectjava
```

Create a _Main.java_ file in the project's root folder

```bash
touch Main.java
```

Update the java source file

```java
public class Main {

  public static void main(String...args){
    System.out.println("Hello, world");
  }
}
```

Compile the source file

```bash
javac -cp . Main.java
```

Run the Java binary

```bash
java -cp . Main
Hello, world
```

Now move the _Main.java_ file into a package _some/folder_

```java
package some.folder;

public class Main {

  public static void main(String...args){
    System.out.println("Hello, world");
  }
}
```

Rebuild taking into consideration the package

```bash
javac -cp . some/folder/Main.java
```

Rerun again taking consideration the package 

Run the Java binary

```bash
java -cp . some.folder.Main
Hello, world
```

The complexity of using Java's _classpath_ explodes exponentially when dealing with external dependencies, and this is where the use of more specialized tooling becomes absolutely necessary.

> On a more serious note now

Now, any self-respecting java developer would typically not roll out a project and build manually as show above. One will typically use a build toolchain. I will illustrate an example using gradle, which happens to be my favorite in the wild.

0. If you haven't installed java and gradle at this point, pause everything and do that first, before continuing. __hint: use sdkman__

1. Create the project's root folder, observing maven conventions for the soource directory

```bash
mkdir java-app/src/main/java/app
cd java-app
```

This command creates a project called _java-app_, then creates the source directory _src/main/java_, then creates a source package _app_, and then navigates to the _project's root folder_.

2. Create a settings file and add the project name 

```bash
rootProject.name = 'java-app' > settings.gradle
```

Any intelligent Java IDE will detect you have created a gradle project at this point, and it will configure itself to assist you better.

3. Create a build file and configure a java plugin, which will contains tasks that are necessary for managing the project, like compiling, bundling, testing, cleaning etc.

```bash
cat > build.gradle << EOF
plugins {
    id 'java'
}
EOF
```

Any intelligent Java IDE will detect the java plugin, and it will reveal the tasks available in the plugin for you to use.

4. Configure _gradlew_ so that running any gradle tasks in the project will execute uniformly regardless of underlying platform

```bash
gradle wrapper
```

5. Add a java class into the _app_ package (inside _src/main/java/app_)

```java
package app;

public class HelloWorld{

    public static void main(String[] args) {
        System.out.println("Hello World!");
    }
}
```

6. Build the project using gradle's _clean_ and the _build_ tasks

```bash
./gradlew clean build
```

7. Finally, you can execute the HelloWorld class in the bundled jar file

```bash
java -cp build/libs/java-app.jar app.HelloWorld
Hello World!
```

8. You could alternatively skip all the steps discussed above and use gradle's project generation cli command instead, and then follow the prompts generated. This method however, requires a bit more finess with gradle, so it might not necessarily be the best option for a complete beginner.

```bash
mkdir java-app
cd java-app
gradle init
```

With gradle, the complexity of managing of a project's build lifecycle is greatly diminished, even as the project continues to grow bigger in terms of classes or modules (sub-projects), or as the project keeps racking on more dependencies.

:::info[GO]
Hello world
:::

Create a project folder

```bash
mkdir projectgo
cd projectgo
```

Initialize a new project module

```bash
go mod init example/hello

# go: creating new go.mod: module example/hello
```

The go mod init command creates a go.mod file to track your code's dependencies. Now create a _hello.go_ file in the _main_
package and having a _main_ method, in the project's root directory

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Run the code to see it in action

```bash
go run hello.go

# Hello, World!
```

Now create another file _greetings.go_ in the same location (project's root directory) in the package _greetings_ and exporting
the function _Hello_. A function is exported from a package by simply making its name begin with an _uppercase_ letter.

```go
package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
```

Update the _hello.go_ file to make use of this new functionality. The idea is to import the _greetings_ package from the "example/hello"
module

```go
package main

import (
	greetings "example/hello"
	"fmt"
)

func main(){
  fmt.Println("Hello, world")
  fmt.Println(greetings.Hello("Jim"))
}
```

Rerun the main method again. This time, there is a problem,

```bash
go run hello.go 
# hello.go:4:2: found packages greetings (greetings.go) and main (hello.go) in <project_root_dir>
```

These two file define two different packages within the same folder, which Go doesn't like. Move greetings into its own folder
and update the main function to reflect the change made by moving the greetings file

```go
package main

import (
	"example/hello/greetings"
	"fmt"
)

func main(){
  fmt.Println("Hello, world")
  fmt.Println(greetings.Hello("Jim"))
}
```

To further examine the package behavior of Go, create two more file

```go
# example/hello/other/gopher/greet/hello.go

package greet

import "fmt"

func Hello(){
  fmt.Println("Hello, world")
}
```

```go
# example/hello/other/gopher/app/app.go

package main

import (
	"example/hello/greetings"
    "example/hello/other/gopher/greet"
	"fmt"
)

func main(){
  fmt.Println(greetings.Hello("Bosco"))
  greet.Hello()
}
```

Running the app.go file should resolves all dependencies and run successfully

```bash
go run other/gopher/app/app.go 

Hi, Bosco. Welcome!
Hello, world
```

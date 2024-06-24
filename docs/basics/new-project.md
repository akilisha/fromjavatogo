---
sidebar_position: 1
---

# Create a new project

:::info[JAVA]
  New project
:::

In Java, there is no strict project directory layout that is mandated. The only exception is with packages, which are Java's
native way of creating namespaces for source files and compiled artifacts by leveraging the computer's directory structure. 
However, there are some commonly used conventions, and following their wide adoption in the Java universe, these have now 
transformed to become the norm. All these conventions, at the end of the day, come down to handling two key objectives.

- keep testing code, and its associated dependencies, cleanly seperated from the production code and its dependencies for
compilation purposes
- keep both testing and production code is "well-known" classpaths. A classpath is simply a directory location where the 
Java runtime will go to look for java binary file to execute.

You can definitely choose to compile your Java source code directly from the command line, but then you would need to be very 
familiar with the Java compiler options. This is however not necessary because there already exists very powerful tools that 
will do all this and even much more on your behalf. 

In the beginning, the _Ant_ tool was the biggest hammer in the Java universe for compiling and running java programs, and it 
used to be wielded in all sorts of unsavoury and unconventional manner. From the ashes of the carnage wrought on planet Java 
by Ant came forth a savior-figure called _Maven_. Maven was able to reign in all the untethered power that came with Ant, and 
slowly but surely chocked the life out of the once mighty Ant. _Oh, how the mighty have fallen_. 

But Maven quickly became un-scalable due to its very rigid XML structure that did not enable expressiveness during the build 
cycle. However, _the rule of law_ that Maven had established in the Java universe won many hearts and admirers. More affectionately 
known as _convention over configuration_, this ushered in an era of returning to normalcy in managing build artifacts. In parallel 
though, the hunt was on once again for the next better incarnation of Maven, and alas, Gradle was born from the coming together of 
all these bright ideas. Gradle had learnt well from the success of Maven, and used everything that made Maven the peoples' darling, 
but in addition, Gradle replaced the rigid XML with a shinny DSL which enabled expressiveness that was on the same level as full
programming capabilities. _Thank you Groovy and Kotlin_.

So today, in the Java universe, you probably will need to install one of the existing build tools (maven or gradle) and then use 
their specific APIs to create and manage the build lifecycle of an application.

:::info[GO]
New project
:::

Go organizes its code through the use of packages and modules. Java presently also has a module system that was shoehorned into the
platform with Java 9, but which didn't gain much traction thereafter because of the sheer massive amount of Java code already in 
existence that was not designed with Java's module system in mind. Since the module system is radically different from the 
non-modular, classpath-based system, the switch from classpath to modulepath is neither a trivial nor smooth endeavor. The Java 
module system however, was a valiant effort by Java designers to add another dimension of scoping (beyond just public, private and 
default scopes), to control visibility of packages, classes, interfaces, methods and jars during the bundling lifecycle phase, 
to reign in duplication of dependencies, to eliminate code bloat (unused code), to cut down on the size of the generated artifacts, 
and most importantly, to put to rest the infamous classpath hell.

In Go, the standard way to fetch, build, and install Go modules, packages, and commands is accomplished using the same _Go tool_.
Go programs are organized into packages as well, but forget about Java's notions of public, private or default visibility scopes. 
These have no semblance in Go. 

A package is a collection of Go source files in the same directory and have the same package statement at the beginning, that are 
compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files 
within the same package.

A module is a logical namespace for Go source files that is defined using a _.mod_ (module) file, and is used to group together a 
collection of packages into a distributable artifact

A repository contains one or more modules. A module is a collection of related Go packages that are released together. A Go 
repository typically contains only one module, located at the root of the repository. A file named go.mod there declares the 
module path: the import path prefix for all packages within the module. The module contains the packages in the directory containing 
its go.mod file as well as subdirectories of that directory, up to the next subdirectory containing another go.mod file (if any). 

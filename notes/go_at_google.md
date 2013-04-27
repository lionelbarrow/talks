---
author: Rob Pike
date: 2013-4-13
---

# What is Go?

* Open source
* Concurrent
* Garbage collected
* Efficient
* Scalable
* Simple
* Fun
* Boring

Invented to solve Google's software engineering problems. A major one is dependencies. Builds at Google are run on a massive superbuild system but still take forever.

Can the language help?

### What would Go need to look like?

Must work at scale, with large teams, many dependencies. Also should probably be C-like.

But it should be modern. Web-oriented, concurrent, networked.

# Dependencies in Go

Defined syntactically. 'import "encoding/json"' Importing an unused dependency is *illegal*.

Import compilation is a linear process thanks to the way the compiler works. 5 import statements in code means 5 object files are read, not 5 + all of their deps + all of their deps, etc.

Huge difference in compilation fan-out. Go has 40x fanout (expected to decrease) vs 2000x fanout in C++. Additionally, export data in Go is binary -- it doesn't have to be parsed.

No circular imports. Annoying sometimes for developer, but makes things much easier for dependency manager. Type system makes this less annoying than it might be in C.

# Syntax

Go has a very small syntax. 25 keyword, almost regular, easy to predict and reason about.

Distinguishes between declaration and assignment:

var buf bytes.Buffer = bytes.NewBuffer(x)
buf := bytes.NewBuffer(x)

### Functions vs methods

Function on type T:

func Abs(t T) float64

Method of type T:

func (t T) Abs() float64

Closure of type T:

negAbs := func (t T) float64 { return -Abs(t) }

### No default arguments

Why not? Too easy to throw in args to fix design problems, encourages having too many args. That said, Go has built-in support for variadics.

Forcing the programmer to think about interfaces is super important.

### Naming

Uppercase names are exported. Lowercase names are not-exported. (Uppercase and lowercase by unicode.)

Scope hierarchy is simple: universe, package, file, function, block. Boom. Locality of naming is very important. No implicit this.

This has an important consequence: if you import a package, adding names to the package cannot affect your code.

Function lookup is done by name, not type. No resolution of lookup assisted by type. Names must be ambiguous.

# Compare with C/C++

Basically familiar. But many changes. Big ones: concurrency, GC, interfaces.

### Concurrency

Communicating Sequential Processes-style channels. Why CSP? The language should be ordinary and familiarity. Makes sense with a C-style PL.

Go is *not* memory safe. It makes it easy to write safe, concurrent code. It does not stop you from writing bad, unsafe code.

### Garbage collection

Not having GC makes for more leaky abstractions. Memory management happens all the time in non-GC languages. It's a constant source of noise.

Go does give some control of memory to programmer. One big decision: allow interior pointers. Allows for a smaller footprint than e.g. Java.

The implementation remains an active topic. Current algorithm is parallel mark-and-sweep.

### Interfaces

Composition, not inheritance. Go does not have classes or inheritance.

OO is important because it provides a uniformity of interface. Inheritance can be bad when it encourages a non-uniformity of inheritance. You tend to *extend* your parent class.

Design by inheritance is oversold. It generates brittle code, forces early decisions that are hard to change, makes every programmer an interface designer.

Good inheritance programmers see this problem coming, so they shovel in layers and layers of abstraction up front. This is a *smell*.

Go has interfaces. All they are is a set of methods. There is *no* interface implementation declaration.

Example: Reader and Writer interfaces. Allow for many many compositions: files, buffers, networks, encryptions, compression, etc.

Types are composed in a non-hierarchical way, which makes further growth in the system more natural.

Side effect: this lead to function-oriented programs that take interfaces as arguments. This leads to programs that are *nothing like* hierarchical, inheritance-oriented programs.

Good examples:

ReadAll(r io.Reader), LoggingReader(r io.Reader), ErrorInjector(r io.Reader)

# Errors

No try-catch-finally. Instead, there's the error interface: type error interface { Error() string }

Philosophy: errors are not exceptional. They are normal. Treat errors as values and then compute based on them, rather than treating them as unusual.

Makes code more verbose. For sure. But also more robust and graceful.

OK, not all errors are normal. There is a panic() and recover() construct. Their use is not idiomatic except in extreme circumstances.

# Tools

From day 1, Go was built to help make tools easy to help. Std library comes with lexer and parser. Type checker will be included soon.

Gofmt is awesome. Lexes the file and reprints it with idiomatic formatting. Instant community uniformity. It's great.

Gofmt enabled other tools! Go fix, go doc, go build, etc. API checking tools.

Go fix -- automatic refactoring. Pretty cool. Allows you to always have current libraries. Decreased feeling of backwards-compatible weight on the language.

# Conclusion

Go is catching on internally at Google.

* youtube.
* dl.google.com
* golang.org

Externally:

* Heroku
* BBC
* Canonical
* Nokia
* SoundCloud

Still too early to tell if successful or not. Some warts already emerging:

* Declarations can be fussy
* nil is overloaded

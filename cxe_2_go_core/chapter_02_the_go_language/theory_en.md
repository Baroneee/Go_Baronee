## TECHNICAL CONCEPTS MENTIONED

- Build time
- Statically typed language
- Concurrency
- Garbage collector
- Software dependency

## ORIGINS

- There is a legend surrounding the birth of Go. The language was born inside a Google office, emerging from a very long compilation process that took up to 45 minutes.
- This story is told by Rob Pike in [@go-at-google]. It provides us with valuable insight into the motivations behind the creation of Go. Excessively long build times caused headaches... they were forced to find a way to avoid that; that was the starting point of Go.
- Robert Griesemer, Ken Thompson, and Rob Pike are the programmers who began working on Go in 2007. Rob Pike stated that by mid-2008, the language had been “mostly designed and the implementation (compiler, runtime environment) started working.” Later, Ian Lance Taylor and Russ Cox joined the team in 2008 [@pike2009go].
- Go is an open-source programming language maintained by the community and a core team of developers working at Google. March 16, 2011, marked the release of the first Go version (named “r56”). Go version 1 was released on March 28, 2012.

## MOTIVATIONS

- Go (or Golang) was built by Google to solve large-company problems.
- What are the software challenges at large global companies?
  - The codebase of Google's services is massive. Google has millions of lines of code.
  - Those lines of code are written in different languages: C, C++, Java, and others.
  - The build times of those applications “stretched to many minutes, even hours.”
  - Updating certain parts of an application can be very costly.
- The goal of the early Gophers was to make developers' lives easier by:
  - Significantly reducing the build time of programs.
  - Designing a language that is easy to learn, read, and debug for young developers who have been exposed to C, C++, or Java.
  - Designing an efficient dependency management system.
  - Building a language that can produce software that scales well on hardware.

### DEFINITIONS OF SOME CONCEPTS

- Build time: the amount of time required by the compiler to generate a machine-readable executable file.
- Statically Typed Language: Providing an exact definition of this concept at this point is still too early.
- Dependency: a piece of software used by another piece of software.
- Scalability: the ability of a program to handle an increasing amount of tasks to be performed. For example, a website is considered scalable if it can accept an increasing number of requests without crashing (downtime) or increasing load latency.

## CORE FEATURES OF GO

- The creators of Go focused their efforts on several key design choices:
  - Being a compiled language.
  - Having easy-to-understand and easy-to-learn semantics.
  - Statically typed.
  - Having built-in concurrency, a system that makes it easier for developers to work.
  - Robust dependency management.
  - Having a garbage collector.
- The main goal, as Rob Pike presented, is to provide developers with an easy-to-learn language for “engineering large software projects.”

### SOME CONCEPTS

- Concurrency: A program has concurrency when tasks can be executed out of order or in a partial order.
- Garbage collector (usually called GC): When building programs, we need to store data and retrieve data from memory. Memory is not an infinite resource. Therefore, the programmer must ensure that unused elements stored in memory are removed over time. Placing data into memory is called allocation; the opposite action, which involves removing data from memory, is called deallocation. The role of the garbage collector is to reclaim memory when it is no longer in use. When a language does not have a garbage collector, programmers must manage their own garbage and free up memory that is no longer used... Fortunately, Go has a garbage collector.

## SELF-CHECK

- What does concurrency mean?
  - It is when the tasks of a program can be executed at the same time without order or in a partial order.
- On average, does Go have a very long build time? True or False?
  - False, Go was created to solve precisely this problem!

## CORE POINTS TO REMEMBER

- Go was created in 2007.
- Go version 1 was released in 2012.
- The language is easy to understand. Its semantics remain simple.
- It is statically typed.
- It is a compiled language.
- You can write concurrent programs with Go.

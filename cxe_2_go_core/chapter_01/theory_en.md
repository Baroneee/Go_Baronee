## HARDWARE COMPONENTS

- A computer consists of 4 main components: + Memory Unit (MU): stores data and programs. + Arithmetic and Logic Unit (ALU): Performs arithmetic and logical operations on data stored in the MU.
  - Input and Output Unit (I/O): Responsible for loading data into the MU from an input device. And it also sends data from the memory unit to an output device + Control Unit (CU): Receives instructions from a program and controls the operations of other units.
    -> These 4 components represent a structural block diagram of components in a computer.

## MEMORY

- A computer contains 2 types of memory:
  - Main memory
  - Secondary memory
- There are 2 categories of memory:
  - Volatile
  - Non-volatile

### MAIN MEMORY

- Contains 2 types of storage:
  - Random Access Memory (RAM): Requires electrical power to maintain data. When the computer is turned off, data in this memory will be erased. Operating systems and active programs will be loaded into this memory. This type of memory has a Volatile nature.
  - Read-Only Memory (ROM): This is the memory containing essential data for the computer to run correctly. This type of memory has a Non-volatile nature (When the computer is turned off data will not be erased). It is designed to be read-only and the system cannot update it.

### SECONDARY MEMORY

- This type of memory has a Non-volatile nature. When power is lost, stored data will not be erased. Examples: USB, Hard drives, CD-ROM, DVD,...
- Read and write operations for this type of memory are slower compared to RAM.
- Some hard disk drives access memory sequentially -> The system must follow a specific sequence. This takes more time compared to random access mode. Note that some hard disk drives still allow random access.

#### HARD DISK DRIVES

- Hard disk drive or mechanical hard drive (HDD), consists of rotating magnetic platters. Data is read and written via a moving magnetic read-write head. Read and write operations will create rotation and movement of the magnetic head -> thereby consuming time.
- Solid state drive (SSD) is not built like that. There is no magnetic head or magnetic platter. Instead, data is stored in flash memory cells. Accessing data will be faster on this type of drive.

## CPU

- CPU stands for Central Processing Unit. CPU is also called a processor. CPU includes:
  - ALU
  - CU
- CPU is responsible for executing instructions given by a program. For example, a program may request an addition between two numbers. Those numbers will be retrieved from the memory unit and passed to the ALU. A program may also request an input/output operation such as reading data from a hard drive and loading it into RAM for further processing. The CPU will execute those instructions.
- CPU is the central component of a computer.

## WHAT IS A PROGRAM?

- To make a computer perform something, we must provide them with precise instructions. A set of these instructions is called a "program".
- According to a more formal definition, a program is "a combination of computer instructions and data definitions that enable computer hardware to perform computations".
- We will give instructions in human language BUT computers do not understand human sentences. These sentences need to be translated into a language that the machine understands -> What is that language?

## HOW DO WE COMMUNICATE WITH COMPUTERS?

### PROGRAMMING LANGUAGES ARE FORMAL LANGUAGES

- Instructions given to computers are written in programming languages. Programming languages are formal languages. They consist of words constructed from an alphabet (a set of distinct characters). Those words are organized following specific rules. Go is a programming language.
- There are 2 types of programming languages:
  - Low level
  - High level
- Low-level programming languages are closer to processor instructions. High-level languages provide structures that make them easier to learn and easier to use in daily work.
- Some high-level languages are compiled, others are interpreted, and some fall in between. We will see what these two terms mean in the following sections.

### MACHINE LANGUAGE

- To communicate with the computer's processing unit, we can use machine language. Machine language consists entirely of 0s and 1s. An instruction written in machine language is a sequence of 0s and 1s. Each processor (or processor family) will define a list of instructions called an instruction set. There are instructions to add a number, increment by one unit, decrement by one unit, copy data from one location in memory to another...
- Writing computer programs directly in machine language is possible. However, this is not easy.

### ASSEMBLY LANGUAGE

- Assembly language is a low-level programming language. Instructions of a program written in assembly correspond to machine instructions. Assembly language uses short symbolic words (mnemonics) corresponding to a machine instruction. For example, MOV will instruct the computer to move data from one location to another. Developers can also comment in the source code (which is impossible to do with machine language).
- To create a program using assembly language, programmers will write instructions into one or more files. These files are called source files.
- Below is an example of an instruction written in x86 Linux assembly:
  mov eax,1
  int 0x80
- These two lines will execute a system call to close the program (number "1" represents the system call number meaning "exit program"). Note that assembly language varies across different machines. We call it machine-specific.
- An assembler program is used to convert source files written in assembly language into object code files. We call this process assembling the program. A linker will then transform these object code files into an executable file. An executable file contains all necessary computer instructions to launch the program.
  ![Assembly Translation Process](https://www.practical-go-lessons.com/img/assembly_to_executable.926437cf.png)

### HIGH-LEVEL LANGUAGES

- There are many high-level languages on the market, like Go. These languages are not tightly bound to machine architecture. They provide a convenient way to write instructions. For example, if we want to perform a system call to exit the program, in Go we can write:
  os.Exit(1)
- In C language, we can write:
  exit(1)
- In Java, we can write:
  System.exit(1);
- In this example, we do not have to move a number into a register; we use language constructs (functions, packages, methods, variables, data types...). The goal of this book is to provide you with precise and concise definitions of these tools to build Go applications.
- High-level programs are also written into files. These files are called "source files". Generally, programming languages require adding a specific extension to the file name. For Go programs, we will add .go to the end of each file we write. In PHP, the extension is .php.
- When source files are written, the program defined by them cannot be executed immediately. Source files need to be compiled using a compiler. A compiler will transform source files into an executable file. A compiler is also a program. Go is part of the compiled language family.
- Go is a compiled language

  ![Compiler Process](https://www.practical-go-lessons.com/img/compiler.adc1a3b1.png)

#### COMPILED VS INTERPRETED

- Note that some programming languages are interpreted languages. Once source files are written, programmers do not need to compile the source code. With ready source files, the system can execute the program thanks to an interpreter. Each instruction written in the source file is translated and executed by the interpreter. In some cases, interpreters store a compiled version of the program in cache to increase performance (source files are not re-translated every time). PHP, Python, Ruby, Perl are interpreted languages.

## SELF-TEST

### QUESTIONS

##### Where are programs stored?

- In Memory Units (MU)

##### Reading data from a hard drive is slower than reading data from RAM. True or false?

- True because hard drives need time for read heads and magnetic platters to rotate and read, while RAM only needs to fetch from available memory cells.

##### Can you write data to ROM? True or false?

- False because ROM is Read-Only Memory, meaning we can only read data from ROM and cannot write.

##### What are the two types of memory?

- Main memory () and secondary memory

##### What is the definition of "volatile memory"?

- It is when power supply is disconnected, data will be lost.

##### Which program converts code written in assembly language into object code?

- Assembler.

##### Which program converts object code into an executable file?

- Linker

##### State two advantages of high-level languages over low-level languages?

- They provide higher-level constructs that are easier to use.
- Source code will not depend on the technical architecture of a specific machine. We call that portability.

##### Go is an interpreted language? True or false?

- False because Go is a compiled language

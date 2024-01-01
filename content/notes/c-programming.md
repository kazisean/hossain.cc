+++
title = 'C Programming Language'
description = 'My notes on C'
date = 2023-11-15T15:48:50-05:00
draft = false
math = true
ShowToc = true
+++

### junk ideas
C was created by Dennis Ritchie at the [Bell Labs](https://www.bell-labs.com/) along with the unix operating system. Bell labs is mostly famous for inventing information theory, the transistor, the laser, radio astronomy, solar cells, etc. C is a low level programming language mostly used in developing kernals or operating systems. Most of the unix applications are written in C. \

The use of terms such as pritning and what does it mean
Add notes on how to set up and run C programs


### Introduction
Printing hello word in C :

```C
#include <stdio.h>

main (){
    printf("hello, world\n");
}

```
In this case the `\n` is used to create a new line after printing out hello world. Every program must have a main fuction somewhere. In this main is a function that takes in no argument values. Main calls the library function printf with the arguments hello world wich prints the sequence of character hello world. Double quotes are used to signify a character string or string constant.

To compile the following program and run it it will be different for each system. Name the file `hello.c` and run the command `cc hello.c` this will compile and make a file such as `a.out` which would print `hello world` 

C also provides `\t ` for tab, `\b ` for backspace, `\" ` for double quote, and `\\` for backslash. 

When `printf` funtion is passed with an escape sequence that does not exits for example `\v` it will lead to the v printing without the `\`. The output will look something like `v`. 

The C programming Language does not have a heap nor does it have a garbage collector. 

IN C, all veriables must be declared before they are used. The convention is at the beginning of the program. You can add multiple variable names with the same type in a single line by separating them using commas. For example ` int a, b, c` in this all three a, b, c can be used to store int type.

The range of `int` and `float` depends on based on the machine you are using. For 16-bit and 32-bit ints type can be it is between `-32768 to + 32768`. A float is a 32-bit quantity, with at least six digits between $10^{-38}$ to $10^{+38}$

C basic data types are \
• `int` \
• `float` \
• `char` = character containing a single byte\
• `short` = short integer containing two bytes of memory\
• `long` = long integer stores at least 32 bits \
• `double` = precise floating point number 

The size of these data types are machine-dependent. 


# WhatIC

Because command-line interfaces are difficult. Pronounced as "What I See".

[![Build Status](https://travis-ci.com/arcanericky/whatic.svg?branch=master)](https://travis-ci.com/arcanericky/whatic)
[![codecov](https://codecov.io/gh/arcanericky/whatic/branch/master/graph/badge.svg)](https://codecov.io/gh/arcanericky/whatic)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)

## Using

Use it to reveal how applications see command line arguments and options after they've been parsed and transformed by the shell.

```text
$ whatic escaping= \" glob= L* single-quotes= '"'
Command Line
------------
whatic escaping= " glob= LICENSE single-quotes= "

Arguments
---------
0 whatic
1 escaping=
2 "
3 glob=
4 LICENSE
5 single-quotes=
6 "
```

## Purpose

Before [Windows 1.0](https://en.wikipedia.org/wiki/Windows_1.0), the [X Window System](https://en.wikipedia.org/wiki/X_Window_System), [AmigaOS](https://en.wikipedia.org/wiki/AmigaOS), [Lisa OS](https://en.wikipedia.org/wiki/Apple_Lisa#Lisa_OS), and even [Xerox Star](https://en.wikipedia.org/wiki/Xerox_Star), there existed, and always shall exist, the [Command-line Interface](https://en.wikipedia.org/wiki/Command-line_interface) (CLI). At some point in the past, all computers were manipulated through a CLI and in the modern world of the graphical user interface with 64 bazillion colors and the clickety mouse, the use of those CLI manipulation techniques is becoming a lost art.

This very simple application was created for those that find the CLI difficult. All of its arcane incantations using single or double-quotes to isolate arguments, backslashes as an [escape character](https://en.wikipedia.org/wiki/Escape_character), and conjuring up a space character for the middle of an argument can be as mysterious as the code behind the application you are struggling to execute. If you're in this situation, this is the application you need to reveal the wizardry behind the prompt.

This application comes in Linux, Windows, and Darwin (OS X) versions. Download the executable for your platform and use it to discover what the application sees as your arguments after they have been parsed and transformed by your CLI shell. Note that in Windows, the techniques used to quote and escape characters can be different between the [Command Shell](https://en.wikipedia.org/wiki/Cmd.exe) and [PowerShell](https://en.wikipedia.org/wiki/PowerShell). In Linux, these things may also change depending on your shell of choice, whether it be [Bash][2], the old [Bourne shell](https://en.wikipedia.org/wiki/Bourne_shell), or the trendy [fish](https://fishshell.com/). This application will help you discover these differences and practice your CLI-fu.

## Whitespace

A rare edge case is the use of spaces at the beginning or end of an argument or as the whole argument itself. To see these, delimiters are required. Because this application should not accept or interpret any options, enable these by using an environment variable. If the `WHATIC` environment variable is a single character it will be used for both the opening and closing delimiters. If it is two or more characters long the first character will be used for the opening delimiter and the second character for the closing delimiter. In the spirit of WhatIC echoing its input, the value of `WHATIC` will be output.

Some examples:

Linux / Bash

```text
$ WHATIC=\| whatic \ arg1 "   " arg3\ 
Environment
-----------
WHATIC=|

Command Line
------------
|whatic| | arg1| |   | |arg3 |

Arguments
---------
0 |whatic|
1 | arg1|
2 |   |
3 |arg3 |
```

Windows Command Prompt (`cmd.exe`)

```text
C:\>set WHATIC=^|

C:\>whatic.exe arg
Environment
-----------
WHATIC=|

Command Line
------------
|whatic.exe| |arg|

Arguments
---------
0 |whatic.exe|
1 |arg|
```

Windows PowerShell

```text
PS C:\>$Env:WHATIC="|"

PS C:\>.\whatic.exe arg
Environment
-----------
WHATIC=|

Command Line
------------
|C:\whatic.exe| |arg|

Arguments
---------
0 |C:\whatic.exe|
1 |arg|
```

## Examples

The following examples are not full references for the different CLIs. They are only to show how `whatic` can be used to inspect how shells interpret and transform commands. If more information is needed, experiment with `whatic` or search for CLI references.

### Linux / Darwin (OS X)

The [Bash shell][2] is assumed for these examples.

#### Quoting

Parameters can be single or double-quoted.

```text
$ whatic "arg1" 'arg2'
Command Line
------------
whatic arg1 arg2

Arguments
---------
0 whatic
1 arg1
2 arg2
```

#### Escape Characters

To give a single or double-quote as a parameter, use the `\` to escape them

```text
$ whatic \" \' '"'
Command Line
------------
whatic " ' "

Arguments
---------
0 whatic
1 "
2 '
3 "
```

#### Glob Patterns

The shell is responsible for the [glob patterns][1].

```text
$ whatic L* READ??.md
Command Line
------------
whatic LICENSE README.md

Arguments
---------
0 whatic
1 LICENSE
2 README.md
```

### Windows

Because command-line argument parsing is largely left to the individual executables, it can be difficult to determine how to format the arguments for a particular application. While WhatIC won't give you exact answers, it might be useful for figuring out how the underlying application might be twisting your arguments around.

#### Quoting

The Command Prompt (`cmd.exe`) arguments can be double-quoted, but if single-quotes are used, they will be included in the arguments.

```text
C:\>whatic "arg" 'arg'
Command Line
------------
whatic arg 'arg'

Arguments
---------
0 whatic
1 arg
2 'arg'
```

Inserting a double-quote into an argument might take some experimentation. Use the `\` escape character to get that double-quote into the argument. If the double-quote is in the middle of an argument, two double-quotes can work, but be careful as this can surprise you if more parameters are used.

```text
C:\>whatic \" ""arg1"" "arg2""arg3" arg4
Command Line
------------
whatic " arg1 arg2"arg3 arg4

Arguments
---------
0 whatic
1 "
2 arg1
3 arg2"arg3 arg4
```

PowerShell can also be tricky at times. In general the backtick (\` / [grave accent](https://en.wikipedia.org/wiki/Grave_accent)) is used as the escape character. Early versions could use the backtick to escape the double-quote but recent versions don't. Use WhatIC to help refine these character combinations. Some examples are

```text
PS C:\WhatIC> whatic arg1 'arg\"2' arg\`"3 arg4
Command Line
------------
C:\WhatIC\whatic.exe arg1 arg"2 arg"3 arg4

Arguments
---------
0 C:\WhatIC\whatic.exe
1 arg1
2 arg"2
3 arg"3
4 arg4
```

#### Escape Characters

In the Command Prompt (`cmd.exe`), the caret (`^`) is used to escape the `&`, `\`, `<`, `>`, `^`, and `|` symbols

```text
C:\>whatic ^& ^\ ^< ^> ^^ ^|
Command Line
------------
whatic & \ < > ^ |

Arguments
---------
0 whatic
1 &
2 \
3 <
4 >
5 ^
6 |
```

#### Glob Patterns

Notice the Command Prompt (`cmd.exe`) and PowerShell do not handle the [glob patterns][1] like a POSIX shell does. That task is left to the executable being called.

```text
C:\>whatic L*
Command Line
------------
whatic L*

Arguments
---------
0 whatic
1 L*
```

---

## Inspiration

I write a lot of command line applications and utilities. It's not unusual to have bugs filed on my code or get requests for help on how to call my applications because of unfamiliarity with the CLI. In today's world of the GUI this is understandable, but it can be fixed. This little application is my contribution to help people learn how to manipulate and bend the command line and eventually become CLI wizards themselves.

[1]: https://en.wikipedia.org/wiki/Glob_(programming)
[2]: https://en.wikipedia.org/wiki/Bash_(Unix_shell)

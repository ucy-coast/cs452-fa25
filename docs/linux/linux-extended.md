# Introduction to Linux

## Table of Contents
1. [Introduction](#introduction)
   - [Overview](#overview)
   - [History](#history)
   - [Why Linux?](#why-linux)
   - [Files and Processes](#files-and-processes)
1. [Shells](#shells)
   - [Overview](#overview-1)
   - [Exercise: Working with the Shell](#exercise-working-with-the-shell)
   - [Man Pages](#man-pages-short-for-manual-pages)
   - [Directory Navigation](#directory-navigation)
   - [Interacting with Files and Directories](#interacting-with-files-and-directories)
   - [I/O and Redirection](#io-and-redirection)
   - [Exercise: Redirection](#exercise-redirection)
   - [Searching](#searching)
   - [Exercise: Pipes, Sorting, and Counting](#exercises-with-piping-sorting-and-counting)
   - [Job Control](linux-shells.md#job-control)
   - [Environment Variables](linux-shells.md#environment-variables)
1. [Text Editors](#text-editors)
   - [Overview](#overview-2)
   - [Vim](#vim)
   - [Emacs](#emacs)
1. [Accounts](#accounts)
1. [Remote Connections](#remote-connections)
1. [Filesystem](#filesystem)
1. [File Permissions](#file-permissions)
1. [Optional Topics](#optional-topics)


## Introduction

### Overview 

Linux is not a single operating system, but rather a large family of [free and open source](https://en.wikipedia.org/wiki/Free_and_open-source_software) operating systems based on the [Linux kernel](https://en.wikipedia.org/wiki/Linux_kernel). Different variants within this family are referred to as [Linux Distributions](https://en.wikipedia.org/wiki/Linux_distribution). There can be many subtle differences between distributions, but the core functionality of the operating system is the same. Therefore, for simplicity, we will not work with a specific distribution or version in this tutorial. The information, examples, and workflows described here are intended to be generically useful for most (or all) distributions of Linux. There will also be some overlap with the Unix operating system. If you are interested in more information on specific distributions, please see the Linux References.

The Linux operating system is an extremely versatile Unix-like operating system, and has taken a clear lead in the High Performance Computing (HPC) and scientific computing community. Linux is a multi-user, preemptive, multitasking operating system that provides a number of facilities including management of hardware resources, directories, and file systems, as well as the loading and execution of programs. A vast number of utilities and libraries have been developed (mostly free and open source as well) to accompany or extend Linux.

There are two major components of Linux, the kernel and the shell:

1. The **kernel** is the core of the Linux operating system that schedules processes and interfaces directly with the hardware. It manages system and user I/O, processes, devices, files, and memory.
2. The **shell** is a text-only interface to the kernel. Users input commands through the shell, and the kernel receives the tasks from the shell and performs them. The shell tends to do four jobs repeatedly: display a prompt, read a command, process the given command, then execute the command. After which it starts the process all over again.

It is important to note that users of a Linux system typcially do not interact with the kernel directly. Rather, most user interaction is done through the shell or a desktop environment.

[Back to Top](#table-of-contents)


### History

#### Unix

The Unix operating system got its start in 1969 at Bell Laboratories and was written in assembly language. In 1973, Ken Thompson and Dennis Ritchie succeeded in rewriting Unix in their new language C. This was quite an audacious move; at the time, system programming was done in assembly in order to extract maximum performance from the hardware. The concept of a portable operating system was barely a gleam in anyone's eye.

The creation of a portable operating system was very significant in the computing industry, but then came the problem of licensing each type of Unix. Richard Stallman, an American software freedom activist and programmer recognized a need for open source solutions and launched the GNU project in 1983, later founding the [Free Software Foundation](https://www.fsf.org/). His goal was to create a completely free and open source operating system that was Unix-compatible or [Unix-like](https://en.wikipedia.org/wiki/Unix-like).

#### Linux

In 1987, the source code to a minimalistic Unix-like operating system called [MINIX](https://en.wikipedia.org/wiki/MINIX) was released by Andrew Tanenbaum, a professor at Vrije Universiteit, for academic purposes. Linus Torvalds began developing a new operating system [based on MINIX](https://groups.google.com/forum/#!msg/comp.os.minix/dlNtH7RRrGA/SwRavCzVE7gJ) while a student at the University of Helsinki in 1991. In September of 1991, Torvalds released the first version (0.1) of the Linux kernel.

Torvalds greatly enhanced the open source community by releasing the Linux kernel licensed under the [GNU General Public License](https://en.wikipedia.org/wiki/GNU_General_Public_License) so that everyone has access to the source code and can freely make modifications to it. Many components from the GNU project, such as the [GNU Core Utilities](https://en.wikipedia.org/wiki/GNU_Core_Utilities), were then integrated with the Linux kernel, thus completing the first free and open source operating system.

The adoption of Linux has occurred on a variety of different computer systems of many sizes and purposes. Furthermore, different variants of Linux (called Linux distributions) have been [developed over time](https://en.wikipedia.org/wiki/Linux_distribution#History) to meet various needs. There are now hundreds of different Linux distributions available, with a wide variety of features. The most popular operating system in the world is actually Android, which is built on the Linux kernel.

[Back to Top](#table-of-contents)


### Why Linux?

Linux has been so heavily utilized in the HPC and scientific computing community that it has become the standard in many areas of academic and scientific research, particularly those requiring HPC. There have been over 40 years of development in Unix and Linux, with many academic, scientific, and system tools. In fact, as of November 2017, all of the [TOP500](https://www.top500.org/statistics/list/) supercomputers in the world run Linux!

Linux has four essential properties which make it an excellent operating system for the science community:

- **Performance** – Performance of the operating system can be optimized for specific tasks such as running small portable devices or large supercomputers.
- **Functionality** – A number of community-driven scientific applications and libraries have been developed under Linux such as molecular dynamics, linear algebra, and fast-Fourier transforms.
- **Flexibility** – The system is flexible enough to allow users to build applications with a wide array of support tools such as compilers, scientific libraries, debuggers, and network monitors.
- **Portability** – The operating system, utilities, and libraries have been ported to a wide variety of devices including desktops, clusters, supercomputers, mainframes, embedded systems, and smart phones.

[Back to Top](#table-of-contents)


### Files and Processes

Everything in Linux is considered to be either a file or a process:

- A **process** is an executing program identified by a unique process identifier, called a PID. Processes may be short in duration, such as a process that prints a file to the screen, or they may run indefinitely, such as a monitor program.
- A **file** is a collection of data, with a location in the file system called a path. Paths will typically be a series of words (directory names) separated by forward slashes, /. Files are generally created by users via text editors, compilers, or other means.
- A **directory** is a special type of file. Linux uses a directory to hold information about other files, the equivalent of a folder in Windows. You can think of a directory as a container that holds other files or directories.

A file is typically stored on physical storage media such as a disk (hard drive, flash disk, etc.). Every file must have a name because the operating system identifies files by their name. File names may contain any characters, although some special characters (such as spaces, quotes, and parenthesis) can make it difficult to access the file, so you should avoid them in filenames. File names can be as long as 255 characters, so it is convenient to use descriptive names.

Files can hold any sequence of bytes; it is up to the user to choose the appropriate application to correctly interpret the contents in a file. Files can be human readable text organized line by line, a structured sequence only readable by a specific application, or a machine-readable sequence byte by byte. Many programs interpret the contents of a file as having some special structure, such as a pdf or postscript file. In scientific computing, binary files are often used for efficiency in storage and data access. Some other examples include scientific data formats like NetCDF or HDF which have specific formats and provide application programming interfaces (APIs) for reading and writing.

The Linux kernel is responsible for organizing processes and interacting with files; it allocates time and memory to each process and handles the file system and communications in response to system calls. The Linux system uses files to represent everything in the system: devices, internals to the kernel, configurations, etc.

[Back to Top](#table-of-contents)

## Shells

### Overview

A variety of different shells are available for Linux and Unix, each with pros and cons. While **bash** (updated version of sh) and **tcsh** (descended from C-shell/csh) are the most common shells, the choice of shell is entirely up to user preference and availability on the system. In most Linux distributions, bash is the default shell.

#### Purpose
The purpose of a shell is to interpret commands for the Operating System (OS) to execute. Since bash and other shells are scripting languages, a shell can also be used for programming via scripts. The shell is an interactive and customizable environment for the user.

#### Usage
All examples in this tutorial use the **bash shell**. The shell prompt starts with the character `$`. Do not type the `$` when attempting to reproduce the examples.

The shell input to run a command typically follows a certain format:

```
$ <command> <option(s)> <argument(s)>
```

1. **command** – the executable (program or package) that is to be run.
    - If you are running your own application, you must include either the full path or the relative path as part of the command.
    - Most commands that come packaged with the OS or are installed by the package manager (executables often located in `/bin` or `/usr/bin`) do not need the path because they have already been added to the environment variable `$PATH`.
2. **option(s)** – (A.K.A. "flags") optional arguments for the command that alter the behavior.
    - Start with a `-` or `--` (example: `-h` or `--help` for help).
    - Each command may have different options or no options at all.
    - Some options require an argument immediately following.
    - Explore options for commands by reading the Manual Pages.
3. **argument(s)** – depend on the command and the flags selected.
    - Certain flags require an argument.
    - Filename arguments must include a path unless located in the current directory.

#### Basic Examples
To determine which shell you are currently using, you can type the `echo` command followed by the system environment variable `$SHELL` as follows:

```
$ echo $SHELL
/bin/bash
```

Here, echo is the command entered through the shell, and $SHELL is a command argument. The output is showing the location of the installed bash.

#### Tab Completion
The bash shell includes an incredibly useful feature called [tab completion](https://en.wikipedia.org/wiki/Command-line_completion), which enables you to enter part of a command, press the Tab key, and it will complete the command if there is no ambiguity. If there are multiple options, hitting Tab again will display the possible options. Below is an example on Stampede2 where "py" was entered followed by Tab:

```
$ py
pydoc                  pygobject-codegen-2.0  pygtk-demo             pystuck                python2                python2.7-config       python-config
pygmentize             pygtk-codegen-2.0      pyjwt                  python                 python2.7              python2-config         
```

Use tab completion to finish commands, file names, and directory names. Try it out at any point and see how much it simplifies your workflow!

[Back to Top](#table-of-contents)


### Exercise: Working with the Shell

Try these shell commands at the prompt. Many of these commands have extensive additional arguments they can take.

#### Display the $PATH variable
The `$PATH` environment variable stores selected paths to executables; as a result, these executables can be executed without reference to their full paths. Some paths are added to this environment variable at startup, by the system. The user can add additional paths to the environment variable. Executables in directories included in $PATH are often referred to as being "in the path" of the current shell.

```
$ echo $PATH
/usr/local/bin:/usr/bin:/bin 
```

#### List the available shells in the system
The `cat` (concatenate) command is a standard Linux utilities that concatenates and prints the content of a file to standard output (shell output). In this case, *shells* is the name of the file, and */etc/* is the pathname of the directory where this file is stored.

```
$ cat /etc/shells
/bin/sh
/bin/bash
/sbin/nologin
/usr/bin/sh
/usr/bin/bash
/usr/sbin/nologin
/bin/tcsh
/bin/csh
/usr/bin/tmux
/bin/ksh
/bin/rksh
/bin/zsh
```

#### Find the current date and time of the system
Use the date command.

```
$ date
Fri Nov  9 19:23:23 CST 2018
```

#### List all of your own current running processes
Use the `ps` command (process status). In Linux, each process is associated with a PID or process *identification*.

```
$ ps
  PID TTY          TIME CMD
  916 pts/58   00:00:00 bash
 1531 pts/58   00:00:00 ps
```

[Back to Top](#table-of-contents)


### Man Pages (short for Manual Pages)

The easiest way to get more information on a particular Linux command or program is to use the man command followed by the item you want information on:

```
man <program or command>
```

This will bring up the "man pages" for the program within the shell, which have been formatted from the online man pages. These pages can be referenced from any Linux or Unix shell where man is installed, which is most systems. Linux includes a built-in manual for nearly all commands, so these should be your go-to reference.

The manual is divided into a number of sections by type of topic, for example:

Section|	Description
---|---
1|	Executable programs and shell commands
2|	System calls (functions provided by the kernel)
3|	Library calls (functions within program libraries)
4|	Special files
5|	File formats and conventions
6|	Games
7|	Miscellaneous (including macro packages and conventions)
8|	System administration commands (usually only for root)

If you specify a specific section when you issue the command, only that section of the manual will be displayed. For example, `man 2 mkdir` will display the Section 2 man page for the `mkdir` command. Section 1 for any command is displayed by default.

If your terminal does not support scrolling with the mouse, you can navigate the man pages by using the arrow keys to scroll up and down or by using the enter key to advance a line and the space bar to advance a page. Use the `q` key to quit out of the display.

The man pages follow a common layout. Within a man page, sections may include the following topics:

1. **NAME** – a one-line description of what it does.
2. **SYNOPSIS** – basic syntax for the command line.
3. **DESCRIPTION** – describes the program's functionalities.
4. **OPTIONS** – lists command line options that are available for this program.
5. **EXAMPLES** – examples of some of the options available
6. **SEE ALSO** – list of related commands.

Example snippets from the man page for the rm (Remove) command:

```
$ man rm
RM(1)                            User Commands                           RM(1)
    
NAME
    rm - remove files or directories
    
SYNOPSIS
    rm [OPTION]... FILE...
    
DESCRIPTION
    This  man page documents the GNU version of rm.  rm removes each 
    specified file. By default, it does not remove directories.

     If the -I or --interactive=once option is given,  and  there  are  more
     than  three  files  or  the  -r,  -R, or --recursive are given, then rm
     prompts the user for whether to proceed with the entire operation.   If
     the response is not affirmative, the entire command is aborted.
```

Depending on the command, the OPTIONS section can be quite lengthy:

```
OPTIONS
    Remove (unlink) the FILE(s).

       -f, --force
              ignore nonexistent files, never prompt

       -i     prompt before every removal

       -r, -R, --recursive
              remove directories and their contents recursively

       -v, --verbose
              explain what is being done
```

**Fun fact**: there is even a manual entry for the man command. Try:

```
$ man man
```

Issuing the man command with the -k option will print the short man page descriptions for any pages that match the command. For example, if you are wondering if there is a manual entry for the who command:

```
man -k who
```

Since there is a man page listed, you can then display the man page for the `who` command with `man who`.

[Back to Top](#table-of-contents)


### Directory Navigation

#### Directories
In a hierarchical file system like Linux, the **root directory** is the highest directory in the hierarchy, and in Linux this is the / directory. The **home directory** is created for the user, and is typically located at `/home/<username>`. Commonly used shorthands for the home directory are `~` or `$HOME`. The home directory is usually the initial default working directory when you open a shell.

The **absolute path** or **full path** details the entire path through the directory structure to get to a file, starting at `/`. A **relative path** is the path from where you are now (your present working directory) to the file in question. An easy way to know if a path is absolute is to check if it contains the `/` character at the very beginning of the path.

The "`.`" directory is a built-in shortcut for the current directory path and similarly the "`..`" directory is the directory above the current directory. These special shortcuts exist in every directory on the file system, except "`..`" does not exist in the `/` directory because it is at the top. Files that begin with a dot "`.`" (i.e. `.bashrc`) are called **dot files** and are hidden by default during navigation (in the sense that the ls command will not display them), since they are usually used for user preferences or system configuration.

#### Navigating
Here is a list of common commands used for navigating directories:

**`pwd`** – **p**rint **w**orking **d**irectory. Prints the full path to the directory you are in, starting with the root directory. On Stampede2 you might see:

```
$ pwd
/home1/05574/<username>
```

**`ls`** – **l**i**s**ts the contents of a directory.

```
$ ls
test1.txt  test2.txt  test3.txt
```

- Displays the files in the current directory or any directory specified with a path.
- Use the wildcard `*` followed by a file extension to view all files of a specific type (i.e. `ls *.c` to display all C code files).
- Use the `-a` option to display **a**ll files, including dot files.
- There are many options for this command, so be sure to check the man pages. A Stampede2 example:

   ```
   $ ls -lha $SCRATCH
   total 12K
   drwx------ 3 <username> G-819251 4.0K Jul  9 14:57 .
   drwxr-xr-x 5 root       root     4.0K May 14 16:13 ..
   drwx------ 2 <username> G-819251 4.0K Jul  9 14:57 .slurm
   ```

**`cd`** – **c**hange **d**irectory to the directory or path following the command. The following command will take you from your current directory to your home directory on most Linux systems:

```
$ cd ~
```

- This example will take you up one directory, in this case to the root directory / , and then over to the var directory:

    ```
    $ cd ../var
    ```
- With no arguments, cd will take you back to your home directory.

[Back to Top](#table-of-contents)


### Interacting with Files and Directories

Here is a list of common commands used for interacting with files and directories:

**`mkdir`** – **m**a**k**e a new directory of the given name, as permissions allow.

```
$ mkdir Newdir
```

**`mv`** – **m**o**v**e files, directories, or both to a new location.

```
$ mv file1 Newdir
```

- This can also be used to rename files:
   ```
   $ mv file1 file2
   ```
- Use wildcards like `*` to move all files of a specific type to a new location:
   ```
   $ mv *.c ../CodeDir
   ```
- You can always verify that the file was moved with `ls`.

**`cp`** – **c**o**p**y files, directories, or both to a new location.

```
$ cp file1 ~/
```
- You can give the copy a different name than the original in the same command:
   ```
   $ cp file1 file1_copy
   ```
- To copy a directory, use the -r option (recursively). In this case, both the source and the destination are directories, and must already exist.
   ```
   $ cp -r Test1 ~/testresults
   ```

**`rm`** – **r**e**m**oves files or directories ***permanently*** from the system.

```
$ rm file1
```

- **Note**: Linux does not typically have a "trash bin" or equivalent as on other OS, so when you issue rm to remove a file, it is difficult to impossible to recover removed files without resorting to backup restore.
- With the `-r` or `-R` option, it will also remove/delete entire directories recursively and permanently.
   ```
   $ rm -r Junk
   ```
- Avoid using wildcards like *. For instance, the following will remove all of the files and subdirectories within your current directory, so use with caution.
   ```
   rm -r *
   ```
- To remove an empty directory, use `rmdir`

**`touch`** – changes a file's modification timestamp without editing the contents of the file. It is also useful for creating an empty file when the filename given does not exist.

Try these commands to get more familiar with Linux files and directories.

[Back to Top](#table-of-contents)


### I/O and Redirection

#### Input and Output (I/O)

As the title of this section suggests, I/O stands for input/output. Your commands or programs will often have input and/or output. It is important to know how to specify where your input is from, or to redirect where output should go; for example, you may want your output to go to a file rather than printing to the screen. Inputs and outputs of a program are called streams in Linux. There are three types of streams:

- stdin (**st**andar**d** **in**put) – stream data going into a program. By default, this is input from the keyboard.
- stdout (**st**andar**d** **out**put) – output stream where data is written out by a program. By default, this output is sent to the screen.
- stderr (**st**andar**d** **err**or) – another output stream (independent of stdout) where programs output error messages. By default, error output is sent to the screen.

#### Output Redirection
It is often useful to save the output (stdout) from a program to a file. This can be done with the redirection operator `>`.

```
$ example_program > my_output.txt
```

For another example, imagine that you run the ls command on a directory that has so many files that your screen scrolls and you cannot see all of the files listed. You might want to redirect that output to a file so you can open it up in a text editor and look more closely at the output:

```
$ ls > output_file.txt
```

Redirection of this sort will create the named file if it doesn't exist, or else overwrite the existing file of the same name. If you know the file already exists (or even if it does not), you can append the output file instead of rewriting it using the redirection operator >>.

```
$ ls >> output_file
```

#### Input Redirection

Input can also be given to a command from a file instead of typing it in the shell by using the redirection operator `<`.

```
$ mycommand < programinput
```

Alternatively, you can use the "pipe" operator `|` like this:

```
$ cat programinput | mycommand
```

Using the pipe operator `|`, you can link commands together. The *pipe will link stdout from one command to stdin of another command*. In the above example we use the `cat` command to print the file to the screen (stdout), and then we redirect that printing to the command `mycommand`.

#### Error Redirection
When performing normal redirection of the standard output of a program (stdout), stderr will not be redirected because it is a separate stream. Many programmers find it useful to redirect only stderr to a separate file. You might do this to make it easier to find the error messages from your program more easily. Using the shell, this can be accomplished with redirection operator `2>`.

```
$ command 2> my_error_file
```

In addition, you can merge stderr with stdout by using 2>&1.

```
$ command > combined_output_file 2>&1
```

#### Redirect and Save Output
Redirecting the output of a command to a file is useful, but it means that you will not see anything on the screen while it is running. This can be undesirable, especially for long-running commands. To have the output go to both a file and the screen, use the tee command:

```
command | tee outputfile
```

You can also use tee to catch stderr with:

```
command 2>&1 | tee outputfile
```

[Back to Top](#table-of-contents)



### Exercise: Redirection

#### Viewing Output
Use `ls` (list files) and `>` (redirect) to create a file named "mylist" which contains a list of your files.

```
$ ls -l /etc > mylist
```

There are three main methods for viewing a file from the command prompt. Try each of these on your "mylist" file to get a feel for how they work:

1. `cat` shows the contents of the entire file at the terminal, and scrolls automatically to the end.
   ```
   $ cat mylist
   ```
1. `more` shows the contents of the file, pausing when it fills the screen.
   - Note that it reads the entire file before displaying, so it could take a long time to load for large files.
   - Use the spacebar to advance one page at a time.
   ```
   $ more mylist
   ```
1. `less` is similar to `more`, but with more features. It shows the contents of the file, pausing when it fills the screen.
   - Note that `less` is **faster** than `more` on large files because it does not read the entire input file before displaying.
   - Use the spacebar to advance one page at a time, or use the arrow keys to scroll one line at a time. Enter `q` to quit. Entering **`g`** or **`G`** will take you to the beginning or end of the file, respectively.
   - You can also **search** within a file (similar to Vim) by typing / and the word or characters you are searching for (example: **`/foo`** will search for "foo"). `less` will jump to the first match for the word. Move between matches by using n and ? keys.
   ```
   $ less mylist
   ```
**Note**: It may also be useful to explore the man pages for `head` and `tail` and try them out, especially in conjunction with these viewing methods.

#### Combining Redirection and Viewing
Now let's try an exercise where we enter the famous quote "Four score and seven years ago" from Lincoln's Gettysburg address into a file called "lincoln.txt". First, use `cat` to direct stdin to "lincoln.txt":
```
$ cat > lincoln.txt
```

Next, enter the quote above. To end the text input, press **Control-D**.

```
Four score and seven years ago
[Control-D]
```

Finally, you can use `cat` to view the file you just created:

```
$ cat lincoln.txt
Four score and seven years ago
```

Now try adding another line of the famous quote to the existing file:

```
$ cat >> lincoln.txt 
our fathers brought forth on this continent, a new nation
[Control-D]
```
If you wish, you could try appending the rest of the speech to the file. Finally, try viewing the file in both `more` and `less` to test them out. Feel free to test navigation in both and try searching with `less`. If you have a longer file, try viewing that as well so you can get used to scrolling.

#### Creating a Simple Script
We can also redirect input to a script file that we create, and then run the script. First, we will create the script file called "tryme.sh" that contains the cat command without any arguments, forcing it to read from stdin.

```
$ cat > tryme.sh 
#!/bin/sh 
cat 
[Control-D]
```

The first line of the script `#!/bin/sh` indicates which shell interpreter to use. `/bin/sh` is a special sort of file, called a *symlink*, which points at the default interpreter. You can see where it points by:

```
$ ls -l /bin/sh
```

The default is often `/bin/bash`, but you can also specify to use bash (or another shell) directly by replacing the line with the location of bash on your system, which is usually `#!/bin/bash`.

Next, we can execute the script using the `source` command, and redirect the "lincoln.txt" file to stdin. This will cause the script to execute the `cat` command with the contents of "lincoln.txt" as input, consequently printing it to the screen (via stdout):

```
$ source tryme.sh < lincoln.txt 
Four score and seven years ago
our fathers brought forth on this continent, a new nation
```

If you omit the redirection character `<`, the script will try to read from stdin (keyboard input), and then immediately print it back out.

[Back to Top](#table-of-contents)


### Searching

#### Locating Files With `find`

The `find` command provides a wide range of capabilities for searching through directory trees, including executing commands on found files, searching for files based on creation and modification times, and more. It will search any set of directories you specify for files that match the criteria. For example, you might have thousands of files in your home directory and be looking for a file named "foo":

```
$ pwd
/home/jolo

$ find . -name foo
./foo
```

In the example above, the first argument "`.`" indicates for `find` to start searching in the current directory (/home/jolo), and the flag `-name` with the argument `foo` means to search for a file named "foo". Find returns the relative path of the file "foo" when it finds it in the filesystem. In this case, the file was found right in the home directory.

You can also specify more than one location to search:

```
$ find /home/jolo/Project /home/jolo/Results/ . $HOME -name foo
```

This searches for the file name "foo" in the "/home/jolo/Project/", "/home/jolo/Results/", and the current directory.

#### Locating Files With `locate`
Another command provided on most Linux systems is the locate `command`, which builds a file-based database of files and their locations and will match strings. `locate` is usually faster than `find` because it searches the database, rather than looking in each directory and subdirectory. You can use `locate myfile` in order to find where the file is located. Try `locate -h` for a full list of options.

#### Pattern Matching With `grep`
The `grep` (global regular expression print) command is another useful utility which searches the named input file for lines that match the given pattern and prints those matching lines. In the following example, grep searches for instances of the word "bar" in the file "foo":

```
$ cat foo
tool
bar
cats
dogs

$ grep bar foo
bar
```

If there are no matches, `grep` will not print anything to the screen.

[Back to Top](#table-of-contents)

### Exercises With Piping, Sorting, and Counting

#### Piping
Similar to redirection, pipes, using the character `|`, *send the output of one command to the input of another*, thereby chaining simple commands together to perform more complex processing than a single command can do. Most Linux commands will read from stdin and write to stdout instead of only using a file, so `|` can be a very useful tool.

Another useful bash command is `history`, which will print all the bash commands you've entered in the shell on the system to stdout (up to a maximum set by the system administrator). This can be very useful when you only remember one part of a command, or forget exact syntax, but it will quickly become more daunting to search the output the longer you use the shell. You can use `|` combined with `grep` to quickly and easily search this output. First try just the `history` command to view the normal output, then try searching for the `cat` command, which we've used several times, like this:

```
$ history | grep cat
 1901  cat /etc/shells
 1927  cat programinput | mycommand
 1929  cat mylist
 1930  cat > lincoln.txt
 1931  cat lincoln.txt
 1932  cat >> lincoln.txt
 1933  cat > tryme.sh
 1936  cat foo
 1937  history | grep cat
```

As you can see, it is much easier to find specific past commands by this method. The numbers before the commands indicate the line number in the bash history file, which corresponds to when the command was entered, as long as your history has not been cleared. You can even redirect this output to a file by:

```
$ history | grep cat > cat_history.txt
```

This file can then be searched, for example for "lincoln" using less, a text editor (covered later), or by the command:

```
$ cat cat_history.txt | grep lincoln
 1930  cat > lincoln.txt
 1931  cat lincoln.txt
 1932  cat >> lincoln.txt
```

Here's another example. Say you wanted to search the list of processes for which ones were using bash, then you could use `ps -ef` (which outputs all processes in full format) as input to `grep` like so:

```
$ ps -ef | grep bash
```

There are many ways to use `|`, as you have seen in these examples, but feel free to explore more options!

#### Sorting
The `sort` command sorts the content of a file or any stdin, and prints the sorted list to the screen.

```
$ cat temp.txt 
cherry
apple
x-ray
clock
orange
bananna

$ sort temp.txt
apple
bananna
cherry
clock
orange
x-ray
```

To see the sorted list in reverse order, use the `-r` option. `sort -n` will sort the output numerically rather than alphabetically.

```
$ cat temp2.txt
7
48
1
56
8
32

$ sort -nr temp2.txt
56
48
32
8
7
1
```

Note that the two options can be combined by `-nr`, and order does not matter unless there is an input for a particular option. If you were looking for a filename that began with a "w", you may try:

```
$ ls | sort -r
```

#### Counting
The `wc` command reads either stdin or a list of files and generates

1. numbers of lines (by counting the number of newline characters)
1. numbers of words
1. numbers of bytes

Using the file temp.txt from the previous example, we can use `wc` to count the lines, words, and bytes (or characters):

```
$ wc temp.txt 
 6  6 40 temp.txt
```

The output shows that there are **6** lines, **6** words, and **40** bytes (or characters) in the file temp.txt. You can also use the following options with wc to specify certain behavior:

1. Only display line count: -l
   ```
   $ wc -l temp.txt
   6 temp.txt
   ```
1. Only display word count: -w
   ```
   $ wc -w temp.txt
   6 temp.txt
   ```
1. Only display byte count: -c
   ```
   $ wc -c temp.txt
   40 temp.txt
   ```
You can combine `wc -l` with `ls` to list the number of files in a directory:

```
$ ls | wc -l
```

[Back to Top](#table-of-contents)

### Job Control

In addition to starting commands, the shell provides basic job control functions for processes. For shell sessions, such as interactive sessions on Stampede2, it can be useful to see and control processes that run for longer times. Job control allows the user to stop, suspend, and resume jobs from within the shell. This is useful if you have a program which runs longer than desired, does not complete due to a bug, or has other problems. From within a shell session, the ps command will show the current processes running in your shell session.

```
$ ps
  PID TTY          TIME CMD
 4621 pts/7    00:00:00 bash
32273 pts/7    00:00:00 ps
```

From within a running process, using **Ctrl-C** will send an interrupt signal to the process, which will usually cause it to terminate (I/O and other factors can block the interrupt from taking effect immediately). If you have executed a long-running program that you want to complete, but you want to do other things in the same shell while it runs (rather than starting a new shell on the system), you can suspend the process by pressing **Ctrl-Z**. It can be resumed in the **b**ack**g**round with `bg`. Similarly, a process can be invoked and immediately sent to the background by adding an `&` at the end of the command:

```
$ mylong-runningcode.o &
```

A running job can be brought to the foreground with fg:

```
$ ./longscript.sh &
$ fg
./longscript.sh
```
While using job control in the shell you can also use the jobs command to display currently running jobs, similar to `ps`.

```
$ ./longscript.sh &
$ jobs
[1]+  Running                 ./longscript.sh &
$ fg
./longscript.sh
```

In the example above, we invoke `longscript.sh &` and immediately send it to the background. The `jobs` command shows the list of running jobs under shell control. Using `fg`, we can bring "longscript.sh" back to the foreground.

You can also use the `top` command to view details about running processes. The program `htop` is a common, more interactive, alternative to `top` that is not installed by default on most Linux systems, but is worth exploring. And finally, use the `kill` command followed by a PID to stop a process. For more information on these job control commands, see their respective man pages.

[Back to Top](#table-of-contents)

### Environment Variables

As mentioned previously, the environment variable `$PATH` stores a list of directory paths which the shell searches when you issue a command. You can view this list with `echo $PATH`. If the command you issue is not found in any of the listed paths (having been added either by the OS, a system administrator, or you), then the shell will not be able to execute it, since it is not found in any of the directories in the PATH environment variable. You can change the directories in the list using the export command. To add directories to your path, you can use:

```
$ export PATH=$PATH:/path/to/new/command
```

The : joins directories in the path variable together. If you wanted to completely replace the list with a different path:

```
$ export PATH=/path/to/replacement/directory
```

Please keep in mind that the previous list will be erased.

Try `env` to list all environment variables. Examples such as `SHELL`, `HOME`, and `PATH` are built-in shell variables, but any variable can be declared to the shell by typing

```
MYVAR=something
```

It is standard convention to declare shell variables in all capitals. The variable can then be used in bash commands or scripts by inserting `$MYVAR`. Note the use of a preceding `$` when using the variable which is not present when setting it. If you want to insert the variable inside a string, this can be done by `${MYVAR}` with the beginning and ending portions of the string on either side. For example:

```
$ DATAPATH=/home/jolo/Project
$ ls $DATAPATH
$ cp newdatafile.txt ${DATAPATH}/todaysdatafile.txt
$ ls $DATAPATH
todaysdatafile.txt
```

[Back to Top](#table-of-contents)


## Text Editors

### Overview

A text editor is a tool to assist the user with creating and editing files. There is no "best" text editor; it depends on personal preferences. Regardless of your typical workflow, you will likely need to be proficient in using at least one common text editor if you are using Linux for scientific computing or similar work. Two of the most widely used command line editors are **Vim** and **Emacs**, both of which are available on Stampede2 via the `vim` and `emacs` commands respectively.

Another commonly used text editor is Vi (command is `vi`), the predecessor of Vim. Most modern systems actually alias `vi` to `vim`, so that you are using `vim` whenever you enter the `vi` command. You can determine if this is the case by entering:

```
$ which vi
alias vi='vim'
    /bin/vim
```

The line `alias vi='vim'` tells you that vim will be executed whenever the command `vi` is entered. The above output is actually from Stampede2. For this reason, we will focus on Vim in this tutorial and not Vi. We will also provide a basic overview of Emacs.

#### Major Differences

Each text editor in Linux has a designed workflow to assist you in editing, and some workflows work better than others depending on your preferences. For example, Emacs relies heavily on key-chords (or multiple key strokes), while Vim uses distinct editing [modes](https://en.wikibooks.org/wiki/Learning_the_vi_Editor/Vim/Modes#Modes). Vim users tend to enter and exit the editor repeatedly, and use the shell for complex tasks, while Emacs users typically remain within the editor and use Emacs itself for [complex tasks](https://www.gnu.org/software/emacs/manual/html_node/emacs/Shell.html). Most users develop a preference for one text editor and stick with it.

Following is a brief introduction to the basics of both Vim and Emacs. It is recommended that you try each one and work through testing each of the commands to select which one works best for your workflow. Additional editors on Stampede2 you may want to consider are [nano](https://www.howtogeek.com/howto/42980/the-beginners-guide-to-nano-the-linux-command-line-text-editor/), a simple text editor designed for new users, or [gedit](https://help.gnome.org/users/gedit/stable/), a general-purpose text editor focused on simplicity and ease of use, with a simple GUI. There are also many more, so feel free to explore other options.

[Back to Top](#table-of-contents)


### Vim

#### Basic Functions
- **Open** an existing file by entering vim in the shell followed by the name of the file.
- **Create** a new file in the same way as opening a file by specifying the new filename. The new file will not be saved unless specified.
- **Save** a file that is currently open by entering the :w command.
- **Quit** an open file by entering the :q command. If you have made any edits without saving, you will see an error message. If you wish to quit without saving the edits, use :q!.
- **Save and Quit** at the same time by combining the commands: :wq.
- **Edit** the file by entering insert mode to add and remove text. Entering into normal mode will allow you to easily copy, paste, and delete (as well as other functionality).
- **Cancel** a command before completely entering it by hitting Esc twice.

#### Insert Mode

When you first open a document, you will always start in normal mode and have to enter insert mode. To enter insert mode where the cursor is currently located in the file, press the letter **`i`** or the **`Insert`** key. Additionally, you can press the letter **`a`** (for append) if you would like to enter insert mode at the character after the cursor. To exit insert mode, press the `Esc` key. When in insert mode, `-- INSERT --` will be visible at the bottom of the shell. Navigation in insert mode is done with the standard arrow keys.

#### Normal (Command) Mode
Vim starts in normal mode, and returns to normal mode whenever you exit another mode. When in normal mode, there is no text at the bottom of the shell, except the commands you are entering.

#### Navigation
Navigation in normal mode has a large number of shortcuts and extra features, which we will only cover some of here. Basic movement can be done using the arrow keys or using the letter keys in the following table:

Move|	Key
---|---
 ← | h 
 ↓ | j
 ↑ | k
 → | l

The benefits of using the alternate keys is that you do not have to move your hand back-and-forth to the arrow keys while in this mode, and can more effectively enter Vim commands (once you are practiced). Some other examples of navigation shortcuts include:

- Move to the **beginning of the line**: `0`
- Move to the **end of the line**: `$`
- Move to the **beginning of the next word**: `w` This can also be used with a number to move multiple words at once (i.e. `5w` moves 5 words forward).
- Move to the **end of the current word**: `e` This can be used with a number in the same way that `w` can to move multiple words at once.

These extra navigation shortcuts become powerful when combined with other Vim functions, allowing you to edit text and navigate through the file without changing modes.

#### Editing Features
Here are some important commands to know:

- **Undo** the previous command, even the last edit in insert mode, with the command u
- **Redo** the previous command (after undo) with Ctrl-R
- **Copy** (yank) characters, words, or lines:
   - **`yl`** to copy a single character under the cursor
   - **`yy`** to copy the current line
   - **`y#y`** or `#yy` where `#` is replaced with the number of lines you want to copy (i.e. `y25y` will copy 25 lines).
- **Paste** (put) characters, words, or lines:
   - **`p`** will paste after the cursor for characters and words, or on the next line (regardless of the cursor location within a line) if you are pasting lines.
   - **`P`** will paste before the cursor for characters and words, or on the preceding line (regardless of the cursor location within a line) if you are pasting lines.
- **Delete** or **Cut** characters, words, or lines (that can then be pasted elsewhere):
   - **`x`** to delete a single character under the cursor
   - **`dd`** to delete the current line
   - **`d#d`** or `#dd` where `#` is replaced with the number of lines you want to delete (similar to copy).
- **Search** for strings throughout a file and optionally **replace**:
   - A basic search for a word is simply **`/word`** followed by `Enter`. This will jump to the first occurrence of the word after the cursor. Phrases can also be used.
   - Once a search is active, you can use **`n`** to jump to the next occurrence and **`N`** to jump to the previous occurrence.
   - [Search and replace](https://vim.wikia.com/wiki/Search_and_replace) has many options, but one example is to find all occurrences of "foo" in the file and replace (substitute) them with "bar" with the command: :%s/foo/bar/g
- **Split** the screen vertically or horizontally to view multiple files at once in the same shell:
   - **`:sp <filename>`** will open the specified file above the current active file and split the screen horizontally.
   - **`:vsp <filename>`** will open the specified file to the left of the current active file and split the screen vertically
   - Navigate between split-screen files by pressing **`Ctrl-W`** followed by navigation keys (i.e. `Ctrl-W h` or `Ctrl-W ←` to move to the left file)
   - Also note that you can open several documents at once from the shell using appropriate flags. See `man vim` for more information.

Any of the editing commands can easily be combined with navigation commands. For example, `5de` will delete the next 5 words, or `y$` will copy from the current cursor location to the end of the line. There are a large number of combinations and possible commands. Note that copying, pasting, and deleting can also be done efficiently using visual mode.

#### Visual Mode
From normal mode, press the **`v`** key to enter visual mode. This mode enables you to highlight words in sections to perform commands on them, such as copy or delete. Navigation in visual mode is done with the normal mode navigation keys or the standard arrow keys. For example, if you are in normal mode and you want to copy a few words from a single line and paste them on another line:

1. Navigate to the first character of the first word you want to copy
1. Enter visual mode by `v`
1. Navigate to the last character of the last word you want to copy (this should highlight all the words you want)
1. Enter `y` to copy the words
1. Navigate to where you want to paste the words
1. Enter `p` to paste

Note that step 6 will paste after the cursor instead of on the next line even if you have copied several lines. You can also replace that step with `P` to paste before the cursor.

#### Vim and the Shell
Working with Vim regularly can mean switching back and forth between it and the shell, but there are two ways to simplify this. From normal mode, you can use the command `:!` followed by any shell command to execute a single command without closing the file. For example, `:!ls` will display the contents of the current directory. This will appear to background Vim while executing the command (so you can see the shell and output), and display the following message upon completion:

```
Press ENTER or type command to continue
```

Pressing Enter will return you to your open file. Alternatively, you can simply background the file while in normal mode with **`Ctrl-Z`** to view the shell and issue commands. When you want to return to the file, use the foreground command **`fg`**. In this way, you can actually have a number of files open (with or without splitting the screen) all in the same shell, and easily switch between them. Note that if you background multiple files, the foreground command will bring them up in reverse order (most recent file accessed first).

#### Customization
Vim uses a `.vimrc` file for customizations. Essentially, this file is to consist of Vim commands that you would like issued each time you open Vim to customize your experience. One example of a command you will likely want is `syntax on`, which provides syntax highlighting for programming languages. There are also a number of commands you can explore to customize the coloring of the syntax. Here is an example of a simple `.vimrc` file that you may use:

```
syntax on
set tabstop=4
set expandtab
set number
set hls
```

In addition to syntax highlighting, the above customizations will set tabs to be 4 characters wide, replace tabs with spaces, show line numbers along the left-hand side of the screen, and highlight matching words when searching. There is a global `vimrc` file that sets system-wide Vim initializations (you will not have access to this on Stampede2), and each user has their own located at `~/.vimrc` wich can be used for personal initializations.

#### A Hands-On Tutorial
One of the most effective ways to learn Vim is through the built-in hands-on tutorial that can be accessed via the shell by the command **`vimtutor`**. This command will open a text file in Vim that will walk you through all the major functionalities of Vim as well as a few useful tips and tricks. If you plan to use Vim even occasionally, it is a great resource. Furthermore, the above list of features and commands is not exhaustive, and the interested new Vim user should certainly explore the man pages and online resources to discover more Vim features.

[Back to Top](#table-of-contents)


### Emacs 

#### Basic Functionality
- **Open** an existing file by entering **`emacs`** in the shell followed by the name of the file. This will default to running Emacs in a GUI, but it can also be run within the shell (`emacs -nw`). Note that to use the GUI with a remote connection such as Stampede2, you must use X11 forwarding (covered in the Remote Connections section), otherwise the `emacs` command will open within the shell. If you want to run the GUI and keep the shell free, you can open and background Emacs with `emacs &`. Use `Ctrl-x f` to open a file from within Emacs.
- **Create** a new file in the same way as opening a file by specifying the new filename. The new file will not be saved unless specified.
- **Save** a file that is currently open by entering the **`Ctrl-x Ctrl-s`** command.
Quit by entering **`Ctrl-x Ctrl-c`**.
- **Save and Quit** is the same command as quitting, except that when you have unsaved files it will ask if you would like to save each one. To save, enter `y`.
- **Edit** a file by simply entering and removing text.
- **Cancel** a command before completely entering it or a command that is executing with **`Ctrl-g`** or by hitting **`Esc`** 3 times.

#### Navigation
Similar to Vim, navigation in Emacs has shortcuts and extra features. Basic movement can be done using the arrow keys or using the letter keys in the following table:

Move|	Key
---|---
 ← | Ctrl-b 
 ↓ | Ctrl-n
 ↑ | Ctrl-p
 → | Ctrl-f

The benefits of using the alternate keys is that you do not have to move your hand back-and-forth to the arrow keys, and can more effectively enter Emacs commands (once you are practiced). Some other examples of navigation shortcuts include:

- Move to the **next screen view**: `Ctrl-v`
- Move to the **previous screen view**: `Alt-v`
- Move to the **next word**: `Alt-f` This can also be used with a number to move multiple words at once (i.e. `Alt-5f` moves 5 words forward).
- Move to the **previous word**: `Alt-b` This can be used with a number in the same way to move multiple words at once.
- Move to the **beginning of the line**: `Ctrl-a`
- Move to the **end of the line**: `Ctrl-e`
- Move to the **beginning of a sentence**: `Alt-a`
- Move to the **end of a sentence**: `Alt-e`

Note that the more customary keys `Page Up`, `Page Down`, `Home`, and `End` all work as expected.

### Editing Features
Here are some important commands to know:

- **Undo** the previous command with the command **`Ctrl-x u`**
- **Redo** the previous command (after undo) by performing a non-editing command (such as `Ctrl-f`), and then undo the undo with `Ctrl-x u`
- **Delete or Cut** characters, words, or lines (that can then be pasted elsewhere):
   - **`Backspace`** to delete a single character before the cursor
   - **`Ctrl-d`** to delete a single character after the cursor
   - **`Alt-Backspace`** to delete the word before the cursor
   - **`Alt-d`** to delete the word after the cursor
   - **`Ctrl-k`** to delete from the cursor to end of the line
   - **`Alt-k`** to delete from the cursor to end of the sentence
- **Paste** characters, words, or lines:
   - **`Ctrl-y`** pastes the most recent deleted text
   - **`Alt-y`** pastes the deleted text before the most recent
- **Copy** characters, words, or lines: The easiest way to copy is actually to cut the text and then paste it back where it was. Then it can be pasted in a new location also.
- **Search** for strings throughout a file and optionally **replace**:
   - **`Ctrl-s`** starts a forward search that is incremental (each character you enter updates the search). Entering Ctrl-s again skips to the next occurrence. Enter ends the search.
   - **`Ctrl-r`** starts a backwards search that behaves similarly to the forward search.
   - [Search and replace](https://www.gnu.org/software/emacs/manual/html_node/emacs/Search.html#Search) has many options, but one example is to find all occurrences of "foo" in the file and replace them with "bar" with the command: **`Alt-x replace-string foo Enter bar`**
   - You can use tab-completion for entering commands after typing `Alt -x`. For example, type `Alt-x`, then `rep`, then hit `Tab` twice to see a list of matching commands. Since similar commands are named similarly, you will find other useful related commands, such as `replace-regexp`.
- **Split** the screen vertically or horizontally to view multiple files at once in emacs:
   - **`Ctrl-x 3`** will split the screen horizontally
   - **`Ctrl-x 2`** will split the screen vertically
   - **`Ctrl-x 1`** closes all panes except the active one
   - **`Ctrl-x 0`** closes a pane

#### Highlighting Mode
This mode enables you to highlight words in sections to perform commands on them, such as copy or delete. For example, if you want to copy a few words from a single line and paste them on another line:

1. Navigate to the first character of the first word you want to copy
1. Enter highlighting mode by `Ctrl-Space`
1. Navigate to the last character of the last word you want to copy (this should highlight all the words you want)
1. Enter `Alt-w` to copy the words
1. Navigate to where you want to paste the words
1. Enter `Ctrl-y` to paste

#### Emacs and the Shell

There are several options for running shell commands from Emacs. To execute a single shell command while in Emacs, use the command **`Alt-!`** followed by the shell command and hit `Enter`. The output of the command will display in a portion of the screen called an [echo area](https://www.gnu.org/software/emacs/manual/html_node/emacs/Echo-Area.html). There are several more features for running shell commands, including running an interactive shell inside Emacs (we recommend [ansi-term](https://www.emacswiki.org/emacs/AnsiTerm)). For more about these features, please see the [Emacs documentation](https://www.gnu.org/software/emacs/manual/html_node/emacs/Shell.html) on the topic. Alternatively, you can suspend Emacs with the command **`Ctrl-z`**. As with suspending Vim, you can execute commands in the shell, and then return to Emacs with the foreground command **`fg`**.

#### Customization
Emacs is customizable in many ways including changing the [key bindings](https://www.gnu.org/software/emacs/manual/html_node/emacs/Key-Bindings.html) for commands, the color scheme (themes), and more. Due to the breadth of options, we refer you to existing documentation on [customization](https://www.gnu.org/software/emacs/manual/html_node/emacs/Customization.html).

A Hands-On Tutorial
One of the most effective ways to learn Emacs is through the built-in hands-on tutorial that can be accessed by opening Emacs without any filename input. It will walk you through all the major functionalities of Emacs as well as a few useful tips and tricks. If you plan to use Emacs even occasionally, it is a great resource. Furthermore, the above list of features and commands is not exhaustive, and the interested new Emacs user should certainly explore the man pages and online resources to discover more Emacs features. In particular, [buffers](https://www.gnu.org/software/emacs/manual/html_node/emacs/Buffers.html) are a useful concept to understand when using Emacs, but are not covered here.

Additionally, you may want to consider looking into [spacemacs](http://spacemacs.org/) if you are familiar with Vim key bindings or would like to continue using emacs with more customization. Users can install spacemacs to their local directory on Stampede2 using git.

[Back to Top](#table-of-contents)

## Accounts

A user account is required for a user to log into any Linux system. An account typically includes identity information such as username, password, user id (UID), and group identifier (GIDs) so that the system can identify the user. An account will also include a set of resources such as accessible disk space that the user can work in, typically called a home directory and information about the default shell preference.

### Username
Every account in a Linux system is associated with a unique username, which is typically a sequence of alphanumeric characters at least three characters in length. It is case sensitive; for example, `Apple01` is a different username than `apple01`. A unique integer, the user id (UID), is assigned to each username. Linux uses the UID rather than the username to manage user accounts, since it is more efficient to process numbers than strings. However, you don't necessarily need to know your UID.

### Group
Linux also has the notion of a group of users who need to share files and processes. Each account is assigned a primary group with a numerical group id (GID) that corresponds to the particular group. A single account can belong to many groups, but you may have only one primary group. Groups can also be used to assign certain permissions to users on the system.

### Password
Each username requires a password. The username is the account identifier and the password is the authenticator. A password is required for operations such as logging into the system and accessing files. On XSEDE, passwords must be a minimum of 8 characters with at least 3 of the following character classes:

- lower-case letters
- upper-case letters
- numerical digits
- punctuation

When you enter your password, the system encrypts it and compares it to a stored string. This ensures that even the Operating System does not know your plain text password. This method is frequently used on websites and servers, especially those that run Linux. Once you have your XSEDE username and password, you are ready to login to a remote XSEDE system using a secure shell (`ssh`).

[Back to Top](#table-of-contents)


## Remote Connections

### Connect Remotely with ssh

**S**ecure **SH**ell (SSH) is designed to be a secure way to connect to remote resources from your local machine over an unsecured network. The following example uses an account with username "jolo" using SSH to log into a machine named "foo.edu":

```
$ ssh jolo@foo.edu
```

This will open a connection to the remote machine "foo.edu" and log in as the user if authentication is successful. If you log into Stampede2 or other XSEDE resources via SSH, a password and/or private key will be required for authentication. The above example is the most straight-forward version of the command, but there are many additional options. For example, to use a [key pair](https://www.ssh.com/ssh/public-key-authentication) (where `my_key` is the name of the private key file) to login, then the command will look like:

```
$ ssh -i my_key jolo@foo.edu
```

Another common option is [X11](https://en.wikipedia.org/wiki/X_Window_System) forwarding, which can be achieved using the `-X` or `-Y` flags. X11 forwarding is useful when you are going to use applications that open up outside of the shell. For more information on this and other options, see the man page for `ssh`.

In order to access a Linux system via `ssh`, you will need an ssh client and a terminal program on your system. Sometimes these are included in a single application for simplicity. There are many different terminals available, but here are a few examples:

- On Linux, simply open your terminal emulator and enter `ssh` commands
- On Mac OS, the Terminal app is included with the system, and ssh can be invoked from the command line in the Terminal app
- On Windows:
  - The Linux Bash Shell is available as the [Windows Subsystem for Linux](https://docs.microsoft.com/en-us/windows/wsl/about) and supports many Linux commands, including `ssh`
  - One commonly-used terminal and ssh client combo is [PuTTY](https://www.chiark.greenend.org.uk/~sgtatham/putty/latest.html)
  - Another terminal and ssh client combo is [MobaXterm](https://mobaxterm.mobatek.net/)

### Securely Copy with scp

**S**ecure **C**opy **P**rotocol (SCP) is based on the SSH protocol, and is used for securely copying files across the network. Say you have a file "code.c" located in your current directory on your local machine that you want to copy **to a remote resource** in the "Project" directory under your home directory on the remote machine (we'll stick with the user "jolo" and "foo.edu"). This can be done using scp as follows:

```
scp code.c jolo@foo.edu:~/Project
```

Alternatively, if you want to copy the file "output.txt" **from a remote resource** located in the "Project" directory to the directory "Results" on your local machine and rename the file to "Run12_data.txt" during the move:

```
scp jolo@foo.edu:~/Project/output.txt ./Results/Run12_data.txt
```

Similar syntax can be used to copy from a remote host to another remote host as well. The -r option can be used to copy full directories recursively. For more options, see the scp man page.

To use `scp` with a remote system, similar to `ssh`, you will need a program to support it. Here are a few examples:

- On Linux, open your terminal emulator and enter `scp` commands
- On Mac OS, open the Terminal app and enter `scp` commands
- On Windows:
  - The [Windows Subsystem for Linux](https://docs.microsoft.com/en-us/windows/wsl/about) supports many Linux commands, including `scp`
  - From the developers of PuTTY, you can use [PSCP](https://www.chiark.greenend.org.uk/~sgtatham/putty/latest.html)
  - [MobaXterm](https://mobaxterm.mobatek.net/) comes with a built-in SCP client

If you are expecting to copy large files to or from remote locations, note that File and Directory Compression will be covered later in this tutorial, under [Optional Topics](linux-optional.md).

[Back to Top](#table-of-contents)


## Filesystem

As we discussed previously, Linux has a hierarchical filesystem. The files and directories form a tree structure, in accordance with the [Filesystem Hierarchy Standard (FHS)](https://en.wikipedia.org/wiki/Filesystem_Hierarchy_Standard). The topmost directory is the root directory `/` and all directories are contained within or below this directory in the hierarchy. There are several directories within the root directory – called subdirectories – that are generated upon installation of a Linux distribution. Many of these are used exclusively by the system. There are also some generated for use by users, where subdirectories can be created without elevated permissions. A sample portion of this structure is depicted in Figure 1.

![Figure 1: A sample portion of the filesystem structure tree](linux-filesystem_structure.jpg)

The FHS includes descriptions of the core directories in the hierarchy, causing this structure to be relatively standard across Linux systems. Table 1 provides a list of the major subdirectories of the root directory `/`. There is no need to remember the purpose of every directory unless you are working at a lower level within a Linux system. Rather, this table should give you an idea of the basic layout of a Linux filesystem, and possibly serve as a useful reference in the future.

Directory	| Contents
---|---
`bin`	| Binary files for command execution
`boot` | Files for the boot loader
`dev`	| Device files for interacting with devices connected to the system
`etc`	| System configuration files
`home` | User home directories
`lib`	| System shared libraries needed by binaries in bin and sbin
`media`	| Location for temporarily mounting filesystems from replaceable media
`mnt`	| Location for temporarily mounting filesystems
`opt`	| Optional application software packages
`proc` | Virtual filesystem for process and system information
`root` | Home directory of root user
`run`	| Run-time variable data since last boot
`sbin` | System binary files for command execution
`srv`	| Data for services provided by the system
`sys`	| Virtual directory for system information
`tmp`	| Temporary files
`usr`	| Read-only user data for all users; Some important subdirectories include: /usr/bin - program binaries /usr/include - include files /usr/lib - libraries for binaries in /usr/bin and /usr/sbin /usr/local - local host data /usr/sbin - Non-essential system binaries /usr/share - shared data, such as documentation /usr/src - kernel source code and headers
`var`	| Variable data

Most of the work you do will likely be performed in your home directory while on a Linux system, while programs you use will reside in other locations as explained above. You may want to familiarize yourself with the Stampede2 Filesystem as well as how to navigate it if you are planning on doing work there. Also feel free to peruse [Optional Topics](linux-optional.md) in a later section for more information on the root user and mounts.

[Back to Top](#table-of-contents)


## File Permissions

### Overview

Linux is a multi-user environment where many users can run programs and share data. File permissions are used to protect users and system files by controlling who can **read**, **write**, and **execute** files. The types of permissions a file can have are:

Read Permissions | Write Permissions | Execute Permissions
-----------------|-------------------|--------------------
r                |  w 	             | x

Furthermore, files and directories have 3 levels of permissions: **User**, **Group** and **World**. When displayed, permissions are arranged into three sets of three characters each. The first set is the User (owner) permissions, the second is Group permissions, and finally permissions for Others or everyone else on the system (World). In the following example, the owner can read and write the file, while group and all others have read access only.

User (owner) | Group | Others (everyone else)
-------------|-------|-----------------------
rw-          | r--	 | r--

#### Displaying File Permissions
You can view a file's permissions by using the "long list" option `ls -l`, which outputs the permissions as a character string at the beginning of the row for each file or directory. The string will begin with a d for a directory or a - for a file. The next nine characters refer to the file permissions in the order discussed above. Other information included per row of the output is (in order) links to the file, username of the owner, group, file size, date and time of last edit, and filename. For example:

```
$ ls -l $HOME
-rw-r--r-- 1 jdoe jdoe            796631 2009-11-20 14:25 image_data.dat
-rwxrwxr-- 1 jdoe community_group 355    2010-02-18 15:50 my_script.sh
```

In this example, user "jdoe" owns the two files: "image_data.dat" and "my_script.sh". For the first file, we can tell that "jdoe" has read and write access (but not execute permissions) because of the `rw-` in the -rw-r--r-- character string on that row. Similarly, we can see that the group only has read access (-rw-r--r--) and all others on the system only have read access (-rw-r--r--). The second file can be read, written, and executed by "jdoe" and others who are in the "community_group".

#### Changing File Permissions
You can use the `chmod` command to change permissions on a file or directory (use `chmod -R` for recursive). This command changes the file permission bits of each file according to a given mode, which can be either a symbolic representation (characters) of changes to be made or an octal number representing the bit pattern for the new mode bits.

#### Symbolic Mode
The syntax of the command in symbolic mode is

```
chmod [references][operator][modes] file
```

- **references** can be "u" for user, "g" for group, "o" for others, or "a" for all three types
- **operator** can be "+" to add, "-" to remove permissions, and "=" to set the modes exactly
- **modes** can be "r" for read, "w" for write, and "x" for execute

In the following example, we are giving the owner read, write, and execute permissions, while the group and everyone else is given no permissions.

```
$ chmod u+rwx my_script.sh

$ ls -l my_script.sh
-rwx------ 1 jdoe community_group     355 2010-02-18 15:50 mmy_script.sh
```

The `u+` adds permissions for the user, and the `rwx` specifies which permissions to add. A common use for this method is to make a script that you have written executable. The command chmod u+x my_script.sh will make the script executable by the owner. Once you have changed the permissions, you can run the script by issuing ./my_script.sh.

Alternatively, you can run a script with the `source` command, in which case it is not necessary for the script file to be executable. However, be aware that doing `source my_script.sh` will run the commands from `my_script.sh` as if you were typing them into the current shell. Thus, any variables defined or changed in the script will remain defined or changed in your current shell environment, unlike what happens when you run an executable script, which does not affect your current environment.

#### Numeric Mode
Numeric mode uses numbers from one to four octal digits (0-7). The rightmost digit selects permissions for the World, the second digit for other users in the group, and the third digit (leftmost) is for the owner. The fourth digit is rarely used.

The value for each digit is derived by adding up the bits with values 4 (read only), 2 (write only), and 1 (execute only). For example, to give read and write permissions, but not execute permissions, you would use a 6. The value 0 removes all permission for the specified set, whereas the value 7 turns on all permissions (read, write, and execute).

Let's say you have an executable that you would like others in your group to be able to read and execute, but you do not want anybody else to be able to have any access. First you need to set the read, write, and execute permission for yourself (7), then give read and execute to your group (5), and finally no permissions for everybody else (0). So the full number you would use is 750.

```
$ ls -l my_script.sh
rw-r--r-- 1 jdoe community_group     355 2010-02-18 15:50 my_script.sh

$ chmod 750 my_script.sh

$ ls -l my_script.sh
-rwxr-x--- 1 jdoe community_group    355 2010-02-18 15:50 my_script.sh
```

For more on user permissions, see *Root and Sudo* later in [Optional Topics](linux-optional.md).

[Back to Top](#table-of-contents)


## Optional Topics

### File and Directory Compression
Compression in Linux typically deals with collections of files into an archive, using the tar command which gets its name from **t**ape **ar**chive. Files or directories can be packed into a single tar file, as well as compressed further either the `-z` option to tar or other programs. The `-c` flag is used to **c**reate an archive and the `-x` flag is for e**x**traction of an archive. The `-v` option enables **v**erbose output, and `-f` specifies to store as an archive **f**ile. By default, directories are added recursively, unless otherwise specified. Here is an example of creating an archive or tar file:

```
$ tar -cvf my_archive.tar file1 file2 file3
file1
file2
file3
```

And to extract the same archive (not verbose):

```
$ tar -xf my_archive.tar
```

A program commonly used along with `tar` is `gzip`, which creates archives with the extension `.gz`. A file can be compressed simply by `gzip file` (with an added `-r` for a directory) or a `.tar.gz` file can be created (or extracted) by adding the `-z` option to a `tar` command. For example, the same command from above to extract with `gzip`:

```
$ tar -xzf my_archive.tar.gz
Another common extension for a gzipped tar file is .tgz.
```
For more on compression, see this [detailed article](https://www.digitalocean.com/community/tutorials/an-introduction-to-file-compression-tools-on-linux-servers).

### Symbolic Links
Symbolic links are a special type of file which refer to another file in the filesystem. The symbolic link contains the location of the target file. Symbolic links are used to provide pointers to files in more than one place and can be used to facilitate program execution, make navigating on the system easier, and are frequently used to manage system library versions. To make a symbolic link:

```
$ ln -s data/file/thats/far/away righthere
```

See the man pages for `ln` for more information on linking files.

### Root and Sudo

The **root** user on any system is the administrative account with the highest level of permissions and access. This account is sometimes referred to as the [superuser](https://en.wikipedia.org/wiki/Superuser). By default, most Linux systems have a single root account when installed and user accounts have to be set up. The root account has a UID of 0, and the system will treat any user with a UID of 0 as root.

If you have access to a root account on any Linux system, best practice is to only use this account when the privileges are needed to perform your work (such as installing packages), and to use a user account for all of your other work. Note that the root directory is not the home directory of the root user, but rather the root of the filesystem. The home directory of the root user is actually located at `/root`.

The program **sudo** allows users to run commands with the equivalent privileges of another user. The default privileges selected are the root user's, but any user can be selected. A user with sudo privileges can run commands with root privileges without logging in as root (must enter user's password) by putting `sudo` in front the command. The first user account created on some Linux distributions is given sudo privileges by default, but most distributions require you to specifically give sudo privileges to a user. This is typically done by editing the `/etc/sudoers` file (requires either root or sudo access), or running a command like `usermod`.

### Package Managers

The root user and any user with sudo privileges have full access to the features of a [package manager](https://en.wikipedia.org/wiki/Package_manager). In short, [packages](https://en.wikipedia.org/wiki/Package_format) are archives of software and associated data, and a package manager is used to install, uninstall, and manage packages on a system. They are used in the shell or through a GUI, and have varying features. Most Linux distributions have a default package manager installed with the system. Some common package managers available are:

- APT, which includes:
   - apt
   - aptitude
   - apt-get
  
  The Debian-recommended CLI choice is `apt`; see this article for a detailed explanation.

- Synaptic - a GUI for APT
- dpkg
- yum
- pacman

Commands for these package managers can be found in their supporting documentation or via the man pages. Note that if you are only using Linux on an XSEDE managed resource, the availability of user software is typically managed through the Module Utility.

### Mounts
The `mount` command can be used to attach the filesystem of another device at a specified place in the directory tree for easy read/write access. `mount` with no arguments is useful for seeing what devices are mounted. Typically, you must specify the [type](https://en.wikipedia.org/wiki/File_system#Types_of_file_systems) of the filesystem, name of the [device](https://www.dell.com/support/article/us/en/04/sln151767/ubuntu-linux-terms-for-your-hard-drive-and-devices-explained?lang=en#Linux_device_naming_convention), and the path to where you want to mount it:

```
mount -t [type] [device] [path]
```

Use the `umount` command to unmount a device's filesystem. It has similar options to `mount`, and both commands have thorough man pages. Another way to mount a device is to use `fstab`, which automates the process. Network shares can be mounted as well, so long as appropriate credentials are supplied to connect.

[Back to Top](#table-of-contents)

## Acknowledgements

This tutorial is based on material from [An Introduction to Linux](https://cvw.cac.cornell.edu/Linux/default) by the [Cornell University Center for Advanced Computing](https://www.cac.cornell.edu/). All contents copyright © Cornell University. All rights reserved.

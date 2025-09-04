# Introduction to Linux

## Table of Contents

- [Getting Started](#getting-started)
- [Navigating](#navigating)
- [Manipulating files and directories](#manipulating-files-and-directories)
- [Uploading/downloading/editing files ](#uploadingdownloadingediting-files)
- [Beyond the basics](#beyond-the-basics)
- [Common pitfalls](#common-pitfalls)

## Getting Started

### Getting help

Use the documentation available on the system, through the `help`, `info` and `man` commands (use `q` to exit).

```
help cd
info ls
man cp
```

### Basic terminal usage

The basic interface is the so-called shell prompt, typically ending with `$` (for `bash` shells).

You use the shell by executing commands, and hitting `<enter>`. For example:

```
$ echo hello
hello
```

You can go to the start or end of the command line using `Ctrl-A` or `Ctrl-E`.

To go through previous commands, use `<up>` and `<down>`, rather than retyping them.

#### Command history

A powerful feature is that you can ”search” through your command history, either using the
history command, or using `Ctrl-R`:

```
$ history
    1 echo hello

# hit Ctrl-R, type ’echo’
(reverse-i-search)‘echo’: echo hello
```

#### Stopping commands

If for any reason you want to stop a command from executing, press `Ctrl-C`. For example, if a
command is taking too long, or you want to rerun it with different arguments.

### Variables

At the prompt we also have access to shell variables, which have both a *name* and a *value*.

They can be thought of as placeholders for things we need to remember.

For example, to print the path to your home directory, we can use the shell variable named `HOME`:

```
$ echo $HOME
/users/alice
```

This prints the value of this variable.

#### Defining variables
There are several variables already defined for you when you start your session, such as `$HOME`
which contains the path to your home directory.

For a full overview of defined environment variables in your current session, you can use the `env`
command. You can sort this output with `sort` to make it easier to search in:

```
$ env | sort
...
HOME=/users/alice
...
```

You can also use the `grep` command to search for a piece of text. The following command will
output all SSH-specific variable names and their values:

```
$ env | sort | grep SSH
```

But we can also define our own. this is done with the export command (note: variables are
always all-caps as a convention):

```
$ export MYVARIABLE="value"
```

It is important you don’t include spaces around the = sign. Also note the lack of `$` sign in front
of the variable name.

If we then do

```
$ echo $MYVARIABLE
```

this will output `value`. Note that the quotes are not included, they were only used when defining
the variable to escape potential spaces in the value.

##### Changing your prompt using $PS1

You can change what your prompt looks like by redefining the special-purpose variable `$PS1`.

For example: to include the current location in your prompt:

```
$ export PS1=’\w $’
~ $ cd test
~/test $
```

Note that `~` is short representation of your home directory.

To make this persistent across session, you can define this custom value for `$PS1` in your `.profile` startup script:

```
$ echo ’export PS1="\w $ " ’ >> ∼/.profile
```

#### Using non-defined variables

One common pitfall is the (accidental) use of non-defined variables. Contrary to what you may
expect, this does *not* result in error messages, but the variable is considered to be *empty* instead.

This may lead to surprising results, for example:

```
$ export WORKDIR=/tmp/test
$ cd $WROKDIR
$ pwd
/users/alice
$ echo $HOME
/users/alice
```

To understand what’s going on here, see the section on `cd` below.

The moral here is: **be very careful to not use empty variables unintentionally**.

**Tip for job scripts: use set -e -u to avoid using empty variables accidentally**.

The `-e` option will result in the script getting stopped if any command fails.

The `-u` option will result in the script getting stopped if empty variables are used. (see [this](https://ss64.com/bash/set.html) for a more detailed explanation and more options)

More information can be found at [here](http://www.tldp.org/LDP/abs/html/variables.html).

#### Restoring your default environment

If you’ve made a mess of your environment, you shouldn’t waste too much time trying to fix it.
Just log out and log in again and you will be given a pristine environment.

### Basic system information

Basic information about the system you are logged into can be obtained in a variety of ways.

We limit ourselves to determining the hostname:

```
$ hostname
node0.alice-129004.ucy-cs499-dc-pg0.wisc.cloudlab.us
$ echo $HOSTNAME
node0.alice-129004.ucy-cs499-dc-pg0.wisc.cloudlab.us
```

And querying some basic information about the Linux kernel:

```
$ uname -a
Linux node0.alice-129004.ucy-cs499-dc-pg0.wisc.cloudlab.us 4.15.0-169-generic #177-Ubuntu SMP Thu Feb 3 10:50:38 UTC 2022 x86_64 x86_64 x86_64 GNU/Linux
```

### Exercises

- Print the full path to your home directory
- Determine the name of the environment variable to your personal scratch directory
- What’s the name of the system you’re logged into? Is it the same for everyone?
- Figure out how to print the value of a variable without including a newline
- How do you get help on using the `man` command?

[Back to Top](#table-of-contents)


## Navigating

### Current directory: `pwd` and `$PWD`

To print the current directory, use pwd or $PWD:

```
$ cd $HOME
$ pwd
/user/alice
$ echo "The current directory is: $PWD"
The current directory is: /user/alice
```

### Listing files and directories: `ls`
A very basic and commonly used command is `ls`, which can be used to list files and directories.

In it’s basic usage, it just prints the names of files and directories in the current directory. For
example:

```
$ ls
afile.txt some_directory
```

When provided an argument, it can be used to list the contents of a directory:

```
$ ls some_directory
one.txt two.txt
```

A couple of commonly used options include:
- detailed listing using `ls -l`:
   ```
   $ ls -l
   total 4224
   -rw-rw-r-- 1 alice ucy-cs499-dc-PG0 2157404 Apr 12 13:17 afile.txt
   drwxrwxr-x 2 alice ucy-cs499-dc-PG0     512 Apr 12 12:51 some_directory
   ```
- to print the size information in human-readable form, use the `-h` flag:
   ```
   $ ls -lh
   total 4.1M
   -rw-rw-r-- 1 alice ucy-cs499-dc-PG0 2.1M Apr 12 13:16 afile.txt
   drwxrwxr-x 2 alice ucy-cs499-dc-PG0  512 Apr 12 12:51 some_directory
   ```

- also listing hidden files using the `-a` flag:

   ```
   $ ls -lah
   total 3.9M
   drwxrwxr-x   3 alice ucy-cs499-dc-PG0  512 Apr 12 13:11 .
   drwx------ 188 alice ucy-cs499-dc-PG0 128K Apr 12 12:41 ..
   -rw-rw-r--   1 alice ucy-cs499-dc-PG0 1.8M Apr 12 13:12 afile.txt
   -rw-rw-r--   1 alice ucy-cs499-dc-PG0    0 Apr 12 13:11 .hidden_file.txt
   drwxrwxr-x   2 alice ucy-cs499-dc-PG0  512 Apr 12 12:51 some_directory
   ```

- ordering files by the most recent change using -rt:

   ```
   $ ls -lrth
   total 4.0M
   drwxrwxr-x 2 vsc10002 vsc10002 512 Apr 12 12:51 some_directory
   -rw-rw-r-- 1 vsc10002 vsc10002 2.0M Apr 12 13:15 afile.txt
   ```

If you try to use `ls` on a file that doesn’t exist, you will get a clear error message:

```
$ ls nosuchfile
ls: cannot access nosuchfile: No such file or directory
```

### Changing directory: `cd`
To change to a different directory, you can use the `cd` command:

```
$ cd some_directory
```

To change back to the previous directory you were in, there’s a shortcut: `cd -`

Using `cd` without an argument results in returning back to your home directory:

```
$ cd
$ pwd
/users/alice
```

### Inspecting file type: `file`

The file command can be used to inspect what type of file you’re dealing with:

```
$ file afile.txt
afile.txt: ASCII text

$ file some_directory
some_directory: directory
```

### Absolute vs relative file paths
An *absolute* filepath starts with `/` (or a variable which value starts with `/`), which is also called
the root of the filesystem.

Example: absolute path to your home directory: `/users/alice`.

A *relative* path starts from the current directory, and points to another location up or down the
filesystem hierarchy.

Example: `some_directory/one.txt` points to the file `one.txt` that is located in the subdirectory named `some_directory` of the current directory.

There are two special relative paths worth mentioning:
- `.` is a shorthand for the current directory
- `..` is a shorthand for the parent of the current directory

You can also use `..` when constructing relative paths, for example:

```
$ cd $HOME/some_directory
$ ls ../afile.txt
../afile.txt
```

### Permissions
Each file and directory has particular *permissions* set on it, which can be queried using `ls -l`.

For example:
```
$ ls -l afile.txt
-rw-rw-r-- 1 vsc10002 agroup 2929176 Apr 12 13:29 afile.txt
```

The `-rwxrw-r--` specifies both the type of file (`-` for files, `d` for directories (see first character)), and the permissions for user/group/others:
1. each triple of characters indicates whether the read (`r`), write (`w`), execute (`x`) permission
bits are set or not
2. the 1st part `rwx` indicates that the *owner* “alice” of the file has all the rights
3. the 2nd part `rw-` indicates the members of the *group* “agroup” only have read/write permissions (not execute)
4. the 3rd part `r--` indicates that *other* users only have read permissions

The default permission settings for new files/directories are determined by the so-called *umask*
setting, and are by default:
1. read-write permission on files for user/group (no execute), read-only for others (no write/execute)
2. read-write-execute permission for directories on user/group, read/execute-only for others
(no write)

See also the [chmod](#changing-permissions-chmod) command later in this manual.

### Finding files/directories: `find`
`find` will crawl a series of directories and lists files matching given criteria.

For example, to look for the file named one.txt:

```
$ cd $HOME
$ find . -name one.txt
./some_directory/one.txt
```
To look for files using incomplete names, you can use a wildcard `*`; note that you need to escape
the `*` to avoid that Bash *expands* it into `afile.txt` by adding double quotes:

```
$ find . -name "*.txt"
./.hidden_file.txt
./afile.txt
./some_directory/one.txt
./some_directory/two.txt
```
A more advanced use of the `find` command is to use the `-exec` flag to perform actions on the
found file(s), rather than just printing their paths (see `man find`).

### Exercises
- Go to `/tmp`, then back to your home directory. How many different ways to do this can
you come up with?
- When was your home directory created or last changed?
- Determine the name of the last changed file in `/tmp`.
- See how home directories are organised. Can you access the home directory of other users?

[Back to Top](#table-of-contents)


## Manipulating files and directories

Being able to manage your data is an important part of using the HPC infrastructure. The bread
and butter commands for doing this are mentioned here. It might seem annoyingly terse at first,
but with practice you will realise that it’s very practical to have such common commands short
to type.

### File contents: `cat`, `head`, `tail`, `less`, `more`

To print the contents of an entire file, you can use `cat`; to only see the first or last N lines, you
can use `head` or `tail`:

```
$ cat one.txt
1
2
3
4
5
$ head -2 one.txt
1
2
$ tail -2 one.txt
4
5
```

To check the contents of long text files, you can use the less or more commands which support
scrolling with `<up>`, `<down>`, `<space>`, etc.

### Copying files: “cp”

```
$ cp source target
```

This is the `cp` command, which copies a file from source to target. To copy a directory, we use
the `-r` option:

```
$ cp -r sourceDirectory target
```

A last more complicated example:

```
$ cp -a sourceDirectory target
```

Here we used the same `cp` command, but instead we gave it the `-a` option which tells `cp` to copy
all the files and keep timestamps and permissions.


### Creating directories: `mkdir`
```
$ mkdir directory
```
which will create a directory with the given name inside the current directory.

### Renaming/moving files: `mv`
```
$ mv source target
```
`mv` will move the source path to the destination path. Works for both directories as files.

### Removing files: `rm`
**Note: there are NO backups, there is no ’trash bin’. If you remove files/directories,
they are gone.**

```
$ rm filename
```
`rm` will remove a file or directory. (`rm -rf directory` will remove every file inside a given
directory). WARNING: files removed will be lost forever, there are no backups, so beware when
using this command!

#### Removing a directory: `rmdir`

You can remove directories using `rm -r` directory, however, this is error prone and can ruin
your day if you make a mistake in typing. To prevent this type of error, you can remove the
contents of a directory using `rm` and then finally removing the directory with:

```
$ rmdir directory
```

#### Changing permissions: `chmod`

Every file, directory, and link has a set of permissions. These permissions consist of permission
groups and permission types. The permission groups are:
1. User - a particular user (account)
2. Group - a particular group of users (may be user-specific group with only one member)
3. Other - other users in the system

The permission types are:
1. Read - For files, this gives permission to read the contents of a file
2. Write - For files, this gives permission to write data to the file. For directories it allows
users to add or remove files to a directory.
3. Execute - For files this gives permission to execute a file as through it were a script. For
directories, it allows users to open the directory and look at the contents.

Any time you run `ls -l` you’ll see a familiar line of `-rwx------` or similar combination of the
letters `r`, `w`, `x` and `-` (dashes). These are the permissions for the file or directory. (See also the [previous section on permissions](#permissions))

```
$ ls -l
total 1
-rw-r--r--. 1 alice ucy-cs499-dc-PG0 4283648 Apr 12 15:13 articleTable.csv
drwxr-x---. 2 alice ucy-cs499-dc-PG0      40 Apr 12 15:00 Project_GoldenDragon
```

Here, we see that `articleTable.csv` is a file (beginning the line with `-`) has read and write
permission for the user `alice` (`rw-`), and read permission for the group `ucy-cs499-dc-PG0` as well
as all other users (r-- and r--).

The next entry is `Project_GoldenDragon`. We see it is a directory because the line begins
with a `d`. It also has read, write, and execute permission for the `alice` user (rwx). So that
user can look into the directory and add or remove files. Users in the `ucy-cs499-dc-PG0` can also look into the directory and read the files. But they can’t add or remove files (`r-x`). Finally, other
users can read files in the directory, but other users have no permissions to look in the directory
at all (`---`).

Maybe we have a colleague who wants to be able to add files to the directory. We use chmod to
change the modifiers to the directory to let people in the group write to the directory:

```
$ chmod g+w Project_GoldenDragon
$ ls -l
total 1
-rw-r--r--. 1 alice ucy-cs499-dc-PG0 4283648 Apr 12 15:13 articleTable.csv
drwxrwx---. 2 alice ucy-cs499-dc-PG0      40 Apr 12 15:00 Project_GoldenDragon
```

The syntax used here is `g+x` which means group was given write permission. To revoke it again,
we use `g-w`. The other roles are `u` for user and `o` for other.

You can put multiple changes on the same line: `chmod o-rwx,g-rxw,u+rx,u-w somefile`
will take everyone’s permission away except the user’s ability to read or execute the file.

You can also use the `-R` flag to affect all the files within a directory, but this is dangerous. It’s
best to refine your search using `find` and then pass the resulting list to `chmod` since it’s not
usual for all files in a directory structure to have the same permissions.

#### Access control lists (ACLs)
However, this means that all users in `ucy-cs499-dc-PG0` can add or remove files. This could be problematic if you only wanted one person to be allowed to help you administer the files in the project. We need a new group. To do this in the HPC environment, we need to use access control lists
(ACLs):

```
$ setfacl -m u:otheruser:w Project_GoldenDragon
$ ls -l Project_GoldenDragon
drwxr-x---+ 2 alice ucy-cs499-dc-PG0 40 Apr 12 15:00 Project_GoldenDragon
```

This will give the user `otheruser` permissions to write to `Project_GoldenDragon`

Now there is a `+` at the end of the line. This means there is an ACL attached to the directory.
`getfacl Project_GoldenDragon` will print the ACLs for the directory.

Note: most people don’t use ACLs, but it’s sometimes the right thing and you should be aware
it exists.

See [setfacl](https://linux.die.net/man/1/setfacl) for more information.

### Zipping: `gzip`/`gunzip`, `zip`/`unzip`
Files should usually be stored in a compressed file if they’re not being used frequently. This
means they will use less space and thus you get more out of your quota. Some types of files (e.g.,
CSV files with a lot of numbers) compress as much as 9:1. The most commonly used compression
format on Linux is gzip. To compress a file using gzip, we use:

```
$ ls -lh myfile
-rw-r--r--. 1 alice ucy-cs499-dc-PG0 4.1M Dec 2 11:14 myfile
$ gzip myfile
$ ls -lh myfile.gz
-rw-r--r--. 1 alice ucy-cs499-dc-PG0 1.1M Dec 2 11:14 myfile.gz
```
Note: if you zip a file, the original file will be removed. If you unzip a file, the compressed file
will be removed. To keep both, we send the data to stdout and redirect it to the target file:

```
$ gzip -c myfile > myfile.gz
$ gunzip -c myfile.gz > myfile
```

#### `zip` and `unzip`
Windows and macOS seem to favour the zip file format, so it’s also important to know how to
unpack those. We do this using unzip:

```
$ unzip myfile.zip
```

If we would like to make our own zip archive, we use zip:

```
$ zip myfiles.zip myfile1 myfile2 myfile3
```

### Working with tarballs: `tar`
Tar stands for “tape archive” and is a way to bundle files together in a bigger file.

You will normally want to unpack these files more often than you make them. To unpack a `.tar`
file you use:

```
$ tar -xf tarfile.tar
```
Often, you will find gzip compressed .tar files on the web. These are called tarballs. You can
recognize them by the filename ending in .tar.gz. You can uncompress these using gunzip
and then unpacking them using tar. But tar knows how to open them using the -z option:

```
$ tar -zxf tarfile.tar.gz
$ tar -zxf tarfile.tgz
```

#### Order of arguments
Note: Archive programs like zip, tar, and jar use arguments in the ”opposite direction” of
copy commands.

```
# cp, ln: <source(s)> <target>
$ cp source1 source2 source3 target
$ ln -s source target

# zip, tar: <target> <source(s)>
$ zip zipfile.zip source1 source2 source3
$ tar -cf tarfile.tar source1 source2 source3
```

If you use tar with the source files first then the first file will be overwritten. You can control
the order of arguments of `tar` if it helps you remember:

```
$ tar -c source1 source2 source3 -f tarfile.tar
```

### Exercises
1. Create a subdirectory in your home directory named test containing a single, empty file
named `one.txt`.
2. Copy `/etc/hostname` into the test directory and then check what’s in it. Rename the
file to `hostname.txt`.
3. Make a new directory named `another` and copy the entire `test` directory to it. `another
/test/one.txt` should then be an empty file.
4. Remove the `another/test` directory with a single command.
5. Rename `test` to `test2`. Move `test2/hostname.txt` to your home directory.
6. Change the permission of `test2` so only you can access it.
7. Create an empty `job` script named `job.sh`, and make it executable.
8. gzip `hostname.txt`, see how much smaller it becomes, then unzip it again.

[Back to Top](#table-of-contents)

## Uploading/downloading/editing files

### Uploading/downloading files


SCP (secure copy) is a command-line utility that allows you to securely copy files and directories between two locations.

To copy a file from a local to a remote system run the following command:

```
scp file.txt remote_username@10.10.0.2:/remote/directory
```

Where `file.txt` is the name of the file we want to copy, `remote_username` is the user on the remote server, `10.10.0.2` is the server IP address. The `/remote/directory` is the path to the directory you want to copy the file to. If you don’t specify a remote directory, the file will be copied to the remote user home directory.

To copy a directory from a local to remote system, use the `-r` option:

```
scp -r /local/directory remote_username@10.10.0.2:/remote/directory
```

To copy a file from a remote to a local system, use the remote location as a source and local location as the destination.

For example to copy a file named `file.txt` from a remote server with IP `10.10.0.2` run the following command:

```
scp remote_username@10.10.0.2:/remote/file.txt /local/directory
```

### Symlinks for data/scratch

As we end up in the home directory when connecting, it would be convenient if we could access
our data and VO storage. To facilitate this we will create symlinks to them in our home directory. This will create 4 symbolic links (they’re like “shortcuts” on your desktop) pointing to the
respective storages:

```
$ cd $HOME
$ ln -s $VSC_SCRATCH scratch
$ ln -s $VSC_DATA data
$ ls -l scratch data
lrwxrwxrwx 1 alice ucy-cs499-dc-PG0 31 Mar 27 2009 data -> /data/brussel/100/vsc10002
lrwxrwxrwx 1 alice ucy-cs499-dc-PG0 34 Jun 5 2012 scratch -> /scratch/brussel/100/vsc10002
```

### Editing with nano
Nano is the simplest editor available on Linux. To open Nano, just type `nano`. To edit a file,
you use `nano the_file_to_edit.txt`. You will be presented with the contents of the file
and a menu at the bottom with commands like `^O Write Out` The `^` is the Control key. So
`^O` means `Ctrl-O`. The main commands are:
1. Open ("Read"): `^R`
2. Save ("Write Out"): `^O`
3. Exit: `^X`

More advanced editors (beyond the scope of this page) are `vim` and `emacs`. A simple tutorial
on how to get started with `vim` can be found at [https://www.openvim.com/](https://www.openvim.com/).

### Copying faster with rsync

`rsync` is a fast and versatile copying tool. It can be much faster than `scp` when copying large
datasets. It’s famous for its “delta-transfer algorithm”, which reduces the amount of data sent
over the network by only sending the differences between files.

You will need to run `rsync` from a computer where it is installed. Installing `rsync` is the easiest
on Linux: it comes pre-installed with a lot of distributions.

For example, to copy a folder with lots of CSV files:

```
$ rsync -rzv testfolder alice@hydra.vub.ac.be:data/
```

will copy the folder `testfolder` and its contents to `$VSC_DATA` on the `VUB-HPC`, assuming
the data symlink is present in your home directory, see section [above](#symlinks-for-datascratch).

The `-r` flag means “recursively”, the -z flag means that compression is enabled (this is especially handy when dealing with CSV files because they compress well) and the -v enables more
verbosity (more details about what’s going on).

To copy large files using `rsync`, you can use the `-P` flag: it enables both showing of progress
and resuming partially downloaded files.

To copy files from the VUB-HPC to your local computer, you can also use rsync:

```
$ rsync -rzv alice@hydra.vub.ac.be:data/bioset local_folder
```

This will copy the folder bioset and its contents that on $VSC_DATA of the VUB-HPC to a
local folder named local_folder.

See `man rsync` or https://linux.die.net/man/1/rsync for more information about `rsync`.

[Back to Top](#table-of-contents)

### Exercises
1. Download the file `/etc/hostname` to your local computer.
2. Upload a file to a subdirectory of your personal `$VSC_DATA` space.
3. Create a file named `hello.txt` and edit it using nano.

[Back to Top](#table-of-contents)


## Beyond the basics

Now that you’ve seen some of the more basic commands, let’s take a look at some of the deeper
concepts and commands.

### Input/output
To redirect output to files, you can use the redirection operators: `>`, `>>`, `&>`, and `<`.

First, it’s important to make a distinction between two different output channels:
1. `stdout`: standard output channel, for regular output
2. `stderr`: standard error channel, for errors and warnings

#### Redirecting `stdout`
`>` writes the (`stdout`) output of a command to a file and overwrites whatever was in the file
before.
```
$ echo hello > somefile
$ cat somefile
hello
$ echo hello2 > somefile
$ cat somefile
hello2
```

`>>` appends the (`stdout`) output of a command to a file; it does not clobber whatever was in
the file before:
```
$ echo hello > somefile
$ cat somefile
hello
$ echo hello2 >> somefile
$ cat somefile
hello
hello2
```

#### Reading from `stdin`
`<` reads a file from standard input (piped or typed input). So you would use this to simulate
typing into a terminal. `< somefile.txt` is largely equivalent to `cat somefile.txt |` .

One common use might be to take the the results of a long running command and store the
results in a file so you don’t have to repeat it while you refine your command line. For example,
if you have a large directory structure you might save a list of all the files you’re interested in
and then reading in the file list when you are done:

```
$ find . -name .txt > files
$ xargs grep banana < files
```

#### Redirecting `stderr`
To redirect the `stderr` output (warnings, messages), you can use `2>`, just like `>`

```
$ ls one.txt nosuchfile.txt 2> errors.txt
one.txt
$ cat errors.txt
ls: nosuchfile.txt: No such file or directory
```

#### Combining `stdout` and `stderr`
To combine both output channels (`stdout` and `stderr`) and redirect them to a single file, you
can use `&>`
```
$ ls one.txt nosuchfile.txt &> ls.out
$ cat ls.out
ls: nosuchfile.txt: No such file or directory
one.txt
```

### Command piping
Part of the power of the command line is to string multiple commands together to create useful
results. The core of these is the pipe: `|`. For example to see the number of files in a directory,
we can pipe the (`stdout`) output of `ls` to `wc` (word count, but can also be used to count the
number of lines with the `-l` flag).
```
$ ls | wc -l
42
```

A common pattern is to to pipe the output of a command to less so you can examine or search
the output:
```
$ find . | less
```
Or to look through your command history:

```
$ history | less
```
You can put multiple pipes in the same line. For example, which `cp` commands have we run?
```
$ history | grep cp | less
```

### Shell expansion
The shell will expand certain things, including:
1. `*` wildcard: for example `ls t*txt` will list all files starting with ’t’ and ending in ’txt’
2. tab completion: hit the `<tab>` key to make the shell complete your command line; works
for completing file names, command names, etc.
3. `$...` or `${...}`: environment variables will be replaced with their value; example: `echo
"I am $USER"` or `echo "I am ${USER}"`
4. square brackets can be used to list a number of options for a particular characters; example:
`ls *.[oe][0-9]`. This will list all files starting with whatever characters (`*`), then a
dot (`.`), then either an ‘o’ or an ‘e’ (`[oe]`), then a character from ‘0’ to ‘9’ (so any digit)
(`[0-9]`). So this filename will match: `anything.o5`, but this one won’t: `anything.o52`.

### Process information

#### `ps` and `pstree`
`ps` lists processes running. By default, it will only show you the processes running in the local
shell. To see all of your processes running on the system, use:

```
$ ps -fu $USER
```
To see all the processes

```
$ ps -elf
```
To see all the processes in a forest view, use:

```
$ ps auxf
```

The last two will spit out a lot of data, so get in the habit of piping it to `less`.

`pstree` is another way to dump a tree/forest view. It looks better than `ps auxf` but it has
much less information so its value is limited.

`pgrep` will find all the processes where the name matches the pattern and print the process IDs
(PID). This is used in piping the processes together as we will see in the next section.

#### `kill`
`ps` isn’t very useful unless you can manipulate the processes. We do this using the kill
command. Kill will send a message (`SIGINT`) to the process to ask it to stop.

```
$ kill 1234
$ kill $(pgrep misbehaving_process)
```

Usually this ends the process, giving it the opportunity to flush data to files, etc. However, if
the process ignored your signal, you can send it a different message (`SIGKILL`) which the OS
will use to unceremoniously terminate the process:

```
$ kill -9 1234
```

#### `top`
`top` is a tool to see the current status of the system. You’ve probably used something similar
in Task Manager on Windows or Activity Monitor in macOS. `top` will update every second and
has a few interesting commands.

To see only your processes, type `u` and your username after starting `top`, (you can also do this
with `top -u $USER`). The default is to sort the display by `%CPU`. To change the sort order,
use `<` and `>` like arrow keys.

There are a lot of configuration options in `top`, but if you’re interested in seeing a nicer view,
you can run `htop` instead. Be aware that it’s not installed everywhere, while top is.

To exit top, use `q` (for ’quit’).

For more information, see Brendan Gregg’s excellent site dedicated to performance analysis.

#### `ulimit`

`ulimit` is a utility to get or set the user limits on the machine. For example, you may be limited
to a certain number of processes. To see all the limits that have been set, use:

```
$ ulimit -a
```

### Counting: `wc`

To count the number of lines, words and characters (or bytes) in a file, use `wc` (word count):

```
$ wc example.txt
90 468 3189 example.txt
```
The output indicates that the file named `example.txt` contains 90 lines, 468 words and 3189
characters/bytes.

To only count the number of lines, use `wc -l`:

```
$ wc -l example.txt
90 example.txt
```

### Searching file contents: `grep`
`grep` is an important command. It was originally an abbreviation for “globally search a regular
expression and print” but it’s entered the common computing lexicon and people use ’grep’ to
mean searching for anything. To use grep, you give a pattern and a list of files.

```
$ grep banana fruit.txt
$ grep banana fruit_bowl1.txt fruit_bowl2.txt
$ grep banana fruit*txt
```

`grep` also lets you search for [Regular Expressions](https://en.wikipedia.org/wiki/Regular_expression), but these are not in scope for this introductory text.

### cut
`cut` is used to pull fields out of files or pipes streams. It’s a useful glue when you mix it with
`grep` because `grep` can find the lines where a string occurs and cut can pull out a particular
field. For example, to pull the first column (`-f 1`, the first field) from (an unquoted) CSV
(comma-separated values, so `-d ’,’`: delimited by `,`) file, you can use the following:

```
$ cut -f 1 -d ’,’ mydata.csv
```

### sed

`sed` is the stream editor. It is used to replace text in a file or piped stream. In this way it works
like grep, but instead of just searching, it can also edit files. This is like “Search and Replace” in
a text editor. `sed` has a lot of features, but most everyone uses the extremely basic version of
string replacement:

```
$ sed ’s/oldtext/newtext/g’ myfile.txt
```

By default, `sed` will just print the results. If you want to edit the file inplace, use `-i`, but be very careful that the results will be what you want before you go around destroying your data!

### awk
`awk` is a basic language that builds on sed to do much more advanced stream editing. Going in
depth is far out of scope of this tutorial, but there are two examples that are worth knowing.

First, `cut` is very limited in pulling fields apart based on whitespace. For example, if you have
padded fields then `cut -f 4 -d ’ ’` will almost certainly give you a headache as there might be an uncertain number of spaces between each field. `awk` does better whitespace splitting. So, pulling out the fourth field in a whitespace delimited file is as follows:

```
$ awk ’{print $4}’ mydata.dat
```

You can use `-F ’:’` to change the delimiter (F for field separator).

The next example is used to sum numbers from a field:
```
$ awk -F ’,’ ’{sum += $1} END {print sum}’ mydata.csv
```

### Basic Shell Scripting

The basic premise of a script is to execute automate the execution of multiple commands. If
you find yourself repeating the same commands over and over again, you should consider writing
one script to do the same. A script is nothing special, it is just a text file like any other. Any
commands you put in there will be executed from the top to bottom.

However there are some rules you need to abide by.

Here is a [very detailed guide](http://www.tldp.org/LDP/Bash-Beginners-Guide/html/) should you need more information.

#### Shebang
The first line of the script is the so called shebang (`#` is sometimes called hash and `!` is sometimes called bang). This line tells the shell which command should execute the script. In the most
cases this will simply be the shell itself. The line itself looks a bit weird, but you can copy paste
this line as you need not worry about it further. It is however very important this is the very
first line of the script! These are all valid shebangs, but you should only use one of them:

```
#!/bin/sh
```

```
#!/bin/bash
```

```
#!/usr/bin/env bash
```

#### Conditionals

Sometimes you only want certain commands to be executed when a certain condition is met. For
example, only move files to a directory if that directory exists. The syntax:

```
if [ -d directory ] && [ -f file ]
then
  mv file directory
fi
```

Or you only want to do something if a file exists:

```
if [ -f filename ]
then
  echo "it exists"
fi
```

Or only if a certain variable is bigger than one

```
if [ $AMOUNT -gt 1 ]
then
  echo "More than one"
  # more commands
fi
```

Several pitfalls exist with this syntax. You need spaces surrounding the brackets, the `then` needs
to be on the beginning of a line. It is best to just copy this example and modify it.

In the initial example we used `-d` to test if a directory existed. There are [several more checks](http://tldp.org/LDP/Bash-Beginners-Guide/html/sect_07_01.html).

Another useful example, to test if a variable contains a value (so it’s not empty):

```
if [ -z $PBS_ARRAYID ]
then
  echo "Not an array job, quitting."
  exit 1
fi
```
the `-z` will check if the length of the variable’s value is greater than zero.

#### Loops
Are you copy pasting commands? Are you doing the same thing with just different options? You
most likely can simplify your script by using a loop.

Let’s look at a simple example:

```
for i in 1 2 3
do
  echo $i
done
```

#### Subcommands
Subcommands are used all the time in shell scripts. What they basically do is storing the output
of a command in a variable. So this can later be used in a conditional or a loop for example.

```
CURRENTDIR=‘pwd‘ # using backticks
CURRENTDIR=$(pwd) # recommended (easier to type)
```

In the above example you can see the 2 different methods of using a subcommand. `pwd` will output the current working directory, and its output will be stored in the `CURRENTDIR` variable.
The recommend way to use subcommands is with the `$()` syntax.

#### Errors
Sometimes some things go wrong and a command or script you ran causes an error. How do you
properly deal with these situations?

Firstly a useful thing to know for debugging and testing is that you can run any command like this:
```
command 2>&1 output.log # one single output file, both output and errors
```

If you add `2>&1 output.log` at the end of any command, it will combine `stdout` and `stderr`, outputting it into a single file named `output.log`.

If you want regular and error output separated you can use:

```
command > output.log 2> output.err # errors in a separate file
```
this will write regular output to output.log and error output to output.err.
You can then look for the errors with less or search for specific text with grep.

In scripts you can use
```
set -e
```
this will tell the shell to stop executing any subsequent commands when a single command in the script fails. This is most convenient as most likely this causes the rest of the script to fail as well.

##### Advanced error checking
Sometimes you want to control all the error checking yourself, this is also possible. Everytime you
run a command, a special variable `$?` is used to denote successful completion of the command. A value other than zero signifies something went wrong. So an example use case:

```
command_with_possible_error
exit_code=$? # capture exit code of last command
if [ $exit_code -ne 0 ]
then
  echo "something went wrong"
fi
```

### `.bashrc` login script
If you have certain commands executed every time you log in (which includes every time a job
starts), you can add them to your `$HOME/.bashrc` file. This file is a shell script that gets
executed every time you log in.

Examples include:
- modifying your `$PS1` (to tweak your shell prompt)
- printing information about the current/jobs environment (echoing environment variables,
etc.)
- selecting a specific cluster to run on with `module swap cluster/...`

Some recommendations:
- Avoid using module load statements in your `$HOME/.bashrc` file
- Don’t directly edit your `.bashrc` file: if there’s an error in your `.bashrc` file, you might
not be able to log in again. In order to prevent that, use another file to test your changes,
then copy them over when you tested the script.

## Common Pitfalls

### Files
####Location
If you receive an error message which contains something like the following:
```
No such file or directory
```
It probably means that you haven’t placed your files in the correct directory or you have mistyped
the file name or path.

Try and figure out the correct location using `ls`, `cd` and using the different `$VSC_*` variables.

#### Spaces
Filenames should not contain any spaces! If you have a long filename you should use underscores
or dashes (e.g., very_long_filename).
```
$ cat some file
No such file or directory ’some’
```
Spaces are permitted, however they result in surprising behaviour. To cat the file `some file`
as above, you can escape the space with a backslash (“\ ”) or you can put the filename in quotes:

```
$ cat some\ file
...
$ cat "some file"
...
```

This is especially error prone if you are piping results of `find`:

```
$ find . -type f | xargs cat
No such file or directory name ’some’
No such file or directory name ’file’
```

This can be worked around using the `-print0` flag:

```
$ find . -type f -print0 | xargs -0 cat
...
```

But, this is tedious and you can prevent errors by simply colouring within the lines and not using
spaces in filenames.

#### Missing/mistyped environment variables
If you use a command like `rm -r` with environment variables you need to be careful to make
sure that the environment variable exists. If you mistype an environment variable then it will
resolve to a blank string. This means the following resolves to `rm -r ~/*` which will remove
every file in your home directory!

```
$ rm -r ∼/$PROJETC/*
```

#### Typing dangerous commands
A good habit when typing dangerous commands is to precede the line with `#`, the comment
character. This will let you type out the command without fear of accidentally hitting enter and
running something unintended.

```
$ #rm -r ∼/$POROJETC/*
```

Then you can go back to the beginning of the line (`Ctrl-A`) and remove the first character (`Ctrl
-D`) to run the command. You can also just press enter to put the command in your history so
you can come back to it later (e.g., while you go check the spelling of your environment variables).

#### Permissions

```
$ ls -l script.sh # File with correct permissions
-rwxr-xr-x 1 alice ucy-cs499-dc-pg0 2983 Jan 30 09:13 script.sh
$ ls -l script.sh # File with incorrect permissions
-rw-r--r-- 1 alice ucy-cs499-dc-pg0 2983 Jan 30 09:13 script.sh
```

Before submitting the script, you’ll need to add execute permissions to make sure it can be
executed:

```
$ chmod +x script_name.sh
```

#### Help
If you stumble upon an error, don’t panic! Read the error output, it might contain a clue as to
what went wrong. You can copy the error message into Google (selecting a small part of the
error without filenames). It can help if you surround your search terms in double quotes (for
example "`No such file or directory`"), that way Google will consider the error as one
thing, and won’t show results just containing these words in random order.

If you need help about a certain command, you should consult its so called “man page”:

```
$ man command
```

This will open the manual of this command. This manual contains detailed explanation of all
the options the command has. Exiting the manual is done by pressing ’q’.


## ACK

https://hpcugent.github.io/vsc_user_docs/pdf/intro-Linux-windows-gent.pdf
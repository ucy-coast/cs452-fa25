## Secure Shell (ssh)

Secure Shell (ssh) is the standard way to connect to a remote machine’s shell. We can use `ssh` to log in, establish an interactive session, and then execute commands to perform various administrative tasks on the remote machine. However, using interactive SSH sessions can often get tedious as we need to take the separate actions of connecting and authenticating, interactively running the command, and then disconnecting. 

`ssh` allows us to execute a command on a remote machine without logging into that machine. 

To run a single command on a remote machine instead of spawning a shell session, you can add the command after the connection information, like this:

```
ssh username@remote_host command_to_run 
```

This will connect to the remote host, authenticate with your credentials, and execute the command you specified. The connection will immediately close afterwards.

When you skip the username, `ssh` will attempt to remote connect to the machines using your current user name.

### Single command execution

Let us execute single command `date` to fetch from the remote machine:

```
ssh node1 date
```

### Command redirection

It’s possible to redirect the output of a command executed on the remote server to the local machine using the redirection operator(>):

```
ssh node1 'dmesg' > /tmp/dmesg-node1
```

### Multi-command execution

To execute multiple commands, each command needs to be separated using a semicolon(;) to be enclosed within a single quote or double quote:

```
ssh node1 'date;uname;who'
```

### Multi-command execution using pipe

Using an unnamed pipe, let us try to see when the root user last logged onto the remote server,

```
ssh node1 'lastlog | grep root'
```

### Multi-command execution from a local file	

It’s possible to run commands from a local file. The `-s` option of bash helps to read the executable command from the standard input:

```
ssh node1 'bash -s' < scripts/get_host_info.sh
```

### Command execution with privileges

At times, we may need to execute commands with elevated privileges. The user `alice` in the remote server doesn’t have to write permission in the /etc/ directory.

```
ssh node1 'touch /etc/config'
```

If you need to elevate your privileges on the far side of the SSH connection with sudo, then force the use of a pseudo-terminal with -t. Use this if sudo will challenge you for a password. The command looks like this:

```
ssh -t node1 'sudo touch /etc/config'
```

### Command execution on multiple nodes

It is common to execute commands on many nodes/hosts via SSH for managing a cluster. Generally, to run commands on many nodes, there are two modes: serial mode and parallel mode. In serial mode, the command is executed on the node one by one. In parallel mode, the command is executed on many nodes together. The serial mode is easy to reason about with and debug while the parallel mode is usually much faster. For a short-running task this might not matter much, but if a task needs an hour to complete and you need to run it on 20 hosts, parallel execution beats serial by a mile.

Run commands in serial order (one by one) using Bash over SSH

```
for h in node{1..3} ; do
    ssh $h hostname
done
```

Run commands in parallel using Bash over SSH

```
for h in node{1..3} ; do
    ssh $h hostname &
done
wait
```


## parallel-ssh

[`parallel-ssh`](https://manpages.org/parallel-ssh) is a program for executing ssh in parallel on a number of hosts. 

`parallel-ssh` allows you to execute commands on multiple nodes in parallel at the same time. While you could accomplish this workflow using a combination of `ssh` and a bash `for` loop as [described above](#command-execution-on-multiple-nodes), `parallel-ssh` provides for a more natural experience.

On Debian based distributions, you can install `parallel-ssh` using aptitude.

```
sudo apt -y install pssh
```

You can create a text file called hosts file which can be used as an input to `parallel-ssh`. The syntax is pretty simple. Each line in the host file are of the form [user@]host[:port] and can include blank lines and comments lines beginning with “#”. Here is a sample file named `pssh_hosts`:

```
node1
node2
node3
```

You can run `parallel-ssh` for all the hosts from this file to execute the `date` command. You can use the `-i` option to display standard output and standard error as each host completes.

```
parallel-ssh -i -h pssh_hosts date
```

Sample output:

```
[1] 06:48:49 [SUCCESS] node1
Wed Jul  6 06:48:49 CDT 2022
[2] 06:48:49 [SUCCESS] node2
Wed Jul  6 06:48:49 CDT 2022
[3] 06:48:49 [SUCCESS] node3
Wed Jul  6 06:48:49 CDT 2022
```

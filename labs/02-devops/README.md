# Lab: DevOps

In this lab, you’ll get hands-on experience with core command-line tools used to automate and manage infrastructure, including Git, SSH, Bash, and parallel-ssh.

These tools are foundational for modern software delivery and serve as essential building blocks for DevOps workflows, even in their simplest form. While we won’t build a full DevOps pipeline in this lab, the practices you learn here, including scripting, version control, and remote automation, are central to the DevOps philosophy of repeatable, reliable, and collaborative deployment.

By the end of this lab, you will have:
- Written and versioned infrastructure scripts using Git
- Used SSH to remotely configure servers
- Automated multi-node deployments with parallel-ssh

This foundation will support more advanced labs later in the semester, where we introduce Docker, CI/CD, and orchestration tools.

## Background

### DevOps

Traditional software delivery models often lead to friction between development and operations teams. Developers want to release features quickly, while operations prioritize stability and reliability. This mismatch in priorities results in delays and errors during deployment.

DevOps aims to bridge this gap by fostering a collaborative culture and integrating automated tools into the development and deployment pipeline. This approach supports Continuous Integration (CI) and Continuous Delivery (CD), enabling faster, safer, and more consistent software releases.

DevOps is not a single tool or practice, it is a culture shift supported by:

- Practices: such as CI/CD, Infrastructure as Code, and automated testing.
- Tools: like Git, Jenkins, Docker, Kubernetes, Ansible, and more.
- Philosophy: emphasizing collaboration, transparency, and automation.

These components help streamline the software delivery lifecycle from development to production. 

### From Manual to Automated: A Deployment Example

To see DevOps principles in action, we'll study a simple example: deploying an nginx web server.

Here’s a basic shell script `nginx.sh` that installs and starts nginx:

```bash
#!/bin/bash
sudo apt update
sudo apt install -y nginx
sudo systemctl enable nginx
sudo systemctl start nginx
```

This works well on a single machine, but it has limitations:
- You must run it manually each time
- There’s no version control, so it’s hard to track changes or share with others
- It doesn’t scale well when deploying to multiple machines

To improve this process, we’ll bring in key tools that help make deployment more structured, repeatable, and easier to scale.

Tool | Role in DevOps | How we’ll use it
:-----|:-----------------|:------------------
Git  | Version control | Track changes to scripts and static files
SSH  | Remote access | Execute commands and transfer files on remote servers
parallel-ssh | Parallel remote execution | Run scripts across multiple nodes simultaneously

## Prerequisites

### Setting up the Experiment Environment in Cloudlab

For this tutorial, you will be using a small CloudLab cluster. 

Start a new experiment on CloudLab using the `multi-node-cluster` profile in the `UCY-COAST-TEACH` project, configured with four (4) physical machine nodes. 

### Brashing up essential shell skills

Before working with remote systems and automation tools, make sure you're comfortable with the Unix command line. These skills are essential for scripting, remote access, and troubleshooting.

Try the following commands on your `node0` terminal:

```bash
cd ~
mkdir lab-test        # Create a test directory
cd lab-test           # Enter the directory

# Create a simple Bash script using a here-document
cat << EOF > hello.sh
#!/bin/bash
echo "Hello, world"
EOF

cat hello.sh          # Display the file content
chmod +x hello.sh     # Make the script executable
./hello.sh            # Run the script (prints: Hello, world)
```

This quick exercise gives you a feel for:
- Creating and running Bash scripts
- Using `chmod` to make files executable
- Writing multi-line files directly from the command line with `cat << EOF`

Other useful commands to know:
- `ls`, `pwd`, `rm`, `man`, `sudo`
- `curl` to check if a web server is running
- `ps`, `top` to inspect running processes

See the [Linux Survival Guide](../../docs/linux/linux-basic.md) for a quick reference on Bash commands, file permissions, and more.

## Part 1: Version Control with Git

In a DevOps workflow, infrastructure is treated as code. That means it’s just as important to track changes to your deployment scripts, configuration files, and automation playbooks as it is to track application source code.

This is where Git comes in, a version control system that plays a foundational role in DevOps.
- Track infrastructure changes: Git allows you to maintain a clear history of how your infrastructure code evolves over time.
- Collaborate efficiently: Multiple team members can work on scripts or configurations without overwriting each other’s changes, thanks to Git’s branching and merging capabilities.
- Recover quickly: If something breaks, you can easily revert to a known good state using Git’s commit history.

Using Git ensures transparency, accountability, and consistency across your deployment process—core values in any DevOps practice.

A typical workflow involves editing and working on your content in your local repository on your computer, and then sending your changes to the remote repository on GitHub.

At this point, it’s important to distinguish between Git and GitHub:

- [Git](https://git-scm.com/) is the version control tool itself. It tracks changes to your files and allows you to manage different versions of your code.
- [GitHub](https://github.com/) is a cloud-based platform for hosting Git repositories. It adds collaboration features like pull requests, issues, and project management tools. You can think of it as a social and collaborative layer on top of Git.

If you don’t have a GitHub account yet, go [here](https://github.com/signup) and create one, it will be essential for the rest of this lab.

### Getting started

Open a remote SSH terminal session to `node0`. Make sure you enable forwarding of the authentication agent connection.

Make sure you have already installed Git on your working machine (`node0`). To check if git is already installed, open the Terminal and hit:

```
git --version
```

To install the latest version of git:

```
sudo apt -y install git
```

When you're using Git for the first time on a machine, you should set your global username and email:

```bash
git config --global user.name "Your Name"
git config --global user.email "you@example.com"
```

This ensures that all your commits on that system are properly identified, no matter which repository you're working in.

### Working locally 

To use Git you will have to setup a repository. You can take an existing directory to make a Git repository, or create an empty directory.

```bash
# create a new directory, and initialize it with git-specific functions
git init my-devops-lab

# change into the `my-devops-lab` directory
cd my-devops-lab

# git isn't aware of the file, stage it
git add nginx.sh

# take a snapshot of the staging area
git commit -m "Initial commit: NGINX setup script"
```

As your project evolves, you can continue tracking additional files with Git. For example, let’s add a simple static web page.

First, create a file named index.html:

```html
<h1>Hello from DevOps!</h1>
```

Then, stage and commit it:

```bash
# Stage the new HTML file
git add index.html

# Commit it with a message describing the change
git commit -m "Add static index.html"
```

Before you can push to GitHub for the first time, you’ll need to create a remote repository and connect it to your local Git project.

### Hosting your source code on GitHub

You need to [create](https://github.com/new) a repository for your project. **Do not** initialize the repository with a README, .gitignore or License file. This empty repository will await your code.

<figure>
  <p align="center"><img src="assets/images/github-create-repo-public.png" width="60%"></p>
  <figcaption><p align="center">Figure. Creating a new GitHub repository </p></figcaption>
</figure>

Before you can push commits made on your local branch to a remote repository, you will need to provide the path for the repository you created on GitHub and rename your local branch.

#### Providing the path to the remote repository 

You can push commits made on your local branch to a remote repository identified by a remote URL. A remote URL is Git's fancy way of saying "the place where your code is stored". That URL could be your repository on GitHub, or another user's fork, or even on a completely different server.

You can only push to two types of URL addresses:
- An HTTPS URL like `https://github.com/user/repo.git`. The `https://` URLs are available on all repositories, regardless of visibility. `https://` URLs work even if you are behind a firewall or proxy.
- An SSH URL, like `git@github.com:user/repo.git`. SSH URLs provide access to a Git repository via SSH, a secure protocol. To use these URLs, you must generate an SSH keypair on your computer and add the public key to your account on GitHub.com. For more information, see [Connecting to GitHub with SSH](https://docs.github.com/en/github/authenticating-to-github/connecting-to-github-with-ssh).

You can use the git remote add command to match a remote URL with a name. For example, you'd type the following in the command line:

```
git remote add origin  <REMOTE_URL> 
```

This associates the name origin with the `REMOTE_URL`.

You can use the command `git remote set-url` to [change a remote's URL](https://docs.github.com/en/github/getting-started-with-github/managing-remote-repositories).

To provide the path for the repository you created on GitHub using the SSH URL:

```
git remote add origin git@github.com:YOUR-USERNAME/YOUR-REPOSITORY-NAME.git
```

#### Renaming the default branch

Every Git repository has an initial branch, which is the first branch to be created when a new repository is generated. Historically, the default name for this initial branch was `master`, but `main` is an increasingly popular choice. 

Since the default branch name for new repositories created on GitHub is now `main`, you will have to to rename your local `master` branch to `main`:

```
git branch -M main
```

#### Pushing and fetching changes

You're now setup to push changes to the remote repository:

```
git push -u origin main
```

If you're already setup for push as above, then the following will bring changes down and merge them in:

```
git pull
```

What this is doing under-the-hood is running a `git fetch` and then `git merge`.

#### More commands

As you begin working with Git and version control, it's helpful to become more familiar with a [few more key command-line tools](docs/git.md). These will make your workflow more efficient and help you understand what Git is doing under the hood.

## Part 2: Remote Access with SSH

So far you wrote a local bash script to install and configure NGINX, and you used Git to track changes to your code. That gave you a solid foundation for managing infrastructure locally and using version control.

Now, you’ll take the next step in a DevOps workflow: using SSH to run your script on a remote server.

This is essential because real-world infrastructure is almost always deployed and managed remotely, not just on your local machine.

### What is SSH?
SSH (Secure Shell) is a protocol that allows you to securely access and control remote machines over a network. 

SSH is a fundamental tool in DevOps because it enables you to run commands on remote systems just as if you were working on your local terminal. SSH also supports secure file transfer between machines. Many DevOps tools, such as Ansible, Jenkins, and Terraform, rely on SSH to automate tasks like provisioning, configuration, and deployment.

In this part of the lab, we’ll use SSH to install NGINX on a remote machine by running the same script you created earlier, but this time, remotely.

### Running commands remotely with SSH
Here’s how to run your existing script on a remote server using SSH:

```bash
ssh username@remote-host 'bash -s' < nginx.sh
```

Explanation:
- `ssh username@remote-host` connects you to the remote machine.
- `'bash -s'` tells the remote server to run a bash shell and wait for input.
- `< nginx.sh` streams your local script into the remote shell for execution.

This means you don’t need to copy the script over first, and the remote machine installs NGINX as if you were logged in and typing it yourself.

Assuming you're currently working on `node0` and your `nginx.sh` script is available, you can run it directly on `node1` using SSH. 

If log with the same username on both nodes, there's no need to specify the username in the SSH command. You can simply use the hostname `node1` as the remote target.

Run the following command:

```bash
ssh node1 'bash -s' < nginx.sh
```

### Transferring files with SCP
Once NGINX is installed, you’ll want to serve a static web page. Let’s transfer an index.html file to the remote machine.

Using SCP, we can securely copy the file to a temporary directory on the remote host:

```bash
scp index.html node1:/tmp/
```

Using SSH, we can move the file into NGINX’s web directory (/var/www/html) on the remote server. Since this directory  =requires elevated privileges, we need to use `sudo` on the remote side. If sudo prompts for a password, we must also allocate a pseudo-terminal using the `-t` option in the SSH command. This ensures that the remote shell can handle the password prompt properly. The command looks like this:

```bash
ssh -t node1 'sudo cp /tmp/index.html /var/www/html/index.html'
```

Now the file is in the web server’s root directory, and it should be accessible in a browser.

Verify by visiting `http://<remote-host-ip>` and checking that your static page loads.

### Limitations

Using SSH with Bash scripts works well for small setups, but the approach quickly runs into limitations. It requires manual effort for every step, and doesn’t scale when you're managing more than a handful of machines.

For slightly larger environments, tools like [parallel-ssh](docs/parallel-ssh.md) offer a better solution by letting you run the same commands across multiple hosts simultaneously.

## Part 3: Remote Execution with Parallel SSH

For clusters with more than two machines, [`parallel-ssh`](https://manpages.org/parallel-ssh) becomes a highly effective tool for executing commands across multiple hosts simultaneously. It allows you to run the same SSH command in parallel on many machines, making it ideal for managing and automating tasks in larger deployments. For even more advanced orchestration and configuration management, especially at scale, tools like [Ansible](docs/ansible.md) offer additional flexibility and control.

Getting started with `parallel-ssh` is simple. 
1. Choose a machine as your management system and install the tool. Your management node can also be a managed node.
2. Ensure you have passwordless ssh access from your management system to each managed node.
3. Create a hosts file containing an inventory of your nodes.
4. Start using the tool 

We suggest the following:
- For step 1, you use `node0` as the management system.
- For step 2, you use a CloudLab profile that automatically sets up ssh keys for passwordless access. The `multi-node-cluster` profile under `UCY-COAST-TEACH` meets this requirement.

## Installing `parallel-ssh`

On Debian-based systems, you can install `parallel-ssh` using `apt`:

```
$ sudo apt -y install pssh
```

### Running commands

To use `parallel-ssh`, you need to create a text file called `hosts` file that contains a list of all the hosts that you want to have the command executed on:

```
node2
node3
```

To run simple commands like `date` on all hosts in the `hosts` file:

```
parallel-ssh -i -h hosts date
```

The `-i` flag ensures that output from each host is displayed as it completes, including both standard output and standard error.

Sample output:

```
[1] 06:48:49 [SUCCESS] node2
Wed Jul  6 06:48:49 CDT 2022
[2] 06:48:49 [SUCCESS] node3
Wed Jul  6 06:48:49 CDT 2022
```

As another example, you can run `apt` to install `nano` on each host:

```bash
parallel-ssh -i -h hosts -- sudo apt -y install nano
```

### Running a script

You can use input redirection to run a local script like `nginx.sh` remotely on all nodes at once.

Run this command on your management node:

```bash
parallel-ssh -h hosts -i -I < ./nginx.sh
```

This sends the content of `nginx.sh` over SSH to each node, where it’s executed by `bash`. This avoids the need to transfer or store the script on remote machines explicitly.

### Transferring files

Beyond running commands, it's often necessary to distribute files, such as install scripts, config files, or web content, to many machines. `parallel-scp` lets you copy files to multiple hosts at once, making it ideal for tasks like provisioning or deploying static content.

For example, to copy a file called `index.html` to the home directory (~/) of all nodes:

```bash
parallel-scp -h hosts index.html ~/
```

Since writing directly to a privileged path like `/var/www/html/index.html` requires elevated permissions, a safe pattern is to first copy the file to a user-writable location, then move it with sudo:

```bash
parallel-ssh -h hosts 'sudo mv ~/index.html /var/www/html/index.html'
```

This two-step approach ensures compatibility across systems where you might not have sudo access for scp, but do have it via ssh.

With this method, you can quickly deploy web content (or any other resource) to a whole cluster of machines—ideal for testing or distributed environments.

Now the file is in the web server’s root directory, and it should be accessible in a browser.

Verify by visiting `http://<remote-host-ip>` and checking that your static page loads.


## Exercise: Parallel Nginx Log Collection and Search

Objective: Collect and analyze nginx access logs from multiple servers.
- Use `parallel-ssh` to fetch the last 50 lines of the nginx access log (`/var/log/nginx/access.log`) on all hosts.
- Use `parallel-scp` to copy those logs back to your management machine into separate files (`access_node2.log`, `access_node3.log`, etc.).
- Write a simple bash script locally that searches through all the collected logs for HTTP status codes indicating errors (e.g., 4xx or 5xx).
- Summarize the number of errors found per host.

💡 Hint: To generate errors, try requesting missing pages for 404 errors: `curl http://nodeX/nonexistent`


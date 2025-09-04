# Configuration Management with Ansible

Ansible is a modern configuration management tool that facilitates the task of setting up and maintaining remote servers. With a minimalist design intended to get users up and running quickly, it allows you to control one to hundreds of systems from a central location with either ad hoc commands or playbooks.

## Configuration Management

Configuration and resource management is an automated method for maintaining computer systems and software in a known, consistent state. With configuration management, you no longer guess or hope that a configuration is current. It is correct because the configuration management system ensures that it is correct. When combined with automation, configuration management can improve efficiency and reduce risk due to manual error because manual configuration processes are replaced with automated processes. 

Infrastructure as code (IaC) is a form of configuration management that codifies an organization's infrastructure resources into text files. These infrastructure files are then committed to a version control system like Git.

## Ansible Overview 

Ansible is a modern configuration management tool that facilitates the task of setting up and maintaining remote servers, with a minimalist design intended to get users up and running quickly.

Users write Ansible provisioning scripts in [YAML](https://yaml.org/), a user-friendly data serialization standard that is not tied to any particular programming language. This enables users to create sophisticated provisioning scripts more intuitively compared to similar tools in the same category.

Ansible doesn’t require any special software to be installed on the nodes that will be managed with this tool. A control machine is set up with the Ansible software, which then communicates with the nodes via standard SSH.

## Getting started with Ansible

Ansible automates the management of remote systems and controls their desired state. A basic Ansible environment has three main components:

- **Control node**: A system on which Ansible is installed. You run Ansible commands such as `ansible` or `ansible-inventory` on a control node.

- **Managed node**: A remote system, or host, that Ansible controls.

- **Inventory**: A list of managed nodes that are logically organized. You create an inventory on the control node to describe host deployments to Ansible.

<figure>
  <p align="center"><img src="ansible_basic.svg" height=400></p>
  <figcaption><p align="center">Figure. Ansible Architecture.</p></figcaption>
</figure>

### Prerequisite

To follow this tutorial, you will need:

- **One Ansible Control Node**: The Ansible control node is the machine we’ll use to connect to and control the Ansible hosts over SSH. Your Ansible control node can either be your local machine or a server dedicated to running Ansible. Make sure the control node has a non-root user with sudo privileges and an SSH keypair associated with this user.
- **One or more Ansible Hosts**: An Ansible host is any machine that your Ansible control node is configured to automate. Make sure each Ansible host has the Ansible control node’s SSH public key added to the authorized_keys of a system user.

On CloudLab, you can easily meet the above requirements by creating a experiment using the profile `small-lan-basic`. This profile setups a cluster and enables SSH key-based authentication between cluster nodes. Although you can use any cluster node as the Ansible control node, we assume you will use `node0` as the Ansible control node and use the remaining nodes as the Ansible hosts.

### Installing Ansible On Your Control Machine

To take your first steps with Ansible, you first need to install it on your control machine. This is the machine you’ll use to dispatch tasks. On CloudLab, this will be machine `node0`. 

You can install Ansible using standard package managers:

```
sudo apt update
sudo apt -y install ansible
```

Your Ansible control node now has all of the software required to administer your hosts. Next, we’ll go over how to set up an inventory file, so that Ansible can communicate with your managed nodes.

### Setting up Ansible Inventory Files

The Ansible inventory file lists which hosts will receive commands from the control host. The inventory can list individual hosts, or group them under categories you distinguish.

The default location for the inventory file is /etc/ansible/hosts, but it’s possible to create inventory files in any location that better suits your needs. In this case, you’ll need to provide the path to your custom inventory file with the `-i` parameter when running Ansible commands. Using per-project inventory files is a good practice to minimize the risk of running a command on the wrong group of servers.

A typical inventory file can list the managed host either by IP address or by domain names. It is also possible to list one managed host in more than one group. 

Here’s an example of a project inventory file `hosts` listing two hosts under the webservers and dbservers categories.

```ini
[webservers]
node1

[dbservers]
node2
```

To test if all the hosts are discoverable by the inventory file, run the following:

```
ansible all -i ./hosts --list-hosts
```

```
  hosts (2):
    node2
    node1
```

You can also list the hosts by group name:

```
ansible dbservers -i ./hosts --list-hosts
```

```
  hosts (1):
    node2
```

Now that you’ve configured your inventory file, you have everything you need to test the connection to your Ansible hosts.

### Testing Connection

After setting up the inventory file to include your servers, it’s time to check if Ansible is able to connect to these servers and run commands via SSH.

You can use the `-u` argument to specify the remote system user. When not provided, Ansible will try to connect as your current system user on the control node.

From your local machine or Ansible control node, run:

```
ansible all -i ./hosts -m ping -u USER
```

This command will use Ansible’s built-in [`ping` module](https://docs.ansible.com/ansible/latest/modules/ping_module.html) to run a connectivity test on all nodes from your default inventory, connecting as `USER`. The ping module will test: if hosts are accessible; if you have valid SSH credentials; if hosts are able to run Ansible modules using Python.

You should get output similar to this:

```
node1 | SUCCESS => {
    "changed": false, 
    "ping": "pong"
}
node2 | SUCCESS => {
    "changed": false, 
    "ping": "pong"
}
```

Once you get a "pong" reply back from a host, it means you’re ready to run Ansible commands and playbooks on that server.

### Running Ad-Hoc Commands

After confirming that your Ansible control node is able to communicate with your hosts, you can start running ad-hoc commands and playbooks on your servers.

Ad-hoc commands in Ansible are merely those that perform a single command across one or many hosts. Any command that you would normally execute on a remote server over SSH can be run with Ansible on the servers specified in your inventory file. As an example, you can check disk usage on all servers with:

```
ansible all -i ./hosts -a "df -h" -u USER
```

You should get output similar to this:

```
node2 | SUCCESS | rc=0 >>
Filesystem                                   Size  Used Avail Use% Mounted on
udev                                          94G     0   94G   0% /dev
tmpfs                                         19G  1.7M   19G   1% /run
/dev/sda1                                     16G  2.7G   13G  18% /
tmpfs                                         94G     0   94G   0% /dev/shm
tmpfs                                        5.0M     0  5.0M   0% /run/lock
tmpfs                                         94G     0   94G   0% /sys/fs/cgroup
ops.wisc.cloudlab.us:/proj/ucy-cs499-dc-PG0  100G   14G   87G  14% /proj/ucy-cs499-dc-PG0
ops.wisc.cloudlab.us:/share                   50G  2.0G   49G   4% /share
tmpfs                                         19G     0   19G   0% /run/user/20001

node1 | SUCCESS | rc=0 >>
Filesystem                                   Size  Used Avail Use% Mounted on
udev                                          94G     0   94G   0% /dev
tmpfs                                         19G  1.7M   19G   1% /run
/dev/sda1                                     16G  2.7G   13G  18% /
tmpfs                                         94G     0   94G   0% /dev/shm
tmpfs                                        5.0M     0  5.0M   0% /run/lock
tmpfs                                         94G     0   94G   0% /sys/fs/cgroup
ops.wisc.cloudlab.us:/share                   50G  2.0G   49G   4% /share
ops.wisc.cloudlab.us:/proj/ucy-cs499-dc-PG0  100G   14G   87G  14% /proj/ucy-cs499-dc-PG0
tmpfs                                         19G     0   19G   0% /run/user/20001
```

The highlighted command `df -h` can be replaced by any command you’d like.

You can also execute Ansible modules via ad-hoc commands, similarly to what we’ve done before with the ping module for testing connection. 

Modules are discrete units of code that can be used from the command line or in a playbook task. Ansible executes each module, usually on the remote managed node, and collects return values.

When you dispatch a job from a control host to a managed host using an Ansible module, it is known as a task. 

Modules use the available context to determine what actions if any needed to bring the managed host to the desired state and are idempotent, that means if you run the same task again and again, the state of the machine will not change.

To find the list of available modules, use the following command:

```
ansible-doc -l
```

Let’s try to install Nginx on an Ubuntu/Debian host using an ad-hoc command in Ansible:

```
ansible webservers -i ./hosts -b --become-user=root -m shell -a 'apt -y install nginx' -u USER
```

```
node1 | SUCCESS | rc=0 >>
Reading package lists...
Building dependency tree...
Reading state information...
...
...
```

The following flags were used with the above command:

- `-b`: Instruct ansible to become another user to run the command
- `--become-user=root`: Run the command as a root user
- `-m`: Declares which module is used in the command
- `-a`: Declares which arguments are passed to the module

The alternate and preferred way of installing software using an ad-hoc command is to use `apt` module. If your remote managed host is running RHEL/CentOS, then change the module name from `apt` to `yum`.

```
ansible webservers -i hosts -b --become-user=root -m apt -a 'name=nginx state=present update_cache=true'
```

Sample output:
```
node1 | SUCCESS => {
    "cache_update_time": 1657172323, 
    "cache_updated": true, 
    "changed": true, 
    "stderr": "", 
    "stderr_lines": [], 
...
...
```

In the above Ansible command, the `-a` switch passes the arguments to the `apt` module by specifying the name of the package to be installed, the desired state, and whether to update the package repository cache or not.

The line `change: true` in the result section of the above ad-hoc command signifies that the state of the system has been changed. If you run the above ad-hoc command again, the value of changed field will be false, which means the state of the system remains unchanged, because Ansible is aware that Nginx is already present in the system and will not try to alter the state again.

That’s what we call Ansible idempotent. You can run the same ad-hoc command as many times as you’d like and it won’t change anything unless it needs to.

```
node1 | SUCCESS => {
    "cache_update_time": 1657172323, 
    "cache_updated": true, 
    "changed": false
}
```

So far, we have understood the ansible modules and its usages through ad-hoc way, but this is not so useful until we use the modules in ansible playbooks to run multiple tasks in the remote managed host.

### Running Playbooks

An Ansible play is a set of tasks that are run on one or more managed hosts. A play may include one or many different tasks, and the most common way to execute a play is to use a playbook.

Ansible Playbooks are composed of one or more plays and offer more advanced functionality for sending tasks to managed host compared to running many ad-hoc commands.

The tasks in Ansible playbooks are written in Yet Another Markup Language (YAML), which is easier to understand than a JSON or XML file. Each task in the playbook is executed sequentially for each host in the inventory file before moving on to the next task.

Let’s create a simple Ansible playbook example that will install Nginx on the managed hosts that we had already defined in the inventory file.

Create file `nginx.yml`:

```
---
- hosts: webservers
  gather_facts: yes
  become: yes
  become_user: root
  tasks:
   - name: Install Nginx
     apt:
       name: nginx
       state: present
       update_cache: true
```

This playbook does exactly the same as our ad-hoc command above.

The `hosts` tells Ansible on which hosts to run the tasks. The `webservers` group includes a single task to install Nginx.

The `become_user` in both the host section tells ansible to use sudo to run the tasks.

The `gather_facts` option gathers information about managed hosts such as distribution, OS family, and more. In ansible terminology, this information is known as FACTS.

Tasks have a name which helps us track playbook progress and an action. In the action I specify which module I want Ansible to invoke.

Now run the above playbook example using `ansible-playbook`. 

```
ansible-playbook -i ./hosts nginx.yml -u USER
```

```
PLAY [webservers] ****************************************************************

TASK [Gathering Facts] ***********************************************************
ok: [node1]

TASK [Install Nginx] *************************************************************
ok: [node1]

PLAY RECAP ***********************************************************************
node1                      : ok=2    changed=0    unreachable=0    failed=0   
```

We get some useful feedback while this runs, including the Tasks Ansible runs and their result. 

The last line contains information about the current run of the above playbook. The four points of data are:

- ok: The number of tasks that were either executed correctly or didn’t result in a change.
- changed: The number of things that were modified by Ansible.
- unreachable: The number of hosts that were unreachable for some reason.
- failed: The number of tasks failed to execute correctly.

Here we see all ran OK, but nothing was changed. We have Nginx installed already.

Let’s extend our simple Ansible playbook example with a handler to start Nginx as a service after installation.

A handler is the same as a task, but it will be executed when called by another task. It is like an event-driven system. A handler will run a task only when it is called by an event it listens for.

Create file `nginx-service.yml`:

```
---
- hosts: webservers
  gather_facts: yes
  become: yes
  become_user: root
  tasks:
  - name: Install Nginx
    apt:
      name: nginx
      state: present
      update_cache: true
    notify:
    - Start Nginx
  handlers:
  - name: Start Nginx
    service: 
      name: nginx
      state: started
```

Here we add a notify directive to the installation task. This notifies any handler named "Start Nginx" after the task is run. 

Then we create the handler called "Start Nginx". This handler is the task called when "Start Nginx" is notified.

This particular handler uses the `service` module, which can start, stop, restart, reload (and so on) system services. In this case, we tell Ansible that we want Nginx to be started.

Note that Ansible has us define the ***state*** you wish the service to be in, rather than defining the ***change*** you want. Ansible will decide if a change is needed, we just tell it the desired result.

Let's run this playbook again:

```
PLAY [webservers] ****************************************************************

TASK [Gathering Facts] ***********************************************************
ok: [node1]

TASK [Install Nginx] *************************************************************
ok: [node1]

RUNNING HANDLER [Start Nginx] ****************************************************
ok: [node1]

PLAY RECAP ***********************************************************************
node1                      : ok=2    changed=0    unreachable=0    failed=0   
```

We get the similar output, but this time the Handler was run.

Notifiers are only run if the Task is run. If I already had Nginx installed, the Install Nginx Task would not be run and the notifier would not be called.

We can use playbooks to run multiple tasks, add in variables, define other settings and even include other playbooks.

Let’s extend our simple Ansible playbook example to install Nginx and a MySQL server on the managed hosts that we had already defined in the inventory file.

To be more precise, we want Nginx installed on hosts in the `webservers` group and a MySQL server installed on hosts in the `dbservers` group.

Create file `nginx-mysql.yml`:

```
---
- hosts: webservers
  gather_facts: yes
  become: yes
  become_user: root
  tasks:
  - name: Install Nginx
    apt:
      name: nginx
      state: present
      update_cache: true
    notify:
    - Start Nginx
  handlers:
  - name: Start Nginx
    service: 
      name: nginx
      state: started

- hosts: dbservers
  become: yes
  become_user: root
  tasks:
  - name: Install mysql
    apt: 
      pkg: mysql-server
      state: present       
```

The `hosts` tells Ansible on which hosts to run the tasks. The above Ansible playbook includes two host groups from the inventory file. The tasks for `webservers` group are to install Nginx and enable Nginx during boot, and the `dbservers` group includes a single task to install MySQL.

### Advanced features

#### Variables

In Ansible, variables are similar to variables in any programming language—they let you input values and numbers dynamically into your playbook. Variables simplify operations by allowing you define and declare them throughout all the various roles and tasks you want to perform.

There are few places where you can define variables in an Ansible playbook.

- In the playbook
- In the inventory file
- In a separate variable file
- Using `group_vars`

To define variables in a playbook, use `vars` key just above the task where you want to use the variable. Once declared, you can use it inside the `{{ }}` tag. Let’s declare a variable by the name `pkgname` and assign it the value of the package name that we want to install, which is `nginx`. Once done, we can use the variable in a task.

```
---
- hosts: webservers
  gather_facts: yes
  become: yes
  become_user: root
  
  vars:
    pkgname: nginx
  
  tasks:
  - name: Install "{{ pkgname }}"
    apt:
      pkg: "{{ pkgname }}"
      state: present
      update_cache: true
```

It is also possible to declare a variable in the inventory file using the syntax [host_group_name:vars]. Let’s define the variable pkgname in the inventory file.

```
[webservers:vars]
pkgname=nginx
```

Now the variable pkgname can be used anywhere in the webservers hosts section in the playbook.

You can also define variables in a separate variable file and import it into the playbook. Create a variable file using vi another text editor and define the variable pkgname here.

Create file `ansible_vars.yml`

```
---
pkgname: nginx
```

To use the variable `pkgname`, import the above file using the `vars_files` keyword in the playbook.

```
---
- hosts: webservers
  gather_facts: yes
  become: yes
  become_user: root
  
  vars_files:
  - ./ansible_vars.yml 
...
...
```

Another preferred way of managing variables is to create a group_vars directory inside your Ansible working directory. Ansible will load any YAML files in this directory with the name of any Ansible group.

Create the directory `group_vars` in your Ansible working directory, and then create the variable files matching with the group name from the inventory file. In our example, this would be `webservers` and `dbservers`. This allows you to separate variables according to host groups, which can make everything easier to manage.

Create file `group_vars/webservers`:

```
---
pkgname: nginx
```

Create file `group_vars/dbservers`:

```
pkgname: mysql-server
```

You don’t need to declare the variable in your playbook, as Ansible will automatically pull the variables from each `group_vars` files and will substitute them during runtime.

Now suppose you want to have variables that will apply to all the host groups mentioned in the inventory file. To accomplish it, name a file by the name `all` inside `group_vars` directory. The `group_vars/all` files are used to set variables for every host that Ansible connects to.

#### Conditionals

In Ansible, conditionals are analogous to an `if` statement in any programming language. You use a conditional when you want to execute a task based on certain conditions.

In our last playbook example, we installed Nginx, so let’s extend that by creating a task that installs Nginx when Apache is not present on the host. We can add another task to the playbook we’ve already built.

```
...
...
  tasks:
  - name: Check if Apache is already installed
    shell: dpkg -s apache2 | grep Status
    register: apache2_is_installed  
    failed_when: no
  - name: Install "{{ pkgname }}"
    apt: pkg="{{ pkgname }}" state=present
    when: apache2_is_installed.rc == 1
    notify:
    - restart nginx
...
...
```

The first task in the above playbook checks if Apache is installed using `dpkg -s` command and stores the output of the task to `apache2_is_installed` variable. The return value of the task will be a non-zero value if Apache is not installed on the host.

Usually, Ansible would stop executing other tasks because of this non-zero value, but the `failed_when: no` gives Ansible permission to continue with the next set of tasks when it encounters a non-zero value.

The second task will install Nginx only when the return value of `rc` is equal to one, which is declared via when: `apache2_is_installed.rc == 1`.

#### Loops

All programming languages provide a way to iterate over data to perform some repetitive task. Ansible also provides a way to do the same using a concept called looping, which is supplied by Ansible lookup plugins. With loops, a single task in one playbook can be used to create multiple users, install many packages, and more.

While there are many ways to use loops in Ansible, we’ll cover just one of them to get you started. The easiest way to use loops in ansible is to use `with_items` keyword, which is used to iterate over an item list to perform some repetitive tasks. The following playbook includes a task which installs packages in a loop using the keyword `with_items`.

```
---
- hosts: webservers
  gather_facts: yes
  become_user: root
    
  tasks:
 
  - name: Installing packages using loops
    apt: pkg={{ item }} state=present update_cache=yes
    with_items:
      - sysstat
      - htop
      - git    
```

Run the above playbook from your command line, and you’ll see that you’ve installed all three packages on the remote host with a single task!

#### Tags
Tags allow you to run only specific tasks from your playbook via the command line. Just add the `tags` keyword for each task and run only the task(s) that you want by using `--tags` switch at the end of the ansible command. In the following playbook, we have added tags at the end of each task, thereby allowing us to run tasks separately from a single playbook.

```
---
- hosts: webservers
  gather_facts: yes
  become_user: root
 
  tasks:
  - name: Check if Apache is already installed
    shell: dpkg -s apache2 | grep Status
    register: apache2_is_installed  
    failed_when: no
  - name: Install "{{ pkgname }}"
    apt: pkg="{{ pkgname }}" state=present
    when: apache2_is_installed.rc == 1
    notify:
    - restart nginx
  - name: ensure nginx is running and enable it at boot
    service: name=nginx state=started enabled=yes
    tags:
    - mytag1

  handlers:
    - name: restart nginx
      service: name=nginx state=restarted
    
- hosts: dbservers
  become_user: root
  tasks:
  - name: Install mysql
    apt: pkg="{{ pkgname }}" state=present
    tags:
    - mytag2
```

Now run any of the tasks by specifying tag name at the end of ansible command.

```
ansible-playbook -i ./hosts playbook.yml -u USER --tags 'mytag2'
```

#### Templates

Typically, after installing a web server like Nginx, you need to configure a virtual hosts file to properly serve a given website on your VPS. Instead of using SSH to log into your VPS to configure it after running Ansible, or using Ansible’s `copy` module to copy many unique configuration files individually, you can take advantage of Ansible’s templates features.

A template file contains all of the configuration parameters you need, such as the Nginx virtual host settings, and uses variables, which are replaced by the appropriate values when the playbook is executed. Template files usually end with the .j2 extension that denotes the Jinja2 templating engine.

To begin working with templates, create a directory for template files in your Ansible working directory.

```
mkdir templates
```

Create two template files. The first template file will be the default index.html file for each site, and the second template file will contain configuration settings for the Nginx virtual host.

Create file `templates/index.html.j2`:

```
<html>
You are visiting {{ domain_name }} !
</html>
```

Similarly, create a template file for the Nginx virtual host. 

Create file `templates/nginx-vh.j2`:

```
server {
        listen       80;
        server_name  {{ domain_name }};
        client_max_body_size 20m;
        index index.php index.html index.htm;
        root   /var/www/html/{{ domain_name }};

        location / {
                    try_files $uri $uri/ /index.html?q=$uri&$args;
        }
        location ~* \.(js|css|png|jpg|jpeg|gif|ico|woff|ttf|svg|otf)$ {
               expires 30d;
               add_header Pragma public;
               add_header Cache-Control "public";
               access_log off;
    }
}
```

Notice that the variables `domain_name` in the above two template files are enclosed within `{{ }}`, which means they will be substituted during runtime by the value of this variable. To define the variable `domain_name`, navigate to the `group_vars` directory and edit the file `webservers` and add the following lines in it.

```
domain_name: SUBDOMAIN.DOMAIN.TLD
```

Finally, edit the ansible playbook to create a root folder for sites, copy the index.html file to the site’s root folder, and copy the virtual host file to the Nginx virtual host directory /etc/nginx/sites-enabled one by one.

```
---
- hosts: webservers
  gather_facts: yes
  become_user: root
 
  tasks:
  - name: Check if Apache is already installed
    shell: dpkg -s apache2 | grep Status
    register: apache2_is_installed  
    failed_when: no

  - name: Install "{{ pkgname }}"
    apt: pkg="{{ pkgname }}" state=present
    when: apache2_is_installed.rc == 1
    notify:
    - restart nginx

  - name: ensure nginx is running and enable it at boot
    service: name=nginx state=started enabled=yes
 
  - name: create virtual host root directory
    file: name=/var/www/html/{{ domain_name }} state=directory

  - name: Copying index file to webroot
    template:
      src: templates/index.html.j2
      dest: /var/www/html/{{ domain_name }}/index.html
 
  - name: Enables nginx virtual host
    template:
      src: templates/nginx-vh.j2
      dest: /etc/nginx/sites-enabled/{{ domain_name }}

  - name: restart nginx
    service: name=nginx state=restarted

    tags:
    - mytag1

  handlers:
    - name: restart nginx
      service: name=nginx state=restarted
    
- hosts: dbservers
  become_user: root
  tasks:
  - name: Install mysql
    apt: pkg="{{ pkgname }}" state=present
    tags:
    - mytag2
```

The template task in the above Ansible playbook takes two mandatory parameters `src` and `dest`. There are also a few optional parameters that can be specified in a template task but is not required at this stage.

- The `src` parameter specifies the name of the template file from templates directory that Ansible will copy to the remote server. In our case, the two templates files that we have created are `index.html.j2` and `nginx-vh.j2`
- The `dest` parameter is the path in the remote server where the file should be placed.

Finally, run the playbook from your ansible working directory:

```
ansible-playbook -i ./hosts playbook.yml -u USER
```

#### Blocks

Blocks, which were introduced in version 2.0, allow you to logically group tasks and better handle errors, which is useful when you want to execute multiple tasks under a single condition.

To end the block, use the `when` keyword once you’re done defining all the tasks you want to be executed. If the evaluation of the `when` condition returns `true`, then all the tasks within the blocks will be executed one by one. All tasks within the blocks will inherit the common data or directives that you set just after the ‘when’ keyword.

```
---
- hosts: webservers

  tasks:
  - name: Install Nginx
    
    block:
    - apt: pkg=nginx state=present
    - service: name=nginx state=started enabled=yes

    when: ansible_distribution == 'Ubuntu'
    become: true
    become_user: root
```

The `block` section in the above playbook includes two related tasks to install `nginx` and start/enable it. The `when` evaluation specifies that these tasks should only be run when the remote managed host is using Ubuntu as its operating system. Both the tasks will inherit the privilege escalation directives after the ‘when’ keyword.

You can also use blocks to handle failures, similar to exceptions in most programming languages. The aim is to gracefully handle failures within the `block` rather than withdrawing the entire deployment.

Here is an example of how to use blocks to handle failures:

```
tasks:  
  - block:  

  - name: Enable Nginx during boot
    service: name=nginx state=started enabled=yes
 
    rescue:  
      - name: This section runs only when there is an error in the block.  
        debug: msg="There was an error in starting/enabling nginx."  
    always:  
      - name: This section will run always.  
        debug: msg="This always executes."`
```

## References 

1. [An Introduction to Configuration Management with Ansible
](https://www.digitalocean.com/community/conceptual_articles/an-introduction-to-configuration-management-with-ansible)

1. [Configuration management with Ansible]
(https://jpmens.net/2012/06/06/configuration-management-with-ansible/)

1. [An Ansible2 Tutorial](https://serversforhackers.com/c/an-ansible2-tutorial)
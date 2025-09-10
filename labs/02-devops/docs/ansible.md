# Automating with Ansible

While `parallel-ssh` makes it easy to run commands across multiple hosts, it doesn't address deeper challenges like tracking system state, managing dependencies, or ensuring repeatable outcomes. To truly address these challenges, DevOps engineers rely on configuration management tools like Ansible, Terraform, or Chef. These tools use SSH under the hood, but provide powerful features for automation, repeatability, scalability, and version-controlled infrastructure.

[Ansible](https://www.ansible.com/) is a modern configuration management tool that facilitates the task of setting up and maintaining remote servers. [Configuration management](https://en.wikipedia.org/wiki/Configuration_management) is an automated method for maintaining computer systems and software in a known, consistent state. 

<figure>
  <p align="center"><img src="../assets/images/ansible-architecture.png" width="60%"></p>
  <figcaption><p align="center">Figure. Ansible Architecture</p></figcaption>
</figure>

Getting started with this tool is simple. 
1. Choose a machine as your management system and install the tool. Your management node can also be a managed node.
2. Ensure you have passwordless ssh access from your management system to each managed node.
3. Create a hosts file containing an inventory of your nodes.
4. Start using the tool 

We suggest the following:
- For step 1, you use `node0` as the management system.
- For step 2, you use a CloudLab profile that automatically sets up ssh keys for passwordless access. All the CloudLab profiles under `UCY-COAST-TEACH` meet this requirement.

On Debian based distributions, you can install `ansible` using `apt`:

```
sudo apt -y install ansible
```

Ansible uses an inventory file, called hosts to determine which nodes you want to manage. This is a plain-text file which lists individual nodes or groups of nodes (e.g. Web servers, Database servers, etc.). 
The default location for the inventory file is /etc/ansible/hosts, but it’s possible to create inventory files in any location that better suits your needs. In this case, you’ll need to provide the path to your custom inventory file with the `-i` parameter when running Ansible commands. 

Using per-project inventory files is a good practice to minimize the risk of running a command on the wrong group of servers. 

To get started, create a file named `hosts` in your current working directory. This file defines the group of target machines for your playbook. For example:

```
[webservers]
node2
node3
```

This defines a group of nodes we call `webservers` with two specified hosts in it.

We can now immediately launch Ansible to see if our setup works:

```bash
ansible webservers -i ./hosts -m ping
```

If everything is set up correctly, you should see output like this:

```json
node2 | SUCCESS => {
    "changed": false, 
    "ping": "pong"
}
node3 | SUCCESS => {
    "changed": false, 
    "ping": "pong"
}
```

This confirms that Ansible was able to connect to each individual node in the `webservers` group using SSH, transmit the required module (here: `ping`), launch the module, and return the module’s output. 

Instead of addressing a group of nodes, we can also specify an individual node or a set of nodes with wildcards:

```bash
ansible node2 -i ./hosts -m ping
ansible "node*" -i ./hosts -m ping
```

Modules use the available context to determine what actions if any needed to bring the managed host to the desired state and are [idempotent](https://en.wikipedia.org/wiki/Idempotence#Computer_science_meaning), that means if you run the same task again and again, the state of the machine will not change. To find the list of available modules, use `ansible-doc -l`

Instead of shooting off individual ansible commands, we can group these together into so-called playbooks which declare a specific configuration we want to apply to a node. Tasks specified are processed in the order we specify. A playbook is expressed in YAML.

Create the following playbook in your current working directory, for example by saving it as `nginx.yml`:

```
---
- name: Setup web server
  hosts: webservers
  become: true
  vars:
    web_root: "/var/www/html"
  tasks:
    - name: Install nginx
      apt:
        name: nginx
        state: present

    - name: Ensure nginx is running
      service:
        name: nginx
        state: started
        enabled: true

    - name: Copy index.html
      copy:
        src: index-ansible.html
        dest: "{{ web_root }}/index.html"
```

The format is very readable. 
With `hosts`, we choose specific hosts and/or groups in our inventory to execute against. 
With `vars`, we define variables we want to pass to tasks. These variables can be used within the playbook and within templates as `{{ var }}`. 
With `tasks`, we define the list of tasks we want to execute. Each task hask a name which helps us track playbook progress and a module we want Ansible to invoke.
Finally, for tasks that require root privileges such as installing packages, we use `become` to ask Ansible to activate privilege escalation and run corresponding tasks as `root` user. 

The playbook uses several types of modules, many of which are self explanatory. 

The `ansible-playbook` utility processes the playbook and instructs the nodes to perform the tasks, starting with an implicit invocation of the setup module, which collects system information for Ansible. Tasks are performed top-down and an error causes Ansible to stop processing tasks for that particular node.

Before you run the playbook, create a file named `index-ansible.html`:

```html
<h1>Hello from Ansible!</h1>
```

You can try to execute the playbook using the command below, assuming the playbook is in file `nginx.yml`:

```
ansible-playbook -i ./hosts nginx.yml
```

Once the playbook is finished, if you go to your browser and access `node2`'s or `node3`'s public hostname or IP address you should see the following page:

```
Hello from Ansible!
```

Alternatively, you can use a `curl` command to GET a remote resource and have it displayed in the terminal:

```
curl http://node2/index.html
```

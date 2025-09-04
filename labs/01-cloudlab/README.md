# Lab: Getting Started with CloudLab 

You will do your labs and assignments for CS452 using CloudLab. CloudLab is a research facility which provides bare-metal access and control over a substantial set of computing, storage, and networking resources. If you haven’t worked in CloudLab before, you need to register a CloudLab account.

This lab walks you through the CloudLab registration process and shows you how to start an experiment in CloudLab.

Most importantly, it introduces our policies on using CloudLab that will be enforced throughout the semester.

## Introduction

CloudLab is a "meta-cloud"—that is, it is not a cloud itself; rather, it is a facility for building clouds. It provides bare-metal access and control over a substantial set of computing, storage, and networking resources; on top of this platform, users can install standard cloud software stacks, modify them, or create entirely new ones.

The current CloudLab deployment consists of more than 25,000 cores distributed across three sites at the University of Wisconsin, Clemson University, and the University of Utah. CloudLab interoperates with existing testbeds including [GENI](http://www.geni.net/) and [Emulab](http://www.emulab.net/), to take advantage of hardware at dozens of sites around the world.

The control software for CloudLab is open source, and is built on the foundation established for [Emulab](http://www.emulab.net/), [GENI](http://www.geni.net/), and [Apt](http://www.aptlab.net/). Pointers to the details of this control system can be found on CloudLab’s [technology page](https://www.cloudlab.us//technology.php).

## Register a CloudLab account

You’ll need an account to use CloudLab. 

If you have an account at one of its federated facilities, like [GENI](http://www.geni.net/) and [Emulab](http://www.emulab.net/), then you already have an account at CloudLab, and you can simply request to join the `UCY-COAST-TEACH` project. 

If not, you can register an account by visiting http://cloudlab.us and creating an account using your University of Cyprus email address as login. Note that an SSH public key is required to access the nodes CloudLab assigns to you; if you are unfamiliar with creating and using ssh keypairs, we recommend taking a look at the [guide to generating SSH keys](../../docs/connect-ssh/generate.md).

To register an account:

1. Start by visiting https://www.cloudlab.us/ in your browser and clicking the **Request an Account** button.

<figure>
  <p align="center"><img src="figures/cloudlab-portal.png" width="60%"></p>
  <figcaption><p align="center">Figure. CloudLab Portal</p></figcaption>
</figure>

2. In the sign-up form, select **Join Existing Project**, and enter `UCY-COAST-TEACH` as the project name. 

<figure>
  <p align="center"><img src="figures/cloudlab-register-account.png" width="60%"></p>
  <figcaption><p align="center">Figure. Request to Join a Project</p></figcaption>
</figure>

3. Fill in your personal information, including username, full name, email, and password.

4. Click on **Choose file** under **SSH Public Key file**. Locate the SSH public key file `.pub` and click **Open**.

5. Finally, click on **Submit Request**. You will get an email notification when the project leader approves your request. 

Once your registration request gets approved, you can start using CloudLab to create experiments.

## Start an Experiment

CloudLab resources are assigned to experiments, which is just a grouping of resources allocated to you. 

To start a new experiment:

1. **Log in**

   Start by pointing your browser at https://www.cloudlab.us/, clicking the **Log In** button, and entering your username and password.

<figure>
  <p align="center"><img src="figures/cloudlab-portal.png" width="60%"></p>
  <figcaption><p align="center">Figure. CloudLab Portal</p></figcaption>
</figure>

2. **Start Experiment**

   Click on the **Experiments** tab in the upper left corner, then select **Start Experiment**. 

<figure>
  <p align="center"><img src="figures/cloudlab-experiments-menu.png" width="60%"></p>
  <figcaption><p align="center">Figure. Experiments Tab</p></figcaption>
</figure>

3. **Experiment Wizard**

   Experiments must be configured before they can be instantiated. A short wizard guides you through the process. The first step is to pick a profile for your experiment. A profile describes [a set of resources](https://docs.cloudlab.us/advanced-topics.html#%28part._rspecs%29) (both hardware and software) that will be used to start your experiment. On the hardware side, the profile will control whether you get [virtual machines](https://docs.cloudlab.us/basic-concepts.html#%28part._virtual-machines%29) or [physical ones](https://docs.cloudlab.us/getting-started.html#:~:text=virtual%20machines%20or-,physical%20ones,-%2C%20how%20many%20there), how many there are, and what the network between them looks like. On the software side, the profile specifies the [operating system and installed software](https://docs.cloudlab.us/advanced-topics.html#%28part._disk-images%29).

   Profiles come from two sources. Some of them are provided by CloudLab itself, and provide standard installation of popular operating systems, software stacks, etc. Others are [created by other researchers](https://docs.cloudlab.us/creating-profiles.html) and may contain research software, artifacts and data used to gather published results, etc. Profiles represent a powerful way to enable [repeatable research](https://docs.cloudlab.us/repeatable-research.html).

   The default profile is `small-lan`, and usually you do not want to use this profile. You should instead click the **Change Profile** button. 
  
   Clicking the **Change Profile** button will let you select the profile that your experiment will be built from.   

<figure>
  <p align="center"><img src="figures/cloudlab-small-lan-select-profile.png" width="60%"></p>
  <figcaption><p align="center">Figure. Experiment Wizard</p></figcaption>
</figure>

4. **Select a Profile**
  
   On the left side is the profile selector which lists the profiles you can choose. The list contains both globally accessible profiles and profiles accessible to the projects you are part of.

   The large display in this dialog box shows the network topology of the profile, and a short description sits below the topology view.

   We provide a few profiles in the `UCY-COAST-TEACH` project, the `multi-node-cluster` profile will give you a small cluster installation with a variable number of bare metal nodes.

   Once you have chosen a profile hit **Select Profile** and then click **Next**.

<figure>
  <p align="center"><img src="figures/cloudlab-multi-node-cluster-change-profile.png" width="60%"></p>
  <figcaption><p align="center">Figure. Select a Profile</p></figcaption>
</figure>

<figure>
  <p align="center"><img src="figures/cloudlab-multi-node-cluster-select-profile.png" width="60%"></p>
  <figcaption><p align="center">Figure. Selected Profile</p></figcaption>
</figure>

5. **Choose Parameters**

   Some profiles are simple and provide the same topology every time they are instantiated. But others, like the `multi-node-cluster` profile, are parameterized and allow users to make choices about how they are instantiated. The `multi-node-cluster` profile allows you to pick the number of compute nodes, the hardware to use, and many more options. The creator of the profile chooses which options to allow and provides information on what those options mean. Just mouse over a blue ’?’ to see a description of an option. 
  
   For this lab, we will create a single-node experiment. Leave **Number of Nodes** to 1 and click **Next** to continue.

<figure>
  <p align="center"><img src="figures/cloudlab-multi-node-cluster-parameterize.png" width="60%"></p>
  <figcaption><p align="center">Figure. Parameterize</p></figcaption>
</figure>

6. **Pick a Cluster**

   CloudLab can instantiate profiles on several different backend clusters. The cluster selector is located right above the **Create” button; the the cluster most suited to the profile you’ve chosen will be selected by default.

   Here you should name your experiment with `CloudLabLogin-ExperimentName`. The purpose of doing this is to prevent everyone from picking random names and ending up confusing each other since everyone in the `UCY-COAST-TEACH` project can see a full list of experiments created. 
   
   You also need to specify from which cluster you want to start your experiment. CloudLab can instantiate profiles on several different backend clusters. Each cluster has different hardwares. For more information on the hardware CloudLab provides, please refer to [this](https://docs.cloudlab.us/hardware.html). 
   
   Once you select the cluster, click **Next** to continue.  

<figure>
  <p align="center"><img src="figures/cloudlab-multi-node-cluster-finalize-singlenode.png" width="60%"></p>
  <figcaption><p align="center">Figure. Finalize</p></figcaption>
</figure>

7. **Click Finish!**

   You are now ready to run your experiment or schedule the deployment for later. 
  
   To start your experiment immediately, leave **Start on day/time** empty.
   
   To schedule the deployment for later, set a date and time in **Start on day/time**. A typical use for this option, is to schedule your experiment to start shortly after a resource reservation starts. See the manual for more info on [reservations](http://docs.cloudlab.us/reservations.html).
  
   For this lab, leave **Start on day/time** empty and click **Finish** to confirm your deployment. The reservation and deployment processes will run automatically.
   
<figure>
  <p align="center"><img src="figures/cloudlab-multi-node-cluster-schedule.png" width="60%"></p>
  <figcaption><p align="center">Figure. Schedule</p></figcaption>
</figure>

8. **CloudLab instantiates your profile**


   When you click the **Finish** button, CloudLab will start preparing your experiment by selecting nodes, installing software, etc. as described in the profile. What’s going on behind the scenes is that on one (or more) of the machines in one of the CloudLab clusters, a disk is being imaged, VMs and/or physical machines booted, accounts created for you, etc. This process usually takes a couple of minutes.

<figure>
  <p align="center"><img src="figures/cloudlab-multi-node-cluster-provisioning-singlenode.png" width="60%"></p>
  <figcaption><p align="center">Figure. Deployment</p></figcaption>
</figure>
   
9. **Your experiment is ready!**

   When your experiment is ready to use, the progress bar will be complete, and you’ll be given a lot of new options at the bottom of the screen.
   
   The **Topology View** shows the network topology of your experiment (which may be as simple as a single node). Clicking on a node in this view brings up a terminal in your browser that gives you a shell on the node. The **List View** lists all nodes in the topology, and in addition to the in-browser shell, gives you the command to ssh login to the node (if you provided a public key). The **Manifest** tab shows you the technical details of the resources allocated to your experiment. Any open terminals you have to the nodes show up as tabs on this page.

   Clicking on the **Profile Instructions** link (if present) will show instructions provided by the profile’s creator regarding its use.

   Your experiment is yours alone, and you have full “root” access (via the sudo command). No one else has access to the nodes in your experiment, and you may do anything at all inside of it, up to and including making radical changes to the operating system itself. CloudLab will clean it all up when you’re done!

   Your experiment will **terminate automatically after a few hours**. When the experiment terminates, you will **lose anything on disk** on the nodes, so be sure to copy off anything important early and often. You can use the **Extend** button to submit a request to hold it longer, or the **Terminate** button to end it early.

<figure>
  <p align="center"><img src="figures/cloudlab-multi-node-cluster-listview-singlenode.png" width="60%"></p>
  <figcaption><p align="center">Figure. List View</p></figcaption>
</figure>

## Use your experiment

Secure Shell (ssh) is the standard way to connect to a remote machine’s shell. Since you gave CloudLab an ssh public key as part of account creation, you can log in using the ssh client on your laptop or desktop. `node0` is a good place to start. Go to the **List View** on the experiment page to get a full command line for the ssh command and use it to ssh into `node0`. If you are unfamiliar with using an ssh client, we recommend taking a look at this [guide](../../docs/connect-ssh/README.md). 

```
$ ssh -p 22 alice@amd198.utah.cloudlab.us
# Attempts to ssh to CloudLab
```

You may see a warning like this:

```
The authenticity of host 'amd198.utah.cloudlab.us (128.110.219.109)' can't be established.
RSA key fingerprint is SHA256:bB4TBZMIEhrtJN2mb/Mzn8fvSRUBUC+UamxqkhCU7HA.
Are you sure you want to continue connecting (yes/no)? 
```

Verify that the fingerprint in the message you see matches CloudLab's public key fingerprint. If it does, then type yes.

The interactive session is established, and you may now conduct your administrative tasks.

For example, you can run the `lscpu` command to list information about CPUs that are present in the system, including the number of CPUs.

```
$ lscpu
```

### Run a command over SSH on a single machine

What if the only thing you need to do over the SSH connection is execute a single quick command? You might not want to take the separate actions of connecting and authenticating, running the command, and then disconnecting.

SSH allows users to append the desired command directly to the connection attempt. The command executes, and the connection is closed.

The basic syntax is `ssh user@host "command"`.

For example, you could check the installation status of a package:

```
$ ssh alice@amd198.utah.cloudlab.us "dpkg -l | grep nano"
```

You can also run commands from a local file. The `-s` option of `bash` helps to read the executable command from the standard input:

```
ssh alice@amd198.utah.cloudlab.us 'bash -s' < src/get_host_info.sh
```

If you need to elevate your privileges on the far side of the SSH connection with `sudo`, then force the use of a pseudo-terminal with `-t`. Use this if `sudo` will challenge you for a password. The command looks like this:

```
$ ssh -t alice@amd198.utah.cloudlab.us "sudo apt -y install nano"
```

## Terminate your experiment 

Terminate the experiment once you are done. CloudLab is a shared resource, and you should be careful to not waste resources. If we find cases where experiments are not terminated once you are done with them, we will impose a grade penalty.

## Policies on Using CloudLab Resources

**Do not leave your CloudLab experiment instantiated unless you are using it! It is important to be a good citizen of CloudLab.**

The nodes you receive from CloudLab are real hardware machines sitting in different clusters. Therefore, we ask you not to hold the nodes for too long. CloudLab gives users 16 hours to start with, and users can extend it for a longer time. Manage your time efficiently and only hold onto those nodes when you are working on the assignment. Write scripts and use tools to ensure you can easily setup and restart experiments. Remember to back up any data to your local machine so you do not lose data. You should use a private git repository to manage your code, and you must terminate the nodes when you are not using them. If you do have a need to extend the nodes, do not extend them by more than 1 day. We will terminate any cluster running for more than 48 hours.

As a member of the `UCY-COAST-TEACH` project, you have permissions to access another member’s private user space. Stick to your own space and do not access others’ to peek at/copy/use their code, or intentionally/unintentionally overwriting files in others’ workspaces.

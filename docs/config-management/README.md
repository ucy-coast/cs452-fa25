# Configuration Management

This section provides an overview of cluster configuration and management, and tools available for configuring and managing a cluster. It introduces three methods of managing a cluster: Secure Shell (ssh), parallel-ssh, and Ansible. 

### Pre-requisite

In the remaining of this section, we assume that you are already connected to a cluster management node through SSH. Using SSH to connect to the management node gives you the ability to monitor and interact with the cluster. For CloudLab, we assume you are connected to `node0` and you use this node as the cluster managemement node. 

We also assume that you have already setup SSH key-based authentication between cluster nodes and can connect from one cluster node to another without entering a password. For CloudLab, we assume you use a profile that already configures SSH key-based authentication between nodes.

- [Secure Shell (ssh)](ssh.md)
- [parallel-ssh](parallel-ssh.md)
- [Ansible](ansible.md)

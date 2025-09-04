# Testing your SSH connection

After you've set up your SSH key and added it to your account on CloudLab, you can test your connection.

Before adding a new SSH key to your account on CloudLab, you should have:

- [Checked for existing SSH keys](existing-windows-openssh.md)
- [Generated a new SSH key](generate-windows-openssh.md)
- [Added a new SSH key to your CloudLab account](add-windows-openssh.md)

When you test your connection, you'll need to authenticate this action using your password, which is the SSH key passphrase you created earlier. For more information on working with SSH key passphrases, see [Working with SSH key passphrases](passphrases-windows-openssh.md).

1. Create a new experiment on CloudLab.

2. When your experiment is ready to use, the progress bar will be complete, and you’ll be given a lot of new options at the bottom of the screen. The `List View` lists all nodes in the topology, and gives you the command to ssh login to the node (if you provided a public key). 

<figure>
  <p align="center"><img src="assets/images/cloudlab-listview.png"></p>
  <figcaption><p align="center">Figure. CloudLab experiment dashboard.</p></figcaption>
</figure>

3. Click on the SSH command to automatically open a new SSH session, or 

4. Open Windows PowerShell 

5. Enter the SSH command to manually open a new SSH session: 

   ```
   $ ssh -p 22 alice@c220g2-011014.wisc.cloudlab.us
   # Attempts to ssh to CloudLab
   ```

6. You may see a warning like this:

   ```
   The authenticity of host 'c220g2-011014.wisc.cloudlab.us (128.105.145.88)' can't be established.
   RSA key fingerprint is SHA256:/ooH3X5gHoyIFYFol8R2u0XpxMGgmlvRznPYsuBdGuU.
   Are you sure you want to continue connecting (yes/no)? 
   ```

   Verify that the fingerprint in the message you see matches CloudLab's public key fingerprint. If it does, then type yes.

7. Upon the terminal login, you should see a Welcome Message similar to this:

```
Welcome to Ubuntu 20.04 LTS ((GNU/Linux 5.4.0-100-generic x86 64))

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage

 * Super-optimized for small spaces - read how we shrank the memory
   footprint of MicroK8s to make it the smallest full K8s around.

   https://ubuntu.com/blog/microk8s-memory-optimisation

The programs included with the Ubuntu system are free software;
the exact distribution terms for each program are described in
the individual files in /usr/share/doc/*/copyright.

Ubuntu comes with ABSOLUTELY NO WARRANTY, to the extent permitted by
applicable law.

node0:~> 
```
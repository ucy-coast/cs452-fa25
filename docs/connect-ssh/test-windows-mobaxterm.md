# Testing your SSH connection

After you've set up your SSH key and added it to your account on CloudLab, you can test your connection.

Before adding a new SSH key to your account on CloudLab, you should have:

- [Checked for existing SSH keys](existing-windows-mobaxterm.md)
- [Generated a new SSH key](generate-windows-mobaxterm.md)
- [Added a new SSH key to your CloudLab account](add-windows-mobaxterm.md)

When you test your connection, you'll need to authenticate this action using your password, which is the SSH key passphrase you created earlier. For more information on working with SSH key passphrases, see [Working with SSH key passphrases](passphrases-windows-mobaxterm.md).

1. Create a new experiment on CloudLab.

2. When your experiment is ready to use, the progress bar will be complete, and you’ll be given a lot of new options at the bottom of the screen. The `List View` lists all nodes in the topology, and gives you the command to ssh login to the node (if you provided a public key). 

<figure>
  <p align="center"><img src="assets/images/cloudlab-listview.png"></p>
  <figcaption><p align="center">Figure. CloudLab experiment dashboard.</p></figcaption>
</figure>

3. Click on the SSH command to automatically open a new SSH session, or 

4. Open MobaXterm

5. Enter the username and hostname part of the SSH command into the `Remote host` textbox under `Basic SSH settings`.

6. Click `Open` to manually open a new SSH session.

<figure>
  <p align="center"><img src="assets/images/mobaxterm-ssh-1.png"></p>
  <figcaption><p align="center">Figure. MobaXterm SSH Session.</p></figcaption>
</figure>

7. You may see a warning like this:

   ```
   The authenticity of host 'c220g2-011014.wisc.cloudlab.us (128.105.145.88)' can't be established.
   RSA key fingerprint is SHA256:/ooH3X5gHoyIFYFol8R2u0XpxMGgmlvRznPYsuBdGuU.
   Are you sure you want to continue connecting (yes/no)? 
   ```

   Verify that the fingerprint in the message you see matches CloudLab's public key fingerprint. If it does, then click yes.

8. Upon the terminal login, you should see a Welcome Message similar to this:

<figure>
  <p align="center"><img src="assets/images/mobaxterm-ssh-2.png"></p>
  <figcaption><p align="center">Figure. Welcome Message.</p></figcaption>
</figure>

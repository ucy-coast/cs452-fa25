# Working with SSH key passphrases

You can secure your SSH keys and configure an authentication agent so that you won't have to reenter your passphrase every time you use your SSH keys.

With SSH keys, if someone gains access to your computer, they also gain access to every system that uses that key. To add an extra layer of security, you can add a passphrase to your SSH key. You can use `ssh-agent` to securely save your passphrase so you don't have to reenter it.

## Adding or changing a passphrase

You can change the passphrase for an existing private key without regenerating the keypair by typing the following command:

```
$ ssh-keygen -p -f ~/.ssh/id_ed25519
> Enter old passphrase: [Type old passphrase]
> Key has comment 'your_email@example.com'
> Enter new passphrase (empty for no passphrase): [Type new passphrase]
> Enter same passphrase again: [Repeat the new passphrase]
> Your identification has been saved with the new passphrase.
```

If your key already has a passphrase, you will be prompted to enter it before you can change to a new passphrase.

## Manual-launching ssh-agent

You can start `ssh-agent` service manually when opening powershell for the first time. 

1. Open Windows PowerShell.

2. Start SSH Agent:

   ```
   Start-Service ssh-agent
   ```

   If you get an error about being unable to start the ssh-agent service, then you might need to enable the ssh-agent service.

   ```  
   > Get-Service -Name ssh-agent | Set-Service -StartupType Manual
   ```

## Auto-launching ssh-agent

You can configure `ssh-agent` service to automatically start with Windows.

1. Open Windows PowerShell

2. Configure SSH Agent to automatically start.

   ```
   Set-Service ssh-agent -StartupType Automatic 
   ```

3. After that, you need to add your ssh key once:

   ```
   ssh-add %HOMEPATH%\.ssh\id_rsa 
   ```

4. Now everytime the ssh-agent is started, the key will be there. You can check which keys are registered with the ssh-agent:

   ```
   ssh-add -l
   ```
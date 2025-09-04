# Generating a new SSH key and adding it to the ssh-agent

After you've checked for existing SSH keys, you can generate a new SSH key to use for authentication, then add it to the ssh-agent.

## About SSH key generation

If you don't already have an SSH key, you must generate a new SSH key to use for authentication. If you're unsure whether you already have an SSH key, you can check for existing keys. For more information, see [Checking for existing SSH keys](existing-windows-openssh.md).

If you don't want to reenter your passphrase every time you use your SSH key, you can add your key to the SSH agent, which manages your SSH keys and remembers your passphrase.

## Generating a new SSH key

1. Open Windows PowerShell.

2. Paste the text below, substituting in your CloudLab email address.

   ```
   $ ssh-keygen -t ed25519 -C "your_email@example.com"
   ```

   > **Note**: If you are using a legacy system that doesn't support the Ed25519 algorithm, use:
   > 
   > ```
   > $ ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
   > ```

   This creates a new SSH key, using the provided email as a label.

   ```
   > Generating public/private algorithm key pair.
   ```

3. When you're prompted to "Enter a file in which to save the key," press Enter. This accepts the default file location.

   ```
   > Enter a file in which to save the key (/Users/you/.ssh/id_algorithm): [Press enter]
   ```

4. At the prompt, type a secure passphrase. For more information, see "Working with SSH key passphrases."

   ```
   > Enter passphrase (empty for no passphrase): [Type a passphrase]
   > Enter same passphrase again: [Type passphrase again]
   ```

## Adding your SSH key to the ssh-agent

Before adding a new SSH key to the ssh-agent to manage your keys, you should have checked for existing SSH keys and generated a new SSH key.

1. Open Windows PowerShell.

2. Ensure the ssh-agent is running.

   ```
   > Get-Service ssh-agent
   ```
  
   If the ssh-agent is not running, then you can use the "Manual-launching the ssh-agent" or "Auto-launching the ssh-agent" instructions in "[Working with SSH key passphrases](passphrases-windows-openssh.md)".
   
3. Add your SSH private key to the ssh-agent and store your passphrase in the keychain. If you created your key with a different name, or if you are adding an existing key that has a different name, replace *id_ed25519* in the command with the name of your private key file.

    ```
    $ ssh-add ~/.ssh/id_ed25519
    ```

4. Add the SSH key to your account on CloudLab. For more information, see "Adding a new SSH key to your CloudLab account."
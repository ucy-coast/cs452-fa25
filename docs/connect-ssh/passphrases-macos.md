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

## Saving your passphrase in the keychain

On Mac OS X Leopard through OS X El Capitan, these default private key files are handled automatically:

```
.ssh/id_rsa
.ssh/identity
```

The first time you use your key, you will be prompted to enter your passphrase. If you choose to save the passphrase with your keychain, you won't have to enter it again.

Otherwise, you can store your passphrase in the keychain when you add your key to the ssh-agent. For more information, see ["Adding your SSH key to the ssh-agent."](add-macos.md)
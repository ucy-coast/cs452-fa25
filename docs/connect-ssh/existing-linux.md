# Checking for existing SSH keys

Before you generate an SSH key, you can check to see if you have any existing SSH keys.

1. Open Terminal

2. Enter `ls -al ~/.ssh` to see if existing SSH keys are present.

```
$ ls -al ~/.ssh
# Lists the files in your .ssh directory, if they exist
```

3. Check the directory listing to see if you already have a public SSH key. You’re looking for a pair of files named something like id_rsa or id_ecdsa and a matching file with a .pub extension. The .pub file is your public key, and the other file is the corresponding private key.

> :bulb: **Tip**: If you receive an error that ~/.ssh doesn't exist, you do not have an existing SSH key pair in the default location. You can create a new SSH key pair in the next step.

4. Either generate a new SSH key or upload an existing key.

    - If you don't have a supported public and private key pair, or don't wish to use any that are available, generate a new SSH key.

    - If you see an existing public and private key pair listed (for example, id_rsa.pub and id_rsa) that you would like to use to connect to CloudLab, you can add the key to the ssh-agent.

    - For more information about generation of a new SSH key or addition of an existing key to the ssh-agent, see "Generating a new SSH key and adding it to the ssh-agent."
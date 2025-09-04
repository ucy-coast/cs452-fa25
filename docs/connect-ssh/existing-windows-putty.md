# Checking for existing SSH keys

Before you generate an SSH key, you can check to see if you have any existing SSH keys.

1. Open File Explorer

2. Search for private SSH key files with a .ppk extension, which indicates that the private key is in PuTTY's proprietary format. PuTTY doesn't have a default location where placing key files will be picked up by default.

3. Either generate a new SSH key or upload an existing key.

    - If you don't have a supported public and private key pair, or don't wish to use any that are available, generate a new SSH key.

    - If you see a private key file listed (for example, id_rsa.ppk) that you would like to use to connect to CloudLab, you can add the key to pageant, PuTTY's authentication agent.

    - For more information about generation of a new SSH key or addition of an existing key to the pageant, see "[Generating a new SSH key and adding it to pageant](generate-windows-putty.md)"
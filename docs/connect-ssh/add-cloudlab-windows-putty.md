# Adding a new SSH key to your CloudLab account

To configure your account on CloudLab to use your new (or existing) SSH key, you'll also need to add the key to your account.

Before adding a new SSH key to your account on CloudLab, you should have:

- [Checked for existing SSH keys](existing-windows-putty.md)
- [Generating a new SSH key and adding it to the ssh-agent](generate-windows-putty.md)

## Adding a new SSH key when registering for an account

1. Click `Choose file` under `SSH Public Key file` in the sign-up form.

<figure>
  <p align="center"><img src="assets/images/cloudlab-register-account.png" height="500"></p>
  <figcaption><p align="center">Figure. Registering for a CloudLab account.</p></figcaption>
</figure>

2. Locate the SSH public key file `.pub` and click `Open`.

## Adding a new SSH to an existing account

1. Log into the CloudLab portal. Once you are logged in: Click on your username (top right), select `Manage SSH keys` and follow the prompts to load the ssh public key.

<figure>
  <p align="center"><img src="assets/images/cloudlab-profile-menu.png"></p>
  <figcaption><p align="center">Figure. Registering for a CloudLab account.</p></figcaption>
</figure>

2. Upload a SSH public key file `.pub` or paste the public key in the text box under `Add Key`. 

<figure>
  <p align="center"><img src="assets/images/cloudlab-manage-ssh-keys-2.png"></p>
  <figcaption><p align="center">Figure. Registering for a CloudLab account.</p></figcaption>
</figure>

3. Click `Add Key`.

The next time you instantiate an experiment, your ssh public key will be loaded onto all (ssh capable) nodes in your experiment, allowing direct access to these nodes using an ssh client.

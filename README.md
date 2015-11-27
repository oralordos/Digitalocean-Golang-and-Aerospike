# How to setup Go and Aerospike in digitalocean Ubuntu 15.10

## Setup Ubuntu
1. Create a server.
  * Make sure to set the server to ubuntu 15.10.
  * 32-bit is recommended unless you have more than 4gb of RAM.
  * Make sure you turn on private networking so you can connect to your database without using up bandwidth.
  * Make sure you have put in an ssh key here to avoid unsecure password logins.
1. Connect to server's root account.
  * Use ssh on unix machines.
    * `ssh root@<ip_address>`
  * Make sure you have setup ssh keys with your computer.
  * Windows user should use putty.
1. Create a user with sudo access.
  * `adduser <username>`
  * Make sure to put in a good password!
  * You can leave the other settings blank, just keep pressing enter.
  * Give the user sudo access.
    * `gpasswd -a <username> sudo`
1. Add ssh key access to new user account.
  * Flip your access to the new user.
    * `su <username>`
  * Move to the home directory.
    * `cd`
  * Create folder and restrict access to only yourself.
    * `mkdir .ssh`
    * `chmod 700 .ssh`
  * Create a file and add ssh key to it.
    * `nano .ssh/authorized_keys`
    * Paste key into file and save and exit with Ctrl-X.
      * Ctrl-Shift-V for unix users to paste.
      * Right-click in window to paste for putty.
  * Restrict the permissions of the file.
    * `chmod 600 .ssh/authorized_keys`
  * Return to root.
    * `exit`
  * Test if it worked.
    * Connect to your new account with either putty or ssh
      * `ssh <username>@<ip_address>
    * If it asks for your password, something went wrong.
1. Restrict ssh access to root and password connections.
  * As root, edit the settings in the ssh config file.
    * `nano /etc/ssh/sshd_config`
    * Set the line `PermitRootLogin` to no to disable root login.
    * Set the line `PasswordAuthentication` to no to disable logging in with a password.
      * Make sure to uncomment the line as well.
  * Make sure you test if you can still access it with normal connection before you disconnect the root terminal.
  * And test that the root login really is disabled.
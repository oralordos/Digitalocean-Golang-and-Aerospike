# How to setup Go and Aerospike in digitalocean Ubuntu 15.10

## Setup Ubuntu user account
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
            * `ssh <username>@<ip_address>`
        * If it asks for your password, something went wrong.
1. Restrict ssh access to root and password connections.
    * As root, edit the settings in the ssh config file.
        * `nano /etc/ssh/sshd_config`
        * Set the line `PermitRootLogin` to no to disable root login.
        * Set the line `PasswordAuthentication` to no to disable logging in with a password.
            * Make sure to uncomment the line as well.
    * Make sure you test if you can still access it with normal connection before you disconnect the root terminal.
    * And test that the root login really is disabled.

## Setup additional helpful items.
1. Setup firewall.
    * Allow ssh through the firewall.
        * `sudo ufw allow ssh`
        * or `sudo ufw allow 22/tcp`
    * Examine the rules.
        * `sudo ufw show added`
    * If everything looks right, enable the firewall.
        * `sudo ufw enable`
    * Make sure everything is running right.
        * `sudo ufw status`
1. Synchronize the system clock.
    * Set timezone.
        * `sudo dpkg-reconfigure tzdata`
        * A graphical menu will allow you to choose a city to sync time with.
    * Install NTP.
        * If you have not used apt-get yet, run `sudo apt-get update`
        * `sudo apt-get install ntp`
        * ntp will automatically place enable run on boot.
1. Create a swapspace.
    * Reserve the space.
        * `sudo fallocate -l <size> /swapfile`
        * <size> is something like `1G` or `512M`
    * Restrict access to root only.
        * `sudo chmod 600 /swapfile`
    * Configure into a swapfile.
        * `sudo mkswap /swapfile`
    * Start using the swapfile.
        * `sudo swapon /swapfile`
    * Setup automatically using the swapfile on boot.
        * `sudo sh -c 'echo "/swapfile none swap sw 0 0" >> /etc/fstab'`
1. This is a good point to make a snapshot of your server.
    * Shut the server down.
        * `sudo poweroff`
    * Save a snapshot in the digitalocean console.

## Get a Go server running.
1. (optional) install Go.
    * If you have Go 1.5 or newer, you can cross compile most programs and transfer the executable.
    * Some packages still require a native Go install to build though.
    * Download Go.
        * `wget <url>`
        * The url for 1.5.1 is `https://storage.googleapis.com/golang/go1.5.1.linux-386.tar.gz`
    * Extract Go from the archive file.
        * `tar -xzf <filename>`
    * Move Go to the default install location.
        * `sudo mv go /usr/local/go`
    * Change owner to root and alter permissions.
        * `sudo chown root:root /usr/local/go`
        * `sudo chmod 755 /usr/local/go`
    * Create workspace folder.
        * `mkdir <workspace_name>{,/bin,/pkg,/src}`
    * Edit environment variables.
        * Add `export PATH=$PATH:/usr/local/go/bin` to `/etc/profile`
        * Add `export GOPATH=$HOME/<workspace_name>` to `~/.profile`
        * Add `export PATH=$HOME/<workspace_name>/bin:$PATH` to `~/.profile`
    * Delete the go archive file.
        * `rm <filename>`
    * Install git.
        * `sudo apt-get install git`
    * Reconnect to the server to allow environment variables to update.
1. Adjust firewall to allow http connections.
    * Allow http through the firewall.
        * `sudo ufw allow http`
        * or `sudo ufw allow 80/tcp`
    * Allow https through the firewall, if needed.
        * `sudo ufw allow https`
        * or `sudo ufw allow 443/tcp`
    * Reload the firewall.
        * `sudo ufw reload`
1. Setup haproxy.
    * Install haproxy.
        * `sudo apt-get install haproxy`
    * Configure haproxy.
        * Edit `/etc/haproxy/haproxy.cfg`
        * Add `retries 3` to the default section.
        * Add `option redispatch` to the default section.
        * Add the following block to the end of the file:
        ```
        listen serv 0.0.0.0:80
            mode http
            option http-server-close
            timeout http-keep-alive 3000
            server serv 127.0.0.1:9000 check```

    * More information available [here](https://www.digitalocean.com/community/tutorials/how-to-use-haproxy-to-set-up-http-load-balancing-on-an-ubuntu-vps).
1. Get your code onto the server.
    * If you are on windows, use WinSCP
    * If you are on a unix machine, use scp
        * `scp <source> <destination>`
        * Add -rp if it is a folder you are transfering.
        * `scp -rp <source> <destination>`
        * The format for remote connections is `<username>@<ip_address>:<path>`
        * Example: `scp -rp ~/Desktop/testServer daniel@107.170.246.157:~/Desktop/testServer`
    * Make sure it is built, whether on your system or on the server directly.
1. Configure systemd.
    * __TODO__
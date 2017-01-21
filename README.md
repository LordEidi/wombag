WombaG
======

**WombaG** (c) 2017 by [SwordLord - the coding crew](https://www.swordlord.com/)

## Introduction ##

**WombaG** is a lightweight read it later service based on Node.js, using the [Wallabag API](https://v2.wallabag.org/api/doc).

_This is still work in progress! Expect things to crash and burn on a regular basis_

Said that, the current version can save and retrieve entries, does update starred and archived flag and deletes whatever you ask to be deleted. Rest to follow...

If you are looking for a lightweight service to store and manage websites and links in, **WombaG** might be for you:

- serving parts of the Wallabag v2 API.
- based on Node.js.

## Status ##

**WombaG** is still in development and should be approached as such:

- the API is only about 50% supported for now (put, get, delete, patch Entries. No such thing as Attributes and Tags yet. But we are working on it).
- there is no authentication yet (run it in your own, trusted networks with a single user only).
- **WombaG** will not have a web UI for a while (isn't planned, but never say never).

We test the **WombaG** with the iOS App. So YMMV.

## Installation ##

First of all, you need a Node.js installation.

### nodejs on Debian ###

Make sure that you have this line in your /etc/apt/sources.list file:

    deb http://YOURMIRROR.debian.org/debian jessie main

and then run:

    sudo apt-get install nodejs nodejs-legacy npm
    // eventually the next line as well
    // sudo ln -s /usr/lib/nodejs/ /usr/lib/node

### nodejs on OSX with homebrew ###

    brew install node
    brew install npm

### Installation of **WombaG** ###

If you want to run **WombaG** under a specific user (node), do this:

    sudo adduser node
    su node
    cd
    mkdir wombag
    cd wombag

Go into the directory where you want to run your copy of **WombaG** and get the latest and greatest:

    cd /home/node/wombag
    git clone https://github.com/LordEidi/wombag.git

And then with the magic of npm get the required libraries

    npm install

If everything worked according to plan, you should now have a new installation of the latest **WombaG**.

### Use supervisord to run **WombaG** as a service ###

Now we want to make sure that **WombaG** runs forever. First install the required software:

    sudo apt-get install supervisor

Then copy the file utilities/wombag_supervisor.conf into your local supervisor configuration directory. This is usually done like this:
 
    cp utilities/wombag_supervisor.conf /etc/supervisor/conf.d/wombag.conf 
    
Make sure you change the configuration to your local setup.

### How to set up transport security ###

Since **WombaG** does not bring it's own crypto, you may need to install a TLS server in front of **WombaG**. You can do so
with nginx, which is a lightweight http server and proxy.

First prepare your /etc/apt/sources.list file (or just install the standard Debian package, your choice):

    deb http://nginx.org/packages/debian/ jessie nginx
    deb-src http://nginx.org/packages/debian/ jessie nginx

Update apt-cache and install nginx to your system.

    sudo update
    sudo apt-get install nginx

Now configure a proxy configuration so that your instance of nginx will serve / prox the content of / for the
**WombaG** server. To do so, you will need a configuration along this example:

    server {
        listen   443;
        server_name  wombag.yourdomain.tld;

        access_log  /var/www/logs/wombag_access.log combined;
        error_log  /var/www/logs/wombag_error.log;

        root /var/www/pages/;
        index  index.html index.htm;

        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   /var/www/nginx-default;
        }

        location / {
            proxy_pass         http://127.0.0.1:8888;
            proxy_redirect     off;
            proxy_set_header   Host             $host;
            proxy_set_header   X-Real-IP        $remote_addr;
            proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
            proxy_buffering    off;
        }

        ssl  on;
        ssl_certificate  /etc/nginx/certs/yourdomain.tld.pem;
        ssl_certificate_key  /etc/nginx/certs/yourdomain.tld.pem;
        ssl_session_timeout  5m;

        # modern configuration. tweak to your needs.
        ssl_protocols TLSv1.1 TLSv1.2;
        ssl_ciphers 'ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES256-GCM-SHA384:DHE-RSA-AES128-GCM-SHA256:DHE-DSS-AES128-GCM-SHA256:kEDH+AESGCM:ECDHE-RSA-AES128-SHA256:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA:ECDHE-ECDSA-AES256-SHA:DHE-RSA-AES128-SHA256:DHE-RSA-AES128-SHA:DHE-DSS-AES128-SHA256:DHE-RSA-AES256-SHA256:DHE-DSS-AES256-SHA:DHE-RSA-AES256-SHA:!aNULL:!eNULL:!EXPORT:!DES:!RC4:!3DES:!MD5:!PSK';
        ssl_prefer_server_ciphers on;
    
        # HSTS (ngx_http_headers_module is required) (15768000 seconds = 6 months)
        add_header Strict-Transport-Security max-age=15768000;
    }

Please check this site for updates on what TLS settings currently make sense:

[https://mozilla.github.io/server-side-tls/ssl-config-generator](https://mozilla.github.io/server-side-tls/ssl-config-generator)

Now run or reset your nginx and start your instance of **WombaG**.

Thats it, your instance of **WombaG** should run now. All logs are sent to stdout for now. Have a look at */libs/log.js* if you want to change the options.

## Configuration ##

All parameters which can be configured right now are in the file *config.js*. There are not that many parameters yet, indeed.

## How to run ##

Point your Wallabag client to the root of **WombaG**. Use any credential you wish, every user gets logged in. Please stay save, install **WombaG** only in a trusted environment. At least until we include authentication mechanisms.

## Contribution ##

If you know JavaScript and would like to help out, send us a note. There is still much work to be done on **WombaG**.

## Dependencies ##

For now, have a look at the package.json file.


## License ##

**WombaG** is published under the GNU General Public Licence version 3. See the LICENCE file for details.
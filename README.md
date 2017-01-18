WomBag
======

**WomBag** (c) 2017 by [SwordLord - the coding crew](http://www.swordlord.com/)

## Introduction ##

**WomBag** is a lightweight read it later service based on Node.js, using the Wallabag API.

_This is still work in progress! Things crash and burn on a regular basis_

Said that, the current version can save and retrieve entries. Rest to follow

If you are looking for a lightweight service to store and manage websites and links in, **WomBag** might be for you:

- serving (some parts of the) Wallabag v2 API.
- based on Node.js.

## Status ##

**WomBag** is still in development and should be handled as such:

- the API is not 100% supported yet.
- there is no such thing as authentication (run it in your own, trusted networks only).
- not having a web UI for the time coming (but never say never).

We test the current version with the iOS App. So YMMV.

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

### Installation of **WomBag** ###

If you want to run **WomBag** under a specific user (node), do this:

    sudo adduser node
    su node
    cd
    mkdir wombag
    cd wombag

Go into the directory where you want to run your copy of **WomBag** and get the latest and greatest:

    cd /home/node/wombag
    git clone https://github.com/LordEidi/wombag.git

And then with the magic of npm get the required libraries

    npm install

If everything worked according to plan, you should now have a new installation of the latest **WomBag**.

### Use supervisord to run **WomBag** as a service ###

Now we want to make sure that **WomBag** runs forever. First install the required software:

    sudo apt-get install supervisor

Then copy the file utilities/wombag_supervisor.conf into your local supervisor configuration directory. This is usually done like this:
 
    cp utilities/wombag_supervisor.conf /etc/supervisor/conf.d/wombag.conf 
    
Make sure you change the configuration to your local setup.

### How to set up transport security ###

Since **WomBag** does not bring it's own crypto, you may need to install a TLS server in front of **WomBag**. You can do so
with nginx, which is a lightweight http server and proxy.

First prepare your /etc/apt/sources.list file (or just install the standard Debian package, your choice):

    deb http://nginx.org/packages/debian/ jessie nginx
    deb-src http://nginx.org/packages/debian/ jessie nginx

Update apt-cache and install nginx to your system.

    sudo update
    sudo apt-get install nginx

Now configure a proxy configuration so that your instance of nginx will serve / prox the content of / for the
**WomBag** server. To do so, you will need a configuration along this example:

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

Now run or reset your nginx and start your instance of **WomBag**.

Thats it, your instance of **WomBag** should run now. All logs are sent to stdout for now. Have a look at */libs/log.js* if
you want to change the options.

## Configuration ##

All parameters which can be configured right now are in the file *config.js*. There are not much parameters yet, indeed.
But **WomBag** is not ready production anyway. And you are welcome to help out in adding parameters and configuration
options.

## How to run ##

Point your Wallabag client to the root of **WomBag**.

## Contribution ##

If you know JavaScript and would like to help out, send us a note. There is still much work to be done on WomBag.

## Dependencies ##

For now, have a look at the package.json file.


## License ##

**WomBag** is published under the GNU General Public Licence version 3. See the LICENCE file for details.
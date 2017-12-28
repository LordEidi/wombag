Wombag
======

![Wombag](https://raw.githubusercontent.com/LordEidi/wombag/master/wombag_logo.png)

**Wombag** (c) 2017 by [SwordLord - the coding crew](https://www.swordlord.com/)

## Introduction ##

**Wombag** is a lightweight, self-hostable read it later service, supporting the [Wallabag API](https://v2.wallabag.org/api/doc).

_This is still work in progress! Expect things to crash and burn on a regular basis. Said that, the current version can save and retrieve entries, updates starred and archived flags and deletes whatever you ask to be deleted. Rest to follow..._

If you are looking for a lightweight service to store and manage websites and links in, then **Wombag** might be for you:

- **Wombag** supports the core functionality of the _Wallabag v2 API_. So you can use your preferred Wallabag Apps and Clients with **Wombag**
- **Wombag** is not based on a scripting engine but is compiled into a _native binary_ for your platform. This makes **Wombag** _very lightweight_.
- **Wombag** makes use of [Gorm](http://jinzhu.me/gorm/) to store the data in a database. We use SQLite3 as our database of choice. But you may use PostgreSQL, MS SQL Server or MySQL, if you prefer. 
- **Wombag** currently does not have its own web frontend. But there is the wombagcli command line interface to configure users and manage the data.


## Status ##

**Wombag** is still under development and should be approached as such:

- the Wallabag v2 API is only about 70% supported for now (PUT, GET, DELETE, PATCH Entries. No such thing as Attributes and Tags yet. But we are working on it).
- **Wombag** will not have a web UI for a while (isn't planned, but never say never). But there is a CLI interface which helps you in managing your data and users.
- **Wombag** does not (yet) support TLS on its own. Make sure to have a proxy like Nginx in front of Wombag for that.

We test **Wombag** with the iOS and Firefox Wallabag App. YMMV.

## Installation ##

### Installation of **Wombag** ###

Follow this instruction if you want to run Wombag on Linux.

Create the user under which you want to run **Wombag**:

    sudo adduser wombag
    su wombag
    cd

Go into the directory where you want to run your copy of **Wombag** and download the latest version from the Github release page:

    Work in Progress

If everything worked according to plan, you should now have a new installation of the latest **Wombag**.

### Use supervisord to run **Wombag** as a service ###

Now we want to make sure that **Wombag** runs forever. First install the required software:

    sudo apt-get install supervisor

Then copy the file _utilities/wombag_supervisor.conf_ into your local supervisor configuration directory. This is usually done like this:
 
    cp utilities/wombag_supervisor.conf /etc/supervisor/conf.d/wombag.conf 
    
Make sure you change the configuration to your local setup.

### How to set up transport security ###

Since **Wombag** does not bring it's own transport encryption, you should install a TLS server in front of **Wombag**. You can do so with nginx, which is a lightweight http server and proxy.

First prepare your /etc/apt/sources.list file (or just install the standard Debian package, your choice):

    deb http://nginx.org/packages/debian/ stretch nginx
    deb-src http://nginx.org/packages/debian/ stretch nginx

Update apt-cache and install nginx to your system.

    sudo update
    sudo apt-get install nginx

Now configure a proxy configuration so that your instance of nginx will serve / prox the content of / for the
**Wombag** server. To do so, you will need a configuration along this example:

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

Now run or reset your nginx and start your instance of **Wombag**.

Thats it, your instance of **Wombag** should run as expected. All logs are sent to stdout for now. Have a look at *config.json* if you want to change the options.

## Configuration ##

All parameters which can be configured right now are in the file *config.js*.

## How to run ##

Point your Wallabag client to the root of **Wombag**. The rest should work as expected.

If you want to play around with the API for a bit, you might be interested in these curl examples:

* add entry: _curl -X POST 'http://0.0.0.0:8081/api/entries/' --data 'url=http://test' -H 'Content-Type:application/x-www-form-urlencoded' -H "Authorization: Bearer (your access token)"_

* get entries: _curl -X GET 'http://0.0.0.0:8081/api/entries/?page=1&perPage=20' -H "Authorization: Bearer (your access token)"_

* get entry: _curl -X GET 'http://0.0.0.0:8081/api/entries/1' -H "Authorization: Bearer (your access token)"_

* patch entry: _curl -X PATCH 'http://0.0.0.0:8081/api/entries/1' --data 'archive=1&starred=1' -H 'Content-Type:application/x-www-form-urlencoded' -H "Authorization: Bearer (your access token)"_


## Contribution ##

If you know Go (or a bit of Angular for a nifty Web Frontend) and would like to help out, send us a note. There is still much work to be done on **Wombag**.


## License ##

**Wombag** is published under the GNU Affero General Public Licence version 3. See the LICENCE file for details.
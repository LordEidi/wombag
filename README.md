Wombag
======

![Wombag](https://raw.githubusercontent.com/LordEidi/wombag/master/wombag_logo.png)

**Wombag** (c) 2017-20 by [SwordLord - the coding crew](https://www.swordlord.com/)

![Build and Package Wombag](https://github.com/LordEidi/wombag/workflows/Build%20and%20Package%20Wombag/badge.svg?branch=master) ![CodeQL Analysis](https://github.com/LordEidi/wombag/workflows/CodeQL%20Analysis/badge.svg?branch=master)

## Introduction ##

**Wombag** is a lightweight, self-hostable read it later service, supporting the [Wallabag API](https://app.wallabag.it/api/doc).

_This is still work in progress! Expect things to crash and burn on a regular basis. Said that, the current version can save and retrieve entries, updates starred and archived flags, it deletes whatever you ask to be deleted and it handles tags. Don't expect annotations to work somewhen soon, it is a mess (different API URL, different authentication)_

If you are looking for a lightweight service to store and manage websites and links in, then **Wombag** might be for you:

- **Wombag** supports the core functionality of the _Wallabag v2 API_. So that you can use your preferred Wallabag Apps and Clients with **Wombag**.
- **Wombag** is not based on a scripting engine but is compiled into a _native binary_ for your platform. This makes **Wombag** _very lightweight_. Just take the application and run it on your server. The application takes care of creating config files it needs when they are not found on a system.
- **Wombag** makes use of [Gorm](http://jinzhu.me/gorm/) to store its data in a database. We use SQLite3 as our database of choice. But you may use PostgreSQL, MS SQL Server or MySQL, if you prefer or need some more oomph at the data layer. 
- **Wombag** currently does not have its own web frontend. But there is the wombagcli command line interface to configure users and manage the data. Which also means there is no admin UI exposed to the world and dog.

## Components ##

**Wombag** consists of two parts: 

- The **wombagd** daemon is the server part, accepting links, doing the readability magic and serving the stored links to your clients. Run the daemon on a server and point your clients towards it.
- The **wombagcli** admin client, running on the commandline. With this client you can manage your users and devices, as well as the links and websites you want to store in **Wombag**. You could even automate some interaction with **Wombag** with the help of this tool. Like fetching mails and adding the links in them to **Wombag**, or something like that.

## Status ##

**Wombag** is still under development:

- the Wallabag v2 API is about 95% supported for now (PUT, GET, DELETE, PATCH Entries and Tags. No such thing as Attributes. Attributes are a bit special on the Wallabag API anyway. They have a different entry point to the regular API, as example).
- there is also no such thing as multi-user support. While you can configure multiple users and devices, all those users will see the same data. This is definitely a planned feature, but not yet done (you might help out, if you need this quicker).
- **Wombag** will not have a web UI for a while (isn't planned, but never say never). But there is our CLI interface which helps you in managing your data and users right from the commandline.
- **Wombag** does not support TLS on its own. Make sure to have a proxy like Nginx in front of Wombag for that. See below for configuration examples on that.

We mostly test and use **Wombag** with the Firefox Wallabag App as well as with [Wallabag Pro](https://itunes.apple.com/gb/app/wallabag-pro/id1187619443) on iOS. YMMV if you use different clients.


## Installation ##

### Installation of **Wombag** ###

Follow this instruction if you want to run Wombag on Linux.

Create the user under which you want to run **Wombag**:

    sudo adduser wombag
    su wombag
    cd

Go into the directory where you want to run your copy of **Wombag** and download the latest version from the Github release page:

https://github.com/LordEidi/wombag/releases

### Add your first user with **wombagcli** ###

Run these in a terminal

    > wombagcli user add testuser testpassword
    > wombagcli device add testdevice password testuser

You can now authenticate with that device on **wombagd**

### Use systemd to run **wombagd** as a service ###

Have a look at the _utilities/wombagd.service_ file.

The commands to install and run **wombagd** as a systemd service are:

    cp utilities/wombagd.service /lib/systemd/system/.
    chmod 755 /lib/systemd/system/wombagd.service
    systemctl enable wombagd.service
    systemctl start wombagd.service
    
Please see journalctl for errors in your log.

### Use supervisord to run **wombagd** as a service ###

If you do not like systemd or if this is not an option, you might want to run it with supervisord. First install the required software:

    sudo apt install supervisor

Then copy the file _utilities/wombag_supervisor.conf_ into your local supervisor configuration directory. This is usually done like this:
 
    cp utilities/wombag_supervisor.conf /etc/supervisor/conf.d/wombag.conf 
    
Make sure you change the configuration to your local setup.

### How to set up transport security ###

Since **Wombag** does not bring it's own transport encryption, you should install a TLS server in front of **Wombag**. You can do so with nginx, which is a lightweight http server and proxy.

Install nginx to your system (if you did not already do so).

    sudo apt update
    sudo apt install nginx

Now configure a proxy configuration so that your instance of nginx will serve / prox the content of / for the **Wombag** server. To do so, you will need a configuration along this example:

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

Thats it, your instance of **Wombag** should run as expected. All logs are sent to stdout for now. Have a look at *wombag.config.json* if you want to change the options.

## Configuration ##

All parameters which can be configured right now are in the file *wombag.config.js*. A default configuration file will be written on the first run of the application.

## How to run ##

Point your Wallabag client to the root of **Wombag**. The rest should work as expected.

If you want to play around with the API for a bit, you might be interested in these curl examples (replace 0.0.0.0 with your domain or IP):

* add entry: _curl -X POST 'http://0.0.0.0:8081/api/entries/' --data 'url=http://test' -H 'Content-Type:application/x-www-form-urlencoded' -H "Authorization: Bearer (your access token)"_

* get entries: _curl -X GET 'http://0.0.0.0:8081/api/entries/?page=1&perPage=20' -H "Authorization: Bearer (your access token)"_

* get entry: _curl -X GET 'http://0.0.0.0:8081/api/entries/1' -H "Authorization: Bearer (your access token)"_

* patch entry: _curl -X PATCH 'http://0.0.0.0:8081/api/entries/1' --data 'archive=1&starred=1' -H 'Content-Type:application/x-www-form-urlencoded' -H "Authorization: Bearer (your access token)"_


## Contribution ##

If you know Go (or a bit of Angular for a nifty Web Frontend) and would like to help out, send us a note. There is still much work to be done on **Wombag**.

## Dependencies ##

Dependencies are managed in the go.mod file.

## License ##

**Wombag** is published under the GNU Affero General Public Licence version 3. See the LICENCE file for details.
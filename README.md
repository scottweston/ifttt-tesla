# ifttt-tesla

A simple golang based service to allow ifttt to talk to the tesla owner api

# Build

```
$ docker build -t ifttt-tesla .
```

# Configure

Create a config file:

```
$ mkdir -p ~/.config
$ cat <<EOM >~/.config/tesla.yml
client_id: "e4a9949fcfa04068f59abb5a658f2bac0a3428e4652315490b659d5ab3f35a9e"
client_secret: "c75f14bbadc8bee3a7594412c31416f8300256d7668ea7e6e7f06727bfb9d220"
username: "your@tesla.login"
password: "your_tesla_password"
auth_tokens:
  - some_random_auth_token
  - another_random_auth_tkn
$ chmod 600 ~/.config/tesla.yml
```

# Run

```
$ docker run --name=ifttt-tesla --restart=always -p 127.0.0.1:3514:3514 -v ~/.config/tesla.yml:/tesla.yml -d ifttt-tesla
```

# Reverse Proxy

Setting up a reverse proxy is outside the scope of this document however you
will need to create a reverse proxy http server (I pity the fool that doesn't
SSL this service!) to map access through to your service. If you want a quick
and easy way to do this with automatic LetsEncrypt SSL protection look into
using [caddy](https://caddyserver.com/)

## Caddy example

Setup your FQDN and then create a `Caddyfile` like:

```
api.tesla.my.domain {
  proxy / localhost:3514
}
```

Caddy will then try and obtain and use LetsEncrypt certs for your service.
As the only protection to calling the API is the AuthToken it is important
that you at least SSL protect the service so that the AuthToken is not able
to be packet captured in transit. Not SSLing your service open it up to
potential abuse. I have specifically not put in the ability to remote
**start** a Tesla because even with SSL there is a non-zero chance of this
API being exploited.

# Setting up IFTTT

The following endpoints are supported:

  * `/honk/{vehicle}`
  * `/unlock/{vehicle}`
  * `/lock/{vehicle}`
  * `/set_charge_limit/{vehicle}/{limit}`
  * `/start_charge/{vehicle}`
  * `/stop_charge/{vehicle}`
  * `/flash/{vehicle}`
  * `/open_charge_port/{vehicle}`
  * `/set_temperature/{vehicle}/{temp}`

If you only own 1 vehicle then `{vehicle}` will be `0`

Now create IFTTT applets that looks like:

![ifttt applet](https://raw.githubusercontent.com/scottweston/ifttt-tesla/master/ifttt.com_applets_43679679d.png)

Obviously you'll need to replace `api.tesla.my.domain` and `some_random_auth_token`
with your own values for your installation.

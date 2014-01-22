Fudge Edge Router
=================

A HTTP/HTTPS websocket capable edge router for a Fudge installation.  This is meant to be installed as a container
on all endpoint host nodes.

It's loosely based on [hipache-nginx](https://github.com/samalba/hipache-nginx) from Sam Alba.

** !Notice ** This is still a work in progress and isn't compatible with the Redis format used by Hipache.

Contains:
---------

* Nginx for SSL termination
* Varnish
* Supervisor startup

Depends On:
-----------

* [stackbrew/ubuntu:saucy](https://index.docker.io/u/stackbrew/ubuntu/)

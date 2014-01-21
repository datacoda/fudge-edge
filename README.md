Fudge Edge Router
=================

Contains:
---------

* Nginx for SSL termination
* Varnish
* Supervisor startup

Depends On:
-----------

* [stackbrew/ubuntu:saucy](https://index.docker.io/u/stackbrew/ubuntu/)


Usage:
------

A HTTP/HTTPS websocket capable edge router for a fudge installation.  This is meant to be installed as a container
on all endpoint host nodes.


How It Works:
-------------

This container is part of a bigger system.  The Fudge network assumes that webapp containers are publicly exposed
to their host nodes via -P.

eg:

```bash
docker run -d -P -expose 80 dataferret/websocket-echo 80
```

This binds a randomized port to the container's 80/tcp.  Use a firewall on the host node to suppress public access to
the randomized ports.  Cross-connecting nodes can be done via VPC or tinc.


Credit:
-------

* https://github.com/samalba/hipache-nginx/network
* https://github.com/marcusramberg/hipache-nginx
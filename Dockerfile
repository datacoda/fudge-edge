#
# Image Name:: dataferret/fudge-router
#
# Copyright 2014, Nephila Graphic.
#
# Licensed under the Apache License, Version 2.0 (the 'License');
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an 'AS IS' BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#


FROM stackbrew/ubuntu:saucy
MAINTAINER Ted Chen <ted@nephilagraphic.com>

# Enable the necessary sources and upgrade to latest
RUN echo "deb http://archive.ubuntu.com/ubuntu saucy main universe" > /etc/apt/sources.list
RUN echo "deb http://archive.ubuntu.com/ubuntu saucy-security main universe" >> /etc/apt/sources.list
RUN apt-get update && apt-get upgrade -y -o DPkg::Options::=--force-confold

# Install packages
RUN apt-get update && apt-get install supervisor varnish nginx-extras redis-server -y

# Cleanup
# RUN apt-get clean && rm -rf /var/cache/apt/* && rm -rf /var/lib/apt/lists/*

ADD ./varnish/varnish_opts /etc/default/varnish
ADD ./supervisor/ /etc/supervisor/conf.d/

EXPOSE 80 443

CMD ["/usr/bin/supervisord", "-n"]

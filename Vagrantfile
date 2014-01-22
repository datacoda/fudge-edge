# -*- mode: ruby -*-
# vi: set ft=ruby :

$script = <<SCRIPT
#echo "deb http://archive.ubuntu.com/ubuntu saucy main universe" > /etc/apt/sources.list
#echo "deb http://archive.ubuntu.com/ubuntu saucy-security main universe" >> /etc/apt/sources.list
#dpkg --configure -a
#apt-get update && apt-get upgrade -y -o DPkg::Options::=--force-confold
add-apt-repository ppa:nginx/stable

# Install packages
apt-get install supervisor varnish nginx-extras redis-server lua-nginx-redis -y
apt-get install git nano wget curl golang-go -y

# For vagrant development, use the built-in user
mkdir /home/vagrant/go -p
chown vagrant.vagrant /home/vagrant/go
echo "export GOPATH=/home/vagrant/go" >> /home/vagrant/.profile
su - vagrant -c 'go install github.com/codegangsta/martini'
su - vagrant -c 'go install github.com/garyburd/redigo/redis'
git clone https://github.com/agentzh/lua-resty-redis.git /opt/lua-resty-redis
SCRIPT


Vagrant.configure("2") do |config|

  # Required Ubuntu 13.10
  config.vm.box = "saucy64"
  config.vm.box_url = "http://cloud-images.ubuntu.com/vagrant/saucy/current/saucy-server-cloudimg-amd64-vagrant-disk1.box"


  config.vm.hostname = "docker-fudge-edge"
  config.vm.network :private_network, ip: "33.33.33.25"

  config.vm.synced_folder "./", "/home/vagrant/go/src/github.com/dataferret/fudge-edge"

  config.vm.provider :virtualbox do |vb|
     # Use VBoxManage to customize the VM. For example to change memory:
     vb.customize ["modifyvm", :id, "--memory", "512"]
  end


  config.vm.provision "shell", inline: $script
end

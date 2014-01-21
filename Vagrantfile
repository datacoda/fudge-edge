# -*- mode: ruby -*-
# vi: set ft=ruby :

$script = <<SCRIPT
echo "deb http://archive.ubuntu.com/ubuntu saucy main universe" > /etc/apt/sources.list
echo "deb http://archive.ubuntu.com/ubuntu saucy-security main universe" >> /etc/apt/sources.list
dpkg --configure -a
apt-get update && apt-get upgrade -y -o DPkg::Options::=--force-confold

# Install packages
apt-get install supervisor varnish nginx-extras redis-server lua-nginx-redis -y
apt-get install git nano wget curl -y

SCRIPT


Vagrant.configure("2") do |config|

  # Required Ubuntu 13.10
  config.vm.box = "saucy64"
  config.vm.box_url = "http://cloud-images.ubuntu.com/vagrant/saucy/current/saucy-server-cloudimg-amd64-vagrant-disk1.box"


  config.vm.hostname = "docker-fudge-edge"
  config.vm.network :private_network, ip: "33.33.33.25"


  config.vm.provider :virtualbox do |vb|
     # Don't boot with headless mode
     # vb.gui = true

     # Use VBoxManage to customize the VM. For example to change memory:
     vb.customize ["modifyvm", :id, "--memory", "128"]
  end


  config.vm.provision "shell", inline: $script
end

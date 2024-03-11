# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "centos/7"
  config.vm.synced_folder ".", "/code" , type: "virtualbox"
  config.vm.provision "shell", inline: <<-SHELL
    sudo yum update -y
    sudo yum install -y vim
    sudo yum install -y epel-release
    sudo yum install -y golang
    echo "" >> ~/.bashrc
    echo "cd /code/src" >> ~/.bashrc
  SHELL
end

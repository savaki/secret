require File.expand_path(File.dirname(__FILE__) + '/tasks/env')

Vagrant.configure("2") do |config|
  config.vm.box       = "precise64"
  config.vm.box_url   = "http://files.vagrantup.com/precise64.box"

  config.vm.network :public_network

  config.vm.provision :shell, :inline => "sudo dpkg -i /vagrant/secret-tool_#{VERSION}_amd64.deb"

  config.vm.provider :virtualbox do |vb|
  end
end

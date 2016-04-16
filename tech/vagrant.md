## vagrant
box = image

[Official Vagrant document](https://www.vagrantup.com/)

- download vargrant at https://www.vagrantup.com/downloads.html
- download VirtualBox at https://www.virtualbox.org/wiki/Downloads


### vagrant box
- list , show boxs list
- add [box name] [url] , add box to local
- remove [box name] , remove a box



### vagrant init [box name]
init project

### vagrant up 
start project

### vagrant ssh [hostname]
connect to vm

### vagrant halt [vm name]
shutdown vm

### vagrant reload [vm name]
reload vm for vagrant file

### vagrant package
package box


### vagrant file
```
config.vm.define :app do |app_config|
    app_config.vm.customize ["modifyvm", :id, "--name", "app", "--memory", "512"]
    app_config.vm.box = "ubuntu-12-10"
    app_config.vm.host_name = "app"
    app_config.vm.network :hostonly, "33.33.13.10"
end
config.vm.define :db do |db_config|
    db_config.vm.customize ["modifyvm", :id, "--name", "db", "--memory", "512"]
    db_config.vm.box = "ubuntu-12-10"
    db_config.vm.host_name = "db"
    db_config.vm.network "public_network"
end
config.vm.define :db do |db_config|
    db_config.vm.customize ["modifyvm", :id, "--name", "db", "--memory", "512"]
    db_config.vm.box = "ubuntu-12-10"
    db_config.vm.host_name = "db"
    db_config.vm.network :bridge
end



```

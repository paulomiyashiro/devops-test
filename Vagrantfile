

$nodescript = <<-SCRIPT
	cat /vagrant/id_rsa.pub >> /home/vagrant/.ssh/authorized_keys
SCRIPT

$controllerscript = <<-SCRIPT
	sudo cp -r /vagrant/id_rsa /home/vagrant/.ssh/id_rsa
	sudo chmod 400  /home/vagrant/.ssh/id_rsa
	sudo chown vagrant:vagrant /home/vagrant/.ssh/id_rsa
SCRIPT


Vagrant.configure("2") do |config|

	config.vm.box = "centos/7"
	
	config.vm.define "node1" do |machine|

		machine.ssh.insert_key = false

		machine.vm.network "private_network", ip: "172.17.177.21"

		machine.vm.network "forwarded_port", guest: 8080, host: 8080

		machine.vm.network "forwarded_port", guest: 8443, host: 8443

		machine.vm.provision "shell", inline: $nodescript

	end

	config.vm.define 'controller' do |machine|

		machine.ssh.insert_key = false
	
		machine.vm.network "private_network", ip: "172.17.177.11"

		machine.vm.provision "shell", inline: $controllerscript

		machine.vm.provision "provision", type: "ansible_local", run: "once" do |ansible|

			ansible.vault_password_file = "vaultpass"
			ansible.playbook 			= "provision.yml"		
			ansible.verbose        		= true
			ansible.install        		= true
			ansible.limit          		= "all"
			ansible.inventory_path 		= "inventory"
			ansible.provisioning_path 	= "/vagrant"
		end

		machine.vm.provision "build", type: "ansible_local", run: "never" do |ansible|

			ansible.vault_password_file = "vaultpass"
			ansible.playbook 			= "build.yml"		
			ansible.verbose        		= true
			ansible.install        		= true
			ansible.limit          		= "controller"
			ansible.inventory_path 		= "inventory"
			ansible.provisioning_path 	= "/vagrant"

			ansible.extra_vars = {
				tag: "v1"
			}
		end

		machine.vm.provision "deploy", type: "ansible_local", run: "never" do |ansible|

			ansible.vault_password_file = "vaultpass"
			ansible.playbook 			= "deploy.yml"		
			ansible.verbose        		= false
			ansible.install        		= true
			ansible.limit          		= "nodes"
			ansible.inventory_path 		= "inventory"
			ansible.provisioning_path 	= "/vagrant"

			ansible.extra_vars = {
				tag: "v1"
			}
		end
	end
end

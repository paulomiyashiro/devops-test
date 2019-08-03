# Devops Test
Projeto de avaliação de conhecimentos de automação com Ansible
## Pré-requisitos
Para executar o projeto é necessária a instalação do Vagrant e Virtual Box.
### Instalação dos Pré-requisitos
As instruções para instalação estão descritas em:
* [Vagrant](https://www.vagrantup.com/docs/installation/) - Instalação do Vagrant
* [Virtual Box](https://www.virtualbox.org/manual/UserManual.html#intro-installing) - Instalação do Virtualbox
## Preparação do Ambiente
A preparação do ambiente deverá ser realizada em dois passos que consistem em incluir os certificados, chaves e senhas na raiz do projeto e provisionar o ambiente no vagrant.
### Inclusão dos Certificados e Chaves
As chaves, senhas e certificados não estão adicionadas neste pacote e deverão ser descompactadas no diretório raiz do projeto.

Deverão ser descompactados na raiz do projeto os arquivos descritos abaixo: 
* id_rsa - chave privada que será utilizada para comunicação ssl entre as instâncias do Vagrant
* id_rsa.pub - chave publica que será utilizada para comunicação ssl entre as instâncias do Vagrant
* server.crt - certificado que será publicado na aplicação para servir requisições https
* server.key - chave utilizada na aplicação para servir requisições https
* vaultpass - senha para desencriptografar as credencias definidas no ansible-vault

## Quick Start
Para disponibilizar a aplicação após a cópia dos arquivos adicionais digite:
```
vagrant up
vagrant up --provision-with build,deploy
```

## Ciclo de Vida da Aplicação
É possível através dos playbooks criados gerenciar o ciclo de vida da aplicação.

### Provision
Após a clonagem do repositório será necessário provisionar o ambiente no vagrant.

Serão criadas duas instâncias no Vagrant:
* Controller
Esta instância contém o ansible e é reponsável por executar os playbooks.
* Node1
Esta instância é contém a aplicação publicada em Docker.

Na raiz do projeto digite:
```
vagrant up
```
Após essa execução o ambiente estará disponível para publicação do serviço.
### Build
No processo de build o script ansible compila a aplicação, gera imagem docker, realiza o login no Dockerhub e publica a imagem.

Para executar o procedimento digite:
```
vagrant up --provision-with build
```
### Deploy
No processo de deploy o script ansible realiza o login no Dockerhub, realiza o download da imagem e inicializa a instância Docker.
```
vagrant up --provision-with deploy
```
### Update
Para atualizar o projeto após realizar as alterações no código fonte.

- Altere o arquivo Vagrantfile na área de provision do build e deploy para a versão desejada.
```
machine.vm.provision "build", type: "ansible_local", run: "never" do |ansible|
...
  ansible.extra_vars = {
    tag: "v1" //Ex: Troque para v2 
  }
...
end
machine.vm.provision "deploy", type: "ansible_local", run: "never" do |ansible|
...
  ansible.extra_vars = {
    tag: "v1" //Ex: Troque para v2 
  }
...
end
```
- Recarregue os arquivos do projeto, execute o build e deploy
```
vagrant reload --provision-with build,deploy
```
## Estrutura do projeto
```
|-group_vars
|--all
|---all.yml   //arquivo de configurações que associa as credenciais definidas no arquivo vault
|---vault     //arquivo criptografado das credenciais do docker hub
|-roles
|--docker-build     //geração de tags e imagens e publicação no docker hub
|--docker-deploy    //baixa imagem do dockerhub e cria instância docker
|--docker-install   //instala o docker nas maquinas do Vargrant
|--docker-build     //responsável pela compilação, geração de tags e imagens e publicação no docker hub
|--docker-login     //login no docker hub
|--epel-repository  //adiciona repositórios necessários para instalação do docker
|-src //código fonte da aplicação
|-ansible.cfg //configurações do ansible
|-provision.yml //playbook responsável pelo provisionamento das instâncias
|-build.yml //playbook responsável pela geração da imagem docker e armazenamento no dockerhub 
|-deploy.yml //playbook responsável pela criação da instância no docker e realização de testes
|-inventory //arquivo com definição dos hosts gerenciado pelo ansible
|-Vagrantfile //arquivo com descrição das configurações do Vagrant
|-Dockerfile //arquivo de configurações do Docker, compila a aplicação e disponibiliza os serviços
```
# Autor
Paulo Miyashiro


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
* id_rsa
* id_rsa.pub
* server.crt
* server.key
* vaultpass
### Provisionamento do ambiente
Após a clonagem do repositório será necessário provisionar o ambiente no vagrant.

Na raiz do projeto digite:
```
vagrant up
```
Após essa execução o ambiente estará disponível para publicação do serviço.

## Ciclo de Vida da Aplicação
É possível através dos playbooks criados gerenciar o ciclo de vida da aplicação.
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

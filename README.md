# api

## Instalando as depedências

É necessário instalar os serviçoes mongo e redis. Recomendamos a utilização do docker, exemplo:

**docker pull mongo**

**docker run --name ufpblor-api-mongo -p 0.0.0.0:27017:27017 -d mongo**

## Executando os testes

Após a instalação das depedências, executar o comando para rodar os testes.

**make test-all**

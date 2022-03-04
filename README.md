# kemper.com.br

Este projeto cria um site experimental para aulas de programação.

Como rodar o site localmente:

```shell
  make build
```

## Como gerar o site

A forma mais prática de montar o site e todos os plugins é com o comando `make build`.

Ele irá gerar um `build` do executável principal e `build` para todos os `plugins` do site.

Os principais comandos são:

```shell
  # Monta o site e todos os plugins
  make build
  
  # Monta apenas o site
  make build-site
  
  # Monta todos os plugins
  make plugins
  
  # Monta o plugin mongodb-language 
  make plugin-mongodb-language
  
  # Monta o plugin mongodb-menu
  make plugin-mongodb-menu
  
  # Monta o plugin mongodb-user
  make plugin-mongodb-user
  
  # Monta o plugin sqlite-languages
  make plugin-sqlite-languages
  
  # Monta o plugin sqlite-menu
  make plugin-sqlite-menu
  
  # Monta o plugin sqlite-user
  make plugin-sqlite-user
  
  # Apaga todos os binários, site e plugins
  make clean
  
  # Apaga todos os plugins
  make plugins-clean
  
  # Apaga o build do site
  make build-clean
```

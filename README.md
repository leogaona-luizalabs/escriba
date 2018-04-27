# Escriba
Slackbot criado para organizar os artigos escritos para o blog do luizalabs.

Inicialmente, o escriba possui os seguintes comandos:
    - *help* - _help_
    - *hello* - _healthcheck do bot_
    - *pending* *reviews* - _lista artigos pendentes de revisão_
    - *pending* *publications* - _lista artigos pendentes de publicação_
    - *add* `url` - _adiciona artigo_
    - *approve* `url` - _aprova um artigo_
    - *publish* `url` - _marca um artigo como publicado_

O projeto foi criado em Go `1.10.1` e utiliza uma base de dados MySQL para persistir os links e aprovações. Caso você deseje subir um container mysql localmente, use o comando:
```shell
docker run -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=escriba -e MYSQL_USER={$USER} -e MYSQL_PASSWORD={$PASSWORD} --name mysql -d -p 3306:3306 mysql:5.7.21
```

O diretório `migrations` possui todos os scripts para criar a base de dados corretamente.

## Variáveis de ambiente

### slack_token
tipo: string
descrição: token de integração do slack

### mysql_dsn
tipo: string
descrição: data source name para conexão mysql (exemplo: `USER:PASSWORD@tcp(HOST:PORTA)/DATABASE`)

### draft_approvals
tipo: int
descrição: valor mínimo de aceites necessário para que um artigo seja considerado pronto para publicação
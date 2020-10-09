# Banking-go

Essa é uma aplicação escrita em go, que faz uso de apenas algumas libs e implementa uma transação bancárias.

Para definir arquitetura da aplicação, foram utilizados conceitos da arquitetura hexagonal o qual me permite ter baixissimo acoplamento entre as partes do sistema.
Dito isso, para uma melhor organização do codigo, foram utilizadas as seguintes nomenclaturas: 
   - driving (o que dirige a regra de negócio), rest endpoints, infra.
   - domain (regras de negócio)
   - driven (o que é dirigido pela regra de negócio), client http endpoints, repository, infra.

Para exemplificar:
 Driving ou user-side : onde o usuário interage com a aplicação. 
 Domain ou bussiness logic: onde as regras de negócio devem estar puramente explicitadas nessa camada.
 Driven ou server-side: onde estão os códigos de infra que são dirigidos pela aplicação, nessa camada normalmente temos interação com banco de dados, definição de serviços externos, normalmente são os "atores" gerenciados pela aplicação.
 
 veja mais [aqui](https://blog.octo.com/en/hexagonal-architecture-three-principles-and-an-implementation-example/).
 
Uma vez dito isso, defini a seguinte estrutura:

```shell
banking-go:
|------> domain (regras de negócio com suas entidades e value objects)
|------|-------> entity 
|------\-------> vo
|------> driven (o que é dirigido pelas regras de negócio)
|------\-------> repository 
|------> driving (regras de interação com usuario, http handlers, validação de payload e afins)
|------|-------> handlers
|------|-------> middlewares
|------|-------> request
|------\-------> response
|------> rsources (arquivos e recursos para subir a aplicação)
|------\------> migration
```

## Instalação

Para execução e utilização do sistema é necessário ter docker e docker-compose.
A aplicação e banco de dados irão subir respectivamentes nas portas `8080` e `3306`.

Para executar a aplicação e o banco de dados rode o seguinte comando.

```shell script
make start-all
```

Para  executar os testes unitários e ou abrir o coverage no browser utilize os comandos localizados no Makefile.

Obs: este projeto faz uso de hot reload, ou seja, cada alteração nos arquivos .go, o Daemon rodando no container irá rebuildar o projeto.

## Documentação da api

A documentção atende todos os casos de uso da aplicação, tais como:
 - Criar uma conta;
 - Encontrar uma conta;
 - Criar uma transação para uma conta especifica;
 
Ela foi escrita na ferramenta postman, e pode ser encontrada [aqui](https://documenter.getpostman.com/view/359751/TVRj793K#b7524462-3340-4932-89c1-ef9c4b6d486e).

## Definições do domínio

Existe uma diferença chave entre o requisito e o dominio implementado, para nosso modelo de transação foram definidos os seguintes tipos/operações

| id | descrição       | natureza/modo  |
| ---|:---------------:| -------:|
| 1  | Compra à vista  | Débito  |
| 2  | Compra parcelada| Débito |
| 3  | Saque           | Débito |
| 4  | Pagamento       | Crédito |

Deste modo, nenhuma transação é registrada com valor monetário negativo.  

## todo

- [ ] terminar os testes unitários 
- [ ] implementar um caso de teste ponta à ponta

# go-arq-hexagonal

# Sobre o projeto
Este projeto é uma implementação da arquitetura hexagonal em Go. A arquitetura hexagonal, também conhecida como Ports and Adapters, permite a criação de aplicações com componentes desacoplados, facilitando testes e manutenção.


## O que é a Arquitetura Hexagonal?
A arquitetura hexagonal, é uma abordagem para criar aplicações de software de maneira que as regras de negócio (domínio) sejam independentes das interfaces externas, como bancos de dados, interfaces de usuário, serviços web, entre outros. A ideia principal é que o núcleo da aplicação (domínio) deve ser isolado de detalhes técnicos e infraestruturais.

## Importância da Arquitetura Hexagonal
- Desacoplamento: Facilita a substituição de componentes externos sem afetar a lógica de negócios.
- Testabilidade: Permite testar a lógica de negócios independentemente dos componentes externos.
- Manutenibilidade: Torna o código mais organizado e fácil de manter.
- Flexibilidade: Facilita a adição de novos componentes externos, como diferentes tipos de interfaces de usuário ou bancos de dados.

# Como rodar
## Requisitos Obrigatórios ##

- Go 1.16 ou superior
- Docker

## Instalação
```bash
git clone https://github.com/Igorcand/go-arq-hexagonal.git
cd go-arq-hexagonal
```

## Uso
```bash
docker-compose up --build -d
```

# Detalhes
## Application
A pasta /application contém a regra de negócio da aplicação. É aqui que estão implementados os serviços e casos de uso que representam a lógica de negócios central da aplicação. Esta camada é independente dos detalhes de infraestrutura e dos adaptadores, garantindo que a lógica de negócios permaneça desacoplada e testável.

- services: Contém os serviços que implementam a lógica de negócios. Esses serviços são utilizados pelos adaptadores para executar operações específicas do domínio.


## Adapters
### CLI Adapter
O adaptador CLI utiliza a biblioteca cobra para criar comandos de linha de comando. Ele permite que os usuários interajam com a aplicação através do terminal, executando comandos específicos para diferentes funcionalidades. Este adaptador facilita a execução de tarefas administrativas e automação de processos.

### DB Adapter
O adaptador DB utiliza a biblioteca sql para se comunicar com o banco de dados. Ele implementa os métodos definidos nas interfaces do domínio para realizar operações de leitura e escrita no banco de dados. Este adaptador garante que a lógica de acesso a dados esteja separada da lógica de negócios, seguindo o princípio de desacoplamento da arquitetura hexagonal.

### Web Adapter
O adaptador Web utiliza a biblioteca gorilla/mux para lidar com as requisições HTTP. Ele define os endpoints da aplicação e mapeia as requisições para os casos de uso apropriados no domínio. Este adaptador permite que a aplicação exponha uma API web, facilitando a interação com outros sistemas e clientes.


# Ajuda

## Contribuição
Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests.
1. Fork o projeto
2. Crie sua feature branch (git checkout -b feature/nova-feature)
3. Commit suas mudanças (git commit -m 'Adiciona nova feature')
4. Push para a branch (git push origin feature/nova-feature)
5. Abra um pull request


# Autor

Igor Cândido Rodrigues

https://www.linkedin.com/in/igorc%C3%A2ndido/
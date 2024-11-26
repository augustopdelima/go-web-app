# Currículo Web App

Este é um projeto de sistema web em Go para gerenciamento de currículos. Ele permite cadastrar, listar e visualizar detalhes de currículos por meio de uma interface web. 

## Grupo
- Augusto Pasquali, Gabriel Eberhardt e Lucas Pacheco


## Estrutura do Projeto

- **`main.go`**: Contém o ponto de entrada da aplicação.
- **`app/`**: Define o ambiente da aplicação.
- **`db/`**: Configurações para inicialização do banco de dados.
- **`middleware/`**: Middleware para rotas e proteção.
- **`models/sqlite/`**: Implementação do modelo de currículos usando SQLite.
- **`templates/`**: Arquivos HTML para renderização do front-end.
- **`static/`**: Arquivos estáticos como CSS.

## Funcionalidades

- **Cadastro de Currículos**: Permite criar um novo currículo.
- **Listagem de Currículos**: Exibe todos os currículos cadastrados.
- **Detalhes de Currículos**: Visualiza as informações completas de um currículo.
- **Proteção Contra CSRF**: Usa o pacote `gorilla/csrf` para proteger formulários.
- **Headers Seguros**: Adiciona cabeçalhos de segurança às respostas HTTP.

## Dependências

Este projeto utiliza as seguintes bibliotecas:

- [gorilla/csrf](https://github.com/gorilla/csrf): Middleware para proteção contra ataques CSRF.
- [bulma/css](https://github.com/jgthms/bulma): Framework css
- [modernc/sqlite](https://pkg.go.dev/modernc.org/sqlite#section-readme): Driver sqlite

## Pré-requisitos

- [Go](https://go.dev/) versão 1.22.3 ou superior.


## Como Executar

1. **Clone o repositório**:
   ```bash
   git clone https://github.com/augustopdelima/go-web-app.git
   cd resume-web-app
   ```

2. **Baixar depêndencias**
   ```bash
   go mod tidy
   ```

3. **Rodar o projeto**
   ```bash
   go run main.go
   ```
   
4. **Fazer o build (Opcional)**
   ```bash
   go build main.go
   ./main
   ```

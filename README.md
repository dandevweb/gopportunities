# gopportunities

`gopportunities` é uma aplicação baseada em Go projetada para gerenciar oportunidades de forma eficiente.

## Funcionalidades

- Gestão de oportunidades (CRUD simples para estudo)

## Instalação

1. Clone o repositório:
   ```sh
   git clone https://github.com/dandevweb/gopportunities.git
   ```
2. Navegue até o diretório do projeto:
   ```sh
   cd gopportunities
   ```
3. Instale as dependências:
   ```sh
   go mod tidy
   ```

## Configuração

Defina as seguintes variáveis de ambiente no seu shell ou arquivo `.env`:

```sh
DB_HOST=localhost
DB_PORT=3306
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_DATABASE=seu_banco_de_dados
```

## Executando a Aplicação

Para executar a aplicação, execute o seguinte comando:

```sh
go run main.go
```

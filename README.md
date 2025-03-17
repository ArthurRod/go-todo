# go-todo

Todo list de terminal que salva e consome os dados em um arquivo `.json`, feita para fins didáticos

---

## 🌍 **Instruções em Português**

### Pré-requisitos

1. **Go 1.18+** - Você precisa ter o Go instalado em sua máquina. Caso não tenha, pode instalar o Go [aqui](https://golang.org/dl/).
2. **Git** - Para clonar o repositório.

### Como rodar o projeto

1. **Clone o repositório**

   Clone este repositório usando o comando:

   ```bash
   git clone https://github.com/username/go-todo.git
   ```

2. **Navegue até o diretório do projeto**

   cd go-todo

   ```bash
   git clone https://github.com/username/go-todo.git
   ```

3. **Baixe as dependências**

   No diretório raiz do projeto, execute:

   ```bash
   go mod tidy
   ```

4. **Execute o projeto**

   Para rodar o projeto localmente, execute o seguinte comando:

   ```bash
   go run cmd/main.go
   ```

   Isso irá iniciar o aplicativo de To-Do List no terminal, e você poderá interagir com ele.

### Como compilar o projeto

Para criar um binário do seu aplicativo, siga os passos abaixo:

1. **Compile o projeto**

   Execute o comando abaixo para compilar o código:

   ```bash
   go build -o go-todo ./cmd
   ```

   Isso criará um binário chamado `go-todo` no diretório raiz do projeto.

2. **Execute o binário**

   Após compilar, você pode rodar o binário diretamente:

   ```bash
   ./go-todo
   ```

3. **(Opcional) Mova para o diretório de binários do sistema**

   Se quiser usar o comando go-todo de qualquer lugar no seu sistema,
   mova o binário para um diretório presente no PATH, como `/usr/local/bin` (Linux), ou
   adicione o caminho do binário na variável PATH (Windows):

   ```bash
   sudo mv go-todo /usr/local/bin/
   ```

Agora você pode rodar o aplicativo de qualquer lugar do terminal com o comando:

    ```bash
    go-todo
    ```

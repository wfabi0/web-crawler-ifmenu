# Projeto Web Crawler em Go

Este é um projeto de web crawler desenvolvido em [Go](https://go.dev/) para extrair o cardápio do dia no site do [IFMG Campus São João Evangelista](https://solucoes.sje.ifmg.edu.br/CardapioUan/cardapio.php).

## Desenvolvimento Local

Para instalar e executar este projeto, siga as etapas abaixo:

1. Certifique-se de ter o Go instalado em sua máquina. Caso contrário, faça o download e instale-o a partir do site oficial do Go (https://golang.org).

2. Clone este repositório para o seu ambiente local:

    ```
    git clone https://github.com/wfabi0/web-crawler-ifmenu.git
    ```

3. Acesse o diretório do projeto:

    ```
    cd web-crawler-ifmenu
    ```

4. Instale as dependências do projeto:

    ```
    go mod download
    ```

5. Compile o código fonte:

    ```
    go build -o out/web-crawler-ifmenu.exe
    ```

6. Basta executar o executável criado na pasta out.

## Download do executável

Você também pode baixar o executável do Web Crawler diretamente na seção de [releases](https://github.com/wfabi0/web-crawler-ifmenu/releases) deste repositório.

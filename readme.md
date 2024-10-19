# Stress Test

## Objetivo
Criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas. O sistema gerará um relatório com informações específicas após a execução dos testes.

## Entrada de Parâmetros via CLI
Os parâmetros podem ser fornecidos ao executar a aplicação da seguinte forma:

- `--url`: URL do serviço a ser testado.
- `--requests`: Número total de requests a serem enviados.
- `--concurrency`: Número de chamadas simultâneas.

## Exemplo de Uso
- Baixar o Repositório

  ```bash
  git clone https://github.com/guirialli/stress_test
  ```

- Compile o projeto

  ```bash
  docker build -t carga-tester .
  ```

- Exemplo de execução

  ```bash
  docker run carga-tester --url=http://google.com --requests=1000 --concurrency=10
  ```

  
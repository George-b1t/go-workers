# go-workers

Este projeto implementa um servidor master em Go, responsável por gerenciar workers e distribuir tarefas para processamento. O sistema simula falhas de execução nos workers e reatribui as tarefas automaticamente em caso de falha, proporcionando resiliência e desempenho distribuído.


## Funcionalidades

- **Servidor Master**: Gerencia workers e clientes, distribuindo tarefas.
- **Workers**: Processam tarefas com 20% de chance de falha simulada.
- **Cliente Interativo**: Envia tarefas via linha de comando.
- **Reatribuição de Tarefas**: Falhas resultam em nova atribuição automática.
- **Timeout**: 10 segundos para resposta dos workers.



## Estrutura do Projeto

- `internal/`: Contém o código do cliente interativo que envia as tarefas.
    - `app/`
      - `client/`: Codigos refernte ao cliente.
        - `client.go` : Implementação do client
      - `worker/`: Codigos refernte ao worker.
        - `worker.go` : Implementação do worker
      - `server/`: Codigos refernte ao server.
        - `server.go` : Implementação do server

- `cmd/`
  - `client/`: Codigos refernte ao cliente.
    - `main.go` : Entrypoint do client
  - `worker/`: Codigos refernte ao worker.
    - `main.go` : Entrypoint do worker
  - `server/`: Codigos refernte ao server.
    - `main.go` : Entrypoint do server

- `go.mod/`: Arquivo de gerenciamento de dependências do Go.
- `go.sum/`: Arquivo de checksum para garantir a integridade das dependências.


## Execução

Siga os passos abaixo para executar o servidor, o worker e o cliente:

1. **Abra 3 terminais** (T1, T2, T3).
2. **Inicie o servidor (T1)**:
   No terminal T1, execute o comando abaixo para iniciar o servidor master:
   ```bash
   go run cmd/server/main.go
   ```
3. **Inicie o Worker (T2)**
  No terminal T2, execute o comando abaixo para iniciar o worker:
   ```bash
   go run cmd/worker/main.go
   ```
4. **Inicie o Cliente (T3)**
  No terminal T3, execute o comando abaixo para iniciar o cliente interativo:
   ```bash
   go run cmd/worker/main.go
   ```
5. **Teste o cliente (T3)**
   - No Terminal T3, inicie o cliente e digite a tarefa a ser processada:
      ```bash
      Digite a tarefa a ser processada (ou 'bye' para sair): tarefa_teste
      Resultado recebido: resultado do 'tarefa_teste' processado
      ```

    - No Terminal T1 (Logs do Servidor Master), o servidor master receberá a tarefa e a atribuirá a um worker:
      ```bash
      025/03/15 09:46:35 Tarefa recebida do cliente 127.0.0.1:55345: tarefa_teste
      2025/03/15 09:46:35 Atribuindo tarefa 'tarefa_teste' ao worker 127.0.0.1:55342
      ```

    - No Terminal T2 (Logs do Worker), o worker receberá e processará a tarefa::
      ```bash
      2025/03/15 09:46:35 Tarefa recebida: tarefa_teste
      2025/03/15 09:46:38 Tarefa concluída: resultado do 'tarefa_teste' processado
      ```   

###### Dica
>Você pode abrir mais de um worker, se desejar, para aumentar a capacidade de processamento do servidor.

## 🤝 Colaboradores

Agradecemos às seguintes pessoas que contribuíram para este projeto:

<table>
  <tr>
    <td align="center">
      <a href="https://github.com/YuriGarciaRibeiro">
        <img src="https://avatars.githubusercontent.com/u/81641949?v=4" width="100px;" alt="Yuri Garcia"/><br>
        <sub><b>Yuri Garcia</b></sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/George-b1t">
        <img src="https://avatars.githubusercontent.com/u/67129166?v=4" width="100px;" alt="George Soares"/><br>
        <sub><b>George Soares</b></sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/MaykeESA">
        <img src="https://avatars.githubusercontent.com/u/81484737?v=4" width="100px;" alt="Mayke Erick"/><br>
        <sub><b>Mayke Erick</b></sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/GugaAAndrade">
        <img src="https://avatars.githubusercontent.com/u/105755546?v=4" width="100px;" alt="Gustavo Andrade"/><br>
        <sub><b>Gustavo Andrade</b></sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/jmcanario1">
        <img src="https://avatars.githubusercontent.com/u/79545726?v=4" width="100px;" alt="João Marcelo"/><br>
        <sub><b>João Marcelo</b></sub>
      </a>
    </td>
  </tr>
</table>


## Licença
Distribuído sob licença MIT.
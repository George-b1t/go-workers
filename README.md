# go-workers

Este projeto implementa um servidor master que gerencia workers para executar tarefas recebidas de clientes.  
As conexões são reconhecidas como “workers” ou “clientes” e processadas conforme suas funções.

## Visão Geral
- O servidor fica escutando na porta `:12345`.
- Clientes podem enviar tarefas de forma contínua, até enviarem `bye`.
- Workers ficam persistentes esperando tarefas, respondendo o resultado ou “fail”.

## Execução
1. Compile o projeto com `go build`.
2. Inicie o servidor:  
   ```
   .\go-workers.exe
   ```
3. Conecte-se como worker ou cliente (ex: via telnet ou netcat).

## Estrutura
- `main.go`: Lógica principal do servidor, incluindo registro de workers, enfileiramento de tarefas e gestão de conexões.

## 🤝 Colaboradores

Agradecemos às seguintes pessoas que contribuíram para este projeto:

<table>
  <tr>
    <td align="center">
      <a href="#">
        <a href="https://github.com/YuriGarciaRibeiro">
          <img src="https://avatars.githubusercontent.com/u/81641949?v=4" width="100px;" alt="Foto do Brenno Oliveira no GitHub"/><br>
        </a>
        <br>
        <sub>
          <b>Yuri Garcia</b>
        </sub>
      </a>
    </td>
    <td align="center">
      <a href="#">
        <a href="https://github.com/George-b1t">
          <img src="https://avatars.githubusercontent.com/u/67129166?v=4" width="100px;" alt="Foto do Brenno Oliveira no GitHub"/><br>
        </a>
        <br>
        <sub>
          <b>George Soares</b>
        </sub>
      </a>
    </td>
    <td align="center">
      <a href="#">
        <a href="https://github.com/MaykeESA">
          <img src="https://avatars.githubusercontent.com/u/81484737?v=4" width="100px;" alt="Foto do Mayke Erick no GitHub"/><br>
        </a>
        <br>
        <sub>
          <b>Mayke Erick</b>
        </sub>
      </a>
    </td>
    <td align="center">
      <a href="#">
        <a href="https://github.com/GugaAAndrade">
          <img src="https://avatars.githubusercontent.com/u/105755546?v=4v=4" width="100px;" alt="Foto do Brenno Oliveira no GitHub"/><br>
        </a>
        <br>
        <sub>
          <b>Gustavo Andrade</b>
        </sub>
      </a>
    </td>
    <td align="center">
      <a href="#">
        <a href="https://github.com/jmcanario1">
          <img src="https://avatars.githubusercontent.com/u/79545726?v=4" width="100px;" alt="Foto do Brenno Oliveira no GitHub"/><br>
        </a>
        <br>
        <sub>
          <b>João Marcelo</b>
        </sub>
      </a>
    </td>
  </tr>
</table>

## Licença
Distribuído sob licença MIT.
# Ping-Pong em Go

Este é um programa simples em Go que simula um jogo de "ping-pong" usando goroutines e canais para demonstrar concorrência.

## Como funciona

O programa consiste em três funções principais:

1. **`ping`**: Envia a mensagem `"ping"` para um canal em um loop infinito.
2. **`pong`**: Envia a mensagem `"pong"` para o mesmo canal em um loop infinito.
3. **`show`**: Recebe as mensagens do canal e as imprime no console, com um intervalo de 1 segundo entre cada mensagem.

O programa usa goroutines para executar as funções `ping`, `pong` e `show` simultaneamente, e um canal (`chan string`) para comunicação entre elas.


## Como executar

1. Certifique-se de ter o Go instalado em sua máquina. Você pode verificar isso executando:
   ```bash
   go version
   ```

2. Clone este repositório ou copie o código do `main.go` para um arquivo com a extensão `.go`.

3. Navegue até o diretório onde o arquivo está localizado e execute o seguinte comando:
   ```bash
   go run main.go
   ```

4. O programa começará a imprimir `"ping"` e `"pong"` alternadamente no console, com um intervalo de 1 segundo entre cada mensagem.

5. Para encerrar o programa, pressione `Enter` no terminal.

## Exemplo de saída

```
ping
pong
ping
pong
ping
pong
...
```

## Observações

- O programa usa `fmt.Scanln(&input)` para manter a função `main` em execução. Sem isso, o programa terminaria imediatamente após iniciar as goroutines.
- Este é um exemplo básico de concorrência em Go, utilizando canais para comunicação entre goroutines.

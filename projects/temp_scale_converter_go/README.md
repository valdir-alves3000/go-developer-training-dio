# Temperature Scale Converter

## Descrição
Este projeto em Go realiza a conversão de escalas termométricas, exibindo os pontos de fusão e ebulição da água nas escalas **Kelvin (K)º**, **Celsius (C)º** e **Fahrenheit (F)º**.

## Como Funciona
O programa define constantes para valores de referência:
- **Ponto de fusão da água:** 273 Kº
- **Ponto de ebulição da água:** 373 Kº
- Converte esses valores para Celsius e Fahrenheit.
- Exibe os resultados formatados no terminal com cores para facilitar a visualização.

## Tecnologias Utilizadas
- **Linguagem:** Go
- **Formatação no terminal:** Códigos ANSI para colorização

## Como Executar o Projeto
1. Certifique-se de ter o Go instalado em sua máquina.
2. Clone este repositório:
   ```sh
   git clone https://github.com/valdir-alves3000/temp_scale_converter_go.git
   ```
3. Acesse a pasta do projeto:
   ```sh
   cd temp_scale_converter_go
   ```
4. Execute o programa:
   ```sh
   go run main.go
   ```

## Exemplo de Saída
```sh
Ponto de Ebulição da Água:
Kelvin: 373 Kº
Celsius: 100 Cº
Fahrenheit: 212 Fº

Ponto de Fusão da Água:
Kelvin: 273 Kº
Celsius: 0 Cº
Fahrenheit: 32 Fº
```
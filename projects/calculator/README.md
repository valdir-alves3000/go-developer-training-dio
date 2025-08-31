# 📐 Go Math Functions

[![Build Status](https://img.shields.io/badge/tests-passing-brightgreen)](#)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](#)
[![Made with Go](https://img.shields.io/badge/made%20with-Go-00ADD8.svg)](https://golang.org/)

Projeto simples em **Go** implementando funções matemáticas básicas (Soma, Subtração, Multiplicação e Divisão) com testes unitários cobrindo casos de uso e borda.

---

## 🚀 Funcionalidades

- `Sum(nums ...int) int` → Soma uma lista de números.
- `Subtract(nums ...int) int` → Subtrai os números na ordem recebida.
- `Multiply(nums ...int) int` → Multiplica uma lista de números.
- `Divide(a, b int) int` → Divide dois números, com verificação de divisão por zero.

---

## 🧪 Executando os Testes

Para rodar todos os testes com detalhes:

```bash
go test -v
```

## ⚠️ Tratamento de Erros

A função `Divide` dispara um **panic** com a mensagem:

```
Divisão por zero não é permitida
```

se o segundo parâmetro for `0`.
O teste `TestDivideByZero` valida exatamente essa mensagem.
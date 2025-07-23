# Grandma Clients API

API RESTful para gerenciamento de usuários com dados de endereço, construída em Go com GORM e SQLite.

## Visão Geral

Esta API permite criar, ler, atualizar e deletar usuários e seus respectivos endereços, usando arquitetura limpa com camadas de entidades, repositórios, casos de uso e handlers HTTP.

### Tecnologias principais

* Linguagem: Go
* Banco de Dados: SQLite via GORM ORM
* Router HTTP: Gorilla Mux

---

## Funcionalidades

* CRUD completo de usuários e endereços (Create, Read, Update, Delete)
* Associação 1:1 entre Usuário e Endereço com relacionamento via GORM
* API REST com JSON
* Suporte a listagem e filtros simples via frontend (exemplo com template HTML)

---

## Endpoints principais

| Método | Endpoint          | Descrição             | Corpo (JSON) / Resposta         |
| ------ | ----------------- | --------------------- | ------------------------------- |
| POST   | `/api/users`      | Criar usuário         | `CreateUserRequest` JSON        |
| GET    | `/api/users`      | Listar usuários       | Array de `UserResponse` JSON    |
| GET    | `/api/users/{id}` | Buscar usuário por ID | `UserResponse` JSON             |
| PUT    | `/api/users/{id}` | Atualizar usuário     | `UpdateUserRequest` JSON        |
| DELETE | `/api/users/{id}` | Deletar usuário       | HTTP 204 No Content |


---

## Modelo de resposta esperado

```json
{
  "id": "921981ad-6864-4b79-a709-ab70d5f7426a",
  "name": "Luigi",
  "age": 40,
  "address": {
    "city": "Rio de Janeiro",
    "state": "RJ"
  }
}
```
---

## Como rodar localmente

Pré-requisitos: Go 1.20+, SQLite

```bash
git clone https://github.com/valdir-alves3000/go-developer-training-dio.git
cd projects/grandma-clients-api

go mod tidy
go run main.go
```

A API estará disponível em [http://localhost:8080](http://localhost:8080).
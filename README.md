# GoApi

API em Go para consulta de endereço por CEP (Código de Endereçamento Postal). O projeto usa duas fontes externas (ViaCEP e BrasilAPI) com fallback: se a primeira falhar, a segunda é tentada automaticamente.

## Estrutura do projeto (DDD)

- **cmd/api** — Ponto de entrada. Permite rodar como servidor HTTP ou como menu CLI, via flags.
- **cmd/cli** — Menu interativo no terminal para consultar CEP.
- **cmd/server** — Servidor HTTP que expõe a rota de consulta por CEP.
- **internal/application** — Caso de uso: buscar endereço por CEP (orquestra as fontes com fallback).
- **internal/infrastructure** — Clientes HTTP (ViaCEP e BrasilAPI) e mappers que normalizam as respostas para um único tipo `Address`.

## Pré-requisitos

- Go 1.22 ou superior (para rotas com `{zipcode}` no path).

## Como rodar

### Servidor HTTP (para frontend ou API)

```bash
go run ./cmd/api --server
```

O servidor sobe na porta **8080**. Exemplo de requisição:

```bash
curl http://localhost:8080/zipcode/01001000
```

A rota é `GET /zipcode/{zipcode}` (CEP no path, não na query string).

### Menu CLI (terminal)

```bash
go run ./cmd/api --cli
```

Menu interativo: opção 1 para consultar CEP, opção 2 para sair.

### Padrão (sem flag)

Se nenhuma flag for passada, o programa sobe o **servidor** por padrão.

## Normalização do CEP

As duas APIs retornam o CEP em formatos diferentes (com ou sem hífen). O projeto normaliza para **apenas dígitos** (ex.: `01001000`) em todos os retornos.

## Licença

Uso livre para estudo e projetos pessoais.

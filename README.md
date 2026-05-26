# 🦟 MosquiLab — Projeto de Educação Sanitária

Landing page e painel administrativo para o MosquiLab, projeto de extensão da Univille,
financiado pela FAPESC e Governo do Estado de Santa Catarina.

---

## Stack

| Camada      | Tecnologia            |
|-------------|----------------------|
| Frontend    | React + TypeScript + Vite + Tailwind CSS |
| Backend     | Go + Gin             |
| Banco       | PostgreSQL           |
| Auth        | JWT (httpOnly cookie + header) |

---

## Estrutura

```
mosquilab/          ← Frontend (React)
mosquilab-backend/  ← Backend (Go)
```

---

## Setup — Frontend

```bash
cd mosquilab
cp .env.example .env         # configure VITE_API_URL
npm install
npm run dev                  # http://localhost:5173
```

### Imagens necessárias

Coloque o arquivo `instagram.png` em:
```
mosquilab/public/img/instagram.png
```

---

## Setup — Backend

### Pré-requisitos
- Go 1.22+
- PostgreSQL rodando

```bash
cd mosquilab-backend
cp .env.example .env         # configure DATABASE_URL, JWT_SECRET, etc.
go mod tidy
go run ./cmd/main.go
```

O servidor sobe em `http://localhost:8080`.

Na primeira execução, o banco é migrado automaticamente e um usuário admin padrão é criado:
- **E-mail**: `admin@mosquilab.univille.br`
- **Senha**: `mosquilab2024` (altere via variável `ADMIN_PASSWORD`)

---

## Rotas da Aplicação

| Rota              | Descrição                          |
|-------------------|------------------------------------|
| `/`               | Landing page pública               |
| `/admin`          | Redireciona para `/admin/login`    |
| `/admin/login`    | Tela de login (oculta do público)  |
| `/admin/dashboard`| CRUD de eventos (requer login)     |

---

## Endpoints da API

| Método | Rota                    | Auth | Descrição                |
|--------|-------------------------|------|--------------------------|
| GET    | `/api/agenda`           | —    | Lista eventos futuros    |
| POST   | `/api/auth/login`       | —    | Login admin              |
| GET    | `/api/admin/agenda`     | JWT  | Lista todos os eventos   |
| POST   | `/api/admin/agenda`     | JWT  | Cria evento              |
| PUT    | `/api/admin/agenda/:id` | JWT  | Edita evento             |
| DELETE | `/api/admin/agenda/:id` | JWT  | Remove evento            |

---

## Parceiros

- **Univille** — Universidade da Região de Joinville (60 anos)
- **FAPESC** — Fundação de Apoio à Pesquisa Científica e Tecnológica do Estado de Santa Catarina
- **Governo de SC** — Secretaria de Estado da Ciência, Tecnologia e Inovação

---

## Identidade Visual

```css
--green-dark:  #1a5c2a;
--green-light: #6abf3e;
--sand:        #f0ead6;
--orange:      #e87722;
--teal:        #3aada8;
--navy:        #1a4f8a;
```

Tipografia: **Bebas Neue** (títulos) + **Nunito** (corpo)

---

## Contato

Instagram: [@mosquilab.univille](https://www.instagram.com/mosquilab.univille/)

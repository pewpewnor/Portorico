# Portorico

Portorico is a web app prototype for creating and publishing simple personal or
company websites from reusable templates. Users can register, sign in, create
and edit website entries, choose a template, and view a published page. The
current frontend includes a basic landing template; the backend exposes account,
session, and website endpoints.

## Stack

- Frontend: Next.js 14, React, TypeScript, NextUI, and Axios
- Backend: Go, Fiber, and PostgreSQL
- Database changes: SQL migration files run by golang-migrate

The frontend currently calls the backend at
[http://localhost:8000](http://localhost:8000).

## Requirements

Install Node.js (or Bun), Go 1.21 or newer, and PostgreSQL.

## Run locally

Create a PostgreSQL database, then create `.env.local` in the repository root:

~~~dotenv
DB_URI=postgres://USER:PASSWORD@localhost:5432/portorico?sslmode=disable
SERVER_PORT=8000
~~~

Start the API from the repository root:

~~~sh
go run ./server/src
~~~

In another terminal, install frontend dependencies and start Next.js:

~~~sh
npm install
npm run dev
~~~

Open [http://localhost:3000](http://localhost:3000). The frontend is configured
for the API on port 8000; keep that port or update the Axios base URL in the
frontend before changing it.

## Development note

The backend currently forces a migration version and runs a down/up migration
cycle during startup. Use a disposable local database until that startup
behavior is reviewed; do not point this prototype at data you need to keep.

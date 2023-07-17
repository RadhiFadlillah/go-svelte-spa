# Go + Svelte SPA

This is a template for developing SPA using Go and Svelte.

## Recommended IDE Setup

[VS Code][vscode] + [Go][go-vscode] + [Svelte][svelte-vscode].

## Installation

To make it easier, you might want to use [`degit`][degit] like this:

```
degit RadhiFadlillah/go-svelte-spa new-project
```

## Features

-   Based on official Svelte + Vite template.
-   Live reload Go binary using [`air`][air].
-   Hash based routing using [`svelte-spa-router`][svelte-spa-router].
-   Automatically embed the assets to Go app when building for production.

## Workflow

When in development, run this command in terminal:

```
npm run web-dev
```

The command above will run `vite` in live reload mode, which will build our Svelte frontend whenever there are changes. After that, open separate terminal then run this command:

```
npm run go-dev
```

The command above will run `air` which will build our Go backend whenever there are changes, and also serve our app.

Once development finished, you can build the final binary by running:

```
npm run build
```

## Directory Structure

```
go-svelte-spa
├── backend/ ────────────── the Go code for our backend
├── build/ ──────────────── the directory which contain development build
├── dist/ ───────────────── the directory which contain production build
├── frontend/ ───────────── the Svelte code for our frontend
├── node_modules/ ───────── node dependencies
├── .air.toml ───────────── config for `air`, live reload for Go
├── .eslintignore ───────┬─ config for `eslint`
├── .eslintrc.cjs ───────┘
├── .gitignore ──────────── items to be ignored by Git
├── .prettierignore ─────┬─ config for `prettier`
├── .prettierrc ─────────┘
├── app-assets_dev.go ───┬─ our `main` Go code
├── app-assets_prod.go   │
├── app-log_dev.go       │
├── app.go ──────────────┘
├── go.mod ──────────────┬─ Go module
├── go.sum ──────────────┘
├── jsconfig.json ───────── config for JS language server
├── package-lock.json ───┬─ config for NPM
├── package.json ────────┘
├── README.md ───────────── this readme file
├── svelte.config.js ────── config for `svelte`
└── vite.config.js ──────── config for `vite`
```

[degit]: https://github.com/Rich-Harris/degit
[air]: https://github.com/cosmtrek/air
[svelte-spa-router]: https://github.com/ItalyPaleAle/svelte-spa-router
[vscode]: https://code.visualstudio.com/
[go-vscode]: https://marketplace.visualstudio.com/items?itemName=golang.Go
[svelte-vscode]: https://marketplace.visualstudio.com/items?itemName=svelte.svelte-vscode

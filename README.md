# vimrg

Little tool to support glob for fzf-lua in nvim

## Requirements

-   go
-   rg

## Install

```sh
go get github.com/dinhhuy258/vimrg
```

## Syntax:

Find `myservice` in all files

```sh
vimrg myservice
```

Find `myservice` in test files

```sh
vimrg myservice > *test*
```

Find `myservice` in non-test files

```sh
vimrg myservice > !*test*
```

Find `myservice` in non-test controller files

```sh
vimrg myservice > !*test* *controller*
```

Find `myservice` not in `/test/` directories

```sh
vimrg myservice > !*/test/*
```

## Configure with fzf-nvim

-   Install [fzf-lua](https://github.com/ibhagwan/fzf-lua)
-   Use `:lua require('fzf-lua').live_grep({ cmd = "vimrg" })` to open the fzf live_grep supporting globs

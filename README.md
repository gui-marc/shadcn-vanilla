# shadcn-vanilla

[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Issues](https://img.shields.io/github/issues/gui-marc/shadcn-vanilla)](https://github.com/gui-marc/shadcn-vanilla/issues)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](CONTRIBUTING.md)  
![Blade](https://img.shields.io/badge/Blade-Laravel-red)
![Nunjucks](https://img.shields.io/badge/Nunjucks-yellow)
![Go%20Templates](https://img.shields.io/badge/Go%20Templates-blue)
![Go%20Templ](https://img.shields.io/badge/Go%20Templ-lightgrey)

---

**shadcn-vanilla** is a framework-agnostic port of [shadcn/ui](https://ui.shadcn.com), built to work with **any fullstack framework** — Laravel, Django, Go, Node.js, and more.  
It ships TailwindCSS + AlpineJS components that can be output in multiple templating syntaxes such as **Blade**, **Nunjucks**, **Go Templates**, and **Go Templ**.

---

## ✨ What is shadcn-vanilla?

The goal of shadcn-vanilla is to bring the flexibility and quality of shadcn/ui components to **non-React projects**.  
Instead of relying on a single frontend stack, shadcn-vanilla:

- Provides **prebuilt UI components** using TailwindCSS & AlpineJS.
- Outputs components in **different template syntaxes** via adapters.
- Installs components directly into your project via a **Go-based CLI**.
- Keeps the same **customization freedom** as shadcn/ui — you own the code.

---

## ⚙️ How It Works

1. **Canonical Components**  
   Each component is written once using a neutral placeholder syntax.

2. **Adapters**  
   Adapters transform the canonical version into different templating syntaxes:
   - Blade (Laravel)
   - Nunjucks
   - Go Templates
   - Go Templ

3. **CLI**  
   The `shadcn-vanilla` CLI installs the chosen component in your project with:
   - Correct template syntax
   - TailwindCSS classes
   - AlpineJS behavior (if applicable)


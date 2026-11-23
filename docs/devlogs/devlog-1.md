# Devlog

## 2026-11-23

- Built: initialized the project tooling with a Makefile, `.editorconfig`, `.gitignore`, Lefthook, and a GitHub Actions CI workflow.
- Learned: a Go module can expose several binaries later while sharing internal packages; GitHub Actions builds and tests code on pushes and pull requests.
- Confused: initially, shell commands appended to `lefthook.yml` caused YAML parsing errors; Lefthook configuration must contain only valid YAML.

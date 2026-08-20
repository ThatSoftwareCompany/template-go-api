# AI agent guidance

`AGENTS.md` is the canonical repository instruction file for Codex and other coding agents. `CLAUDE.md`, `GEMINI.md`, and `.github/copilot-instructions.md` should point contributors back to the same rules.

Agents must inspect Git status, remotes, branches, and existing files before modifying the repository. They must preserve unrelated work, keep backend and frontend repositories separate, avoid secrets and legal-license decisions, run relevant validation, and leave public API/OpenAPI documentation synchronized.

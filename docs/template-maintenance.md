# Template maintenance

The manifest is the source of truth for the template identity, version, source repository, compatibility, and generated origin.

## Implemented lifecycle safeguards

- `scripts/setup.sh` records `template_commit` from the source Git commit when available.
- `scripts/setup.sh` records `generated_from` from the derived repository identifier or module path.
- `cmd/template -command validate` validates the manifest contract.
- `scripts/validate-template.sh` verifies required template files and blocks a real `.env` file.
- CI runs the lifecycle validation and writes build outputs outside the repository root.

## Planned update mechanism

The future maintenance workflow will:

1. Detect a newer `template_version` or template commit from `ThatSoftwareCompany/template-go-api`.
2. Compare the generated repository's recorded `generated_from` and compatibility fields.
3. Create an automatic pull request in each derived repository with the template changes.
4. Run the derived repository's Go, integration, Docker, and compatibility checks.
5. Record incompatible or breaking changes in the template changelog and PR description.
6. Leave small file conflicts for manual resolution; it must never overwrite unrelated application code automatically.

The update workflow/script is intentionally planned rather than enabled in this foundation. It must use least-privilege GitHub permissions and must not merge generated pull requests automatically.

## Compatibility policy

- Patch and compatible minor template updates may be proposed automatically when the declared Go and PostgreSQL compatibility ranges remain valid.
- Changes to public routes, configuration semantics, database migrations, dependency major versions, or security behavior require an explicit compatibility note.
- Breaking changes require a new documented template version and a migration note for derived repositories.

## Post-merge repository checklist

After the backend and separate frontend PRs are reviewed and merged:

- [x] Mark `ThatSoftwareCompany/template-go-api` as a GitHub Template Repository in `Settings -> General -> Template repository`.
- [ ] Mark the frontend template repository as a GitHub Template Repository in `Settings -> General -> Template repository`.
- [ ] Confirm that template generation preserves the manifest and setup-script behavior.

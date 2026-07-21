# Repository Guidelines

## Project Structure & Module Organization

Root-level `watch*.ahk` scripts define hotkeys and recording workflows for different languages or contexts. `paste.ahk` handles paste behavior, while `prompt.txt` and `transtriptionPrompt*.txt` provide prompt variants. The Go helper lives in `whisper-autohotkey/`; its executable package is under `cmd/whisper-autohotkey/`. Packaged third-party tools and release binaries belong in `bin/`. GitHub Actions workflows are in `.github/workflows/`.

Treat `config.template.json` as the public configuration example. Keep local secrets and generated artifacts such as `config.json`, `rec.mp3`, and `script.ahk` untracked.

## Build, Test, and Development Commands

- `task build` builds the Go executable from the repository root.
- `task package` builds, then copies the executable into `bin/whisper-autohotkey/`.
- `.\whisper-autohotkey\whisper-autohotkey.exe` runs the built helper from the repository root, where it can read `config.json` and `rec.mp3`.
- `cd whisper-autohotkey; go test ./...` compiles and runs all Go tests. No test files exist yet, so add them with behavior changes.
- `cd whisper-autohotkey; go vet ./...` performs standard Go static checks.

Use AutoHotkey v1.1 for `.ahk` scripts; this project is not written for AutoHotkey v2.

## Coding Style & Naming Conventions

Format Go changes with `gofmt`. Follow idiomatic Go naming: exported identifiers use `PascalCase`, internal functions use `camelCase`, and test files end in `_test.go`. Keep command-specific Go files focused (`ahk.go`, `gpt4.go`, `whisper.go`).

For AutoHotkey, preserve existing v1 command syntax and keep language/context variants named consistently, such as `watch_ru.ahk` or `watch_context.ahk`. Avoid editing vendored files in `bin/` unless intentionally updating a packaged dependency.

## Testing Guidelines

Place Go tests beside the code they cover and prefer table-driven tests for text transformations. Name tests after behavior, for example `TestToSnakeCase`. For AutoHotkey changes, manually verify the F8 start/stop flow, recording creation, transcription, clipboard output, and paste behavior on Windows.

## Commit & Pull Request Guidelines

History uses short, descriptive subjects such as `Fixed Unicode Support` and `Update README.md`. Keep subjects imperative, focused, and under roughly 72 characters; separate unrelated changes.

Pull requests should explain user-visible behavior, list verification commands, and note Windows/AutoHotkey versions tested. Link relevant issues. Include screenshots or logs only when UI behavior or failure output benefits from them, and redact API keys and dictated content.

## Security & Configuration

Never commit `config.json` or an OpenAI API key. Add new options to `config.template.json` with safe defaults and document them in `README.md`.

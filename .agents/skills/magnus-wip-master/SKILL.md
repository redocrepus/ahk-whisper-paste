---
name: magnus-wip-master
description: Manage the WIP handoff file at <target-repo>/wip/wip.md. At significant checkpoints update current state, scope, next steps and references to useful info such as file:line references and commit SHAs.
---

# Magnus WIP Master

Manages `<target-repo>/wip/wip.md` — the live handoff document tracking current scope, state, and next steps across agent sessions.

## When to invoke

1. Before starting working on any development task.
2. Before committing.
3. After completing a plan item.
4. When scope or next steps change materially.
5. When you discover useful info that can help next agent that picks up work.
6. When surfacing a blocker or deferred issue.

## Target repo

If current work is related to a specific repo, or that repo and its submodules, use `<target-repo>` to refer to that repo.
If the work is across multiple repos, use `<target-repo>` to refer to this WS.

## How to maintain `<target-repo>/wip/wip.md`

1. Read `<target-repo>/wip/wip.md`.
2. Update **Scope** if the task or SOW changed.
3. Update **Current State** with what was just done — include file:line references and commit SHAs where useful.
4. Update **Next Steps** — mark completed `[x]`, in-progress `[>]`, pending `[ ]`.
5. Remove stale info. Move completed items to main docs when applicable.
6. Write the file to `<target-repo>/wip/wip.md`.

## Formatting rules

1. All lists numbered.
2. Use `[ ]` for pending, `[>]` for in-progress, `[x]` for done.
3. Include file:line references and commit SHAs where they help orient the next agent.
4. No stale content — remove or archive when done.
5. Keep it short. This is a handoff, not a journal.

## Templates

If `wip` directory does not exist, or does not contain `README.md`, create it with the following content:

```markdown
# WIP Directory

This directory contains work-in-progress artifacts.
It must be removed before merging to target branch.
```

If `<target-repo>/wip/wip.md` does not exist, create it:

```markdown
# WIP Handoff

## Scope

## Current State

## Next Steps
```

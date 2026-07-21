---
name: describe-pr
description: Create concise pull request titles and descriptions by comparing the current branch against a specified remote branch (for example, origin/main). Use when asked to describe/summarize a PR, branch diff, changelist, or draft PR metadata.
---

# Describe PR

Write a concise PR title and description from git history and diff, grouped into logical change clusters.
Use parallel agents for two tasks: 1. Analyze changes 2. Maintain PR title rules.

## Maintaining PR Title rules

1. Maintain PR Title rules in this WS's knowledge base.
2. Use Google's Open Knowledge Format (OKF).
3. Use skills for OKF if available.
4. If you don't know where the knowledge base is, start searching by checking `.okf`, `.okf/repos/<target-repo>/pr-title-rules.md`
5. If rules older than 6 months, review PR titles of current repo in the last 6 months, infer rules and update rules accordingly.
6. If there are not enough PRs to infer rules, use this definition for title:

```
<TYPE>: <change>[, <change>...]. [<TYPE>: <change>[, <change>...]. ...]

where <TYPE> ::= NEW | FIXED | CHANGED | IMPROVED | REMOVED | INTERNAL | TESTS
```

where NEW is a new user facing feature, FIXED is a bug fix, CHANGED is a change to existing functionality, IMPROVED is an improvement to existing functionality, REMOVED is a removal of existing functionality, INTERNAL is an internal change that does not affect users, and TESTS is a change to tests.
(User can be a developer or end user, depending on the context).

## Workflow

1. Resolve comparison target.
   Use the branch explicitly requested by the user.
   If user did not specify, resolve default remote target branch and ask user to confirm.
   Talk to user until valid target branch is confirmed.
2. In parallel:
   1. Review/Maintain PR title rules
   2. Review PR contents for creating title and description.
      1. Warn user about any uncommitted changes or commits that are not pushed to the remote branch.
         Read all commit messages of current branch since common ancestor with the target branch.
         Read the code changes that will be introduced by merging the current branch into the target branch like in:
         `git diff origin/<target-branch>...origin/<pr-branch>`

      2. Build logical clusters.
         Group changes by meaning, not by commit.
         Typical clusters examples: - 3D view UI improvements - Optimized milling algorithm - Tightened metal deposition testing - Formatting/cleanup/refactoring

3. Compose a single-line title that captures the primary changes according to the rules. Keep it specific and brief (target <= 100 characters).
4. Compose change description as concise numbered bullets. Use one bullet per logical change cluster. Keep each bullets as short as possible while still conveying the change and its impact.
5. Ask `magnus-wip-master` skill to save the result as `<target-repo>/wip/PR.md`.
   If `magnus-wip-master` skill is not available, do it yourself, and add `<target-repo>/wip/README.md` with the following content:

```markdown
# WIP Directory

This directory contains work-in-progress artifacts.
It must be removed before merging to target branch.
```

## PR.md Format

Use this exact style:

```markdown
# <concise PR title>

## Changes:

1. <Change title>: <concise summary of related changes and impact>.
2. <Change title>: <concise summary of related changes and impact>.
```

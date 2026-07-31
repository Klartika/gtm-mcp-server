# AGENTS.md

Guidance for working on `Klartika/gtm-mcp-server` — a Go MCP server for Google Tag
Manager. See `README.md` for what it does and `ARCHITECTURE.md` for how it is built.

## This is a fork

Upstream is `sprawz/gtm-mcp-server`. Bare `gh issue list` / `gh pr create` resolve to
**upstream**, not here — always pass `--repo Klartika/gtm-mcp-server`. Upstream's issue
tracker is not ours; nothing we defer belongs there.

## Issue tracking

GitHub issues are the authoritative tracker for individual work items. Work with them
continuously, without being asked.

- **One issue per deferred item**, assigned to a milestone where the repo uses them. If
  it isn't an issue, it isn't tracked — never keep a parallel list in a doc or in memory.
- **Write issues to stand alone**: the evidence, the ground truth you verified, and the
  concrete fix. A future reader won't have the conversation.
- **Never defer silently.** If work is dropped, narrowed, or postponed, open an issue in
  the same turn and say so.
- **Close issues in the PR that resolves them** (`Closes #N`), not afterwards.
- **Check for staleness proactively**, especially issues predating recent decisions.
  Prefer a **dated update comment** over rewriting the body — the original framing plus
  an update is more useful than a silently edited issue.
- **Docs ship with the work.** Any doc whose claims a PR changes is updated in that PR,
  not a follow-up. Verify doc claims against ground truth (code, data, computed output) —
  never against another doc, since both can be stale.

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

**The two cadences are different, and deliberately so.**

- **Keep GitHub aligned at all times.** The tracker is the live state: open, update and
  close issues as the work happens, not batched at the end of a session. A decision that
  changes an issue, a milestone or the order is reflected the moment it is made.
- **Docs are aligned at PR time** — in the PR that changes what they claim, never
  continuously and never in a follow-up. Verify their claims against ground truth (code,
  data, computed output), never against another doc, since both can be stale. The one
  exception is a build phase long enough that the docs would be wrong for days, where an
  interim docs commit earns its keep.

**Where the repo uses milestones, the milestone title carries the order** — a two-digit
sort prefix and a middot: `01 · Eval Harness`. Zero-pad, or `10` sorts before `2`. GitHub
can only sort milestones by title, due date or completion, and we set no due dates, so
the title is the only place ordering can live without inventing dates that go stale.
Standing buckets outside the build sequence take a `~ ` prefix and no number
(`~ Public Readiness`); they sort last and claim no place in the sequence. Reordering
means renumbering the affected titles in the same turn.

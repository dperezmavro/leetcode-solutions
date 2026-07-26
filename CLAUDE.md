# CLAUDE.md

## Repo layout
- Solutions live under `easy/`, `medium/`, `hard/`.
- Go solutions: each problem has a `go/` subdirectory (or sits at the problem root for older problems).
- importpath convention: `github.com/dperezmavro/leetcode/<difficulty>/<problem-slug>[/go]`

## Build system
- Bazel with Bzlmod. Single Go module at the root.
- After adding new Go files: run `bazel run //:gazelle` to regenerate BUILD files.
- After changing `go.mod`: run `bazel mod deps && bazel mod tidy`.
- Run all tests: `make test-go`
- Run one test: `make test-problem TARGET=//path/to:target_test`

## Local vs remote (BuildBuddy)
- Builds run **locally on the host toolchain by default**. Remote execution (RE)
  and remote cache (CAS) are **opt-in** via Bazel configs — nothing hits the
  network unless you ask.
- `--config=remote`: RE + CAS on BuildBuddy Linux/amd64 workers.
- `--config=remote-cache`: local execution, shared BuildBuddy cache only.
- Via Make (default = local; flags apply to `test`, `test-go/rust/python`,
  `lint-*`, `fmt-rust-check`, `test-problem`):
  - `make test` — local
  - `make test REMOTE=1` — remote execution + cache (same as `make test-remote`)
  - `make test CACHE=1` — local execution + remote cache
  - `make build-remote` — build everything remotely
- RE prerequisites (already wired): a linux/amd64 Go SDK in `MODULE.bazel`, the
  `toolchains_buildbuddy` module (Linux platform + C++ toolchain), and a
  BuildBuddy API key as `--remote_header=x-buildbuddy-api-key=...` in
  `~/.bazelrc`. RE pins platforms to `@toolchains_buildbuddy//platforms:linux_x86_64`
  (NOT `@buildbuddy_toolchain//:platform`, whose alias resolves to the darwin host).

## External Go deps
- Declared in `/go.mod`. Bazel resolves them via the `go_deps` extension in MODULE.bazel.
- Do NOT create per-problem go.mod files for new problems. The existing ones in 703 and 735 are legacy.

## Naming conventions
- Library target name = problem slug (e.g. `703-kth-largest-element-in-a-stream`)
- Test target name = problem slug + `_test`
- embed in go_test must match the library target name exactly.

# Agent instructions
Adhere to the following rules when responding:
- avoid prose
- use technically correct language
- use as few words as possible
- do not be verbose

Adhere to the following rules when generating code:
- avoid very big changes
- avoid anything that relies on string or naming manipulation - instead propose an architectural review and highlight the cause of the issue
- do not commit to git, this is a user-level concern
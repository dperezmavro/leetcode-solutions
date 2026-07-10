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
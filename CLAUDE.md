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
Scaffold a new Go LeetCode problem in this monorepo.

Arguments: `$ARGUMENTS` — the full problem path, e.g. `easy/1234-two-sum` or `medium/456-my-problem`.

## Steps

1. Parse the argument to determine:
   - `difficulty`: the first path segment (`easy`, `medium`, or `hard`)
   - `slug`: the second path segment (e.g. `1234-two-sum`)
   - `full_path`: `<difficulty>/<slug>/go`
   - `importpath`: `github.com/dperezmavro/leetcode/<difficulty>/<slug>/go`
   - `pkg`: a valid Go package name derived from the slug — strip the leading number and dashes, convert to snake_case. For example `1234-two-sum` → `two_sum`.

2. Create the directory `<difficulty>/<slug>/go/`.

3. Write `<difficulty>/<slug>/go/solution.go`:
```go
package <pkg>
```

4. Write `<difficulty>/<slug>/go/solution_test.go`:
```go
package <pkg>

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.True(t, true) // replace with real tests
}
```

5. Write `<difficulty>/<slug>/go/BUILD.bazel`:
```starlark
load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "<slug>",
    srcs = ["solution.go"],
    importpath = "<importpath>",
    visibility = ["//visibility:public"],
)

go_test(
    name = "<slug>_test",
    srcs = ["solution_test.go"],
    embed = [":<slug>"],
    size = "small",
    deps = ["@com_github_stretchr_testify//assert:go_default_library"],
)
```

6. Run `bazel run //:gazelle` to let Gazelle verify and potentially adjust the BUILD file.

7. Run `make test-problem TARGET=//<difficulty>/<slug>/go:<slug>_test` to confirm the scaffold builds and the stub test passes.

8. Report the target path to the user so they can start solving the problem.

## Notes
- Do NOT create a `go.mod` or `go.sum` inside the new directory. All external deps are managed at the repo root via `go.mod` + the `go_deps` extension in `MODULE.bazel`.
- If the user has not specified a difficulty prefix (i.e. only a slug was given), ask which difficulty directory to use before proceeding.
- The `go/` subdirectory is intentional — it reserves space for `py/`, `rs/` siblings later.

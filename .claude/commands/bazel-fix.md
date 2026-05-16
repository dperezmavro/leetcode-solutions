Diagnose and fix a broken Bazel build or test in this repo.

Arguments (optional): `$ARGUMENTS` — a specific Bazel target to investigate, e.g. `//easy/703-kth-largest-element-in-a-stream:703-kth-largest-element-in-a-stream_test`. If omitted, run `make test-go` and fix whatever is failing.

## Diagnosis steps

Run the failing target (or `make test-go` if no target given) and capture output. Then work through these failure categories in order:

### 1. "no such target" / "no such package"
- The directory is probably in `.bazelignore`. Remove the entry and re-run.
- Or the BUILD.bazel target name doesn't match what's being referenced. Read the BUILD.bazel file and fix the label.

### 2. External dependency not found (`@com_github_...` or similar)
- The `go_deps` extension in `MODULE.bazel` is missing the dependency.
- Check if the dep is in `/go.mod`. If not, add it and run `go mod tidy` (or add manually with correct checksums to `go.sum`).
- Run `bazel mod deps` to refresh `MODULE.bazel.lock`.
- Run `bazel mod tidy` — it will auto-correct the `use_repo()` call in `MODULE.bazel`.
- Re-run the target.

### 3. `embed` label mismatch
- The `go_test` `embed` field references a target that doesn't exist. Read the BUILD.bazel and align the embed label with the `go_library` name.

### 4. Import cycle or missing dep in `deps`
- The compiler error will name the missing import. Find the Bazel label for that package (`bazel query 'kind("go_library", //...)'` or look for `@com_github_...` labels).
- Add it to the `deps` list in the BUILD.bazel.

### 5. Test logic failure (not a build failure)
- Read the test log at the path shown in the Bazel output.
- Read the source file and test file side by side.
- Look for common Go gotchas: `for i := range slice` gives indices, not values; nil pointer on uninitialized struct fields; off-by-one in heap/window logic.
- Fix the test or the solution as appropriate.

### 6. MODULE.bazel.lock out of date
- If Bazel complains about a lock file mismatch, run `bazel mod deps` to regenerate.

## After each fix
Re-run the target (or `make test-go`) immediately to confirm the fix worked before moving to the next issue.

## Done condition
Report success only when `make test-go` exits 0 with all targets passing.

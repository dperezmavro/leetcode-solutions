.PHONY: test test-go test-rust test-python test-remote build-remote test-problem setup tidy lint-rust fmt-rust-check fmt-rust lint-go fmt-go-check fmt-go list-tests

# Remote execution + cache is OPT-IN (see .bazelrc). Everything runs locally on
# the host toolchain by default. Toggle per invocation with a variable:
#   make test REMOTE=1   # BuildBuddy remote execution + cache (Linux workers)
#   make test CACHE=1    # local execution, shared BuildBuddy cache only (CAS)
# REMOTE takes precedence if both are set.
BAZEL_CONFIG := $(if $(REMOTE),--config=remote,$(if $(CACHE),--config=remote-cache))

# Run every Go test in the repo.
# Bazel discovers all go_test targets under all packages via //...
# The python package in 1115 is excluded via .bazelignore (it has its own workspace).
test:
	bazel test $(BAZEL_CONFIG) //...

# Run the full suite on BuildBuddy remote execution (equivalent to REMOTE=1).
test-remote:
	bazel test --config=remote //...

# Build everything on BuildBuddy remote execution.
build-remote:
	bazel build --config=remote //...

# Run only Go tests (useful if non-Go rules are added later).
# The kind() filter selects targets whose rule type matches 'go_test'.
test-go:
	bazel test $(BAZEL_CONFIG) --build_tag_filters='' \
		$(shell bazel query 'kind("go_test", //...)')

test-rust:
	bazel test $(BAZEL_CONFIG) --build_tag_filters='' \
		$(shell bazel query 'kind("rust_test", //...)')

test-python:
	bazel test $(BAZEL_CONFIG) --build_tag_filters='' \
		$(shell bazel query 'kind("py_test", //...)')

# Run clippy over every rust target via the bazel aspect.
# Fails the build on any clippy finding (see .bazelrc config:clippy).
lint-rust:
	bazel build $(BAZEL_CONFIG) --config=clippy //...

# Verify every rust source is rustfmt-clean (non-mutating).
fmt-rust-check:
	bazel build $(BAZEL_CONFIG) --config=rustfmt //...

# Rewrite every rust source in place with rustfmt.
fmt-rust:
	bazel run @rules_rust//:rustfmt

# Run `go vet` analyzers over every go target via nogo. nogo runs during
# compilation, so building the go_library targets is what triggers it; the
# build fails on any vet finding.
lint-go:
	bazel build $(BAZEL_CONFIG) --build_tag_filters='' \
		$(shell bazel query 'kind("go_library", //...)')

# Verify every go source is gofmt-clean (non-mutating). Uses the hermetic
# gofmt from the bazel-managed Go SDK; go.work workspace mode breaks
# `go fmt ./...`, so gofmt runs over the source tree directly.
fmt-go-check:
	@out=$$("$$(bazel run -- @rules_go//go env GOROOT 2>/dev/null)/bin/gofmt" -l .); \
	if [ -n "$$out" ]; then echo "gofmt needed on:"; echo "$$out"; exit 1; fi; \
	echo "gofmt: all go sources formatted"

# Rewrite every go source in place with the hermetic gofmt.
fmt-go:
	"$$(bazel run -- @rules_go//go env GOROOT 2>/dev/null)/bin/gofmt" -w .

# Run the tests for a single problem.
# Usage: make test-problem TARGET=//medium/438-find-all-anagrams-in-a-string:find_all_anagrams_test
TARGET ?= //medium/438-find-all-anagrams-in-a-string:find_all_anagrams_test
test-problem:
	bazel test $(BAZEL_CONFIG) $(TARGET)

# Regenerate MODULE.bazel.lock after changing MODULE.bazel dependencies.
# Run this whenever you add a new bazel_dep or change a dep version.
setup:
	bazel mod deps
	make tidy

tidy:
	bazel mod tidy
	bazel run //:gazelle

# Show all test targets without running them.
list-tests:
	bazel query 'kind("go_test|rust_test", //...)'
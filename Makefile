.PHONY: test test-go test-problem setup

# Run every Go test in the repo.
# Bazel discovers all go_test targets under all packages via //...
# The python package in 1115 is excluded via .bazelignore (it has its own workspace).
test:
	bazel test //...

# Run only Go tests (useful if non-Go rules are added later).
# The kind() filter selects targets whose rule type matches 'go_test'.
test-go:
	bazel test --build_tag_filters='' \
		$(shell bazel query 'kind("go_test", //...)')

# Run the tests for a single problem.
# Usage: make test-problem TARGET=//medium/438-find-all-anagrams-in-a-string:find_all_anagrams_test
TARGET ?= //medium/438-find-all-anagrams-in-a-string:find_all_anagrams_test
test-problem:
	bazel test $(TARGET)

# Regenerate MODULE.bazel.lock after changing MODULE.bazel dependencies.
# Run this whenever you add a new bazel_dep or change a dep version.
setup:
	bazel mod deps

# golden
Opinionated golden file testing.

## Getting Started

Run `go get github.com/elielamora/golden`

add `github.com/elielamora/golden` to your test

Then just assert the value with `golden.Assert(t, "golden")`.
It will pick up the file name based on the test name
and represent nested tests as folder dependencies.

Run with the env var `UPDATE_GOLDEN` to update the golden files. Test suite will pass and you will
see the new test cases as snapshots in your git diff.

## TODO:
- Automate removing orphaned golden files when tests are renamed (for now just clean up manually).
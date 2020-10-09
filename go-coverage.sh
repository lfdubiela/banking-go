#!/bin/bash
# Generate test coverage statistics for Go packages.
#
# Works around the fact that `go test -coverprofile` currently does not work
# with multiple packages, see https://code.google.com/p/go/issues/detail?id=6909

set -e

workdir=.cover
profile="$workdir/cover.out"
mode=count

generate_cover_data() {
    rm -rf "$workdir"
    mkdir "$workdir"

    for pkg in "$@"; do
        f="$workdir/$(echo $pkg | tr / -).cover"
        go test -covermode="$mode" -coverprofile="$f" "$pkg"
    done

    echo "mode: $mode" >"$profile"
    grep -h -v "^mode:" "$workdir"/*.cover >>"$profile"

    generate_report
}

generate_report() {
    report_dir=$(pwd)/report

    mkdir -p $report_dir
    chmod +x $report_dir

    go tool cover -html="$profile" -o $report_dir/coverage.html
    coverage_percent=$(go tool cover -func="$profile" | grep 'total:' | grep -oe '[0-9.].*')

    echo "Cobertura de testes:"  $coverage_percent

    echo ${coverage_percent//%/""} >> $report_dir/coverage.txt
    chmod +x $report_dir/coverage.txt
}

generate_cover_data $(go list ./...)


#!/bin/bash
TEST_COVERAGE_THRESHOLD=75
echo "Threshold: $TEST_COVERAGE_THRESHOLD%"
make test
test_coverage=`go tool cover -func cover.out | grep total | grep -Eo '[0-9]+\.[0.9]+'`
echo "Current test coverage: $test_coverage%"
if (( $(echo "$test_coverage $TEST_COVERAGE_THRESHOLD" | awk '{print ($1  > $2)}') )); then
  echo "Ok"
else
  echo "Current test coverage is below the threshold. Please add more unit tests or adjust threshold to a lower value"
  echo "Failed"
  exit 1
fi

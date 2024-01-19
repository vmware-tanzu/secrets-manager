# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware, Inc.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */
#
# This file contains the make targets for testing the VSecM project.
#
# Usage:
#   # Run all tests and check coverage against the threshold.
#   make cover
coverage_file := coverage.out
threshold = 70

cover:
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=$(coverage_file) ./... > /dev/null
	@echo "Checking test coverage..."
	@coverage=$$(go tool cover -func=$(coverage_file) | grep total | grep -Eo '[0-9]+\.[0-9]+' || echo "0.0"); \
	echo "Test coverage: $$coverage"; \
	echo "Test Threshold: $(threshold)"; \
    if [ "$$coverage" != "" ] && awk 'BEGIN{exit !('"$$coverage"'<=$(threshold))}'; then \
    	echo "Test coverage is less than $(threshold)"; \
		exit 0; \
	fi
	@echo "Test coverage is greater than $(threshold)"
	@rm -f $(coverage_file)



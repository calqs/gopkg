MODULE_DIRS := . ./dt ./router ./crypt ./env
GO_TEST_CMD=go test -count=1 -coverprofile=coverage.out ./...	
GO_TIDY_CMD=go mod tidy
GO_HTML_COVERAGE=go tool cover -html=coverage.out

tidy_all:
	@for dir in $(MODULE_DIRS); do \
		echo "Tidying $$dir"; \
		(cd $$dir && ${GO_TIDY_CMD}); \
	done

test_coverage:
	cd $$dir && ${GO_TEST_CMD} && ${GO_HTML_COVERAGE}

test_all:
	@for dir in $(MODULE_DIRS); do \
		echo "Testing $$dir"; \
		(cd $$dir && ${GO_TEST_CMD}); \
	done

coverage_all:
	@for dir in $(MODULE_DIRS); do \
		echo "Testing $$dir"; \
		(cd $$dir && ${GO_HTML_COVERAGE}); \
	done

test_coverage_all: test_all coverage_all

update_go_version_all:
	./_scripts/update-go-mods.sh
	$(MAKE) tidy_all
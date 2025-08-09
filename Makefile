MODULE_DIRS := . ./dt ./httpserver ./crypt
GO_TEST_CMD=go test -count=1 -v ./...	
GO_TIDY_CMD=go mod tidy

tidy_all:
	@for dir in $(MODULE_DIRS); do \
		echo "Tidying $$dir"; \
		(cd $$dir && ${GO_TIDY_CMD}); \
	done


test_all:
	@for dir in $(MODULE_DIRS); do \
		echo "Testing $$dir"; \
		(cd $$dir && ${GO_TEST_CMD}); \
	done

update_go_version_all:
	./_scripts/update-go-mods.sh
	$(MAKE) tidy_all
COVERAGE_DIR ?= .coverage

# cp from: https://github.com/yyle88/formatgo/blob/875f8d978776d192da40c60f67c33830b19a4dc9/Makefile#L4
test:
	@-rm -r $(COVERAGE_DIR)
	@mkdir $(COVERAGE_DIR)
	make test-with-flags TEST_FLAGS='-v -race -covermode atomic -coverprofile $$(COVERAGE_DIR)/combined.txt -bench=. -benchmem -timeout 20m'

test-with-flags:
	@go test $(TEST_FLAGS) ./...

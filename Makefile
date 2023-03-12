.PHONY: dev
dev:
	@cargo tauri dev

.PHONY: build
build:
	@cargo tauri build

.PHONY: clean
clean:
	@rm -rf src-tauri/target

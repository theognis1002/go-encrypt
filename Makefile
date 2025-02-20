.PHONY: encrypt decrypt clean test all

KEY = 1234567890123456
INPUT_FILE = example.csv
OUTPUT_DIR = output

encrypt:
	@echo "Encrypting the file..."
	@mkdir -p $(OUTPUT_DIR)
	go run main.go -key "$(KEY)" -input "$(INPUT_FILE)" -encrypt "encrypted.csv"

decrypt:
	@echo "Decrypting the file..."
	@mkdir -p $(OUTPUT_DIR)
	go run main.go -key "$(KEY)" -input "$(OUTPUT_DIR)/encrypted.csv" -decrypt "decrypted.csv"

clean:
	@find $(OUTPUT_DIR)/ -type f ! -name '.gitkeep' -delete
	@echo "Squeaky clean!"

test:
	@echo "Running tests..."
	go test -v ./...

test-coverage:
	@echo "Running tests with coverage..."
	@mkdir -p $(OUTPUT_DIR)
	go test -v -coverprofile=$(OUTPUT_DIR)/coverage.out ./...
	go tool cover -html=$(OUTPUT_DIR)/coverage.out -o $(OUTPUT_DIR)/coverage.html

all: test

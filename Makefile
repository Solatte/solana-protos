# Define the output directory for generated Go files
OUTPUT_DIR := . 

# Find all proto files in subdirectories
PROTO_FILES := $(shell find . -type f -name '*.proto')

# Define the rule to generate .pb.go files from .proto files
gen: $(PROTO_FILES)
	@mkdir -p $(OUTPUT_DIR)
	@for proto in $^; do \
		dir=$$(dirname $$proto); \
		protoc --proto_path=$$dir --go_out=$(OUTPUT_DIR) --go-grpc_out=$(OUTPUT_DIR) $$proto --experimental_allow_proto3_optional; \
	done

# Clean up generated Go files
clean:
	rm -f $(OUTPUT_DIR)/*.pb.go

# Default target
.PHONY: gen clean


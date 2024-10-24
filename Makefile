# Define the output directory for generated Go files
OUTPUT_DIR := . 

# Find all proto files in subdirectories
PROTO_FILES := $(shell find . -type f -name '*.proto')

# Define the rule to generate .pb.go files from .proto files
gen: $(PROTO_FILES)
	@mkdir -p $(OUTPUT_DIR)
	@for proto in $^; do \
		dir=$$(dirname $$proto); \
		protoc --proto_path=$$dir --go_out=$(OUTPUT_DIR) $$proto --experimental_allow_proto3_optional; \
	done

# Clean up generated Go files
clean:
	rm -f $(OUTPUT_DIR)/*.pb.go

# Default target
.PHONY: gen clean


# # Define the base directory
# BASE_DIR := .

# # Find all .proto files in all subdirectories
# PROTO_FILES := $(shell find $(BASE_DIR) -type f -path '*/protos/*.proto')

# # Define the corresponding .pb.go files for each .proto file
# PB_GO_FILES := $(PROTO_FILES:%.proto=%.pb.go)
# PB_GO_FILES := $(subst protos,pb, $(PB_GO_FILES))

# # Specify the Go compiler and options
# GO_OUT_FLAGS = --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --proto_path=protos

# # Default target
# .PHONY: all
# all: $(PB_GO_FILES)

# # Pattern rule to generate .pb.go files from .proto files
# %.pb.go: %.proto
# 	@echo $(PB_GO_FILES)
# 	@echo "Compiling $< to $@"
# 	@mkdir -p $(dir $@)  # Create the output directory if it doesn't exist
# 	protoc $(GO_OUT_FLAGS) $< --experimental_allow_proto3_optional

# # Clean target to remove generated files
# .PHONY: clean
# clean:
# 	rm -rf pb/*.pb.go
# 	rm -rf */pb/*.pb.go

# # Add a target to check for changes in proto files and recompile if necessary
# .PHONY: proto
# proto: all

# .PHONY: debug
# debug:
# 	@echo "Proto files: $(PROTO_FILES)"
# 	@echo "PB Go files: $(PB_GO_FILES)"
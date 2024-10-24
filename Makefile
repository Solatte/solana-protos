
# Find all proto files in subdirectories
PROTO_FILES := $(shell find . -type f -name '*.proto')

# Define the rule to generate .pb.go files from .proto files
gen: $(PROTO_FILES)
	@for proto in $^; do \
		dir=$$(dirname $$proto); \
		OUTPUT_DIR=$$(echo $$dir | sed 's/protos/pb/'); \
 		protoc --proto_path=$$dir --go_out=$$OUTPUT_DIR --go_opt=paths=source_relative --go-grpc_out=$$OUTPUT_DIR --go-grpc_opt=paths=source_relative $$proto --experimental_allow_proto3_optional; \
	done

# Clean up generated Go files
clean:
	rm -f */pb/*.pb.go

# Default target
.PHONY: gen clean


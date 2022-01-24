default: build

ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

clean:
	@bazelisk clean || { echo "Clean failed, check above for errors!"; exit 1; }
	rm -rf $(ROOT_DIR)/work

build:
	@bazelisk run //:gazelle -- update-repos -from_file=go.mod -to_macro=BUILD.golang.bzl%go_dependencies -prune || { echo "Unable to run dependency update, check above for errors!"; exit 1; }
	@bazelisk run //:gazelle || { echo "Unable to run gazelle, check above for errors!"; exit 1; }
	@bazelisk build //... || { echo "Build is failing, check above for errors!"; exit 1; }
	# @go fmt $(go list ./...)|| { echo "Unable format go files, check above for errors!"; exit 1; }
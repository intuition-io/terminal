# Makefile
# Hackliff, 2014-08-19 23:04
# vim:ft=make

all:
	godep go build

# start documentation server on localhost:6060
doc:
	@echo "Connect at localhost:6060 to see package documentation"
	@sh -c "godoc -http=:6060"

bin:
	@sh -c "$(CURDIR)/scripts/build.sh"

dev:
	@TF_DEV=1 sh -c "$(CURDIR)/scripts/build.sh"

dist:
	@sh -c "$(CURDIR)/scripts/dist.sh $(VERSION)"

install:
	go get -u github.com/bndr/gopencils
	go get -u github.com/codegangsta/cli
	go install

devinstall:
	go get -u github.com/tools/godep
	go get -u github.com/mitchellh/gox
	go get -u github.com/axw/gocov/gocov
	go get -u  github.com/mattn/goveralls
	godep save

test:
	TF_ACC= godep go test $(TEST) $(TESTARGS) -timeout=10s

testrace:
	TF_ACC= godep go test -race $(TEST) $(TESTARGS)

coverage:
	# TODO Check for $COVERALLS_TOKEN
	goveralls -v -service drone.io -repotoken $(COVERALLS_TOKEN)

clean:
	@sh -c "rm bin/*"

.PHONY: default updatedeps dist bin clean test

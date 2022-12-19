GO_CMD=go
GO_TEST=$(GO_CMD) test -count=1 -v -cover


.PHONY:test
test:
	@-$(GO_TEST) ./... ||:
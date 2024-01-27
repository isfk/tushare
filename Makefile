xlsx:
	cd tool; go run main.go
	buf format proto -w

buf:
	buf generate

all:
	go build


test:
	go build
	@mkdir -p out ref
	./datamatrix-decode testdata/test1.png >out/test1.out
	@diff out/test1.out ref
	@echo PASS


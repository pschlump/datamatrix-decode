
all:
	go build


test: test1 test2 

test1:
	go build
	@mkdir -p out ref
	./datamatrix-decode testdata/test1.png >out/test1.out
	@diff out/test1.out ref
	@echo PASS

test2:
	go build
	@mkdir -p out ref
	./datamatrix-decode testdata/test2.png >out/test2.out
	@diff out/test2.out ref
	@echo PASS


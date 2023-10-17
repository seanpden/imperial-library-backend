compile:
	go build -C cmd/ -o ../build/main

run: compile 
	./build/main

clean:
	rm build/main

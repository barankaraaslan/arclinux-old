build:
	mkdir -p target
	docker run --rm -it -v $$PWD/../:/app -w /app --platform linux/amd64 gcc gcc -I lib1/ -L lib1/target -l1 bin1/main.c -o bin1/target/bin1

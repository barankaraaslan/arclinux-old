build:
	mkdir -p target
	docker run --rm -it -v $$PWD:/app -w /app --platform linux/amd64 gcc gcc hw.c -shared -o target/lib1.so

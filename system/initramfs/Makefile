BUILD := ../../build
ROOT := $(BUILD)/initramfs-root
OUTPUT := $(BUILD)/initramfs.cpio.gz

.PHONY: all clean

all:
	rm -rf $(ROOT)
	mkdir -p $(ROOT)
	cp $(BUILD)/kittyinit $(ROOT)/init
	chmod +x $(ROOT)/init
	cd $(ROOT) && find . -print0 \
		| cpio --null -o --format=newc \
		| gzip -9 > ../initramfs.cpio.gz

clean:
	rm -rf $(ROOT) $(OUTPUT)

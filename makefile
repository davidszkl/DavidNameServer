PROJECT = DavidNameServer

all:
	make -C src PROJECT=$(PROJECT)

clean:
	rm -f $(PROJECT)
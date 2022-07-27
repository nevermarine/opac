opac: opac.go 
	go build -v -ldflags "-s" opac.go
clean: 
	go clean 
install: opac
	install -d $(DESTDIR)/usr/bin
	install -t $(DESTDIR)/usr/bin ./opac
uninstall: opac
	rm -r $(DESTDIR)/usr/bin/opac

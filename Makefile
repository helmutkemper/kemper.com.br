GO ?= go

.PHONY: plugin-mongodb-language
##
plugin-mongodb-language:
	$(MAKE) -C ./plugin/kemper.com.br.plugin.mongodb.languages

.PHONY: plugin-mongodb-menu
##
plugin-mongodb-menu:
	$(MAKE) -C ./plugin/kemper.com.br.plugin.mongodb.menu

.PHONY: plugin-mongodb-user
##
plugin-mongodb-user:
	$(MAKE) -C ./plugin/kemper.com.br.plugin.mongodb.user

.PHONY: plugin-sqlite-languages
##
plugin-sqlite-languages:
	$(MAKE) -C ./plugin/kemper.com.br.plugin.sqlite.languages

.PHONY: plugin-sqlite-menu
##
plugin-sqlite-menu:
	$(MAKE) -C ./plugin/kemper.com.br.plugin.sqlite.menu

.PHONY: plugin-sqlite-user
##
plugin-sqlite-user:
	$(MAKE) -C ./plugin/kemper.com.br.plugin.sqlite.user

.PHONY: plugins
##
plugins: plugins-clean plugin-mongodb-language plugin-mongodb-menu plugin-mongodb-user plugin-sqlite-languages plugin-sqlite-menu plugin-sqlite-user

.PHONY: clean
##
clean: build-clean plugins-clean

.PHONY: build-clean
##
build-clean:
	@if [ -f ./build/site.so ]; then rm ./build/site.so; fi

.PHONY: plugins-clean
##
plugins-clean:
	@if [ -f ./cmd/plugin/languages.mongodb.so ]; then rm ./cmd/plugin/languages.mongodb.so; fi
	@if [ -f ./cmd/plugin/languages.sqlite.so ]; then rm ./cmd/plugin/languages.sqlite.so; fi
	@if [ -f ./cmd/plugin/menu.sqlite.so ]; then rm ./cmd/plugin/menu.sqlite.so; fi
	@if [ -f ./cmd/plugin/menu.sqmongodb.so ]; then rm ./cmd/plugin/menu.sqmongodb.so; fi
	@if [ -f ./cmd/plugin/user.mongodb.so ]; then rm ./cmd/plugin/user.mongodb.so; fi
	@if [ -f ./cmd/plugin/user.sqlite.so ]; then rm ./cmd/plugin/user.sqlite.so; fi

.PHONY: build
##
build:
	$(MAKE) build-site
	$(MAKE) plugins
	$(MAKE) open-site

.PHONY: open-site
##
open-site:
	@if [ -f ./build/site.so ]; then echo site found; else $(MAKE) build-site; $(MAKE) plugins; fi
	@/usr/bin/open -a "/Applications/Google Chrome.app" 'http://localhost:3000/static/'
	./build/site.so

.PHONY: build-site
##
build-site:
	$(MAKE) -C ./cmd tidy
	$(MAKE) -C ./cmd build
	@chmod +X ./build/site.so

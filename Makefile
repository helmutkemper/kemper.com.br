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
	@rm ./build/site.so

.PHONY: plugins-clean
##
plugins-clean:
	@rm ./main.site/plugin/languages.mongodb.so
	@rm ./main.site/plugin/languages.sqlite.so
	@rm ./main.site/plugin/menu.sqlite.so
	@rm ./main.site/plugin/menu.sqmongodb.so
	@rm ./main.site/plugin/user.mongodb.so
	@rm ./main.site/plugin/user.sqlite.so

.PHONY: build
##
build:
	$(MAKE) -C ./main.site tidy
	$(MAKE) -C ./main.site build
	$(MAKE) plugins
	@chmod +X ./build/site.so
	@/usr/bin/open -a "/Applications/Google Chrome.app" 'http://localhost:3000/static/'
	./build/site.so
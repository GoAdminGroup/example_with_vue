BACKEND=./backend

all: serve

serve:
	NODE_ENV=production npm --prefix ./src/adminlte run build
	NODE_ENV=production $(BACKEND)

serve-quiet:
	NODE_ENV=production npm --prefix ./src/adminlte run build
	NODE_ENV=production $(BACKEND) --debug=true --quiet=true

serve-element:
	NODE_ENV=production npm --prefix ./src/element run build
	NODE_ENV=production $(BACKEND) --debug=true --theme=sword --ui_theme=element

build:
	npm --prefix=./src run build

help:
	$(BACKEND) --help

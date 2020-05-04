BACKEND=./backend

all: serve

serve:
	$(BACKEND)

serve-quiet:
	$(BACKEND) --debug=true --quiet=true

serve-element:
	$(BACKEND) --debug=true --theme=sword --ui_theme=element

build:
	npm --prefix=./src run build

help:
	$(BACKEND) --help

BACKEND=./backend

all: serve

serve:
	$(BACKEND) --debug=true

serve-quiet:
	$(BACKEND) --debug=true --quiet=true

build:
	npm --prefix=./src run build

help:
	$(BACKEND) --help

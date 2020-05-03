BACKEND=./backend

all: serve

serve:
	$(BACKEND)

serve-quiet:
	$(BACKEND) --debug=true --quiet=true

build:
	npm --prefix=./src run build

help:
	$(BACKEND) --help

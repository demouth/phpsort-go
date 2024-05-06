.PHONY: run-php8 run-php7 help

run-php8:
	docker run -it --rm -v $(PWD):/usr/src/ -w /usr/src/ php:8.3.2-cli php $(file)

run-php7:
	docker run -it --rm -v $(PWD):/usr/src/ -w /usr/src/ php:7.4.33-cli php $(file)

help:
	@echo "Usage: make run-php8 file=<file.php>"
	@echo "Example: make run-php8 file=test1.php"

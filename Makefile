run:
	if [ ! -e .env ]; \
	then \
		cp .env_default .env; \
	fi; \

	docker-compose up $(args)
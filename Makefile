.PHONY: setup run down logs rebuild

setup:
	docker compose up -d
	docker exec -it ollama ollama pull phi3:mini

run:
	docker compose up --build

down:
	docker compose down

logs:
	docker compose logs -f

rebuild:
	docker compose build --no-cache
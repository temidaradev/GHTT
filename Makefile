run: build
	@./bin/app

build:
	@go build -o bin/app .


css:
	npx tailwindcss -i views/css/app.css -o public/styles.css --watch   
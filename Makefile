.PHONY: build templ tailwind test clean

build: templ tailwind
	@go build -o main cmd/api/main.go

templ: $(wildcard cmd/web/*.templ)
	@templ generate

tailwind: cmd/web/assets/css/input.css
	@tailwindcss \
		-i cmd/web/assets/css/input.css \
		-o cmd/web/assets/css/output.css

test:
	@go test ./... -v

clean:
	@rm cmd/web/assets/css/output.css
	@find . -type f -name "*_templ.go" -delete
	@rm -f main

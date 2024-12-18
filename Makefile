TW_SOURCE=./src/tw/input.css
TW_TARGET=./static/css/main.css

tw-watch:
	@pnpm exec tailwindcss -i $(TW_SOURCE) -o $(TW_TARGET) --watch

dev:
	@~/go/bin/air


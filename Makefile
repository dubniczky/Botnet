.PHONY: commander
# Start the commander server
commander::
	@echo "Starting commander..."
	@cd commander \
		&& node main.js
	
.PHONY: bot
# Start the bot
bot::
	@echo "Starting bot..."
	@cd bot \
		&& go run main.go

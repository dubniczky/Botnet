.PHONY: commander
# Start the commander server
commander::
	@echo "Starting commander server..."
	@cd commander \
		&& node main.js
	
.PHONY: bot
# Start the bot
bot::
	@echo "Starting commander server..."
	@cd bot \
		&& go run main.go
.DEFAULT_GOAL = help


#.PHONY: ent-init
#ent-init: ## inits schemas for Product and Basket.
#	go run entgo.io/ent/cmd/ent init --target internal/ent/schema

.PHONY: ent-generate
ent-generate: ## Generates full, graphs based, entity for the schemas.
	go run -mod=mod entgo.io/ent/cmd/ent generate ./internal/data/schema

.PHONY: help
help: ## Prints this message.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	sort | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

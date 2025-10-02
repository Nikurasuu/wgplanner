package server

import (
	"wgplanner/internal/handler"

	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
)

func addGroupRoutes(api *fuego.Server, groupHandler *handler.GroupHandler) {
	groups := fuego.Group(api, "/groups",
		option.Tags("Groups"),
	)

	fuego.Post(groups, "", groupHandler.CreateGroup,
		option.Tags("Create"),
		option.Summary("Create a new group"),
	)
	fuego.Get(groups, "/{id...}", groupHandler.GetGroupFromID,
		option.Tags("Get"),
		option.Summary("Get a group by ID"),
	)
}

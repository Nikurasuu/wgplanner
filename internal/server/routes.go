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
		option.Summary("Create a new group"),
	)

	group := fuego.Group(groups, "/{groupId}")

	fuego.Get(group, "", groupHandler.GetGroupFromID,
		option.Summary("Get a group by ID"),
	)
	fuego.Post(group, "/members", groupHandler.AddMemberToGroup,
		option.Summary("Add a new member to the group"),
	)
}

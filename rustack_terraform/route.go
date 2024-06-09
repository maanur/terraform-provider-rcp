package rustack_terraform

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (args *Arguments) injectResultRoute() {
	args.merge(Arguments{
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Id of the route",
		},
		"destination": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Destination CIDR of the route",
		},
		"nexthop": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Next hop of the route",
		},
	})
}

func (args *Arguments) injectCreateRoute() {
	args.merge(Arguments{
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Id of the route",
		},
		"destination": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Destination CIDR of the route",
		},
		"nexthop": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Next hop of the route",
		},
	})
}

func (args *Arguments) injectCreateListRoutes() {
	Route := Defaults()
	Route.injectCreateRoute()

	args.merge(Arguments{
		"routes": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     &schema.Resource{Schema: Route},
		},
	})
}

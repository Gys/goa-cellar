package design

import (
	. "github.com/Gys/goa/design"
	. "github.com/Gys/goa/design/apidsl"
)

var _ = Resource("health", func() {

	BasePath("/_ah")

	Action("health", func() {
		Routing(
			GET("/health"),
		)
		Description("Perform health check.")
		Response(OK, "text/plain")
	})
})

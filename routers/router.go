package routers

type Routers interface {
	MapRoutes()
}

func MapRoutes(routers ...Routers) {
	for _, router := range routers {
		router.MapRoutes()
	}
}

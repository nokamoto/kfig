package main

func removedRoutes(present []Route, config []Route) []Route {
	res := make([]Route, 0)
	for _, route := range present {
		found := false

		for _, c := range config {
			if c.functionallyEqual(route) {
				found = true
				break
			}
		}

		if !found {
			res = append(res, route)
		}
	}
	return res
}

func addedRoutes(present []Route, config []Route) []Route {
	res := make([]Route, 0)
	for _, c := range config {
		found := false

		for _, route := range present {
			if route.functionallyEqual(c) {
				found = true
				break
			}
		}

		if !found {
			res = append(res, c)
		}
	}
	return res
}
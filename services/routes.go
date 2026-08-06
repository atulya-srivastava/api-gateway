package services

// var Routes = map[string] string{
// 	"/users":    "http://localhost:8081",
// 	"/products": "http://localhost:8082",
// 	"/orders":   "http://localhost:8083",
// }

var Routes = map[string][]string{

	"/users": {
		"http://localhost:8081",
		"http://localhost:8084",
		"http://localhost:8085",
	},

	"/products": {
		"http://localhost:8082",
	},

	"/orders": {
		"http://localhost:8083",
	},
}
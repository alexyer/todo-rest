package main

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"TodoIndex", "GET", "/todos", TodoIndex},
	Route{"TodoCreate", "POST", "/todos", TodoCreate},
	Route{"TodoShow", "GET", "/todos/{todoId}", TodoShow},
}

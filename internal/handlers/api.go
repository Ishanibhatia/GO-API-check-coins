package handlers

import(
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/Ishanibhatia/GO-API-check-coins/internal/middleware"
)


//Middleware is essantially a function that gets called before the primary function which handles the endpoint 

func Handler(r *chi.Mux){
	//Global middleware 
	//you can add middleware to a route by using use function 
	//global middleware means this middleware is going to apply to all the endpoints 
	r.Use(chimiddle.StripSlashes)
	//this strip slashes function will make sure that fucntion will make sure that the trailing slashes will always be ignored ..[1]-Notion


	//now let's setup our route 
	r.Route("/account", func(router chi.Router){
		//we can use the parametrized router to define our get method 
		//but first we're creating "middleware.Authorization" to check if the person is authorized enough to access the account or not
		router.Use(middleware.Authorization)
		//we'll create this authorization function in our middleware package later
		router.Get("/coins", GetCoinBalance)
		//we're defining get method with coins path which will be handled by GetCoinBalance function
	})
}
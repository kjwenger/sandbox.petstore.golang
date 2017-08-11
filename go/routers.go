package routers

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v2/",
		Index,
	},

	Route{
		"AddPet",
		"POST",
		"/v2/pets",
		AddPet,
	},

	Route{
		"DeletePet",
		"DELETE",
		"/v2/pets/{petId}",
		DeletePet,
	},

	Route{
		"FindPetsByStatus",
		"GET",
		"/v2/pet/findByStatus",
		FindPetsByStatus,
	},

	Route{
		"FindPetsByTags",
		"GET",
		"/v2/pets/findByTags",
		FindPetsByTags,
	},

	Route{
		"GetPetById",
		"GET",
		"/v2/pets/{petId}",
		GetPetById,
	},

	Route{
		"UpdatePet",
		"PUT",
		"/v2/pets",
		UpdatePet,
	},

	Route{
		"UpdatePetWithForm",
		"POST",
		"/v2/pets/{petId}",
		UpdatePetWithForm,
	},

	Route{
		"UploadFile",
		"POST",
		"/v2/pets/{petId}/uploadImage",
		UploadFile,
	},

	Route{
		"DeleteOrder",
		"DELETE",
		"/v2/stores/orders/{orderId}",
		DeleteOrder,
	},

	Route{
		"GetInventory",
		"GET",
		"/v2/stores/inventory",
		GetInventory,
	},

	Route{
		"GetOrderById",
		"GET",
		"/v2/stores/orders/{orderId}",
		GetOrderById,
	},

	Route{
		"PlaceOrder",
		"POST",
		"/v2/stores/orders",
		PlaceOrder,
	},

	Route{
		"CreateUser",
		"POST",
		"/v2/users",
		CreateUser,
	},

	Route{
		"CreateUsersWithArrayInput",
		"POST",
		"/v2/users/createWithArray",
		CreateUsersWithArrayInput,
	},

	Route{
		"CreateUsersWithListInput",
		"POST",
		"/v2/users/createWithList",
		CreateUsersWithListInput,
	},

	Route{
		"DeleteUser",
		"DELETE",
		"/v2/users/{username}",
		DeleteUser,
	},

	Route{
		"GetUserByName",
		"GET",
		"/v2/users/{username}",
		GetUserByName,
	},

	Route{
		"LoginUser",
		"GET",
		"/v2/users/login",
		LoginUser,
	},

	Route{
		"LogoutUser",
		"GET",
		"/v2/users/logout",
		LogoutUser,
	},

	Route{
		"UpdateUser",
		"PUT",
		"/v2/users/{username}",
		UpdateUser,
	},

}

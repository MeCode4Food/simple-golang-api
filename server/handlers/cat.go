package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-backend/server/models/cat"
	"simple-backend/server/services/db"
	"simple-backend/server/utils/handlerutil"

	"github.com/gorilla/mux"
)

// GetOneCatHandler export
func GetOneCatHandler(database *db.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var response interface{}
		fmt.Println(mux.Vars(r))
		name := mux.Vars(r)["name"]

		for _, cat := range database.CatStore {
			if cat.Name == name {
				response = cat
			}
		}
		if response == nil {
			json.Unmarshal([]byte(`{}`), &response)
		}
		handlerutil.RespondWithJSON(w, http.StatusOK, response)
	}
}

// GetAllCatHandler export
func GetAllCatHandler(database *db.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(database.CatStore)
		handlerutil.RespondWithJSON(w, http.StatusOK, database.CatStore)
	}
}

// PostCatHandler export
func PostCatHandler(database *db.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		catExists := false
		var catInput cat.Cat
		var response interface{}

		err := json.NewDecoder(r.Body).Decode(&catInput)
		if err != nil {
			fmt.Println(err.Error())
			json.Unmarshal([]byte(`{ "error": "invalid body" }`), &response)
			handlerutil.RespondWithJSON(w, http.StatusBadRequest, response)
		}

		// check if cat exists
		for _, cat := range database.CatStore {
			if cat.Name == catInput.Name {
				catExists = true
				break
			}
		}

		if catExists {
			json.Unmarshal([]byte(`{ "error": "cat name conflict" }`), &response)
			handlerutil.RespondWithJSON(w, http.StatusConflict, response)
		} else {
			// if does not exist, append cat to "database"
			database.CatStore = append(database.CatStore, catInput)
			handlerutil.RespondWithJSON(w, http.StatusCreated, catInput)
		}
	}
}

// PutCatHandler export
func PutCatHandler(database *db.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		catUpdated := false
		var catInput cat.Cat
		var response interface{}

		err := json.NewDecoder(r.Body).Decode(&catInput)
		if err != nil {
			fmt.Println(err.Error())
			json.Unmarshal([]byte(`{ "error": "invalid body" }`), &response)
			handlerutil.RespondWithJSON(w, http.StatusBadRequest, response)
		}

		// check if cat exists
		for i, cat := range database.CatStore {
			if cat.Name == catInput.Name {
				catUpdated = true
				database.CatStore[i] = catInput
			}
		}

		if catUpdated {
			handlerutil.RespondWithJSON(w, http.StatusOK, catInput)
		} else {
			// if does not exist, append cat to "database"
			database.CatStore = append(database.CatStore, catInput)
			handlerutil.RespondWithJSON(w, http.StatusCreated, catInput)
		}
	}
}

// DeleteCatHandler export
func DeleteCatHandler(database *db.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		nameToRemove := mux.Vars(r)["name"]

		found := false
		var indexToRemove int
		var removedCat cat.Cat

		// check if cat exists
		for i, cat := range database.CatStore {
			if cat.Name == nameToRemove {
				found = true
				indexToRemove = i
				removedCat = cat
			}
		}

		if found {
			// Removal without maintaining order
			database.CatStore[indexToRemove] = database.CatStore[len(database.CatStore)-1] // Copy last element to index i.
			database.CatStore[len(database.CatStore)-1] = cat.Cat{}                        // Erase last element (write zero value).
			database.CatStore = database.CatStore[:len(database.CatStore)-1]               // Truncate slice.
			handlerutil.RespondWithJSON(w, http.StatusOK, removedCat)
		} else {
			var response interface{}
			json.Unmarshal([]byte(fmt.Sprintf(`{ "error": "unable to delete cat with name %s" }`, nameToRemove)), &response)
			handlerutil.RespondWithJSON(w, http.StatusBadRequest, response)
		}
	}
}

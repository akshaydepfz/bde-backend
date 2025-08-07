package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"lantorabde.app/helper"
	"lantorabde.app/models"
)

func Userhandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		PostUsers(w, r)
	} else if r.Method == http.MethodGet {
		GetUsers(w, r)
	} else if r.Method == http.MethodDelete {
		DeleteUsers(w, r)
	} else if r.Method == http.MethodPut {
		PutUsers(w, r)
	} else {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
	}

}
func Userhandlerget(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetUserid(w, r)
	} else {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
	}

}
func GetUserid(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	user, err := helper.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return user in JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func PostUsers(w http.ResponseWriter, r *http.Request) {
	var users models.BDEUser
	users.FullName = r.FormValue("full_name")
	users.Email = r.FormValue("email")
	users.Phone = r.FormValue("phone")
	users.PasswordHash = r.FormValue("password")
	users.DrivingLicense = r.FormValue("driving_license")
	users.Role = r.FormValue("role")
	users.Status = r.FormValue("status")

	if users.FullName == "" || users.Email == "" || users.Phone == "" || users.PasswordHash == "" || users.DrivingLicense == "" {
		http.Error(w, "Name,Email,phone,password and driving license are required", http.StatusBadRequest)
		return
	}
	if users.Role == "" {
		users.Role = "BDE"
	} else if users.Role != "BDE" && users.Role != "Manager" && users.Role != "Admin" {
		http.Error(w, "Invalid role: must be 'BDE', 'Manager', or 'Admin'", http.StatusBadRequest)
		return
	}

	if users.Status == "" {
		users.Status = "ACTIVE"
	} else if users.Status != "ACTIVE" && users.Status != "INACTIVE" {
		http.Error(w, "Invalid status: must be 'ACTIVE' or 'INACTIVE'", http.StatusBadRequest)
		return
	}

	id, err := helper.Insertuser(users.FullName, users.Email, users.Phone, users.PasswordHash, users.DrivingLicense, users.Role, users.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users.ID = id
	users.JoinDate = time.Now()
	users.CreatedAt = time.Now()
	users.UpdatedAt = time.Now()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := helper.GetUsers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = helper.DeleteUser(uint(id))
	if err != nil {
		if err.Error() == "User not found" {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to delete User: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("✅ User with ID %d deleted successfully\n", id)
	w.WriteHeader(http.StatusNoContent)

}

func PutUsers(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = r.ParseMultipartForm(10 << 20) // handle form-data (up to 10 MB)
	if err != nil {
		http.Error(w, "Error parsing form data: "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.BDEUser
	user.ID = uint(id)
	user.FullName = r.FormValue("full_name")
	user.Email = r.FormValue("email")
	user.Phone = r.FormValue("phone")
	user.PasswordHash = r.FormValue("password")
	user.DrivingLicense = r.FormValue("driving_license")
	user.Role = r.FormValue("role")
	user.Status = r.FormValue("status")

	err = helper.PutUser(user.ID, user.FullName, user.Email, user.Phone, user.PasswordHash, user.DrivingLicense, user.Role, user.Status)
	if err != nil {
		if err.Error() == "Customer not found" {
			http.Error(w, "Customer not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to update customer: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

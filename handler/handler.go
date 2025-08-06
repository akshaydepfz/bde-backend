package handler
import(
 	"encoding/json"
	// "fmt"
	"net/http"
	// "strconv"
	 "time"
	"lantorabde.app/models"
	"lantorabde.app/helper"
)

func Userhandler(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			PostUsers(w, r)
		} else if r.Method == http.MethodGet {
			GetUsers(w, r)
		}  else {
			http.Error(w, "Invalid request method", http.StatusBadRequest)
		}

}

func PostUsers(w http.ResponseWriter, r *http.Request){
var users models.BDEUser
users.FullName =r.FormValue("full_name")
users.Email=r.FormValue("email")
users.Phone=r.FormValue("phone")
users.PasswordHash=r.FormValue("password")
users.DrivingLicense=r.FormValue("driving_license")
users.Role=r.FormValue("role")
users.Status=r.FormValue("status")

if users.FullName == "" || users.Email == "" || users.Phone =="" || users.PasswordHash ==""|| users.DrivingLicense==""{
			http.Error(w, "Name,Email,phone,password and driving license are required", http.StatusBadRequest)
			return
		}
		if users.Role == "" {
			users.Role = "BDE"
		}
		if users.Status == "" {
			users.Status = "ACTIVE"
		}
		
		id, err := helper.Insertuser(users.FullName,users.Email,users.Phone,users.PasswordHash,users.DrivingLicense,users.Role,users.Status)
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





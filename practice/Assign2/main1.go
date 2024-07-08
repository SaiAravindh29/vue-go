package main

import (
	as "Assign2/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Student struct {
	Name     string      `json:"name"`
	Marks    as.Marks    `json:"marks"`
	School   as.School   `json:"school"`
	Address  as.Address  `json:"address"`
	Personal as.Personal `json:"personal"`
}

var Students []Student

func main() {
	log.Println("Starting...")
	http.HandleFunc("/getStudent", getStudent)
	http.HandleFunc("/createStudent", createStudent)
	http.HandleFunc("/updateStudent", updateStudent)
	http.ListenAndServe(":10000", nil)
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting data ...")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET")
	(w).Header().Set("Content-Type", "application/json")

	name := r.URL.Query().Get("name")
	dtype := r.URL.Query().Get("dtype")

	if name == "" && dtype != "" {

		if dtype == "All" || dtype == "all" {
			Data, err := json.Marshal(Students)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(Data)

		} else if dtype == "Marks" || dtype == "marks" {

			var marksData []as.Marks

			for _, student := range Students {

				marksData = append(marksData, student.Marks)
			}
			Data, err := json.Marshal(marksData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(Data)

		} else if dtype == "School" || dtype == "school" {
			var schoolData []as.School
			for _, student := range Students {
				schoolData = append(schoolData, student.School)
			}
			Data, err := json.Marshal(schoolData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(Data)

		} else if dtype == "Address" || dtype == "address" {
			var addressData []as.Address
			for _, student := range Students {
				addressData = append(addressData, student.Address)
			}
			Data, err := json.Marshal(addressData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(Data)

		} else if dtype == "Personal" || dtype == "personal" {

			var personalData []as.Personal

			for _, student := range Students {
				personalData = append(personalData, student.Personal)
			}

			Data, err := json.Marshal(personalData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(Data)

		} else {
			http.Error(w, "Invalid data type", http.StatusBadRequest)
			return
		}

	} else if name != "" && dtype != "" {

		for _, student := range Students {
			if student.Name == name {
				if dtype == "All" || dtype == "all" {
					Data, err := json.Marshal(student)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					w.Write(Data)
					return
				} else if dtype == "Marks" || dtype == "marks" {
					Data, err := json.Marshal(student.Marks)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					w.Write(Data)
					return
				} else if dtype == "School" || dtype == "school" {
					Data, err := json.Marshal(student.School)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					w.Write(Data)
					return
				} else if dtype == "Address" || dtype == "address" {
					Data, err := json.Marshal(student.Address)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					w.Write(Data)
					return
				} else if dtype == "Personal" || dtype == "personal" {
					Data, err := json.Marshal(student.Personal)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					w.Write(Data)
					return
				} else {
					http.Error(w, "Invalid data type", http.StatusBadRequest)
					return
				}
			}
		}

	} else {
		http.Error(w, "Invalid data type", http.StatusBadRequest)
		return
	}

}

func createStudent(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Creating Student .........")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST")
	(w).Header().Set("Content-Type", "application/json")

	namekey := r.Header.Get("name")

	defer r.Body.Close()
	var newStudent Student

	newStudent.Name = namekey

	err := json.NewDecoder(r.Body).Decode(&newStudent)
	if err != nil {
		log.Println("Error parsing request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if newStudent.Name != "" {
		if newStudent.Marks.M10.English != 0 && newStudent.Marks.M10.Tamil != 0 && newStudent.Marks.M10.Maths != 0 && newStudent.Marks.M10.Science != 0 && newStudent.Marks.M10.Social != 0 {
			if newStudent.School.S10.SchoolName != "" && newStudent.School.S10.Address != "" && newStudent.School.S10.Place != "" && newStudent.School.S10.Pincode != "" && newStudent.School.S10.Type != "" {
				if newStudent.Address.Landmark != "" && newStudent.Address.Pincode != "" && newStudent.Address.StudentAddress != "" {
					if newStudent.Personal.FatherName != "" && newStudent.Personal.MotherName != "" && newStudent.Personal.Age != 0 && newStudent.Personal.Gender != "" {

						Students = append(Students, newStudent)

						fmt.Fprintf(w, "Student %v created", newStudent.Name)
						fmt.Fprintln(w, Students)
						return
					}
				}
			}

		}

	}

	fmt.Fprintln(w, "Student Not created")

	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	log.Println("Error reading request body:", err)
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// err = json.Unmarshal(body, &newStudent)
	// if err != nil {
	// 	log.Println("Error parsing request body:", err)
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// Students = append(Students, newStudent)

}

func updateStudent(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Updating Student .........")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Content-Type", "application/json")

	defer r.Body.Close()

	name := r.URL.Query().Get("name")
	dtype := r.URL.Query().Get("dtype")

	for i := 0; i < len(Students); i++ {

		if Students[i].Name == name {

			if dtype == "Marks" || dtype == "marks" {

				var tempMarks as.Marks

				body, err := ioutil.ReadAll(r.Body)

				if err != nil {
					log.Println("Error reading request body:", err)
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				err = json.Unmarshal(body, &tempMarks)
				if err != nil {
					log.Println("Error parsing request body:", err)
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				Students[i].Marks = tempMarks
				message := Students[i].Name + " Marks is Updated Successfully"
				jsonData, err := json.Marshal(message)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(jsonData)
				return

			} else if dtype == "name" || dtype == "Name" {

				oldname := Students[i].Name

				tempname := r.Header.Get("name")
				Students[i].Name = tempname

				message := oldname + " changed to " + Students[i].Name + " \n Updated Successfully"
				jsonData, err := json.Marshal(message)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(jsonData)
				return

			} else if dtype == "School" || dtype == "school" {

				var tempSchool as.School
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					log.Println("Error reading request body:", err)
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				err = json.Unmarshal(body, &tempSchool)
				if err != nil {
					log.Println("Error parsing request body:", err)
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				Students[i].School = tempSchool
				message := Students[i].Name + " School Details is Updated Successfully"
				jsonData, err := json.Marshal(message)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(jsonData)
				return

			} else if dtype == "Address" || dtype == "address" {

				var tempAddress as.Address
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					log.Println("Error reading request body:", err)
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				err = json.Unmarshal(body, &tempAddress)
				if err != nil {
					log.Println("Error parsing request body:", err)
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				Students[i].Address = tempAddress
				message := Students[i].Name + " Address is Updated Successfully"

				jsonData, err := json.Marshal(message)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(jsonData)
				return

			} else if dtype == "Personal" || dtype == "personal" {

				var tempPersonal as.Personal
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					log.Println("Error reading request body:", err)
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				err = json.Unmarshal(body, &tempPersonal)
				if err != nil {
					log.Println("Error parsing request body:", err)
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				Students[i].Personal = tempPersonal
				message := Students[i].Name + " Personal Details is Updated Successfully"

				jsonData, err := json.Marshal(message)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(jsonData)
				return

			} else {

				http.Error(w, "Invalid data type", http.StatusBadRequest)
				return
			}

		}
	}

}

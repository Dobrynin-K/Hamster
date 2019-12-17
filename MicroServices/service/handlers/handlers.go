package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/matscus/Hamster/MicroServices/service/serv"
	"github.com/matscus/Hamster/Package/Services/service"
)

// //GetServicesHandle -  handle for response all services
// func GetServicesHandle(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	page, ok := params["project"]
// 	if ok {
// 		if len(serv.GetResponseAllData) == 0 {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte("{\"Message\":\"len services slice equally 0\"}"))
// 		} else {
// 			serv.CheckService()
// 			l := len(serv.GetResponseAllData)
// 			res := make([]service.Service, 0, l)
// 			iter := 0
// 			for i := 0; i < l; i++ {
// 				projects := serv.GetResponseAllData[i].Projects
// 				for i := 0; i < len(projects); i++ {
// 					if projects[i] == page {
// 						res = append(res, serv.GetResponseAllData[iter])
// 						break
// 					}
// 				}
// 				iter++
// 			}
// 			err := json.NewEncoder(w).Encode(res)
// 			if err != nil {
// 				w.WriteHeader(http.StatusInternalServerError)
// 				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
// 			}
// 		}
// 	} else {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("{\"Message\":\" params hot found \"}"))
// 	}
// }

//UpdateServiceHandle - handle for update  services data in database
func UpdateServiceHandle(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var service service.Service
	decoder.Decode(&service)
	err := service.Update()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"Status\":\"done\"}"))
	}
}

//NewServiceHandle -  install service to remote host - compress dir of service, use implement scp to copy
//files to remote host and uncompress. SERVICE NOT RUN AFTER INSTALL
func NewServiceHandle(w http.ResponseWriter, r *http.Request) {
	var err error
	decoder := json.NewDecoder(r.Body)
	var service service.Service
	err = decoder.Decode(&service)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	}
	user, _ := serv.HostsAndUsers.Load(service.Host)
	err = service.InstallServiceToRemoteHost(user.(string))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"Status\":\"done\"}"))
	}

}

//DeleteService - hadler from delete service
func DeleteService(w http.ResponseWriter, r *http.Request) {
	var err error
	decoder := json.NewDecoder(r.Body)
	var service service.Service
	err = decoder.Decode(&service)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	}
	user, _ := serv.HostsAndUsers.Load(service.Host)
	err = service.Stop(user.(string))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	}
	err = service.DeleteServiceToRemoteHost(user.(string))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"Status\":\"done\"}"))
	}
}

// //NewOrUpdateGenerator - new or update generator func(new - method post, update - method put)
// func NewOrUpdateGenerator(w http.ResponseWriter, r *http.Request) {
// 	g := hosts.Host{}
// 	err := json.NewEncoder(w).Encode(g)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
// 	}
// 	switch r.Method {
// 	case "GET":
// 		gen, err := client.PGClient{}.New().GetAllGenerators()
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
// 		}
// 		l := len(gen)
// 		temp := make([]hosts.Host, 0, l)
// 		for i := 0; i < l; i++ {
// 			var g hosts.Host
// 			t := gen[i]
// 			id, _ := strconv.Atoi(t[0])
// 			g.ID = int64(id)
// 			g.Host = t[1]
// 			temp = append(temp, g)
// 		}
// 		err = json.NewEncoder(w).Encode(temp)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
// 		}
// 	case "POST":
// 		err = g.Create()
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
// 		}
// 	case "PUT":
// 		err = g.Update()
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
// 		}
// 	}
// }

//NewFile - upload (metoth post) file to bins diectory
func NewFile(w http.ResponseWriter, r *http.Request) {
	servicetype := r.FormValue("servicetype")
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	}
	defer file.Close()
	f, err := os.OpenFile(os.Getenv("BINSDIR")+servicetype+"/"+handler.Filename, os.O_CREATE|os.O_RDWR, os.FileMode(0755))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	}
}

//File - func to download and delete file from bins dir
func File(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		params := mux.Vars(r)
		servicetype, ok := params["servicetype"]
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"Param servicetype not found\"}"))
		}
		servicename, ok := params["servicename"]
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"Param servicename not found\"}"))
			return
		}
		file, err := os.Open(os.Getenv("BINSDIR") + servicetype + "/" + servicename + ".tar.gz")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			return
		}
		_, err = io.Copy(w, file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			return
		}
	case "DELETE":
		params := mux.Vars(r)
		servicetype, ok := params["servicetype"]
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"Param servicetype not found\"}"))
		}
		servicename, ok := params["servicename"]
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"Param servicename not found\"}"))
			return
		}
		err := os.Remove(os.Getenv("BINSDIR") + servicetype + "/" + servicename + ".tar.gz")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			return
		}
	}
}

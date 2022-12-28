package main

import (
	conf "cityService/config"
	filt "cityService/pkg/filter"
	glog "cityService/pkg/glog"
	gs "cityService/pkg/gracefullShutdown"
	rspn "cityService/pkg/response"
	"cityService/pkg/storage"
	"encoding/json"
	f "fmt"
	"net/http"
)

var (
	log   *glog.Logger
	store *storage.Store
	mux   *http.ServeMux
	err   error
	resp  rspn.Response

	storageFileName = conf.GoDotEnvVariable("STORAGE_FILE_NAME")
	hostPort        = conf.GoDotEnvVariable("HOST_PORT")
)

func init() {
	log = glog.Init()

	store, err = storage.Init(storageFileName)

	f.Println("storage loaded:")
	store.Print()

	resp.CheckErr(err)

	mux = http.NewServeMux()

	resp.Res = make(map[string]interface{})
}

func main() {
	gs.CheckInterrupt(store)

	/****** handles ******/
	mux.HandleFunc("/get_city_info", getCityInfo)
	mux.HandleFunc("/add_new_city", addNewCity)
	mux.HandleFunc("/rem_city", removeCity)
	mux.HandleFunc("/upd_city", updateCityInfo)
	mux.HandleFunc("/get_cities", getCities)

	/*** start server ****/
	err := http.ListenAndServe(":"+hostPort, mux)
	resp.CheckErr(err)
}

// получение информации о городе по его id;
// curl -X GET -H "Content-type: application/json" -H "Accept: application/json" -d '{"id":13}' "http://localhost:8080/get_city_info"
// curl -X GET -H "Content-type: application/json" -H "Accept: application/json" -d '{"id":606}' "http://localhost:8080/get_city_info"
func getCityInfo(w http.ResponseWriter, r *http.Request) {
	if !resp.CorrectMeth(r.Method, "GET") {
		return
	}

	var reqCity storage.City
	err := reqCity.UnmarshalCity(&r.Body)
	resp.CheckErr(err)

	resp.Init(w)

	city, ok := store.M[reqCity.Id]
	if !ok {
		resp.Status = http.StatusNoContent
		resp.Send()
		return
	}

	resp.Status = http.StatusOK
	resp.Res["res"] = city
	resp.Send()
}

// добавление новой записи в список городов;
// curl -X PUT -H "Content-type: application/json" -H "Accept: application/json" -d '{"id":999,"population":5,"foundation":2000,"name":"Симферополь","region":"Крым","district":"Южный"}' "http://localhost:8080/add_new_city"
func addNewCity(w http.ResponseWriter, r *http.Request) {
	if !resp.CorrectMeth(r.Method, "PUT") {
		return
	}
	resp.Init(w)

	var reqCity storage.City
	err := reqCity.UnmarshalCity(&r.Body)
	resp.CheckErr(err)

	store.M[reqCity.Id] = &reqCity

	resp.Status = http.StatusCreated
	resp.Res["res"] = resp.RetId(reqCity.Id)
	resp.Send()
}

// удаление информации о городе по указанному id;
// curl -X DELETE -H "Content-type: application/json" -H "Accept: application/json" -d '{"id":999}' "http://localhost:8080/rem_city"
// curl -X DELETE -H "Content-type: application/json" -H "Accept: application/json" -d '{"id":606}' "http://localhost:8080/rem_city"
func removeCity(w http.ResponseWriter, r *http.Request) {
	if !resp.CorrectMeth(r.Method, "DELETE") {
		return
	}
	resp.Init(w)

	var reqCity storage.City
	err := reqCity.UnmarshalCity(&r.Body)
	resp.CheckErr(err)

	if _, ok := store.M[reqCity.Id]; !ok {
		resp.Err("city not exist", http.StatusBadRequest)
		return
	}

	delete(store.M, reqCity.Id)

	resp.Status = http.StatusOK
	resp.Res["res"] = resp.RetId(reqCity.Id)
	resp.Send()
}

// обновление информации о численности населения города по указанному id;
// curl -X POST -H "Content-type: application/json" -H "Accept: application/json" -d '{"id":606,"population":5,"foundation":2000,"name":"Симферополь","region":"Крым","district":"Южный"}' "http://localhost:8080/upd_city"
func updateCityInfo(w http.ResponseWriter, r *http.Request) {
	if !resp.CorrectMeth(r.Method, "POST") {
		return
	}
	resp.Init(w)

	var reqCity storage.City
	err := reqCity.UnmarshalCity(&r.Body)
	resp.CheckErr(err)

	if _, ok := store.M[reqCity.Id]; !ok {
		resp.Err("city not exist", http.StatusBadRequest)
		return
	}

	store.M[reqCity.Id] = &reqCity

	resp.Status = http.StatusOK
	resp.Res["res"] = resp.RetId(reqCity.Id)
	resp.Send()
	log.Println("Обновлена информация по городу #id=", reqCity.Id)
}

// получение списка городов по указанному региону, округу, численности населения, указанному диапазону года основания.
// curl -X GET -H "Content-type: application/json" -H "Accept: application/json" -d '{"district":"Сибирский", "population":[1000000, 2000000]}' "http://localhost:8080/get_cities"
// curl -X GET -H "Content-type: application/json" -H "Accept: application/json" -d '{"region":"Москва"}' "http://localhost:8080/get_cities"
func getCities(w http.ResponseWriter, r *http.Request) {
	resp.Init(w)

	if !resp.CorrectMeth(r.Method, "GET") {
		return
	}

	var reqFilter filt.FilterCity
	err := reqFilter.UnmarshalFilt(&r.Body)
	resp.CheckErr(err)

	res := make([]string, 0, len(store.M)/4)
	keys := store.GetSortedKeys()

	for _, key := range keys {
		city := store.M[key]
		if reqFilter.Region != "" && reqFilter.Region != city.Region {
			continue
		}
		if reqFilter.District != "" && reqFilter.District != city.District {
			continue
		}
		if len(reqFilter.Foundation) != 0 && (reqFilter.Foundation[0] > city.Foundation || reqFilter.Foundation[1] < city.Foundation) {
			continue
		}
		if len(reqFilter.Population) != 0 && (reqFilter.Population[0] > city.Population || reqFilter.Population[1] < city.Population) {
			continue
		}
		cityInfo, err := json.Marshal(city)
		resp.CheckErr(err)
		res = append(res, string(cityInfo))
	}

	if len(res) == 0 {
		resp.Status = http.StatusNoContent
		resp.Send()
		return
	}

	resp.Status = http.StatusOK
	resp.Res["res"] = res
	resp.Send()
}

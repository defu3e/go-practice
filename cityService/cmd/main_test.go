package main

import (
	"bytes"
	conf "cityService/config"
	"io"
	"net/http"
	"testing"
)

/**
* WARNING!
* This test writing for city_orig.csv file
* Before test you should:
* copy cityService/db/city_orig.csv to cityService/db/city.csv
**/

var (
	host = conf.GoDotEnvVariable("HOST")
	port = conf.GoDotEnvVariable("HOST_PORT")
	URL  = host + ":" + port
)

func TestMain(t *testing.T) {
	client := &http.Client{}

	type testCase struct {
		name            string
		url             string
		inputMethod     string
		inputData       string
		expectedStatus  int
		expectedMessage string
	}

	testCases := []testCase{
		// получение информации о городе по его id;
		{
			name:            "get city info by existed city id",
			url:             "/get_city_info",
			inputMethod:     "GET",
			inputData:       `{"id":490}`,
			expectedStatus:  200,
			expectedMessage: `{"res":{"id":490,"population":11514330,"foundation":1147,"name":"Москва","region":"Москва","district":"Центральный"}}`,
		},
		{
			name:            "get city info by unexist city id",
			url:             "/get_city_info",
			inputMethod:     "GET",
			inputData:       `{"id":13}`,
			expectedStatus:  204,
			expectedMessage: ``,
		},
		// добавление новой записи в список городов;
		{
			name:            "add city into db",
			url:             "/add_new_city",
			inputMethod:     "PUT",
			inputData:       `{"id":999,"population":5,"foundation":2000,"name":"Симферополь","region":"Крым","district":"Южный"}`,
			expectedStatus:  201,
			expectedMessage: `{"res":{"id":999}}`,
		},
		// удаление информации о городе по указанному id;
		{
			name:            "correct delete city from db by id",
			url:             "/rem_city",
			inputMethod:     "DELETE",
			inputData:       `{"id":177}`,
			expectedStatus:  200,
			expectedMessage: `{"res":{"id":177}}`,
		},
		{
			name:            "incorrect delete city from db by id (unexist id)",
			url:             "/rem_city",
			inputMethod:     "DELETE",
			inputData:       `{"id":1999}`,
			expectedStatus:  400,
			expectedMessage: `{"err":"city not exist"}`,
		},
		// обновление информации о численности населения города по указанному id;
		{
			name:            "update existing city info",
			url:             "/upd_city",
			inputMethod:     "POST",
			inputData:       `{"id":769,"population":5,"foundation":2000,"name":"Симферополь","region":"Крым","district":"Южный"}`,
			expectedStatus:  200,
			expectedMessage: `{"res":{"id":769}}`,
		},
		{
			name:            "update unexisting city",
			url:             "/upd_city",
			inputMethod:     "POST",
			inputData:       `{"id":1999,"population":5,"foundation":2000,"name":"Симферополь","region":"Крым","district":"Южный"}`,
			expectedStatus:  400,
			expectedMessage: `{"err":"city not exist"}`,
		},
		// получение городов по фильтру
		{
			name:            "get cities by filter: district+population",
			url:             "/get_cities",
			inputMethod:     "GET",
			inputData:       `{"district":"Сибирский", "population":[1000000, 2000000]}`,
			expectedStatus:  200,
			expectedMessage: `{"res":["{\"id\":410,\"population\":1000000,\"foundation\":1628,\"name\":\"Красноярск\",\"region\":\"Красноярский край\",\"district\":\"Сибирский\"}","{\"id\":634,\"population\":1498921,\"foundation\":1893,\"name\":\"Новосибирск\",\"region\":\"Новосибирская область\",\"district\":\"Сибирский\"}","{\"id\":643,\"population\":1154000,\"foundation\":1716,\"name\":\"Омск\",\"region\":\"Омская область\",\"district\":\"Сибирский\"}"]}`,
		},
		{
			name:            "get cities by filter: region",
			url:             "/get_cities",
			inputMethod:     "GET",
			inputData:       `{"region":"Крым"}`,
			expectedStatus:  200,
			expectedMessage: `{"res":["{\"id\":769,\"population\":5,\"foundation\":2000,\"name\":\"Симферополь\",\"region\":\"Крым\",\"district\":\"Южный\"}","{\"id\":999,\"population\":5,\"foundation\":2000,\"name\":\"Симферополь\",\"region\":\"Крым\",\"district\":\"Южный\"}"]}`,
		},
	}

	for _, test := range testCases {

		// make request
		var jsonStr = []byte(test.inputData)
		url := URL + test.url
		req, err := http.NewRequest(test.inputMethod, url, bytes.NewBuffer(jsonStr))
		if err != nil {
			t.Errorf("got error on req: %s", err)
		}

		// do request
		req.Header.Set("content-type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("got error on api request %s", err)
		}
		defer resp.Body.Close()

		// get response
		body, _ := io.ReadAll(resp.Body)

		t.Run(test.name, func(t *testing.T) {
			t.Log(test.name)
			//check status
			if resp.StatusCode != test.expectedStatus {
				t.Errorf("Got unexpected status %s", resp.Status)
			}
			// check body
			if string(body) != test.expectedMessage {
				t.Errorf("Got unexpected result body %s", body)
			}
		})
	}
}

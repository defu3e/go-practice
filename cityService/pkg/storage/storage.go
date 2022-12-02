package storage

import (
	"cityService/pkg/csv"
	log "cityService/pkg/glog"
	"encoding/json"
	f "fmt"
	"io"
	"sort"
	"strconv"
)

type Store struct {
	M  map[uint64]*City
	db string
}

type City struct {
	Id         uint64 `json:"id"`         //уникальный номер
	Population uint64 `json:"population"` //численность населения;
	Foundation uint64 `json:"foundation"` //год основания.
	Name       string `json:"name"`       //название города
	Region     string `json:"region"`     //регион
	District   string `json:"district"`   //округ
}

var (
	glog = log.Init()
)

func Init(dbfile string) (*Store, error) {
	glog.Println("Init storage")
	s := Store{make(map[uint64]*City), dbfile}
	data, err := csv.GetCSVdata(dbfile)
	if err != nil {
		return nil, err
	}

	s.Load(data)

	return &s, nil
}

func (s *Store) Close() {
	glog.Println("Catch close() command")
	glog.Println("Prepare data to csv writer...")

	data := s.prepareToWrite()

	csv.WriteDataToCSV(s.db, data)
}

func (s *Store) GetSortedKeys() []uint64 {
	keys := make([]uint64, 0, len(s.M))
	for key := range s.M {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	return keys
}

func (s *Store) prepareToWrite() [][]string {
	data := make([][]string, 0, len(s.M))
	for i, city := range s.M {
		data[i] = []string{
			strconv.FormatUint(city.Id, 10),
			strconv.FormatUint(city.Population, 10),
			strconv.FormatUint(city.Foundation, 10),
			city.Name,
			city.Region,
			city.District,
		}
	}
	return data
}

func (s *Store) Load(rows [][]string) {
	for _, row := range rows {
		id, _ := strconv.ParseUint(row[0], 10, 64)
		populat, _ := strconv.ParseUint(row[4], 10, 64)
		foundat, _ := strconv.ParseUint(row[5], 10, 64)

		s.M[id] = &City{
			Id:         id,
			Population: populat,
			Foundation: foundat,
			Name:       row[1],
			Region:     row[2],
			District:   row[3],
		}
	}
}

func (c *City) UnmarshalCity(body *io.ReadCloser) error {
	err := json.NewDecoder(*body).Decode(&c)
	if err != nil {
		return err
	}
	defer (*body).Close()

	return nil
}

func (s *Store) Print() {
	//log.Print("INIT STORAGE")
	for _, v := range s.M {
		f.Printf("%+v\n", v)
	}
}

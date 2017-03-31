package slark

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"reflect"

	"github.com/jmoiron/sqlx"

	"fmt"

	"github.com/emostafa/garson"
	"gopkg.in/yaml.v2"
)

var models map[string]*Model
var router *garson.Router
var config *Config

type Config struct {
	Route              string `yaml:"route"`
	DBConnectionString string `yaml:"db_connection_string"`
	DBDriver           string `yaml:"db_driver"`
}

func connectDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect(config.DBDriver, config.DBConnectionString)
	return db, err
}

func readConfigFile() (*Config, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(path.Join(wd, "slark.yaml"))
	conf := &Config{}
	err = yaml.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func init() {
	models = make(map[string]*Model)
	router = garson.New()
	con, err := readConfigFile()
	if err != nil {
		log.Fatal(err)
	}
	config = con
	log.Println(config)
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}

	router.Get(fmt.Sprintf("/%s", config.Route), index(db))
	router.Get(fmt.Sprintf("/%s/:model", config.Route), list(db))
	router.Get(fmt.Sprintf("/%s/:model/new", config.Route), getCreate(db))
	router.Get(fmt.Sprintf("/%s/:model/edit/:id", config.Route), getUpdate(db))
	router.Get(fmt.Sprintf("/%s/:model/:id", config.Route), read(db))

	router.Post(fmt.Sprintf("/%s/:model", config.Route), postCreate(db))
	router.Post(fmt.Sprintf("/%s/:model/edit/:id", config.Route), postUpdate(db))

	router.Delete(fmt.Sprintf("/%s/:model/:id", config.Route), delete(db))
}

// Handle handles http requests
func Handle(w http.ResponseWriter, r *http.Request) {
	defer func() {
		recover()
	}()

	handle, params, err := router.Try(r.URL.Path, r.Method)
	if err != nil {
		http.NotFound(w, r)
	}

	ctx := context.WithValue(r.Context(), "route_params", params)
	r = r.WithContext(ctx)

	handle(w, r)
}

// Register registers a given model to slark
func Register(model interface{}, table string) {
	m := &Model{}
	m.OriginalModel = model
	m.TableName = table
	// get the fields on the model struct
	if t := reflect.TypeOf(model); t.Kind() == reflect.Ptr {
		m.Name = t.Elem().Name()
		m.Reflection = t
	}
	v := reflect.ValueOf(model).Elem()

	m.Fields = make([]Field, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag
		f := Field{}
		f.Name = tag.Get("db")
		f.DBType = tag.Get("type")
		m.Fields[i] = f
	}

	models[m.Name] = m
}

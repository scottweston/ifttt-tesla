package main

import (
	"encoding/json"
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/mux"
	"github.com/jsgoecke/tesla"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
)

var conf *viper.Viper
var vehicles tesla.Vehicles

type AuthToken struct {
	AuthToken string
}

func isValidClient(body io.ReadCloser) bool {
	var auth AuthToken

	b, err := ioutil.ReadAll(body)
	if err != nil {
		return false
	}
	err = json.Unmarshal(b, &auth)
	if err != nil {
		return false
	}
	auth_tokens := conf.GetStringSlice("auth_tokens")
	sort.Strings(auth_tokens)
	i := sort.SearchStrings(auth_tokens, auth.AuthToken)
	if i < len(auth_tokens) && auth_tokens[i] == auth.AuthToken {
		return true
	}
	return false
}

func TeslaHonk(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	v, err := strconv.Atoi(params["vehicle"])
	if err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	if !isValidClient(req.Body) {
		http.Error(w, "unauthorized", 403)
		return
	}

	if len(vehicles) > v {
		vehicle := vehicles[v]
		vehicle.HonkHorn()
	} else {
		http.Error(w, "vehicle not found", 404)
		return
	}
}

func TeslaFlash(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	v, err := strconv.Atoi(params["vehicle"])
	if err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	if !isValidClient(req.Body) {
		http.Error(w, "unauthorized", 403)
		return
	}

	if len(vehicles) > v {
		vehicle := vehicles[v]
		vehicle.FlashLights()
	} else {
		http.Error(w, "vehicle not found", 404)
		return
	}
}

func TeslaStartCharging(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	v, err := strconv.Atoi(params["vehicle"])
	if err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	if !isValidClient(req.Body) {
		http.Error(w, "unauthorized", 403)
		return
	}

	if len(vehicles) > v {
		vehicle := vehicles[v]
		vehicle.StartCharging()
	} else {
		http.Error(w, "vehicle not found", 404)
		return
	}
}

func TeslaStopCharging(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	v, err := strconv.Atoi(params["vehicle"])
	if err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	if !isValidClient(req.Body) {
		http.Error(w, "unauthorized", 403)
		return
	}

	if len(vehicles) > v {
		vehicle := vehicles[v]
		vehicle.StopCharging()
	} else {
		http.Error(w, "vehicle not found", 404)
		return
	}
}

func TeslaLock(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	v, err := strconv.Atoi(params["vehicle"])
	if err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	if !isValidClient(req.Body) {
		http.Error(w, "unauthorized", 403)
		return
	}

	if len(vehicles) > v {
		vehicle := vehicles[v]
		vehicle.LockDoors()
	} else {
		http.Error(w, "vehicle not found", 404)
		return
	}
}

func TeslaUnlock(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	v, err := strconv.Atoi(params["vehicle"])
	if err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	if !isValidClient(req.Body) {
		http.Error(w, "unauthorized", 403)
		return
	}

	if len(vehicles) > v {
		vehicle := vehicles[v]
		vehicle.UnlockDoors()
	} else {
		http.Error(w, "vehicle not found", 404)
		return
	}
}

func TeslaOpenChargePort(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	v, err := strconv.Atoi(params["vehicle"])
	if err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	if !isValidClient(req.Body) {
		http.Error(w, "unauthorized", 403)
		return
	}

	if len(vehicles) > v {
		vehicle := vehicles[v]
		vehicle.OpenChargePort()
	} else {
		http.Error(w, "vehicle not found", 404)
		return
	}
}

func TeslaSetChargeLimit(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	v, err := strconv.Atoi(params["vehicle"])
	if err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	l, err := strconv.Atoi(params["limit"])
	if err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	if !isValidClient(req.Body) {
		http.Error(w, "unauthorized", 403)
		return
	}

	if len(vehicles) > v {
		vehicle := vehicles[v]
		if l > 0 && l <= 100 {
			vehicle.SetChargeLimit(l)
		} else {
			http.Error(w, "invalid charge limit", 400)
			return
		}
	} else {
		http.Error(w, "vehicle not found", 404)
		return
	}
}

func TeslaSetTemperature(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	v, err := strconv.Atoi(params["vehicle"])
	if err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	t, err := strconv.ParseFloat(params["temp"], 64)
	if err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	if !isValidClient(req.Body) {
		http.Error(w, "unauthorized", 403)
		return
	}

	if len(vehicles) > v {
		vehicle := vehicles[v]
		vehicle.SetTemprature(t, t) // TODO: misspelt in support library
	} else {
		http.Error(w, "vehicle not found", 404)
		return
	}
}

func main() {
	router := mux.NewRouter()

	conf = viper.New()
	conf.SetEnvPrefix("tesla")
	conf.SetConfigName("tesla")
	conf.AddConfigPath(".")
	conf.AddConfigPath("$HOME/.config/")
	conf.AddConfigPath("/")

	err := conf.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	conf.WatchConfig()
	conf.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})

	client, err := tesla.NewClient(
		&tesla.Auth{
			ClientID:     conf.GetString("client_id"),
			ClientSecret: conf.GetString("client_secret"),
			Email:        conf.GetString("username"),
			Password:     conf.GetString("password"),
		})
	if err != nil {
		log.Fatal(err)
	}

	vehicles, err = client.Vehicles()
	if err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/honk/{vehicle:[0-9]+}", TeslaHonk).Methods("POST")
	router.HandleFunc("/unlock/{vehicle:[0-9]+}", TeslaUnlock).Methods("POST")
	router.HandleFunc("/lock/{vehicle:[0-9]+}", TeslaLock).Methods("POST")
	router.HandleFunc("/set_charge_limit/{vehicle:[0-9]+}/{limit:[0-9]+}", TeslaSetChargeLimit).Methods("POST")
	router.HandleFunc("/start_charge/{vehicle:[0-9]+}", TeslaStartCharging).Methods("POST")
	router.HandleFunc("/stop_charge/{vehicle:[0-9]+}", TeslaStopCharging).Methods("POST")
	router.HandleFunc("/flash/{vehicle:[0-9]+}", TeslaFlash).Methods("POST")
	router.HandleFunc("/open_charge_port/{vehicle:[0-9]+}", TeslaOpenChargePort).Methods("POST")
	router.HandleFunc("/set_temperature/{vehicle:[0-9]+}/{temp:[0-9]+}", TeslaSetTemperature).Methods("POST")

	log.Fatal(http.ListenAndServe(":3514", router))
}

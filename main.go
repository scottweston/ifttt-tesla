package main

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jsgoecke/tesla"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"
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
		log.Println("missing body, can't authenticate request")
		return false
	}
	err = json.Unmarshal(b, &auth)
	if err != nil {
		log.Println("malformed body, can't authenticate request")
		return false
	}
	auth_tokens := conf.GetStringSlice("auth_tokens")
	sort.Strings(auth_tokens)
	i := sort.SearchStrings(auth_tokens, auth.AuthToken)
	if i < len(auth_tokens) && auth_tokens[i] == auth.AuthToken {
		return true
	}
	log.Println("invalid token: " + auth.AuthToken)
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
		var try = 0
		for try < conf.GetInt("retries") {
			err = vehicle.HonkHorn()
			if err == nil {
				break
			}
			log.Println(err)
			try++
			time.Sleep(1000 * time.Millisecond)
		}
	} else {
		http.Error(w, "vehicle not found", 404)
		return
	}
	fmt.Fprintln(w, "ok")
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
		var try = 0
		for try < conf.GetInt("retries") {
			err = vehicle.FlashLights()
			if err == nil {
				break
			}
			log.Println(err)
			try++
			time.Sleep(1000 * time.Millisecond)
		}
	} else {
		http.Error(w, "vehicle not found", 404)
		return
	}
	fmt.Fprintln(w, "ok")
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
		var try = 0
		for try < conf.GetInt("retries") {
			err = vehicle.StartCharging()
			if err == nil {
				break
			}
			log.Println(err)
			try++
			time.Sleep(1000 * time.Millisecond)
		}
	} else {
		http.Error(w, "vehicle not found", 404)
		return
	}
	fmt.Fprintln(w, "ok")
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
		var try = 0
		for try < conf.GetInt("retries") {
			err = vehicle.StopCharging()
			if err == nil {
				break
			}
			log.Println(err)
			try++
			time.Sleep(1000 * time.Millisecond)
		}
	} else {
		http.Error(w, "vehicle not found", 404)
		return
	}
	fmt.Fprintln(w, "ok")
}

func TeslaStartHvac(w http.ResponseWriter, req *http.Request) {
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
		var try = 0
		for try < conf.GetInt("retries") {
			err = vehicle.StartAirConditioning()
			if err == nil {
				break
			}
			log.Println(err)
			try++
			time.Sleep(1000 * time.Millisecond)
		}
	} else {
		http.Error(w, "vehicle not found", 404)
		return
	}
	fmt.Fprintln(w, "ok")
}

func TeslaStopHvac(w http.ResponseWriter, req *http.Request) {
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
		var try = 0
		for try < conf.GetInt("retries") {
			err = vehicle.StopAirConditioning()
			if err == nil {
				break
			}
			log.Println(err)
			try++
			time.Sleep(1000 * time.Millisecond)
		}
	} else {
		http.Error(w, "vehicle not found", 404)
		return
	}
	fmt.Fprintln(w, "ok")
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
		var try = 0
		for try < conf.GetInt("retries") {
			err = vehicle.LockDoors()
			if err == nil {
				break
			}
			log.Println(err)
			try++
			time.Sleep(1000 * time.Millisecond)
		}
	} else {
		http.Error(w, "vehicle not found", 404)
		return
	}
	fmt.Fprintln(w, "ok")
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
		var try = 0
		for try < conf.GetInt("retries") {
			err = vehicle.UnlockDoors()
			if err == nil {
				break
			}
			log.Println(err)
			try++
			time.Sleep(1000 * time.Millisecond)
		}
	} else {
		http.Error(w, "vehicle not found", 404)
		return
	}
	fmt.Fprintln(w, "ok")
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
		var try = 0
		for try < conf.GetInt("retries") {
			err = vehicle.OpenChargePort()
			if err == nil {
				break
			}
			log.Println(err)
			try++
			time.Sleep(1000 * time.Millisecond)
		}
	} else {
		http.Error(w, "vehicle not found", 404)
		return
	}
	fmt.Fprintln(w, "ok")
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
			var try = 0
			for try < conf.GetInt("retries") {
				err = vehicle.SetChargeLimit(l)
				if err == nil {
					break
				}
				log.Println(err)
				try++
				time.Sleep(1000 * time.Millisecond)
			}
		} else {
			http.Error(w, "invalid charge limit", 400)
			return
		}
	} else {
		http.Error(w, "vehicle not found", 404)
		return
	}
	fmt.Fprintln(w, "ok")
}

func TeslaSetTemperature(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	v, err := strconv.Atoi(params["vehicle"])
	if err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	t, err := strconv.Atoi(params["temp"])
	if err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	if !isValidClient(req.Body) {
		http.Error(w, "unauthorized", 403)
		return
	}

	var temp = float64(t)
	if !conf.GetBool("metric") {
		temp = (temp - 32) / 1.8
	}

	if len(vehicles) > v {
		vehicle := vehicles[v]
		var try = 0
		for try < conf.GetInt("retries") {
			err = vehicle.SetTemprature(temp, temp) // TODO: misspelt in support library
			if err == nil {
				break
			}
			log.Println(err)
			try++
			time.Sleep(1000 * time.Millisecond)
		}
	} else {
		http.Error(w, "vehicle not found", 404)
		return
	}
	fmt.Fprintln(w, "ok")
}

func main() {
	router := mux.NewRouter()

	conf = viper.New()
	conf.SetEnvPrefix("tesla")
	conf.SetConfigName("tesla")
	conf.AddConfigPath(".")
	conf.AddConfigPath("$HOME/.config/")
	conf.AddConfigPath("/")

	conf.SetDefault("metric", true)
	conf.SetDefault("bind", "0.0.0.0")
	conf.SetDefault("port", "3514")
	conf.SetDefault("retries", 3)

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
	router.HandleFunc("/start_hvac/{vehicle:[0-9]+}", TeslaStartHvac).Methods("POST")
	router.HandleFunc("/stop_hvac/{vehicle:[0-9]+}", TeslaStopHvac).Methods("POST")
	router.HandleFunc("/flash/{vehicle:[0-9]+}", TeslaFlash).Methods("POST")
	router.HandleFunc("/open_charge_port/{vehicle:[0-9]+}", TeslaOpenChargePort).Methods("POST")
	router.HandleFunc("/set_temperature/{vehicle:[0-9]+}/{temp:[0-9]+}", TeslaSetTemperature).Methods("POST")

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(conf.GetString("bind")+":"+conf.GetString("port"), loggedRouter))
}

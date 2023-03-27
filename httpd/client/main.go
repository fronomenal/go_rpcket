package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	clrpc "github.com/fronomenal/go_rpcket/httpd/client/client_rpc"
	roc "github.com/fronomenal/go_rpcket/protos/v1"
)

func main() {
	conn, err := clrpc.Connect()
	if err != nil {
		log.Fatalf("could not connect to Server at %s: %v", "51515", err)
	}
	defer conn.Close()

	var handler http.ServeMux
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/rocket":
			rktHandler(w, r)
		case "/":
			if r.Method != "GET" {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Method %s not supported", r.Method)
				return
			}
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "Welcome to Go RPC Rockets MicroService")
		default:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "Invalid Route: Please visit /rocket")
			return

		}

	})

	server := http.Server{Addr: "localhost:8080", Handler: &handler}
	log.Println("Listening on localhost:8080")
	server.ListenAndServe()

}

func rktHandler(w http.ResponseWriter, r *http.Request) {

	getValidArgs := func() (clrpc.Rarg, error) {
		var valargs clrpc.Rarg
		var validId bool

		if num, err := strconv.Atoi(r.FormValue("id")); err == nil {
			valargs.Id = int32(num)
			validId = true
		}
		if len(r.FormValue("name")) > 0 && len(r.FormValue("type")) > 0 {
			valargs.Name = r.FormValue("name")
			valargs.Rkt_type = r.FormValue("type")

			if num, err := strconv.Atoi(r.FormValue("flights")); err == nil {
				valargs.Flights = int32(num)
				valargs.Valid = true
			}
		} else if err := json.NewDecoder(r.Body).Decode(&valargs); err == nil {
			if len(valargs.Name) > 0 && len(valargs.Rkt_type) > 0 {
				valargs.Valid = true
			}
		}

		if !validId {
			return valargs, fmt.Errorf("invalid id passed")
		}

		return valargs, nil
	}
	writeRocketResponse := func(rkt *roc.Rocket) {
		var resp string

		switch r.Header.Get("Accept") {
		case "text/plain":
			w.Header().Add("Content-Type", "text/plain")
			resp = fmt.Sprintf("ID: %d\tName: %s\tType: %s\tFlights: %d", rkt.Id, rkt.Name, rkt.Type, rkt.Flights)
		case "text/html":
			w.Header().Add("Content-Type", "text/html")
			resp = fmt.Sprintf("<p>ID: %d</p><p>Name: %s</p><p>Type: %s</p><p>Flights: %d</p>", rkt.Id, rkt.Name, rkt.Type, rkt.Flights)
		default:
			w.Header().Add("Content-Type", "application/json")
			resp = fmt.Sprintf(`{"ID": %d, "Name": "%s", "Type": "%s", "Flights": %d}`, rkt.Id, rkt.Name, rkt.Type, rkt.Flights)
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, resp)

	}
	writeBadRequest := func(msg string) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, msg)
	}

	switch r.Method {
	case "GET":
		args, err := getValidArgs()
		if err != nil {
			writeBadRequest("Invalid Id passed: Must be an integer")
			return
		}

		rocket, err := clrpc.Get(&args)
		if err != nil {
			writeBadRequest(err.Error())
			return
		}
		writeRocketResponse(rocket)

	case "POST":
		args, _ := getValidArgs()
		if !args.Valid {
			writeBadRequest("Invalid arguments passed: must include name<string>; type<string>; flights<integer>")
			return
		}

		rocket, err := clrpc.Set(&args)
		if err != nil {
			writeBadRequest(err.Error())
			return
		}
		writeRocketResponse(rocket)

	case "DELETE":
		args, err := getValidArgs()
		if err != nil {
			writeBadRequest("Invalid Id passed: Must be an integer")
			return
		}

		status, err := clrpc.Rem(&args)
		if err != nil {
			writeBadRequest(err.Error())
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `"Status": "%s"`, status)

	default:
		writeBadRequest(fmt.Sprintf("Method %s not supported", r.Method))
	}

}

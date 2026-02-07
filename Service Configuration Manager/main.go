package main

import "fmt"

func LoadConfig( env string) (port int, status string){

	switch env {
	case "prod":
		port = 80
	case "dev":
		port = 8080
	case "test":
		port = 9090
	default :
		port = 3000
	}
	if env =="dev" && port == 80 {
		return 0, "Security Alert: Cannot run production port in dev!"
	}

	status= "Config loaded successfully"
	return port, status
}

func main(){
	p, s := LoadConfig("dev")
	fmt.Printf("Port: %d with Status: %s\n", p, s)
}
package main

import (
	"encoding/json"
	"fmt"
    "log"
    "math/rand"
    "net/http"
    "sync"
    "time"
	//"html"
	"qlip/util"
)

type store struct {
    mu       sync.Mutex
    buffer   string
    validOTPs map[string]bool
}

var globalStore = store{
	buffer:   "",
    validOTPs: make(map[string]bool),
}

type PostData struct {
	Qlip string `json:"qlip"`
	Otp  string `json:"otp"`
}


func main() {
	masertKey := util.GetEnvCloud("MASTERKEY")
	if masertKey == "" {
		fmt.Println("Error loading environment variables")
		return
	}



    http.HandleFunc("/getotp", func(w http.ResponseWriter, r *http.Request) {
        pass := r.URL.Query()
		userAuth := false
		for key := range pass {
			if key == masertKey {
				fmt.Println("user is in - horaay", r.RemoteAddr)
				userAuth = true
			}
		}

		if !userAuth {
			fmt.Println("Someone tryed to ENTER: ", r.RemoteAddr)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		newOTP := randomOTP(6)

		globalStore.mu.Lock()
		globalStore.validOTPs[newOTP] = true 
		globalStore.mu.Unlock()

		someData := "Welecome back Master, here is your OTP:\n" //+ "http://192.168.1.68:8080/set?otp=" + newOTP + "&qlip="
		copyPart := `curl "https://qlip.alexandermander.dk/set?otp=` + newOTP + `&qlip=$qlip"`
		someData += copyPart
		fmt.Fprintln(w, someData)
    })

	// mkaie a post http hander called testpost
	http.HandleFunc("/testpost", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// get the qlip and otp params json body
		// get json body
		var postData PostData 
		err := json.NewDecoder(r.Body).Decode(&postData)
		if err != nil {
			http.Error(w, "Invalid JSON body", http.StatusBadRequest)
			return
		}

		qlip := postData.Qlip 
		otp := postData.Otp 

		if qlip == "" || otp == "" {
			http.Error(w, "Missing qlip or otp param", http.StatusBadRequest)
			return
		}
		globalStore.mu.Lock()
		_, exists := globalStore.validOTPs[otp]
		if !exists {
			globalStore.mu.Unlock()
			http.Error(w, "Invalid or used OTP", http.StatusUnauthorized)
			return
		}
		delete(globalStore.validOTPs, otp)
		globalStore.buffer = qlip 
		globalStore.mu.Unlock()
		fmt.Fprintln(w, "Data stored successfully via POST!")
	})

    http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
        qlip := r.URL.Query().Get("qlip")

        if qlip == "" {
            http.Error(w, "Missing qlip param", http.StatusBadRequest)
            return
        }

        otp := r.URL.Query().Get("otp")
        if otp == "" {
            http.Error(w, "Missing otp param", http.StatusBadRequest)
            return
        }

        globalStore.mu.Lock()
        _, exists := globalStore.validOTPs[otp]
        if !exists {
            globalStore.mu.Unlock()
            http.Error(w, "Invalid or used OTP", http.StatusUnauthorized)
            return
        }

        delete(globalStore.validOTPs, otp)

        globalStore.buffer = qlip

        globalStore.mu.Unlock()

        fmt.Fprintln(w, "Data stored successfully!")
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        globalStore.mu.Lock()
        data := globalStore.buffer
        globalStore.mu.Unlock()

		w.Header().Set("Content-Type", "text/html") // Ensure HTML rendering
		w.Write([]byte(data))
    })


	http.Handle("/paste/", http.StripPrefix("/paste/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server is running on port http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}




// randomOTP generates an alphanumeric code of length n
func randomOTP(n int) string {
    const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    rand.Seed(time.Now().UnixNano())

    b := make([]byte, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

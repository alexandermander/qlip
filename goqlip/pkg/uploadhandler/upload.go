package uploadhandler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"os"

	"qlip/pkg/otp"
)

type postData struct {
	Qlip string `json:"qlip"`
	Otp  string `json:"otp"`
}

func UploadToService(uploadPath string) error {
	dat, err := os.ReadFile(uploadPath)
	if err != nil {
		return err
	}

	encodedUploadData := base64.StdEncoding.EncodeToString(dat)
	currntOtp, err := otp.GetOTP()
	if err != nil {
		return err
	}
	
	postbuf := postData{
		Qlip: encodedUploadData,
		Otp: currntOtp,
	}
	
	jsonBytes, err := json.Marshal(postbuf)
	if err != nil {
		return err
	}
	
	newRest, err := http.Post("https://qlip.alexandermander.dk/testpost", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		panic(err)
	}

	defer newRest.Body.Close()
	if newRest.StatusCode != 200 {
		return fmt.Errorf("unexpected status code: %s", newRest.Status)
	}
	fmt.Println("successfully uploaded to the cloud! 🚀💻✨")

	return nil
}

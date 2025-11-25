package otp

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func GetOTP() (string, error) {
	fmt.Println("OTP Password:")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
    if err != nil {
        return "", err
    }

	line = strings.TrimSpace(line)

	resp, err := http.Get("https://qlip.alexandermander.dk/getotp?" + line)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Welecome back Master, here is your OTP:
	// curl "https://qlip.alexandermander.dk/set?otp=Nf3Uhd&qlip=$qlip"
	bodyStr := string(body)

	re := regexp.MustCompile(`otp=([A-Za-z0-9]+)&`)
	match := re.FindStringSubmatch(bodyStr)
	if len(match) < 2 {
		return "", fmt.Errorf("otp not found in response")
	}
	otp := match[1]

	return otp, nil
}

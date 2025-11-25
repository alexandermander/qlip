package download

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadHandler(savePath string) error {
	resp, err := http.Get("https://qlip.alexandermander.dk/")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", resp.Status)
	}

	encBase, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// encodedUploadData := base64.StdEncoding.EncodeToString(dat)
	dcodedData, err := base64.StdEncoding.DecodeString(string(encBase))
	if err != nil {
		return nil
	}

	f, err := os.Create(savePath)
	if err != nil {
		return nil
	}
	f.Write(dcodedData)

	fmt.Printf("Download complete - Successfully!! ☁️💾✔️")
	return nil
}

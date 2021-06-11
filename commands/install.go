package main

import (
    "fmt"
    "github.com/reef-pi/reef-pi/controller/utils"
    "io"
    "log"
    "net/http"
    "os"
    "path"
    "path/filepath"
    "time"
)

const urlTemplate = "https://github.com/hectorespert/planted-pi/releases/download/%s-planted-pi/reef-pi-%s-planted-pi.deb"

func downloadDeb(version string) (string, error) {
	url := fmt.Sprintf(urlTemplate, version, version)
	log.Println("Downloading reef-pi from:", url)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	file := path.Base(resp.Request.URL.Path)
	file = filepath.Join(os.TempDir(), file)
	out, err := os.Create(file)
	if err != nil {
		return "", err
	}

	// Write the body to file
	if _, err = io.Copy(out, resp.Body); err != nil {
		return "", err
	}
	out.Close()
	resp.Body.Close()
	log.Println("reef-pi debian package downloaded at:", file)
	return file, nil
}

func install(version string) error { // version to upgrade to
	log.Println("Executing reef-pi upgrade command")
	time.Sleep(time.Second)
	if _, err := utils.Command("/bin/systemctl", "stop", "reef-pi.service").CombinedOutput(); err != nil {
		log.Println("Failed to stop reef-pi:", err)
	}
	time.Sleep(time.Second)

	file, err := downloadDeb(version)
	if err != nil {
		log.Println("Failed to download reef-pi")
		return err
	}

	if out, err := utils.Command("/usr/bin/dpkg", "-i", file).CombinedOutput(); err != nil {
		log.Println("Failed to install new reef-pi:", err, "\n", string(out))
		return err

	}

	if _, err := utils.Command("/bin/systemctl", "start", "reef-pi.service").CombinedOutput(); err != nil {
		log.Println("Failed to start reef-pi:", err)
		return err
	}
	log.Println("reef-pi upgrade is done")
	return nil
}

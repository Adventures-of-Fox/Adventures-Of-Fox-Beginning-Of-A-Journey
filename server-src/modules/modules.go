package modules

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	g "server/global"

	"github.com/cavaliergopher/grab/v3"
)

func Download(link string, msg string, place string) string {
	// create client
	client := grab.NewClient()
	req, _ := grab.NewRequest(place, link)

	// start download
	g.Logs(msg)
	resp := client.Do(req)
	g.Logs("%v", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			g.Logs("transferred %v / %v bytes (%.2f%%)",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	if err, errMsg := g.CheckError(resp.Err(), "Download failed"); err != nil {
		log.Fatalf(errMsg)
	}

	g.Done("Downloaded %v", resp.Filename)
	return resp.Filename
}

func Unziping(file string, output string) {
	err := unzipSource(file, output)
	if err, errMsg := g.CheckError(err, "Unziping file failed"); err != nil {
		log.Fatalf(errMsg)
	}
	g.Done("Unziped %v", file)
}


func unzipSource(source, destination string) error {
    reader, err := zip.OpenReader(source)
    if err != nil {
        return err
    }
    defer reader.Close()

    destination, err = filepath.Abs(destination)
    if err != nil {
        return err
    }

    for _, f := range reader.File {
        err := unzipFile(f, destination)
        if err != nil {
            return err
        }
    }

    return nil
}

func unzipFile(f *zip.File, destination string) error {
    filePath := filepath.Join(destination, f.Name)
    if !strings.HasPrefix(filePath, filepath.Clean(destination)+string(os.PathSeparator)) {
        return fmt.Errorf("invalid file path: %s", filePath)
    }

    if f.FileInfo().IsDir() {
        if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
            return err
        }
        return nil
    }

    if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
        return err
    }

    destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
    if err != nil {
        return err
    }
    defer destinationFile.Close()

    zippedFile, err := f.Open()
    if err != nil {
        return err
    }
    defer zippedFile.Close()

    if _, err := io.Copy(destinationFile, zippedFile); err != nil {
        return err
    }
    return nil
}

func RunCmd(cmdRun string, cmdArgs ...string) string {
	cmd := exec.Command(cmdRun, cmdArgs...)
	returned, err := cmd.Output()
	if err, errMsg := g.CheckError(err, "Running failed"); err != nil {
		log.Fatalf(errMsg)
	}
	return string(returned)
}
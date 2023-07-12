package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Server().MaxRequestBodySize = 50 * 1024 * 1024

	app.Post("/upload-records", func(c *fiber.Ctx) error {
		requestBody := c.Body()
		fmt.Printf("req body: %d\n", requestBody)
		return c.SendStatus(fiber.StatusOK)
	})

	app.Post("/apk-upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("apk")
		if err != nil {
			return err
		}

		fileContent, err := file.Open()
		if err != nil {
			return err
		}

		defer func(fileContent multipart.File) {
			err := fileContent.Close()
			if err != nil {

			}
		}(fileContent)

		body, err := ioutil.ReadAll(fileContent)
		if err != nil {
			return err
		}

		if err := saveApkFile(file.Filename, body); err != nil {
			return err
		}

		return c.SendString("APK file uploaded successfully")
	})

	app.Get("/get-apk/:name", func(c *fiber.Ctx) error {
		fileName := c.Params("name")
		apkDir := ""

		if strings.Contains(fileName, "Regular") {
			apkDir = "apks/regular/"
		} else if strings.Contains(fileName, "Simulator") {
			apkDir = "apks/simulator/"
		}
		if apkDir == "" {
			return errors.New("invalid APK file name")
		}

		filePath := filepath.Join(apkDir, fileName)
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)

		fileContent, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		return c.Send(fileContent)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		files, err := ioutil.ReadDir("apks/regular/")
		if err != nil {
			return err
		}

		var highestVersion string
		var highestVersionFileName string

		for _, file := range files {
			if !file.IsDir() {
				fileName := file.Name()

				version := extractVersion(fileName)

				if compareVersions(version, highestVersion) > 0 {
					highestVersion = version
					highestVersionFileName = fileName
				}
			}
		}

		if highestVersionFileName == "" {
			return c.Status(fiber.StatusNotFound).SendString("No APK files found")
		}

		filePath := filepath.Join("apks/regular/", highestVersionFileName)
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)

		fileContent, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		return c.Send(fileContent)
	})

	app.Get("/apks-show", func(c *fiber.Ctx) error {
		filesMap := make(map[string][]string)

		regularFiles, err := ioutil.ReadDir("apks/regular")
		if err != nil {
			return err
		}

		var regularFileNames []string
		for _, file := range regularFiles {
			regularFileNames = append(regularFileNames, file.Name())
		}
		filesMap["regular"] = regularFileNames

		simulatorFiles, err := ioutil.ReadDir("apks/simulator")
		if err != nil {
			return err
		}
		var simulatorFileNames []string

		for _, file := range simulatorFiles {
			simulatorFileNames = append(simulatorFileNames, file.Name())
		}
		filesMap["simulator"] = simulatorFileNames

		return c.JSON(filesMap)
	})

	app.Get("/apk-version", func(c *fiber.Ctx) error {
		files, err := ioutil.ReadDir("./apks/regular")
		if err != nil {
			return err
		}
		var highestVersion string

		for _, file := range files {
			if !file.IsDir() {
				fileName := file.Name()
				version := extractVersion(fileName)

				if compareVersions(version, highestVersion) > 0 {
					highestVersion = version
				}
			}
		}

		return c.SendString(highestVersion)
	})

	app.Get("/taggant-settings", func(c *fiber.Ctx) error {
		taggantSettingsList := []TaggantSettings{
			{
				TaggantName:       "Taggant 1",
				SignalFirstRange:  &IntPair{10, 20},
				SignalSecondRange: &IntPair{30, 40},
			},
			{
				TaggantName:       "Taggant 2",
				SignalFirstRange:  &IntPair{50, 60},
				SignalSecondRange: &IntPair{70, 80},
			},
		}

		return c.JSON(taggantSettingsList)
	})

	if err := app.Listen(":3000"); err != nil {
		fmt.Println(err)
	}
}

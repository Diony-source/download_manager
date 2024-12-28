package services

import (
	"download_manager/entities"
	"download_manager/utils"
	"fmt"
	"log"
	"sync"
)

func DownloadApps(apps []entities.App) {
	var wg sync.WaitGroup

	for _, app := range apps {
		wg.Add(1)
		go func(app entities.App) {
			defer wg.Done()
			fmt.Printf("Downloading %s from %s\n", app.Name, app.URL)
			err := utils.DownloadFile(app.Name+".exe",app.URL)
			if err != nil {
				log.Printf("Failed to download %s: %v\n", app.Name, err)
			} else {
				fmt.Printf("%s downloaded successfully!\n", app.Name)
			}
		}(app)
	}

	wg.Wait()
	fmt.Println("All downloads completed.")
}
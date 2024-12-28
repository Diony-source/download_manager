package handlers

import (
	"download_manager/entities"
	"download_manager/services"
	"fmt"
	"strconv"
	"strings"
)

// StartCLI starts the command-line interface for the download manager
func StartCLI() {
	apps := []entities.App{
		{Name: "Google Chrome", URL: "https://www.google.com/chrome/"},
		{Name: "Microsoft Teams", URL: "https://www.microsoft.com/microsoft-teams/download-app"},
		{Name: "Zoom", URL: "https://zoom.us/download"},
		{Name: "WinRAR", URL: "https://www.win-rar.com/download.html"},
		{Name: "VLC Media Player", URL: "https://www.videolan.org/vlc/"},
		{Name: "Spotify", URL: "https://www.spotify.com/download/"},
		{Name: "Notepad++", URL: "https://notepad-plus-plus.org/downloads/"},
		{Name: "Adobe Reader", URL: "https://get.adobe.com/reader/"},
		{Name: "FileZilla", URL: "https://filezilla-project.org/download.php"},
		{Name: "Steam", URL: "https://cdn.fastly.steamstatic.com/client/installer/SteamSetup.exe"},
	}

	fmt.Println("Available Apps:")
	for i, app := range apps {
		fmt.Printf("%d. %s\n", i+1, app.Name)
	}

	var choices string
	fmt.Print("Enter the numbers of the apps to download (comma-separated): ")
	fmt.Scanln(&choices)

	selectedApps := parseChoices(choices, apps)
	services.DownloadApps(selectedApps)
}

// parseChoices parses user input and selects the corresponding apps
func parseChoices(input string, apps []entities.App) []entities.App {
	choices := strings.Split(input, ",")
	selectedApps := []entities.App{}

	for _, choice := range choices {
		// Trim any whitespace and convert to integer
		index, err := strconv.Atoi(strings.TrimSpace(choice))
		if err != nil {
			fmt.Printf("Invalid input: %s. Skipping...\n", choice)
			continue
		}

		// Ensure the index is within bounds
		if index >= 1 && index <= len(apps) {
			selectedApps = append(selectedApps, apps[index-1]) // Corrected index (1-based to 0-based)
		} else {
			fmt.Printf("Invalid choice: %d. Skipping...\n", index)
		}
	}

	return selectedApps
}

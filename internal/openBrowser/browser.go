package openBrowser

import (
	"os/exec"
	"runtime"
)

// OpenURLFunc is a variable that can be set to a different function for testing.
var OpenURLFunc = OpenURL

// OpenURL opens the specified URL in the default browser of the user.
func OpenURL(url string) error {
	var err error
	switch runtime.GOOS {
	case "darwin":
		err = exec.Command("open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	default: // linux
		err = exec.Command("xdg-open", url).Start()
	}
	return err
}

// isWSL checks if the Go program is running inside Windows Subsystem for Linux
// func isWSL() bool {
// 	releaseData, err := exec.Command("uname", "-r").Output()
// 	if err != nil {
// 		return false
// 	}
// 	return strings.Contains(strings.ToLower(string(releaseData)), "microsoft")
// }

package analytic

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

// Data represent analytic data
type Data struct {
	User struct {
		Role  string `json:"role"`
		Email string `json:"email"`
		Nik   string `json:"nik"`
	} `json:"user"`
	Host     string `json:"host"`
	URL      string `json:"url"`
	Level    string `json:"level"`
	Activity string `json:"activity"`
	Object   struct {
		Segment         string `json:"segment"`
		ApplicationType string `json:"applicationType"`
		JobTitle        string `json:"jobTitle"`
	} `json:"object"`
	Label   string `json:"label"`
	Port    int    `json:"port"`
	Action  string `json:"action"`
	Message string `json:"message"`
}

const (
	// AppName label
	AppName = "go-codebase"
)

var (
	// Anaytic logrus custom instance
	Anaytic *log.Logger
)

// InitAnalytic will init logger configuration
// https://github.com/telkomdev/go-stash => s *stash.Stash
func InitAnalytic(s io.Writer) {
	Anaytic = log.New(s, "", 0)

	Anaytic.SetOutput(s)
}

// Log will log all event in info mode
func Log(data *Data) {

	fmt.Printf("%+v", data)
	// this line should be panic
	if Anaytic == nil {
		panic("Analytic not created yet")
	}

	data.Label = AppName

	jsonData, _ := json.Marshal(data)

	Anaytic.Print(string(jsonData))

}

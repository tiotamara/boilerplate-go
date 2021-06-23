package helpers

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

var ProjectFolder = flag.String("folder", "./", "absolute path of project folder")

func WriteLogError(title string, msg string) {
	// overwrite filename
	title = "server"
	WriteToLogfile(*ProjectFolder+"logs/"+title+".log", "Error: "+msg)
}

func WriteToLogfile(filename string, text string) {
	//create your file with desired read/write permissions
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println(text)
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

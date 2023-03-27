package main

import(
	"os"
	"encoding/json"
)

func main() {
	//Get Clips for show by slug, then parse useful info
	clips := GetAllClips("behind-the-bastards")
	list := parseClips(clips)
	
	//Output to File
	data, jsonErr := json.MarshalIndent(list, "", "	")
	if jsonErr != nil {
		panic(jsonErr)
	}
	saveErr := os.WriteFile("Clips.json", data, 0755)
	if saveErr != nil {
		panic(saveErr)
	}
}
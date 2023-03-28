package main

import(
	"os"
	"fmt"
	"strconv"
	"encoding/json"
)

func main() {
	//Get All Clips
	clips := GetAllClips("behind-the-bastards")
	list := parseClips(clips)
	length := len(list)
	
	//Output to File
	data, jsonErr := json.MarshalIndent(list, "", "	")
	if jsonErr != nil {
		panic(jsonErr)
	}
	saveErr := os.WriteFile("Clips.json", data, 0755)
	if saveErr != nil {
		panic(saveErr)
	}
	fmt.Println(strconv.Itoa(length) + " Episodes Found")
}
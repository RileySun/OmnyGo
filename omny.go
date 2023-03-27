package main

import(
	"log"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

const APIURL = "https://api.omny.fm/"

const SHOWURL = "shows/"

/*----------------------------------------------*/

//Clip
type Clip struct {
	AdMarkers []AdMarker
	AudioOptions AudioOptions
	AudioUrl string
	Chapters []Chapter
	ContentRating string
	CustomFieldData map[string]string
	Description string
	DescriptionHtml string
	DurationSeconds float64
	EmbedUrl string
	Episode int32
	EpisodeType string
	ExternalId string
	HasPreRollVideoAd bool
	HasPublishedTranscript bool
	Id string
	ImageUrl string
	ImportedId string
	ModifiedAtUtc string
	OrganizationId string
	PlaylistIds []string
	ProgramId string
	ProgramSlug string
	PublishedAudioSizeInBytes int64
	PublishedUrl string
	PublishedUtc string
	PublishState string
	RecordingMetaData RecordingMetaData
	RssLinkOverride string
	Season int32
	ShareUrl string
	Slug string
	State string
	Summary string
	Tags []string
	Title string
	Visibility string
	WaveformUrl string
}

type AdMarker struct {
	AdMarkerType string
	MaxNumberOfAds int32
	Offset string
}

type AudioOptions struct {
	AutoLevelAudio bool
	IncludeIntroOutro bool
}

type Chapter struct {
	Id string
	Name string
	Position string
	Tags []string
}

type RecordingMetaData struct {
	CaptureEndUtc string
	CaptureStartUtc string
}

//Program
type ProgramSlugResp struct {
	Clips []Clip
	Cursor string
	Total int
}

//Non-API (Formmated Ouput of golang program)
type FormatClip struct {
	Name string
	URL string
	Image string
}

/*----------------------------------------------*/

//Base
func getAPI(url string) []byte {
	//Generic GET request to Open Weather Map API
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	
	req, err := http.NewRequest("GET", url, nil)
	
	if err != nil {
		log.Fatal(err)
	}
	
	resp, err := client.Do(req)
	
	if err != nil {
		log.Fatal(err)
	}
	
	responseData, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		log.Fatal(err)
	}
	
	defer resp.Body.Close()
	
	return responseData
}

//Util
func getProgramID(programSlug string) string {
	url := APIURL + "programs/" + programSlug
	bytes := getAPI(url)
	
	var data map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatal(err)
	}
	
	progID := string(data["Id"][:])
	progID = progID[1:len(progID)-1]
	
	return progID
}

func getOrgID(programSlug string) string {
	url := APIURL + "programs/" + programSlug
	bytes := getAPI(url)
	
	var data map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatal(err)
	}
	
	orgID := string(data["OrganizationId"][:])
	orgID = orgID[1:len(orgID)-1]
	
	return orgID
}

func GetClips(programSlug string) []Clip {
	orgID := getOrgID(programSlug)
	progID := getProgramID(programSlug)

	url := APIURL + "orgs/" + orgID + "/programs/" + progID + "/clips"
	bytes := getAPI(url)
	
	var data ProgramSlugResp 
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatal(err)
	}
	
	return data.Clips
}

//Gets Up to 1000 clips (TODO: can get total clips from program, and then target it specifically)
func GetAllClips(programSlug string) []Clip {
	orgID := getOrgID(programSlug)
	progID := getProgramID(programSlug)

	url := APIURL + "orgs/" + orgID + "/programs/" + progID + "/clips?pageSize=1000"
	bytes := getAPI(url)
	
	var data ProgramSlugResp 
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatal(err)
	}
	
	return data.Clips
}

/*----------------------------------------------*/

//Format
func parseClips(clips []Clip) []*FormatClip {
	var list []*FormatClip
	for _, clip := range clips {
		out := &FormatClip{
			Name:clip.Title,
			URL:clip.AudioUrl,
			Image:clip.ImageUrl,
		}
		list = append(list, out)
	}
	return list
}
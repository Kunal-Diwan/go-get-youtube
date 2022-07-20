package youtubedownloader

import (
	"encoding/json"
	"fmt"
	"go-youtube-downloader/youtubeResponse"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	API_KEY     = "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"
	YOUTUBE_URL = "https://www.youtube.com/youtubei/v1/player?key="
)

// Video contains the information of the Video and the Download option.
type Video struct {
	Id, Title, Author, Keywords, ThumbnailUrl string
	AvgRating                                 float32
	ViewCount, LengthSeconds                  int
	Formats                                   []Format
	Filename                                  string
	ContentLength                             string
}
type Format struct {
	Itag                                   int
	VideoType, Quality, Url, ContentLength string
}

// GetMetaInformation return Video from the Video, given by videoId
func GetMetaInformation(videoId string) (Video, error) {
	videoId = strings.TrimSuffix(videoId, "\r")
	if strings.Contains(videoId, "youtube.com/watch?") {
		videoId, _ = extractId(videoId)
	} else {
		return Video{}, fmt.Errorf("pls enter a valid YouTube link, %s, is not valid \n", videoId)
	}

	// fetch video meta from youtube
	meta, err := fillVideoData(videoId)

	if err != nil {
		return Video{}, err
	}
	return *meta, nil
}

func extractId(input string) (string, error) {
	u, err := url.Parse(input)

	if err != nil {
		return "", err
	}

	queries := u.Query()
	for key, value := range queries {
		if key == "v" {
			return value[0], nil
		}
	}
	return "", fmt.Errorf("no video ID detectable")
}
func fillVideoData(videoId string) (*Video, error) {

	metaInformation, err := getVideoMetaInformation(videoId)
	if err != nil {
		return &Video{}, fmt.Errorf("error fetching metaInformation: %w", err)
	}
	// collate the necessary params
	video := &Video{
		Id:       videoId,
		Title:    metaInformation.VideoDetails.Title,
		Author:   metaInformation.VideoDetails.Author,
		Keywords: fmt.Sprint(metaInformation.VideoDetails.Keywords),
	}
	if len(metaInformation.VideoDetails.Thumbnail.Thumbnails) != 0 {
		video = &Video{ThumbnailUrl: metaInformation.VideoDetails.Thumbnail.Thumbnails[0].URL}
	}

	v, _ := strconv.Atoi(metaInformation.VideoDetails.ViewCount)
	video.ViewCount = v

	l, _ := strconv.Atoi(metaInformation.VideoDetails.LengthSeconds)
	video.LengthSeconds = l

	adaptiveFormats := metaInformation.StreamingData.AdaptiveFormats
	for _, f := range adaptiveFormats {

		video.Formats = append(video.Formats, Format{
			Itag:          f.Itag,
			VideoType:     f.MimeType,
			Quality:       f.Quality,
			Url:           f.URL,
			ContentLength: f.ContentLength,
		})
	}

	return video, nil
}

func getVideoMetaInformation(videoId string) (youtubeResponse youtubeResponse.YoutubeResponse, err error) {

	body := strings.NewReader(
		"{" +
			"  \"videoId\": \"" + videoId + "\"," +
			"  \"context\": {" +
			"    \"client\": {" +
			"      \"hl\": \"en\"," +
			"      \"gl\": \"US\"," +
			"      \"clientName\": \"ANDROID\"," +
			"      \"clientVersion\": \"16.02\"" +
			"    }" +
			"  }" +
			"}")

	request, err := http.NewRequest("POST", YOUTUBE_URL+API_KEY, body)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Errorf("curl error: %w", err)
	}
	err = parseResponseIntoObject(response, &youtubeResponse)
	if err != nil {
		return youtubeResponse, fmt.Errorf("error while parsing response into object: %w", err)
	}
	return youtubeResponse, nil
}

func parseResponseIntoObject(response *http.Response, youtubeResponse *youtubeResponse.YoutubeResponse) error {
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error parsing body: %w", err)
	}
	err = json.Unmarshal(b, &youtubeResponse)
	if err != nil {
		return fmt.Errorf("error while parsing response-body to youtubeResponse object: %w", err)
	}
	return nil
}

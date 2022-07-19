# go-get-youtube v 0.2
A tiny go-library for downloading YouTube videos.
This is just a refactor and updated Version from https://github.com/Kunal-Diwan/go-get-youtube, all credits go to https://github.com/Kunal-Diwan
# Usage
Download and then navigate to the folder where the program is located. 
Open a terminal/cmd window and type in the following:
```
downloader.exe (or ./downloader if you are on linux)
```
Just copy your link in the terminal, and press enter.

# Library

## Example
```go
import (
	youtube "github.com/Kunal-Diwan/go-get-youtube/youtube"
)

func main() {
	// get the video object (with metdata)
	videoUrl:= "https://www.youtube.com/watch?v=C0DPdy98e4c"
	video, err := youtube.GetMetaInformation(videoUrl)

	// download the video and write to file
	option := &youtube.Option{
		Rename: true,  // rename file using video title
		Resume: true,  // resume cancelled download
		Mp3:    true,  // extract audio to MP3
	}
	video.Download(0, "video.mp4", option)
}
```




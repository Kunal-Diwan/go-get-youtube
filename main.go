package main

import (
	"bufio"
	"fmt"
	"go-youtube-downloader/youtubedownloader"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Pls enter your YouTube link")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		video, err := youtubedownloader.GetMetaInformation(text)
		if err != nil {
			fmt.Printf(err.Error())

		} else {
			//func (video *Video) Download(index int, filename string, option *Option) error {
			option := &youtubedownloader.Option{
				Rename: true, // rename file using video title
				Resume: true, // resume cancelled download
				Mp3:    true, // extract audio to MP3
			}
			title := video.Title + ".mp4"
			video.Download(0, title, option)

		}
	}

}

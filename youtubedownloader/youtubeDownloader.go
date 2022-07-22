package youtubedownloader

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	KB float64 = 1 << (10 * (iota + 1))
	MB
	GB
)

// Option Download options
type Option struct {
	Resume bool // resume failed or cancelled download
	Rename bool // rename output file using video title
	Mp3    bool // extract audio using ffmpeg
}

func (video *Video) Download(index int, filename string, option *Option) error {
	var (
		out    *os.File
		err    error
		offset int64
		length int64
	)

	if option.Resume {
		// Resume download from last known offset
		flags := os.O_WRONLY | os.O_CREATE
		out, err = os.OpenFile(filename, flags, 0644)
		if err != nil {
			return fmt.Errorf("unable to open file %q: %s", filename, err)
		}
		offset, err = out.Seek(0, io.SeekEnd)
		if err != nil {
			return fmt.Errorf("unable to seek file %q: %s", filename, err)
		}
		fmt.Printf("Resuming from offset %d (%s)\n", offset, abbr(offset))

	} else {
		// Start new download
		flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
		out, err = os.OpenFile(filename, flags, 0644)
		if err != nil {
			return fmt.Errorf("unable to write to file %q: %s", filename, err)
		}
	}

	defer out.Close()
	url := video.Formats[index].Url
	video.Filename = filename

	// GetMetaInformation video content length
	if len(url) == 0 && len(video.Formats[index].ContentLength) == 0 {
		return fmt.Errorf("head request failed: %s", err)
	} else {
		video.ContentLength = video.Formats[0].ContentLength
		if size := video.Formats[index].ContentLength; len(size) == 0 {
			return errors.New("Content-Length header is missing")
		} else if length, err = strconv.ParseInt(size, 10, 64); err != nil {
			return fmt.Errorf("invalid Content-Length: %s", err)
		}

		if length <= offset {
			fmt.Println("Video file is already downloaded.")
			return nil
		}
	}

	if length > 0 {
		go printProgress(out, offset, length)
	}

	// Not using range requests by default, because YouTube is throttling
	// download speed. Using a single GET request for max speed.
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("request failed: %s", err)
	}
	defer resp.Body.Close()

	if length, err = io.Copy(out, resp.Body); err != nil {
		return err
	}

	// Download stats
	duration := time.Now().Sub(start)
	speed := float64(length) / float64(duration/time.Second)
	if duration > time.Second {
		duration -= duration % time.Second
	} else {
		speed = float64(length)
	}

	if option.Rename {
		// Rename output file using video title
		wspace := regexp.MustCompile(`\W+`)
		fname := strings.Split(filename, ".")[0]
		ext := filepath.Ext(filename)
		title := wspace.ReplaceAllString(video.Title, "-")
		if len(title) > 64 {
			title = title[:64]
		}
		title = strings.TrimRight(strings.ToLower(title), "-")
		video.Filename = fmt.Sprintf("%s-%s%s", fname, title, ext)
		if err := os.Rename(filename, video.Filename); err != nil {
			fmt.Println("Failed to rename output file:", err)
		}
	}

	// Extract audio from downloaded video using ffmpeg
	//TODO CURRENTLY DOES NOT WORK!
	if option.Mp3 {
		if err := out.Close(); err != nil {
			fmt.Println("Error:", err)
		}
		ffmpeg, err := exec.LookPath("ffmpeg")
		if err != nil {
			fmt.Println("ffmpeg not found")
		} else {
			fmt.Println("Extracting audio ..")
			fname := video.Filename
			mp3 := strings.TrimRight(fname, filepath.Ext(fname)) + ".mp3"
			cmd := exec.Command(ffmpeg, "-y", "-loglevel", "quiet", "-i", fname, "-vn", mp3)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Println("Failed to extract audio:", err)
			} else {
				fmt.Println()
				fmt.Println("Extracted audio:", mp3)
			}
		}
	}

	fmt.Printf("Download duration: %s\n", duration)
	fmt.Printf("Average speed: %s/s\n", abbr(int64(speed)))

	return nil
}
func abbr(byteSize int64) string {
	size := float64(byteSize)
	switch {
	case size > GB:
		return fmt.Sprintf("%.1fGB", size/GB)
	case size > MB:
		return fmt.Sprintf("%.1fMB", size/MB)
	case size > KB:
		return fmt.Sprintf("%.1fKB", size/KB)
	}
	return fmt.Sprintf("%d", byteSize)
}

func printProgress(out *os.File, offset, length int64) {
	var clear string
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	start := time.Now()
	tail := offset

	var err error
	for now := range ticker.C {
		duration := now.Sub(start)
		duration -= duration % time.Second
		offset, err = out.Seek(0, io.SeekCurrent)
		if err != nil {
			return
		}
		speed := offset - tail
		percent := int(100 * offset / length)
		progress := fmt.Sprintf(
			"%s%s\t %s/%s\t %d%%\t %s/s",
			clear, duration, abbr(offset), abbr(length), percent, abbr(speed))
		fmt.Println(progress)
		tail = offset
		if tail >= length {
			break
		}
		if clear == "" {
			switch runtime.GOOS {
			case "darwin":
				clear = "\033[A\033[2K\r"
			case "linux":
				clear = "\033[A\033[2K\r"
			case "windows":
			}
		}
	}
}

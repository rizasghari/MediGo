package main

import "github.com/rizasghari/medigo/cmd"

// "flag"
// "fmt"
// "log"
// "path/filepath"
// "time"

// "github.com/rizasghari/medigo/internal/ffmpeg"

func main() {

    cmd.Execute()
    // act := flag.String("action", "convert", "Action to perform: convert or thumbnail")

    // videoPath := flag.String("input", "", "Path to the input video file")
    // timestamp := flag.Int("timestamp", 0, "Timestamp (in seconds) to capture the thumbnail")
    // outputDir := flag.String("output", ".", "Directory to save the thumbnail image")

    // flag.Parse()

    // ffmpeg := ffmpeg.New()

    // switch *act {   
    // case "convert":
    //     // err := ffmpeg.Convert(*videoPath, *outputDir)
    //     // if err != nil {
    //     //     log.Fatal("Error converting video:", err)
    //     // }

    //     // fmt.Println("Video converted successfully")
    // case "thumbnail":
    //     if *videoPath == "" || *timestamp < 0 {
    //         log.Fatal("You must specify a valid video path and a non-negative timestamp.")
    //     }
    //     time := fmt.Sprintf("%d", time.Now().Unix())
    //     outputPath := filepath.Join(*outputDir, fmt.Sprintf("thumbnail-%s.jpg", time))
    //     err := ffmpeg.GenerateThumbnail(*videoPath, *timestamp, outputPath)
    //     if err != nil {
    //         log.Fatal("Error generating thumbnail:", err)
    //     }
    
    //     fmt.Println("Thumbnail generated successfully at", outputPath)
    // } 
}


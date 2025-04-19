package main

import (
	"fmt"
	"log"
	"mcp-yt-download/mcp"
	"mcp-yt-download/youtubedownload"
	"os"

	"github.com/joho/godotenv"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error loading .env file")
	}
	// Load environment variables
	outputDir := os.Getenv("OUTPUT_DIR")
	fmt.Println("Output directory: ", outputDir)
	youtubeDownload := youtubedownload.NewYTDownloader(outputDir)

	s := mcp.NewMCPRouter(youtubeDownload)
	fmt.Println(" มันเริ่มต้นขึ้นแล้ว !!!!")
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

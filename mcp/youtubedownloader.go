package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"mcp-yt-download/entities"

	"github.com/mark3labs/mcp-go/mcp"
)

func YoutubeDownloaderTool() mcp.Tool {
	getVodLongData := mcp.NewTool("youtube_downloader",
		mcp.WithDescription("Download youtube video using only url of youtube video"),
		mcp.WithString("url",
			mcp.Required(),
			mcp.Description("The url of youtube example: https://www.youtube.com/watch?v=t9k1TCVCD2I"),
		),
		mcp.WithString("file_name",
			mcp.Description("Optional: Custom output filename"),
		),
		mcp.WithBoolean("audio_only",
			mcp.Description("If true, download audio only"),
		),
		mcp.WithBoolean("video_only",
			mcp.Description("If true, download video only (no audio)"),
		),
		mcp.WithString("format",
			mcp.Description("Specific format code (e.g., \"best\", \"bestaudio\", \"bestvideo\")"),
		),
		mcp.WithString("quality",
			mcp.Description("Optional: Quality selection (e.g., \"720p\", \"1080p\")"),
		),
		mcp.WithBoolean("embed_subs",
			mcp.Description("If true, embed subtitles if available"),
		),
	)
	return getVodLongData
}

func (h *Handler) YoutubeDownloader(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fmt.Println(" AI มันใช้งาน API เส้นนี้ละ")
	argsBytes, err := json.Marshal(request.Params.Arguments)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), err
	}
	var opts entities.DownloadOptions
	if err := json.Unmarshal(argsBytes, &opts); err != nil {
		return mcp.NewToolResultError(err.Error()), err
	}

	err = h.DownloaderInterface.Download(opts)

	if err != nil {
		return mcp.NewToolResultError(err.Error()), err
	}
	fmt.Println("ทำงานเสร็จละ ")
	return mcp.NewToolResultText("สวัสดีจ้า AI ฉัน ทำงานเสร็จแล้วนะจ๊ะ"), nil
}

type DownloaderInterface interface {
	Download(options entities.DownloadOptions) error
}

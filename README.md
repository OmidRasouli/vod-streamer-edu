# Building a VOD Platform with Go and FFmpeg (Tutorial Series, Article 3)

<p align="center">
  <img src="public/cover.jpg" alt="VOD Streamer EDU Cover" width="100%"/>
</p>

<p align="center">
  <img src="public/logo.png" alt="VOD Streamer EDU" width="200"/>
</p>

<p align="center">
  <a href="https://medium.com/@o.rasouli92"><img src="https://img.shields.io/badge/Medium-@o.rasouli92-violet?logo=medium"></a>
  <img src="https://img.shields.io/badge/go-1.24-blue?logo=Go&logoColor=white">
  <img src="https://img.shields.io/badge/license-MIT-green">
  <img src="https://img.shields.io/badge/architecture-clean-blueviolet" alt="Clean Architecture"/>
  <img src="https://img.shields.io/badge/ffmpeg-enabled-brightgreen?logo=ffmpeg&logoColor=white" alt="FFmpeg"/>
  <img src="https://img.shields.io/badge/HLS-activated-orange?logo=streamlit&logoColor=white" alt="HLS Activated"/>
</p>

---

## Overview

Welcome to **VOD Streamer EDU**, an educational project demonstrating how to build a Video On Demand (VOD) streaming backend in Go using Clean Architecture and FFmpeg.

This repository accompanies the **third article** in my Medium tutorial series:  
👉 [Read more on Medium – @o.rasouli92](https://medium.com/@o.rasouli92)

---

## Table of Contents
- [Building a VOD Platform with Go and FFmpeg (Tutorial Series, Article 3)](#building-a-vod-platform-with-go-and-ffmpeg-tutorial-series-article-3)
  - [Overview](#overview)
  - [Table of Contents](#table-of-contents)
  - [✨ Features](#-features)
  - [🚀 Quick Start](#-quick-start)
  - [📤 Upload a Video](#-upload-a-video)
  - [📺 Test HLS Streaming](#-test-hls-streaming)
  - [🗂️ Project Structure](#️-project-structure)
  - [🔌 API Reference](#-api-reference)
  - [⚙️ Configuration](#️-configuration)
  - [📖 About the Series](#-about-the-series)
  - [📝 License](#-license)

---

## ✨ Features

- Clean Architecture project structure in Go
- Video upload via REST API
- Automatic transcoding to HLS using FFmpeg
- HLS streaming endpoints for playback
- Ready for extension and learning

---

## 🚀 Quick Start

1. **Clone the repository:**
   ```bash
   git clone https://github.com/OmidRasouli/vod-streamer-edu.git
   cd vod-streamer-edu
   ```

2. **Add a test video:**  
   Place a sample video (e.g., `.mp4`, `.mkv`) in `public/test/`.  
   Example provided:  
   `public/test/Pixar.Popcorn.S01E04.1080p.WEB-DL.mkv`

3. **Run the server:**
   ```bash
   go run cmd/server/main.go
   ```

---

## 📤 Upload a Video

Upload a video for processing and HLS conversion:

```bash
curl -X POST http://localhost:8080/upload \
  -F "video=@/vod-streamer-edu/public/test/Pixar.Popcorn.S01E04.1080p.WEB-DL.mkv"
```

> **Tip:** Use the provided curl command to quickly test video uploads!

---

## 📺 Test HLS Streaming

1. After upload, note the returned video UUID.
2. Construct your HLS master playlist URL:
   ```
   http://localhost:8080/stream/{id}/master.m3u8
   ```
   Replace `{id}` with your video UUID.
3. Open [hls.js demo player](https://hlsjs.video-dev.org/demo/?src=) and paste your playlist URL after `src=`, e.g.:
   ```
   https://hlsjs.video-dev.org/demo/?src=http://localhost:8080/stream/your-uuid/master.m3u8
   ```

---

## 🗂️ Project Structure

```
cmd/server/                 # Application entrypoint
configs/                    # Configuration files
internal/
  controller/http/          # HTTP handlers (API endpoints)
  domain/                   # Domain models and ports (interfaces)
  entity/                   # Core entities (e.g., Path)
  infrastructure/           # FFmpeg and storage implementations
  usecase/                  # Application use cases (business logic)
public/                     # Static files and test videos
```

---

## 🔌 API Reference

| Endpoint                                 | Method | Description                                 |
|-------------------------------------------|--------|---------------------------------------------|
| `/upload`                                | POST   | Upload a video file (multipart form, field: `video`) |
| `/stream/{id}/master.m3u8`               | GET    | Get the HLS master playlist for a video     |
| `/stream/{id}/{quality}/{file}`          | GET    | Get a specific HLS segment or playlist      |
| `/health`                                | GET    | Health check endpoint                       |

**Example Upload Response:**
```json
{
  "message": "Video uploaded successfully",
  "filename": "Pixar.Popcorn.S01E04.1080p.WEB-DL.mkv"
}
```

---

## ⚙️ Configuration

- Main configuration file: `configs/config.yaml`
- You can set server port, storage paths, and FFmpeg resource limits here.
- Example:
  ```yaml
  server:
    port: 8080
  storage:
    raw_video_path: ./videos
  ffmpeg:
    cpu_core_limit: 0.8
    cpu_core_request: 0.5
  ```

---

## 📖 About the Series

This is the third article in a hands-on series:

- **Article 1:** [Building a VOD Platform with Go and FFmpeg — Part 1: Foundations](https://medium.com/@o.rasouli92/building-a-vod-platform-with-go-and-ffmpeg-part-1-foundations-771e1e14f79b)
- **Article 2:** [Building a VOD Platform with Go and FFmpeg — Part 2: Deep Dive into HLS and M3U8 Playlists](https://medium.com/@o.rasouli92/building-a-vod-platform-with-go-and-ffmpeg-part-2-deep-dive-into-hls-and-m3u8-playlists-29ffbad7a20a)
- **Article 3 (this branch):** [Building a VOD Platform with Go and FFmpeg - Part 3: Upload, Transcode & Serve](https://medium.com/@o.rasouli92)

Read the full series: [@o.rasouli92 on Medium](https://medium.com/@o.rasouli92)

---

## 📝 License

MIT License © 2025 Omid Rasouli

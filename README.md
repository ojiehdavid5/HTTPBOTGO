# HTTPBOTGO

A Telegram bot built with Go that provides OCR (Optical Character Recognition) capabilities using both OCR.Space API and Tesseract.

## Features

- **Telegram Bot Integration**: Seamless interaction via Telegram
- **OCR Processing**: Extract text from images using:
  - OCR.Space API
  - Tesseract (via gosseract)
- **File Download**: Automatically download images from Telegram
- **Interactive Keyboard**: User-friendly command interface

## Prerequisites

- Go 1.24.1 or higher
- Docker (optional, for containerized deployment)
- Telegram Bot Token (from [BotFather](https://t.me/botfather))
- OCR.Space API Key (free tier available at [ocr.space](https://ocr.space/ocrapi))

## Installation

### Local Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd HTTPBOTGO
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   Create a `.env` file in the project root:
   ```
   TELEGRAM_BOT_TOKEN=your_bot_token_here
   OCR_SPACE_API_KEY=your_ocr_api_key_here
   ```

4. **Build the application**
   ```bash
   go build -o main main.go
   ```

5. **Run the bot**
   ```bash
   ./main
   ```

### Docker Setup

For detailed Docker deployment instructions, see [README.Docker.md](README.Docker.md).

**Quick start with Docker Compose:**
```bash
docker compose up
```

## Configuration

The application requires the following environment variables:

| Variable | Description | Required |
|----------|-------------|----------|
| `TELEGRAM_BOT_TOKEN` | Your Telegram bot token from BotFather | Yes |
| `OCR_SPACE_API_KEY` | API key for OCR.Space service | Yes |

## Usage

1. Start the bot
2. Send an image to the bot
3. The bot will extract text from the image using OCR
4. Receive the extracted text as a response

## Project Structure

```
HTTPBOTGO/
├── main.go              # Main bot logic and OCR integration
├── go.mod              # Go module definition
├── go.sum              # Go module checksums
├── Dockerfile          # Container image definition
├── compose.yaml        # Docker Compose configuration
├── README.md           # This file
└── README.Docker.md    # Docker-specific documentation
```

## Dependencies

- `github.com/go-telegram-bot-api/telegram-bot-api/v5` - Telegram Bot API wrapper
- `github.com/joho/godotenv` - Environment variable management
- `github.com/otiai10/gosseract/v2` - Tesseract OCR binding

## Development

### Building
```bash
go build -o main main.go
```

### Running Tests
```bash
go test ./...
```

## API Integration

### OCR.Space API

The bot uses OCR.Space API for optical character recognition. The service supports:
- Multiple languages (currently configured for English)
- Free tier with reasonable request limits
- Higher accuracy for printed text

### Telegram Bot API

Communicates with Telegram to:
- Receive messages and files
- Send responses
- Manage keyboard interactions

## Troubleshooting

- **Bot not responding**: Verify `TELEGRAM_BOT_TOKEN` is correct
- **OCR errors**: Check `OCR_SPACE_API_KEY` and API rate limits
- **File download issues**: Ensure bot has proper permissions

## License

[Add your license here]

## Contributing

[Add contribution guidelines here]

## Support

For issues and questions, please create an issue in the repository.

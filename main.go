package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func downloadFile(bot *tgbotapi.BotAPI, fileID, destPath string) error {
	file, err := bot.GetFile(tgbotapi.FileConfig{FileID: fileID})
	if err != nil {
		return err
	}

	fileURL := file.Link(bot.Token)
	resp, err := http.Get(fileURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// OCR.Space API integration
func extractTextWithOCRSpace(imgPath, apiKey string) (string, error) {
	file, err := os.Open(imgPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(imgPath))
	if err != nil {
		return "", err
	}
	io.Copy(part, file)

	writer.WriteField("language", "eng")
	writer.WriteField("apikey", apiKey)
	writer.Close()

	req, err := http.NewRequest("POST", "https://api.ocr.space/parse/image", body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		ParsedResults []struct {
			ParsedText string `json:"ParsedText"`
		} `json:"ParsedResults"`
		IsErroredOnProcessing bool   `json:"IsErroredOnProcessing"`
		ErrorMessage          string `json:"ErrorMessage"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if result.IsErroredOnProcessing {
		return "", fmt.Errorf("OCR error: %s", result.ErrorMessage)
	}

	if len(result.ParsedResults) == 0 {
		return "No text found", nil
	}

	return result.ParsedResults[0].ParsedText, nil
}

func main() {
	bot, err := tgbotapi.NewBotAPI("7400994820:AAF3nDB3wwYQP_Cu3v6QeF2uWfDUyvG7A80")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	ocrAPIKey := "K83169488088957" // Replace with your actual OCR.Space key

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Photo != nil {
			fileID := update.Message.Photo[len(update.Message.Photo)-1].FileID
			imgPath := "downloaded_image.jpg"

			// Download image
			if err := downloadFile(bot, fileID, imgPath); err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Failed to download image."))
				continue
			}

			// Extract text via OCR.Space
			text, err := extractTextWithOCRSpace(imgPath, ocrAPIKey)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Failed to extract text: "+err.Error()))
				continue
			}

			// Send extracted text
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Extracted Text:\n"+text)
			bot.Send(msg)

			// Clean up
			os.Remove(imgPath)
			continue
		}

		// Default message
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Send a photo and I'll extract the text for you.")
		bot.Send(msg)
	}
}

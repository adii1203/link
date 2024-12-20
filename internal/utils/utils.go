package utils

import (
	"crypto/rand"
	"math/big"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/adii1203/link/internal/models"
)

func GenerateKey(n int) string {
	result := make([]byte, n)
	base := big.NewInt(int64(len(base62Chars)))

	for i := 0; i < n; i++ {
		randomNum, err := rand.Int(rand.Reader, base)
		if err != nil {
			return ""
		}
		result[i] = base62Chars[randomNum.Int64()]
	}

	return string(result)
}

func GetMetadata(url string) models.Metadata {
	res, err := http.Get(url)
	if err != nil {
		return models.Metadata{
			Title:       "",
			Description: "",
			Image:       "",
		}
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return models.Metadata{
			Title:       "",
			Description: "",
			Image:       "",
		}
	}

	title := getMetaData(doc, "meta[property='og:title']", "content")
	desc := getMetaData(doc, "meta[property='og:description']", "content")
	image := getMetaData(doc, "meta[property='og:image']", "content")

	if title == "" {
		title = doc.Find("title").Text()
	}
	if desc == "" {
		desc = getMetaData(doc, "meta[name='description']", "content")
	}

	return models.Metadata{
		Title:       title,
		Description: desc,
		Image:       image,
	}

}

func getMetaData(doc *goquery.Document, selector, attr string) string {
	val, exists := doc.Find(selector).Attr(attr)
	if !exists {
		return ""
	}

	return strings.TrimSpace(val)
}

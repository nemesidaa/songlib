package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	*http.Client
	APIUrl string
}

type SongDetail struct {
	ReleaseDate time.Time `json:"releaseDate"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
}

func NewClient(url string) *Client {
	return &Client{
		&http.Client{
			Transport: http.DefaultTransport,
			Timeout:   5 * time.Second,
		},
		url,
	}
}

func (s *SongDetail) UnmarshalJSON(data []byte) error {
	// Создаем временную структуру для хранения промежуточных данных
	type Alias SongDetail
	aux := &struct {
		ReleaseDate string `json:"releaseDate"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}

	// Декодируем JSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Парсим дату в формате "16.07.2006"
	layout := "02.01.2006" // Формат даты
	releaseDate, err := time.Parse(layout, aux.ReleaseDate)
	if err != nil {
		return err
	}
	s.ReleaseDate = releaseDate
	return nil
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	return c.Client.Do(req)
}

func (c *Client) GetDataFromAPI(group, song string) (*SongDetail, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/info?group=%s&song=%s", c.APIUrl, group, song), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var songDetail *SongDetail
	if err := json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
		return nil, err
	}

	return songDetail, nil
}

// Было создано для того, чтобы имитировать ответ внешнего API
// Не любитель писать тесты для работы с чёрным ящиком, поэтому могут возникнуть приколы, но раз уж надо,
// то как говорится - хозяин-барин.
func (c *Client) doMock() (*http.Response, error) {
	// Создаем JSON тело, имитирующее ответ внешнего API
	responseBody := `{
		"releaseDate": "16.07.2006",
		"text": "Ooh baby, don't you know I suffer?",
		"link": "https://www.youtube.com/watch?v=Xsp3_a-PMTw"
	}`

	// Преобразуем строку JSON в io.Reader и создаем io.ReadCloser
	body := io.NopCloser(bytes.NewReader([]byte(responseBody)))

	// Возвращаем имитированный http.Response
	return &http.Response{
		StatusCode: 200,
		Body:       body,
		Header:     make(http.Header),
	}, nil
}

func (c *Client) DataMock() (*SongDetail, error) {
	// Создаем JSON тело, имитирующее ответ внешнего API
	resp, _ := c.doMock()

	defer resp.Body.Close()

	var songDetail *SongDetail
	if err := json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
		return nil, err
	}
	return songDetail, nil

}

package deeplapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"translate-remember-repeat/internal/domain"
)

func (d *deepLStruct) SendRequest(ctx context.Context, text []string) (*domain.ResponseOfDeepL, error) {

	req, err := d.createRequest(ctx, text)
	if err != nil {
		return nil, err
	}

	client := http.Client{Timeout: 5 * time.Second}

	resp, err := d.doRequestParseResponse(client, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//

func (d *deepLStruct) createRequest(ctx context.Context, text []string) (*http.Request, error) {

	reqBeforeMarshaling := domain.RequestToDeepL{
		Text:       text,
		TargetLang: domain.TargetLang,
	}

	body, err := json.Marshal(reqBeforeMarshaling)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", domain.Marshaling, err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		d.requestMethod,
		d.defaultLink,
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", domain.CreatingRequest, err)
	}

	req.Header.Set(domain.HeaderKeyContentType, d.contentType)
	req.Header.Set(domain.Authorization, d.deepLAuthKey)

	return req, nil
}

func (d *deepLStruct) doRequestParseResponse(
	client http.Client,
	req *http.Request,
) (
	*domain.ResponseOfDeepL,
	error,
) {

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", domain.ClientDo, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body) // status code is 451, "Unavailable in your region"
		fmt.Println(string(body))        // so need to use VPN every launch )))))))))
		return nil, fmt.Errorf("%w, must be 200, it's %d", domain.StatusCode, resp.StatusCode)
	}

	var deepLResp domain.ResponseOfDeepL

	if err = json.NewDecoder(resp.Body).Decode(&deepLResp); err != nil {
		return nil, fmt.Errorf("%w: %w", domain.Decoding, err)
	}
	return &deepLResp, nil
}

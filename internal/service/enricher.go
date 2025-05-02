package service

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

func fetchAge(name string) (int, error) {
	url := fmt.Sprintf("https://api.agify.io/?name=%s", name)
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch age: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Error("failed to fetch age", "status_code", resp.StatusCode)
		return 0, fmt.Errorf("failed to fetch age: %s", resp.Status)
	}

	slog.Info("api.genderize.io",
		"X-Rate-Limit-Remaining", resp.Header.Get("X-Rate-Limit-Remaining"),
	)

	var result struct {
		Age int `json:"age"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	return result.Age, nil
}

func fetchGender(name string) (string, error) {
	url := fmt.Sprintf("https://api.genderize.io/?name=%s", name)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch age: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Error("failed to fetch age", "status_code", resp.StatusCode)
		return "", fmt.Errorf("failed to fetch nation: %d", resp.StatusCode)

	}

	slog.Info("api.genderize.io",
		"X-Rate-Limit-Remaining", resp.Header.Get("X-Rate-Limit-Remaining"),
	)

	var result struct {
		Gender string `json:"gender"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Gender, nil
}

func fetchNationality(name string) (string, error) {
	url := fmt.Sprintf("https://api.nationalize.io/?name=%s", name)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch nation: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Error("failed to fetch nation", "status_code", resp.StatusCode)
		return "", fmt.Errorf("failed to fetch nation: %d", resp.StatusCode)

	}

	slog.Info("api.genderize.io",
		"X-Rate-Limit-Remaining", resp.Header.Get("X-Rate-Limit-Remaining"),
	)

	var result struct {
		Country []struct {
			CountryID string `json:"country_id"`
		} `json:"country"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result.Country) == 0 {
		return "", fmt.Errorf("no country found")
	}

	return result.Country[0].CountryID, nil
}

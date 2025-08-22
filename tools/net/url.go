package net

import (
	"errors"
	"net/url"
)

func RawURLGetParam(raw, key string) (string, error) {
	r, err := url.Parse(raw)
	if err != nil {
		return "", err
	}

	m := r.Query()
	if v, ok := m[key]; ok && len(v) > 0 {
		return v[0], nil
	}
	return "", errors.New("no param")
}
func RawURLGetParams(raw, key string) ([]string, error) {
	r, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}

	m := r.Query()
	if v, ok := m[key]; ok {
		return v, nil
	}
	return nil, errors.New("no param")
}
func RawURLGetAllParams(raw string) (map[string][]string, error) {
	r, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}
	m := r.Query()
	return m, nil
}
func RawURLAddParam(raw, key, value string) (string, error) {
	r, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	m := r.Query()
	m.Add(key, value)
	r.RawQuery = m.Encode()
	return r.String(), nil
}
func RawURLAddParams(raw string, params map[string]string) (string, error) {
	r, err := url.Parse(raw)
	if err != nil {
		return "", err
	}

	m := r.Query()
	for k, v := range params {
		m.Add(k, v)
	}
	r.RawQuery = m.Encode()
	return r.String(), nil
}
func RawURLDelParams(raw string, keys ...string) (string, error) {
	r, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	m := r.Query()
	for _, v := range keys {
		m.Del(v)
	}
	r.RawQuery = m.Encode()
	return r.String(), nil
}
func RawURLSetParam(raw, key, value string) (string, error) {
	r, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	m := r.Query()
	m.Set(key, value)
	r.RawQuery = m.Encode()
	return r.String(), nil
}

func RawURLSetParams(rawURL string, params map[string]string) (string, error) {
	r, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	m := r.Query()
	for k, v := range params {
		m.Set(k, v)
	}
	r.RawQuery = m.Encode()
	return r.String(), nil
}

func RawURLGetDomain(rawURL string) (string, error) {
	r, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	return r.Hostname(), nil
}

func RawURLGetPort(raw string) (string, error) {
	r, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	return r.Port(), nil
}
func QueryGetParam(query, key string) (string, error) {
	queries, err := url.ParseQuery(query)
	if err != nil {
		return "", err
	}
	if v, ok := queries[key]; ok && len(v) > 0 {
		return v[0], nil
	}
	return "", errors.New("no param")
}
func QueryGetParams(query, key string) ([]string, error) {
	queries, err := url.ParseQuery(query)
	if err != nil {
		return nil, err
	}
	if v, ok := queries[key]; ok {
		return v, nil
	}
	return nil, errors.New("no param")
}

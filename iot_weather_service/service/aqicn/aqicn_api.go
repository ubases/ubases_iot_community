// Copyright 2022 Brian J. Downs
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aqicn

import (
	"errors"
	"net/http"
)

var (
	errInvalidOption     = errors.New("invalid option")
	errInvalidHttpClient = errors.New("invalid http client")
)

// DataUnits represents the character chosen to represent the temperature notation
var (
	//f155ca0e050a7cd07265a40f05ae62db  60ed89a1
	//f155ca0e050a7cd07265a40f05ae62db60ed89a1
	baseURL = "https://api.waqi.info/feed/%s/?token=%s"
)

// Config will hold default settings to be passed into the
// "NewCurrent, NewForecast, etc}" functions.
type Config struct {
	Lang     string // should reference a key in the LangCodes map
	TokenKey string // API Token for connecting to the OWM
}

// CheckAPIKeyExists will see if an API key has been set.
func (c *Config) CheckAPIKeyExists() bool { return len(c.TokenKey) > 1 }

// Settings holds the client settings
type Settings struct {
	client *http.Client
}

// NewSettings returns a new Setting pointer with default http client.
func NewSettings() *Settings {
	return &Settings{
		client: http.DefaultClient,
	}
}

// Optional client settings
type Option func(s *Settings) error

// WithHttpClient sets custom http client when creating a new Client.
func WithHttpClient(c *http.Client) Option {
	return func(s *Settings) error {
		if c == nil {
			return errInvalidHttpClient
		}
		s.client = c
		return nil
	}
}

// setOptions sets Optional client settings to the Settings pointer
func setOptions(settings *Settings, options []Option) error {
	for _, option := range options {
		if option == nil {
			return errInvalidOption
		}
		err := option(settings)
		if err != nil {
			return err
		}
	}
	return nil
}

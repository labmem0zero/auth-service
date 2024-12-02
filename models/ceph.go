package models

type UploadTask struct {
	ReqID  string   `json:"request_id"`
	Export []Export `json:"export"`
}

type Export struct {
	Target string `json:"target"`
	Key    string `json:"key"`
}

type UploadTaskWithProcessing struct {
	ReqID      string                `json:"request_id"`
	Operations []ProcessingOperation `json:"operations"`
}

type ProcessingOperation struct {
	Transforms []any     `json:"transforms"`
	Convert    Convert   `json:"convert"`
	Transmuxe  Transmuxe `json:"transmuxe,omitempty"`
	Export     []Export  `json:"export"`
}

type Transmuxe struct {
	Container string `json:"container"`
}

type Convert struct {
	To      string  `json:"to"`
	Options Options `json:"options"`
}

type Options struct {
	Quality int `json:"quality"`
}

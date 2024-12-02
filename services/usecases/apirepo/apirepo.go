package apirepo

type FileServer interface {
	Upload(reqID string, file any, path string) (url string, err error)
	Download(reqID string, url string) (file []byte, err error)
}

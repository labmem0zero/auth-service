package middlewares

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (mw *Middlewares) MiddlewareRequestLogging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		reqID := uuid.New().String()
		ctx := context.WithValue(r.Context(), "reqID", reqID)
		r = r.WithContext(ctx)

		b, err := readReqBody(r)
		if err != nil {
			mw.l.Error(reqID, err)
		} else {
			var data string
			ct := r.Header.Get("Content-Type")
			switch ct {
			case "application/json":
				fallthrough
			case "plain/text":
				if len(b) < 10_000 {
					data = string(b)
				}
			default:
				data = "Content-Type is " + ct
			}
			mw.l.Debug(reqID, fmt.Sprintf("req len = %v, data = %v", len(b), data))
		}
		h.ServeHTTP(w, r)
		spent := time.Now().Sub(start).Seconds()
		mw.l.Debug(reqID, fmt.Sprintf("seconds spent: %v", spent))
	})
}

func GetReqID(r *http.Request) string {
	reqID, ok := r.Context().Value("reqID").(string)
	if !ok || reqID == "" {
		return "unidentified"
	} else {
		return reqID
	}
}

func readReqBody(r *http.Request) ([]byte, error) {
	buf := &bytes.Buffer{}
	if _, err := buf.ReadFrom(r.Body); err != nil {
		return nil, err
	}
	if err := r.Body.Close(); err != nil {
		return nil, err
	}
	r.Body = io.NopCloser(buf)

	return buf.Bytes(), nil
}

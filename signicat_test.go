package signicat

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Returns a configured client for testing http calls. Handler function is set in test function
func setup() (*Client, *http.ServeMux, func()) {
	mux := http.NewServeMux()

	apiHandler := http.NewServeMux()
	apiHandler.Handle("/", mux)

	server := httptest.NewServer(apiHandler)

	client, err := NewClient(&http.Client{}, server.URL)
	if err != nil {
		panic(fmt.Sprintf("couldnt set up test client: %v", err))
	}

	return client, mux, server.Close
}

func TestSignatureService_CreateDocument(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
		assert.Equal(t, "/signature/documents", req.URL.Path)
		if _, err := io.WriteString(res, `{"documentId":"someDocumentId"}`); err != nil {
			t.Fatal(err)
		}
	})

	document, err := client.Signature.CreateDocument(&CreateDocumentRequest{})
	assert.NoError(t, err)
	assert.Equal(t, "someDocumentId", document.DocumentID)
}

func TestSignatureService_RetrieveDocument(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "/signature/documents/someDocumentId", req.URL.Path)
		if _, err := io.WriteString(res, `{"documentId":"someDocumentId"}`); err != nil {
			t.Fatal(err)
		}
	})

	document, err := client.Signature.RetrieveDocument("someDocumentId")
	assert.NoError(t, err)
	assert.Equal(t, "someDocumentId", document.DocumentID)
}

func TestSignatureService_RetrieveDocumentStatus(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "/signature/documents/someDocumentId/status", req.URL.Path)
		if _, err := io.WriteString(res, `{"documentStatus":"signed"}`); err != nil {
			t.Fatal(err)
		}
	})

	status, err := client.Signature.RetrieveDocumentStatus("someDocumentId")
	assert.NoError(t, err)
	assert.Equal(t, DocumentStatusSigned, status.DocumentStatus)
}

func TestSignatureService_RetrieveFile(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.Method, http.MethodGet)
		assert.Equal(t, "/signature/documents/someDocumentId/files", req.URL.Path)
		assert.Equal(t, "pades", req.URL.Query().Get("fileFormat"))
		assert.Equal(t, "true", req.URL.Query().Get("originalFileName"))
		if _, err := io.WriteString(res, "response"); err != nil {
			t.Fatal(err)
		}
	})

	var buf bytes.Buffer
	err := client.Signature.RetrieveFile("someDocumentId", FileFormatPades, true, &buf)
	assert.NoError(t, err)
	assert.Equal(t, "response", buf.String())
}

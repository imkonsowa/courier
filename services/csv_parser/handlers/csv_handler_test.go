package handlers

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"courier/courierpb"
)

type BasicResponse struct {
	Code    int
	Message string
	Success bool
}

type TestStruct struct {
	Engine *gin.Engine
	Grpc   struct {
		Client     courierpb.CourierServiceClient
		Connection *grpc.ClientConn
	}
}

var TestData TestStruct

func TestMain(m *testing.M) {
	TestData = TestStruct{
		Engine: gin.Default(),
	}

	client, connection := courierpb.NewCourierClient()
	handler := NewCsvHandler(client)

	TestData.Grpc.Connection = connection
	TestData.Grpc.Client = client
	TestData.Engine.POST("/upload", handler.ProcessParcels)

	exitVal := m.Run()

	connection.Close()

	os.Exit(exitVal)
}

func TestCsvHandler_ProcessParcelsFileMissing(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.Close()

	req, _ := http.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()

	TestData.Engine.ServeHTTP(w, req)

	resp := BasicResponse{}

	json.NewDecoder(w.Body).Decode(&resp)

	assert.Equal(t, resp.Code, http.StatusBadRequest)
}

func TestCsvHandler_ProcessParcelsEmptyFile(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	_, _ = writer.CreateFormFile("file", "test.csv")

	writer.Close()

	req, _ := http.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()

	TestData.Engine.ServeHTTP(w, req)

	resp := BasicResponse{}

	json.NewDecoder(w.Body).Decode(&resp)

	assert.Equal(t, resp.Code, http.StatusBadRequest)
}

func TestCsvHandler_ProcessParcelsServerErr(t *testing.T) {
	body := new(bytes.Buffer)
	/*writer := multipart.NewWriter(body)
	_, _ = writer.CreateFormFile("file", "test.csv")

	writer.Close()*/

	req, _ := http.NewRequest("POST", "/upload", body)
	// req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()

	TestData.Engine.ServeHTTP(w, req)

	resp := BasicResponse{}

	json.NewDecoder(w.Body).Decode(&resp)

	assert.Equal(t, resp.Code, http.StatusInternalServerError)
}

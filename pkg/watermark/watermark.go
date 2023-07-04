package watermark

import (
	"context"
	"net/http"

	"github.com/willwchan/go-microservice-app/internal"

	"github.com/go-kit/kit/log"
	"github.com/lithammer/shortuuid/v3"
)

type watermarkService struct{}

func NewService() Service { return &watermarkService{} }

func (w *watermarkService) Get(_ context.Context, filters ...internal.Filter) ([]internal.Document, error) {
	// queries the database using filters and returns the list of documents
	// return error if the filter(key) is invalid and also return error if no item found
	doc := internal.Document{
		Content: "book",
		Title:   "Harry Potter and Half Blood Prince",
		Author:  "J.K. Rowling",
		Topic:   "Fiction and Magic",
	}
	return []internal.Document{doc}, nil
}

func (w *watermarkService) Status(_ context.Context, ticketID string) (internal.Status, error) {
	// query db using the ticketID and return the document info
	// return erro rif ticketId is invalid or no docuemnt exists for that id
	return internal.InProgress, nil
}

func (w *watermarkService) Watermark(_ context.Context, ticketId, mark string) (int, error) {
	// update the db entry with watermark field as not empty
	// first check if the watermark status is alread InProgress, Started, or Finished
	// if yes ,then return invalid request
	// return error if no item found using ticketId
	return http.StatusOK, nil
}

func (w *watermarkService) AddDocument(_ context.Context, doc *internal.Document) (string, error) {
	// add the document entry in the db by calling the db service
	// return error if the doc is invalid and/or the database returns invalid entry error
	newTicketID := shortuuid.New()
	return newTicketID, nil
}

func (w *watermarkService) ServiceStatus(_ context.Context) (int, error) {
	logger.Log("Checking the service health...")
	return http.StatusOK, nil
}

var logger log.logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}

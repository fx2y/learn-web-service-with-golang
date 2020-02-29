package crawler

import (
	"context"
	"github.com/fx2y/learn-web-service-with-golang/internal/pipeline"
	"github.com/fx2y/learn-web-service-with-golang/internal/textindexer/index"
	"time"
)

type textIndexer struct {
	indexer Indexer
}

func newTextIndexer(indexer Indexer) *textIndexer {
	return &textIndexer{
		indexer: indexer,
	}
}

func (i *textIndexer) Process(ctx context.Context, p pipeline.Payload) (pipeline.Payload, error) {
	payload := p.(*crawlerPayload)

	doc := &index.Document{
		LinkID:    payload.LinkID,
		URL:       payload.URL,
		Title:     payload.Title,
		Content:   payload.TextContent,
		IndexedAt: time.Now(),
	}
	if err := i.indexer.Index(doc); err != nil {
		return nil, err
	}

	return p, nil
}

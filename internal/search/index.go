package search

import (
	"github.com/blevesearch/bleve/v2"
)

type IndexManager struct {
	Index bleve.Index
}

func NewIndexManager(indexPath string) (*IndexManager, error) {
	index, err := bleve.Open(indexPath)
	if err == bleve.ErrorIndexPathDoesNotExist {
		mapping := bleve.NewIndexMapping()
		index, err = bleve.New(indexPath, mapping)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	return &IndexManager{Index: index}, nil
}

func (m *IndexManager) IndexMovie(id string, movie interface{}) error {
	return m.Index.Index(id, movie)
}

func (m *IndexManager) Close() error {
	return m.Index.Close()
}

func (s *IndexManager) SearchByTerm(term string) ([]string, error) {
	query := bleve.NewMatchQuery(term)
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, err := s.Index.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	result := []string{}

	for _, hit := range searchResult.Hits {
		result = append(result, hit.ID)
	}

	return result, nil
}

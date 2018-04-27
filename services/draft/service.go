package draft

import (
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/luizalabs/escriba/util"
)

// ServiceIface interface
type ServiceIface interface {
	Add(url string) error
	Approve(url string) error
	ListPendingReviews() ([]Record, error)
	ListPendingPublications() ([]Record, error)
	MarkAsPublished(url string) error
}

// Service struct
type Service struct {
	db           *sql.DB
	minApprovals int
}

// New cria um novo service
func New(db *sql.DB, approvals int) *Service {
	return &Service{
		db:           db,
		minApprovals: approvals,
	}
}

// Add cria um draft no escriba
func (s *Service) Add(url string) error {
	logger := util.GetLogger().WithFields(logrus.Fields{
		"module":         "service",
		"operation_name": "drafts.add",
	})

	start := time.Now()
	_, err := s.db.Exec(insertSQL, url)
	latency := time.Since(start)

	if err != nil {
		logger.WithFields(logrus.Fields{
			"db": logrus.Fields{
				"statement": insertSQL,
				"params":    url,
				"error":     err.Error(),
				"latency":   int(latency / time.Millisecond),
			},
		}).Error()
		return err
	}

	logger.WithFields(logrus.Fields{
		"db": logrus.Fields{
			"statement": insertSQL,
			"params":    url,
			"latency":   int(latency / time.Millisecond),
		},
	}).Info()

	return nil
}

// Approve incrementa a quantidade de aprovações do rascunho em 1
func (s *Service) Approve(url string) error {
	logger := util.GetLogger().WithFields(logrus.Fields{
		"module":         "service",
		"operation_name": "drafts.approve",
	})

	start := time.Now()
	result, err := s.db.Exec(approveSQL, url)
	latency := time.Since(start)

	if err != nil {
		logger.WithFields(logrus.Fields{
			"db": logrus.Fields{
				"statement": insertSQL,
				"params":    url,
				"error":     err.Error(),
				"latency":   int(latency / time.Millisecond),
			},
		}).Error()
		return err
	}

	rows, _ := result.RowsAffected()
	logger = logger.WithFields(logrus.Fields{
		"db": logrus.Fields{
			"statement":       approveSQL,
			"params":          url,
			"latency":         int(latency / time.Millisecond),
			"updated_records": rows,
		},
	})

	if rows == 0 {
		logger.Error()
		return errors.New("no drafts updated")
	}

	logger.Info()
	return nil
}

// ListPendingReviews lista os artigos que ainda não obtiveram o valor mínimo de aprovações
func (s *Service) ListPendingReviews() ([]Record, error) {
	logger := util.GetLogger().WithFields(logrus.Fields{
		"module":         "service",
		"operation_name": "drafts.listPendingReviews",
	})

	items := []Record{}
	start := time.Now()
	result, err := s.db.Query(listReviewSQL, s.minApprovals)
	latency := time.Since(start)

	if err != nil {
		logger.WithFields(logrus.Fields{
			"db": logrus.Fields{
				"statement": listReviewSQL,
				"params":    s.minApprovals,
				"error":     err.Error(),
				"latency":   int(latency / time.Millisecond),
			},
		}).Error()
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var url string
		var approvals int
		var createdAt string

		result.Scan(&url, &approvals, &createdAt)

		items = append(items, Record{
			URL:       url,
			Approvals: approvals,
			CreatedAt: createdAt,
		})
	}

	logger.WithFields(logrus.Fields{
		"db": logrus.Fields{
			"statement": listReviewSQL,
			"params":    s.minApprovals,
			"latency":   int(latency / time.Millisecond),
			"records":   len(items),
		},
	}).Info()

	return items, nil
}

// ListPendingPublications lista os artigos que já sofreram revisão e estão prontos para publicação
func (s *Service) ListPendingPublications() ([]Record, error) {
	logger := util.GetLogger().WithFields(logrus.Fields{
		"module":         "service",
		"operation_name": "drafts.listPendingPublications",
	})

	items := []Record{}
	start := time.Now()
	result, err := s.db.Query(listPublicationSQL, s.minApprovals)
	latency := time.Since(start)

	if err != nil {
		logger.WithFields(logrus.Fields{
			"db": logrus.Fields{
				"statement": listPublicationSQL,
				"params":    s.minApprovals,
				"error":     err.Error(),
				"latency":   int(latency / time.Millisecond),
			},
		}).Error()
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var url string
		var approvals int
		var createdAt string

		result.Scan(&url, &approvals, &createdAt)

		items = append(items, Record{
			URL:       url,
			Approvals: approvals,
			CreatedAt: createdAt,
		})
	}

	logger.WithFields(logrus.Fields{
		"db": logrus.Fields{
			"statement": listReviewSQL,
			"latency":   int(latency / time.Millisecond),
			"records":   len(items),
		},
	}).Info()

	return items, nil
}

// MarkAsPublished atualiza a data de publicação de um artigo
func (s *Service) MarkAsPublished(url string) error {
	logger := util.GetLogger().WithFields(logrus.Fields{
		"module":         "service",
		"operation_name": "drafts.markAsPublished",
	})

	start := time.Now()
	result, err := s.db.Exec(publishSQL, url)
	latency := time.Since(start)

	if err != nil {
		logger.WithFields(logrus.Fields{
			"db": logrus.Fields{
				"statement": publishSQL,
				"params":    url,
				"error":     err.Error(),
				"latency":   int(latency / time.Millisecond),
			},
		}).Error()
		return err
	}

	rows, _ := result.RowsAffected()
	logger = logger.WithFields(logrus.Fields{
		"db": logrus.Fields{
			"statement":       approveSQL,
			"params":          url,
			"latency":         int(latency / time.Millisecond),
			"updated_records": rows,
		},
	})

	if rows == 0 {
		logger.Error()
		return errors.New("no draft updated")
	}

	logger.Info()
	return nil
}

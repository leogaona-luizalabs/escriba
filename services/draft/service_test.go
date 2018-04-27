package draft_test

import (
	"errors"
	"testing"

	"github.com/luizalabs/escriba/services/draft"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestAddOK(t *testing.T) {
	// setup do mock
	draftURL := "http://myblog.com/drafts/mydraft-1211"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		t.Fail()
	}
	defer db.Close()

	mock.ExpectExec("insert into draft").WithArgs(draftURL).WillReturnResult(sqlmock.NewResult(1, 1))

	// chamada do service
	service := draft.New(db, 1)
	err = service.Add(draftURL)

	// assertions
	if err != nil {
		t.Errorf("error returned from service: %s", err.Error())
		t.Fail()
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		t.Fail()
	}

}

func TestAddError(t *testing.T) {
	// setup do mock
	draftURL := "http://myblog.com/drafts/mydraft-1211"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		t.Fail()
	}
	defer db.Close()

	mock.ExpectExec("insert into draft").WithArgs(draftURL).WillReturnError(errors.New("table not found"))

	// chamada do service
	service := draft.New(db, 1)
	err = service.Add(draftURL)

	// assertions
	if err == nil {
		t.Error("no error returned from service")
		t.Fail()
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		t.Fail()
	}

}

func TestApproveOK(t *testing.T) {
	// setup do mock
	draftURL := "http://myblog.com/drafts/mydraft-1211"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		t.Fail()
	}
	defer db.Close()

	mock.ExpectExec("update draft d set approvals").WithArgs(draftURL).WillReturnResult(sqlmock.NewResult(1, 1))

	// chamada do service
	service := draft.New(db, 1)
	err = service.Approve(draftURL)

	// assertions
	if err != nil {
		t.Errorf("error returned from service: %s", err.Error())
		t.Fail()
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		t.Fail()
	}

}

func TestApproveError(t *testing.T) {
	// setup do mock
	draftURL := "http://myblog.com/drafts/mydraft-1211"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		t.Fail()
	}
	defer db.Close()

	mock.ExpectExec("update draft d set approvals").WithArgs(draftURL).WillReturnError(errors.New("table not found"))

	// chamada do service
	service := draft.New(db, 1)
	err = service.Approve(draftURL)

	// assertions
	if err == nil {
		t.Error("no error returned from service")
		t.Fail()
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		t.Fail()
	}

}

func TestApproveNoRowsUpdated(t *testing.T) {
	// setup do mock
	draftURL := "http://myblog.com/drafts/mydraft-1211"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		t.Fail()
	}
	defer db.Close()

	mock.ExpectExec("update draft d set approvals").WithArgs(draftURL).WillReturnResult(sqlmock.NewResult(0, 0))

	// chamada do service
	service := draft.New(db, 1)
	err = service.Approve(draftURL)

	// assertions
	if err == nil {
		t.Error("no error returned from service")
		t.Fail()
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		t.Fail()
	}

}

func TestListPendingReviewsOK(t *testing.T) {
	// setup do mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		t.Fail()
	}
	defer db.Close()
	columns := []string{"url", "approvals"}

	mock.ExpectQuery("select url, approvals from draft where published_at is null and approvals <").WillReturnRows(sqlmock.NewRows(columns).AddRow("url1", 0))

	// chamada do service
	service := draft.New(db, 1)
	items, err := service.ListPendingReviews()

	// assertions
	if err != nil {
		t.Errorf("error returned from service: %s", err.Error())
		t.Fail()
	}

	if len(items) != 1 {
		t.Errorf("wrong results slice lenght. Expected 1 but got %d", len(items))
		t.Fail()
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		t.Fail()
	}

}

func TestListPendingReviewsError(t *testing.T) {
	// setup do mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		t.Fail()
	}
	defer db.Close()
	mock.ExpectQuery("select url, approvals from draft where published_at is null and approvals <").WillReturnError(errors.New("table not found"))

	// chamada do service
	service := draft.New(db, 1)
	items, err := service.ListPendingReviews()

	// assertions
	if err == nil {
		t.Error("no error returned from service")
		t.Fail()
	}

	if items != nil {
		t.Error("non-nil slice returned")
		t.Fail()
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		t.Fail()
	}

}

func TestListPendingPublicationsOK(t *testing.T) {
	// setup do mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		t.Fail()
	}
	defer db.Close()
	columns := []string{"url", "approvals"}

	mock.ExpectQuery("select url, approvals from draft where published_at is null and approvals >=").WillReturnRows(sqlmock.NewRows(columns).AddRow("url1", 5))

	// chamada do service
	service := draft.New(db, 1)
	items, err := service.ListPendingPublications()

	// assertions
	if err != nil {
		t.Errorf("error returned from service: %s", err.Error())
		t.Fail()
	}

	if len(items) != 1 {
		t.Errorf("wrong results slice lenght. Expected 1 but got %d", len(items))
		t.Fail()
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		t.Fail()
	}

}

func TestListPendingPublicationsError(t *testing.T) {
	// setup do mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		t.Fail()
	}
	defer db.Close()
	mock.ExpectQuery("select url, approvals from draft where published_at is null and approvals >=").WillReturnError(errors.New("table not found"))

	// chamada do service
	service := draft.New(db, 1)
	items, err := service.ListPendingPublications()

	// assertions
	if err == nil {
		t.Error("no error returned from service")
		t.Fail()
	}

	if items != nil {
		t.Error("non-nil slice returned")
		t.Fail()
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		t.Fail()
	}

}

func TestMarkAsPublishedOK(t *testing.T) {
	// setup do mock
	draftURL := "http://myblog.com/drafts/mydraft-1211"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		t.Fail()
	}
	defer db.Close()

	mock.ExpectExec("update draft d set published_at = NOW()").WithArgs(draftURL).WillReturnResult(sqlmock.NewResult(1, 1))

	// chamada do service
	service := draft.New(db, 1)
	err = service.MarkAsPublished(draftURL)

	// assertions
	if err != nil {
		t.Errorf("error returned from service: %s", err.Error())
		t.Fail()
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		t.Fail()
	}

}

func TestMarkAsPublishedError(t *testing.T) {
	// setup do mock
	draftURL := "http://myblog.com/drafts/mydraft-1211"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		t.Fail()
	}
	defer db.Close()

	mock.ExpectExec("update draft d set published_at = NOW()").WithArgs(draftURL).WillReturnError(errors.New("table not found"))

	// chamada do service
	service := draft.New(db, 1)
	err = service.MarkAsPublished(draftURL)

	// assertions
	if err == nil {
		t.Error("no error returned from service")
		t.Fail()
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		t.Fail()
	}

}

func TestMarkAsPublishedNoRowsUpdated(t *testing.T) {
	// setup do mock
	draftURL := "http://myblog.com/drafts/mydraft-1211"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
		t.Fail()
	}
	defer db.Close()

	mock.ExpectExec("update draft d set published_at = NOW()").WithArgs(draftURL).WillReturnResult(sqlmock.NewResult(0, 0))

	// chamada do service
	service := draft.New(db, 1)
	err = service.MarkAsPublished(draftURL)

	// assertions
	if err == nil {
		t.Error("no error returned from service")
		t.Fail()
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		t.Fail()
	}

}

package core_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ufpblor/api/core"
	"github.com/ufpblor/api/test"
)

func TestHandlerZen(t *testing.T) {
	db, _ := test.InitDBTest()
	defer db.Session.Close()

	r, err := http.NewRequest("GET", "/zen", nil)
	test.AssertOk(t, err)

	w := httptest.NewRecorder()

	h := &core.Handler{DB: db.Session}

	core.Router(h).ServeHTTP(w, r)

	test.AssertEquals(t, "{\"message\":\"Keep it logically awesome.\"}\n", w.Body.String())
	test.AssertEquals(t, http.StatusOK, w.Code)
}

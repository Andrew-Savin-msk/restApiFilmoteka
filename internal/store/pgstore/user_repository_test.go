package pgstore_test

import (
	"testing"

	model "github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/model/user"
	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/store/pgstore"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	tu := model.TestUser(t)
	db, clear := pgstore.TestStore(t, dbPath)
	defer clear("users")

	st := pgstore.New(db)
	err := st.User().Create(tu)
	assert.NoError(t, err)
	assert.NotNil(t, tu)
}

package pgstore_test

import (
	"testing"

	user "github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/model/user"
	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/store/pgstore"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	tu := user.TestUser(t)
	db, clear := pgstore.TestStore(t, dbPath)
	defer clear("users")

	st := pgstore.New(db)
	err := st.User().Create(tu)
	assert.NoError(t, err)
	assert.NotNil(t, tu)
}

func TestFind(t *testing.T) {
	tu := user.TestUser(t)
	db, clear := pgstore.TestStore(t, dbPath)
	defer clear("users")

	st := pgstore.New(db)

	err := st.User().Create(tu)
	assert.NoError(t, err)
	assert.NotNil(t, tu)

	res, err := st.User().Find(tu.Id)
	assert.NoError(t, err)
	assert.Equal(t, res, tu)
}

func TestFindByEmail(t *testing.T) {
	tu := user.TestUser(t)
	db, clear := pgstore.TestStore(t, dbPath)
	defer clear("users")

	st := pgstore.New(db)

	err := st.User().Create(tu)
	assert.NoError(t, err)
	assert.NotNil(t, tu)

	res, err := st.User().FindByEmail(tu.Email)
	assert.NoError(t, err)
	assert.Equal(t, res, tu)
}

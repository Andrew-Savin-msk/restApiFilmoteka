package pgstore_test

import (
	"testing"

	actor "github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/model/actor"
	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/store/pgstore"
	"github.com/stretchr/testify/assert"
)

func TestActorCreate(t *testing.T) {
	ta := actor.TestActor(t)
	db, clear := pgstore.TestStore(t, dbPath)
	defer clear("actors")

	st := pgstore.New(db)
	err := st.Actor().Create(ta)
	assert.NoError(t, err)
	assert.NotNil(t, ta)
}

func TestFindActor(t *testing.T) {
	ta := actor.TestActor(t)
	db, clear := pgstore.TestStore(t, dbPath)
	defer clear("actors")

	st := pgstore.New(db)
	err := st.Actor().Create(ta)
	assert.NoError(t, err)
	assert.NotNil(t, ta)

	tmp, err := st.Actor().Find(ta.Id)
	assert.NoError(t, err)
	assert.NotNil(t, tmp)

	assert.Equal(t, tmp, ta)
}

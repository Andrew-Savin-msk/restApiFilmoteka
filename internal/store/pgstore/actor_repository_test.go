package pgstore_test

import (
	"testing"

	actor "github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/model/actor"
	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/store"
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

func TestDeleteActor(t *testing.T) {
	ta := actor.TestActor(t)
	db, clear := pgstore.TestStore(t, dbPath)
	defer clear("actors")

	st := pgstore.New(db)
	err := st.Actor().Create(ta)
	assert.NoError(t, err)
	assert.NotNil(t, ta)

	id, err := st.Actor().Delete(ta.Id)
	assert.NoError(t, err)
	assert.Equal(t, id, ta.Id)

	tmp, err := st.Actor().Find(ta.Id)
	assert.Equal(t, err, store.ErrRecordNotFound)
	assert.Nil(t, tmp)
}

// TODO: Test func
func TestOverwrightActor(t *testing.T) {
	ta := actor.TestActor(t)
	db, clear := pgstore.TestStore(t, dbPath)
	defer clear("actors")

	st := pgstore.New(db)
	err := st.Actor().Create(ta)
	assert.NoError(t, err)
	assert.NotNil(t, ta)

	err = st.Actor().Overwright(ta)
	assert.NoError(t, err)

	err = st.Actor().Overwright(ta)
	assert.Equal(t, err, store.ErrRecordNotFound)
}

// TODO: Test func
func TestOverwrightFieldsActor(t *testing.T) {
	ta := actor.TestActor(t)
	db, clear := pgstore.TestStore(t, dbPath)
	defer clear("actors")

	st := pgstore.New(db)
	err := st.Actor().Create(ta)
	assert.NoError(t, err)
	assert.NotNil(t, ta)
}

// TODO: Test func
func TestGetAll(t *testing.T) {
	
}
package dictionary

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		// want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrorDictKeyNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		keyWord := "test"
		want := "this is just a test"

		err := dictionary.Add(keyWord, want)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, keyWord, want)
	})

	t.Run("existing word", func(t *testing.T) {
		keyWord := "test"
		want := "this is just a test"
		dictionary := Dictionary{keyWord: want}

		err := dictionary.Add(keyWord, "new value")

		assertError(t, err, ErrorDictKeyExists)
		assertDefinition(t, dictionary, keyWord, want)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		keyWord := "test"
		definition := "this is just a test"
		dictionary := Dictionary{keyWord: definition}
		newDefinition := "new definition"

		err := dictionary.Update(keyWord, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, keyWord, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		keyWord := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(keyWord, definition)

		assertError(t, err, ErrorDictKeyNotExists)
		// assertDefinition(t, dictionary, keyWord, definition)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}

		err := dictionary.Delete(word)
		assertError(t, err, nil)

		_, err2 := dictionary.Search(word)
		assertError(t, err2, ErrorDictKeyNotFound)
	})

	t.Run("non-existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}

		err := dictionary.Delete(word)
		assertError(t, err, ErrorDictKeyNotExists)
	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if !errors.Is(got, want) {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, keyWord, want string) {
	t.Helper()

	got, err := dictionary.Search(keyWord)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertString(t, got, want)
}

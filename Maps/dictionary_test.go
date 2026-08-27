package Maps

import (
	"errors"
	"testing"
)

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if !errors.Is(got, want) {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word string, value string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("something wrong", err)
	}

	assertStrings(t, got, value)

}

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {

		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get error")
		}
		assertError(t, err, ErrNotFound)

	})

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {

		dict := Dictionary{}
		word := "Test"
		value := "this is just a test"

		err := dict.Add(word, value)

		assertDefinition(t, dict, word, value)
		assertError(t, err, nil)
	})
	t.Run("exisisting word", func(t *testing.T) {
		word := "Test"
		value := "this is just a test"
		dict := Dictionary{word: value}

		err := dict.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, word, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "Test"
		Value := " value of key"
		dict := Dictionary{word: Value}

		newValue := "updated value of key"

		err := dict.Update(word, newValue)

		assertDefinition(t, dict, word, newValue)
		assertError(t, err, nil)
	})
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "Test"
		newValue := "updated value of key"

		err := dict.Update(word, newValue)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {

		dict := Dictionary{}
		word := "Test"
		dict.Delete(word)

		_, err := dict.Search(word)
		assertError(t, err, ErrNotFound)
	})
	t.Run("non-existing word", func(t *testing.T) {
		existWord := "Test"
		Value := "value"
		dict := Dictionary{existWord: Value}

		notExistWord := "new test"
		err := dict.Delete(notExistWord)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

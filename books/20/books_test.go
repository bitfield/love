package books_test

import (
	"books"
	"cmp"
	"net"
	"slices"
	"testing"
)

func TestGetAllBooks_ReturnsAllBooks(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	got := catalog.GetAllBooks()
	assertTestBooks(t, got)
}

func TestOpenCatalog_ReadsSameDataWrittenBySync(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	path := t.TempDir() + "/catalog"
	err := catalog.Sync(path)
	if err != nil {
		t.Fatal(err)
	}
	newCatalog, err := books.OpenCatalog(path)
	if err != nil {
		t.Fatal(err)
	}
	got := newCatalog.GetAllBooks()
	assertTestBooks(t, got)
}

func TestSyncWritesCatalogDataToFile(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	path := t.TempDir() + "/catalog"
	err := catalog.Sync(path)
	if err != nil {
		t.Fatal(err)
	}
	newCatalog, err := books.OpenCatalog(path)
	if err != nil {
		t.Fatal(err)
	}
	got := newCatalog.GetAllBooks()
	assertTestBooks(t, got)
}

func TestNewCatalog_CreatesEmptyCatalog(t *testing.T) {
	t.Parallel()
	catalog := books.NewCatalog()
	books := catalog.GetAllBooks()
	if len(books) > 0 {
		t.Errorf("want empty catalog, got %#v", books)
	}
}

func TestGetBook_FindsBookInCatalogByID(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	got, ok := catalog.GetBook("abc")
	if !ok {
		t.Fatal("book not found")
	}
	if got != ABC {
		t.Fatalf("want %#v, got %#v", ABC, got)
	}
}

func TestGetBook_ReturnsFalseWhenBookNotFound(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := catalog.GetBook("nonexistent ID")
	if ok {
		t.Fatal("want false for nonexistent ID, got true")
	}
}

func TestAddBook_AddsGivenBookToCatalog(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := catalog.GetBook("123")
	if ok {
		t.Fatal("book already present")
	}
	err := catalog.AddBook(books.Book{
		ID:     "123",
		Title:  "The Prize of all the Oceans",
		Author: "Glyn Williams",
		Copies: 2,
	})
	if err != nil {
		t.Fatal(err)
	}
	_, ok = catalog.GetBook("123")
	if !ok {
		t.Fatal("added book not found")
	}
}

func TestAddBook_ReturnsErrorIfIDExists(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := catalog.GetBook("abc")
	if !ok {
		t.Fatal("book not present")
	}
	err := catalog.AddBook(ABC)
	if err == nil {
		t.Fatal("want error for duplicate ID, got nil")
	}
}

func TestSetCopies_SetsNumberOfCopiesToGivenValue(t *testing.T) {
	t.Parallel()
	book := books.Book{
		Copies: 5,
	}
	err := book.SetCopies(12)
	if err != nil {
		t.Fatal(err)
	}
	if book.Copies != 12 {
		t.Errorf("want 12 copies, got %d", book.Copies)
	}
}

func TestSetCopies_ReturnsErrorIfCopiesNegative(t *testing.T) {
	t.Parallel()
	book := books.Book{}
	err := book.SetCopies(-1)
	if err == nil {
		t.Error("want error for negative copies, got nil")
	}
}

func TestSetCopies_OnCatalogModifiesSpecifiedBook(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	copies, err := catalog.GetCopies("abc")
	if err != nil {
		t.Fatal(err)
	}
	if copies != 1 {
		t.Fatalf("want 1 copy before change, got %d", copies)
	}
	err = catalog.SetCopies("abc", 2)
	if err != nil {
		t.Fatal(err)
	}
	copies, err = catalog.GetCopies("abc")
	if err != nil {
		t.Fatal(err)
	}
	if copies != 2 {
		t.Fatalf("want 2 copies after change, got %d", copies)
	}
}

func TestSetCopies_IsRaceFree(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	go func() {
		for range 100 {
			err := catalog.SetCopies("abc", 0)
			if err != nil {
				panic(err)
			}
		}
	}()
	for range 100 {
		_, err := catalog.GetCopies("abc")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestGetAllBooks_OnClientListsAllBooks(t *testing.T) {
	t.Parallel()
	client := getTestClient(t)
	got, err := client.GetAllBooks()
	if err != nil {
		t.Fatal(err)
	}
	assertTestBooks(t, got)
}

func TestGetBook_OnClientFindsBookByID(t *testing.T) {
	t.Parallel()
	client := getTestClient(t)
	got, err := client.GetBook("abc")
	if err != nil {
		t.Fatal(err)
	}
	if got != ABC {
		t.Fatalf("want %#v, got %#v", ABC, got)
	}
}

func TestFindReturnsErrorWhenBookNotFound(t *testing.T) {
	t.Parallel()
	client := getTestClient(t)
	_, err := client.GetBook("bogus")
	if err == nil {
		t.Error("want error when book not found, got nil")
	}
}

func getTestCatalog() *books.Catalog {
	catalog := books.NewCatalog()
	err := catalog.AddBook(ABC)
	if err != nil {
		panic(err)
	}
	err = catalog.AddBook(XYZ)
	if err != nil {
		panic(err)
	}
	return catalog
}

func getTestClient(t *testing.T) *books.Client {
	t.Helper()
	addr := randomLocalAddr(t)
	go func() {
		err := books.ListenAndServe(addr, getTestCatalog())
		if err != nil {
			panic(err)
		}
	}()
	return books.NewClient(addr)
}

func assertTestBooks(t *testing.T, got []books.Book) {
	t.Helper()
	want := []books.Book{ABC, XYZ}
	slices.SortFunc(got, func(a, b books.Book) int {
		return cmp.Compare(a.Author, b.Author)
	})
	if !slices.Equal(want, got) {
		t.Fatalf("want %#v, got %#v", want, got)
	}
}

func randomLocalAddr(t *testing.T) string {
	t.Helper()
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	return l.Addr().String()
}

var (
	ABC = books.Book{
		Title:  "In the Company of Cheerful Ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
		ID:     "abc",
	}

	XYZ = books.Book{
		Title:  "White Heat",
		Author: "Dominic Sandbrook",
		Copies: 2,
		ID:     "xyz",
	}
)

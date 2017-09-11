package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/treacher/step-warrior-api/actions"
	. "github.com/treacher/step-warrior-api/handlers"
	"github.com/treacher/step-warrior-api/models"
)

var (
	chestFetcherInput   string
	responseRecorder    *httptest.ResponseRecorder
	chestCount          int
	lastFetchedChestsAt time.Time
	err                 error
)

type MockChestFetcher struct{}

func (mch *MockChestFetcher) Fetch(user models.FetchedChestsUpdater, steps int) (*actions.FetchChestResponse, error) {
	return &actions.FetchChestResponse{
		ChestCount:          chestCount,
		LastFetchedChestsAt: lastFetchedChestsAt,
	}, err
}

var _ = Describe("CreateChests", func() {
	JustBeforeEach(func() {
		req, err := http.NewRequest("POST", "/chests", strings.NewReader(chestFetcherInput))
		ctx := req.Context()
		ctx = context.WithValue(ctx, "User", models.User{})
		req = req.WithContext(ctx)
		Expect(err).NotTo(HaveOccurred())
		responseRecorder = httptest.NewRecorder()
		handler := http.HandlerFunc(NewCreateChests(&MockChestFetcher{}).ServeHTTP)
		handler.ServeHTTP(responseRecorder, req)
	})

	BeforeEach(func() {
		chestFetcherInput = `{"steps" : 10000}`
		chestCount = 10
		lastFetchedChestsAt = time.Now()
	})

	It("returns the correct json", func() {
		byteStr, _ := lastFetchedChestsAt.MarshalJSON()
		expectedJson := "{\"chestCount\":10,\"lastFetchedChestsAt\":" + string(byteStr) + "}\n"
		Expect(string(responseRecorder.Body.Bytes())).To(Equal(expectedJson))
	})

	It("responds with a 201", func() {
		Expect(responseRecorder.Code).To(Equal(201))
	})
})

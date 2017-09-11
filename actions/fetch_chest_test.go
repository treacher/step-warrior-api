package actions_test

import (
	"errors"
	"time"

	"github.com/bluele/go-timecop"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/treacher/step-warrior-api/actions"
	"gopkg.in/pg.v5"
)

type MockUser struct{}

func (*MockUser) UpdateFetchedChests(steps int, db *pg.DB) (int, time.Time, error) {
	return steps, timecop.Now(), errors.New("fetched_chest_error")
}

var _ = Describe("ChestFetcher", func() {
	var (
		fetchChestAction     *FetchChest
		mockDatabase         *pg.DB
		expectedFetchedChest *FetchChestResponse
		expectedError        = errors.New("fetched_chest_error")
		steps                int
	)

	Describe("Fetch", func() {
		BeforeEach(func() {
			timecop.Freeze(time.Now())
			steps = 10000
			fetchChestAction = NewChestFetcher(mockDatabase)

			expectedFetchedChest = &FetchChestResponse{
				ChestCount:          7,
				LastFetchedChestsAt: timecop.Now(),
			}
		})

		It("converts the response into a struct forwarding on any errors", func() {
			fetchedChests, err := fetchChestAction.Fetch(&MockUser{}, steps)

			Expect(err).To(BeEquivalentTo(expectedError))
			Expect(fetchedChests).To(BeEquivalentTo(expectedFetchedChest))
		})

		Context("step count is over the max amount of steps", func() {
			BeforeEach(func() {
				steps = 25000

				expectedFetchedChest = &FetchChestResponse{
					ChestCount:          13,
					LastFetchedChestsAt: timecop.Now(),
				}
			})

			It("returns the max amount of chests", func() {
				fetchedChests, err := fetchChestAction.Fetch(&MockUser{}, steps)

				Expect(err).To(BeEquivalentTo(expectedError))
				Expect(fetchedChests).To(BeEquivalentTo(expectedFetchedChest))
			})
		})
	})

})

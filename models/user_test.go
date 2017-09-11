package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/treacher/step-warrior-api/models"

	"time"
)

var _ = Describe("User", func() {
	var (
		user                    *User
		userChestCount          int
		userLastFetchedChestsAt time.Time
		databaseUser            *User
	)

	JustBeforeEach(func() {
		user = &User{
			Identifier:          "1010101",
			ChestCount:          userChestCount,
			LastFetchedChestsAt: userLastFetchedChestsAt,
		}

		err := user.Persist(TestDatabase)

		Expect(err).NotTo(HaveOccurred())
	})

	var _ = AfterEach(func() {
		TestDatabase.Exec("DELETE FROM Users;")
	})

	Describe("DecrementChestCount", func() {
		var (
			decrementErr error
		)

		JustBeforeEach(func() {
			decrementErr = user.DecrementChestCount(TestDatabase)

			databaseUser = &User{Id: user.Id}
			TestDatabase.Model(databaseUser).Select()
		})

		BeforeEach(func() {
			userChestCount = 10
		})

		It("decrements the users chest count and doesn't return an error", func() {
			Expect(user.ChestCount).To(Equal(9))
			Expect(decrementErr).NotTo(HaveOccurred())

			Expect(databaseUser.ChestCount).To(Equal(9))
		})
	})

	Describe("UpdateFetchedChests", func() {
		var (
			fetchedChestCount int
			fetchedAt         time.Time
			fetchErr          error
		)

		BeforeEach(func() {
			userChestCount = 0
		})

		JustBeforeEach(func() {
			fetchedChestCount, fetchedAt, fetchErr = user.UpdateFetchedChests(10, TestDatabase)

			databaseUser = &User{Id: user.Id}
			TestDatabase.Model(databaseUser).Select()
		})

		Context("User's lastFetchedAt is zero", func() {
			It("returns the new chest count and lastFetchedChestsAt date", func() {
				Expect(fetchedChestCount).To(Equal(10))
				Expect(fetchedAt).NotTo(BeNil())
				Expect(fetchErr).NotTo(HaveOccurred())

				Expect(databaseUser.ChestCount).To(Equal(10))
				Expect(databaseUser.LastFetchedChestsAt).NotTo(BeNil())
			})
		})

		Context("less than 24 hours since lastFetchedAt", func() {
			BeforeEach(func() {
				userLastFetchedChestsAt = time.Now().Add(-time.Hour)
			})

			It("returns the users current chestCount and lastFetchedAt", func() {
				Expect(fetchedChestCount).To(Equal(0))
				Expect(fetchedAt).To(Equal(userLastFetchedChestsAt))
				Expect(fetchErr).NotTo(HaveOccurred())

				Expect(databaseUser.ChestCount).To(Equal(0))
				Expect(databaseUser.LastFetchedChestsAt).To(
					Equal(userLastFetchedChestsAt.Round(time.Microsecond)))
			})
		})

		Context("greater than 24 hours since lastFetchedAt", func() {
			BeforeEach(func() {
				userLastFetchedChestsAt = time.Now().Add(-(time.Hour * 25))
			})

			It("returns the new chest count and lastFetchedChestsAt date", func() {
				Expect(fetchedChestCount).To(Equal(10))
				Expect(fetchedAt).NotTo(Equal(userLastFetchedChestsAt))
				Expect(fetchErr).NotTo(HaveOccurred())

				Expect(databaseUser.ChestCount).To(Equal(10))
				Expect(databaseUser.LastFetchedChestsAt).NotTo(
					Equal(userLastFetchedChestsAt.Round(time.Microsecond)))
			})
		})
	})

	Describe("Persist", func() {
		JustBeforeEach(func() {
			databaseUser = &User{Id: user.Id}
			TestDatabase.Model(databaseUser).Select()
		})

		It("persists the user and sets the updatedAt and createdAt dates", func() {
			Expect(databaseUser.Identifier).To(Equal(user.Identifier))
			Expect(databaseUser.ChestCount).To(Equal(user.ChestCount))
			Expect(databaseUser.LastFetchedChestsAt).To(Equal(user.LastFetchedChestsAt.Round(time.Microsecond)))
			Expect(databaseUser.CreatedAt).NotTo(BeNil())
			Expect(databaseUser.UpdatedAt).NotTo(BeNil())
		})
	})

	Describe("Update", func() {
		JustBeforeEach(func() {
			user.ChestCount = 5
			err := user.Update(TestDatabase)
			Expect(err).NotTo(HaveOccurred())

			databaseUser = &User{Id: user.Id}
			TestDatabase.Model(databaseUser).Select()
		})

		It("updates the user and sets the updatedAt", func() {
			Expect(databaseUser.ChestCount).To(Equal(user.ChestCount))
			Expect(databaseUser.CreatedAt).To(Equal(user.CreatedAt.Round(time.Microsecond)))
			Expect(databaseUser.UpdatedAt).NotTo(Equal(user.UpdatedAt))
		})
	})
})

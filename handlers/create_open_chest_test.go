package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/treacher/step-warrior-api/handlers"
	"github.com/treacher/step-warrior-api/models"
)

type MockChestOpener struct{}

func (mco *MockChestOpener) Open(user *models.User) *models.Chest {
	material := &models.UserMaterial{Material: &models.Material{Name: "Material", Rarity: "common"}}
	equipment := &models.UserEquipment{
		Equipment: &models.Equipment{
			Name:   "Equipment",
			Rarity: "rare",
		},
		Attributes: []models.UserEquipmentAttribute{{Type: "strength", Value: 10}},
	}
	plan := &models.UserPlan{Plan: &models.Plan{Name: "Plan", Rarity: "ancient"}}

	return &models.Chest{
		UserMaterials: []*models.UserMaterial{material},
		UserEquipment: []*models.UserEquipment{equipment},
		UserPlans:     []*models.UserPlan{plan},
	}
}

var _ = Describe("CreateOpenChest", func() {
	JustBeforeEach(func() {
		req, err := http.NewRequest("POST", "/chest_items", nil)
		ctx := req.Context()
		ctx = context.WithValue(ctx, "User", models.User{ChestCount: 10})
		req = req.WithContext(ctx)
		Expect(err).NotTo(HaveOccurred())
		responseRecorder = httptest.NewRecorder()
		handler := http.HandlerFunc(NewCreateOpenChest(&MockChestOpener{}).ServeHTTP)
		handler.ServeHTTP(responseRecorder, req)
	})

	It("returns the correct json", func() {
		expectedJson := "{\"chestCount\":10,\"materials\":[{\"name\":\"Material\",\"rarity\":\"common\"}],\"equipment\":[{\"name\":\"Equipment\",\"rarity\":\"rare\",\"attributes\":[{\"type\":\"strength\",\"value\":10}]}],\"plans\":[{\"name\":\"Plan\",\"rarity\":\"ancient\"}]}\n"
		Expect(string(responseRecorder.Body.Bytes())).To(Equal(expectedJson))
	})

	It("responds with a 201", func() {
		Expect(responseRecorder.Code).To(Equal(201))
	})

})

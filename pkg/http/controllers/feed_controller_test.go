package controllers

import (
	"bytes"
	"coderrob.com/meraki-cmx-receiver/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestFeedController_ReceiverHandler(t *testing.T) {
	files := []string{
		"../../../test/testdata/v2_wifi_feed.json",
		"../../../test/testdata/v2_bte_feed.json",
		"../../../test/testdata/v3_wifi_feed.json",
	}
	for _, file := range files {
		t.Run(file, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			secret := "supersecret"
			orgID := "12345"
			orgSvc := mocks.NewMockIOrganizationService(ctrl)
			locSvc := mocks.NewMockILocationService(ctrl)
			orgSvc.EXPECT().GetFeedSecret(orgID).Return(secret, nil)
			locSvc.EXPECT().UpdateLocations(gomock.Any()).Return(nil)
			filePath := file
			body, _ := ioutil.ReadFile(filePath)
			req := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))
			req = mux.SetURLVars(req, map[string]string{
				"orgID": orgID,
			})
			res := httptest.NewRecorder()
			feedCtrl := FeedController{
				OrganizationService: orgSvc,
				LocationService:     locSvc,
			}
			feedCtrl.ReceiverHandler(res, req)
		})
	}
}

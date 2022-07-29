package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/openziti-test-kitchen/zrok/controller/store"
	"github.com/openziti-test-kitchen/zrok/rest_model_zrok"
	"github.com/openziti-test-kitchen/zrok/rest_server_zrok/operations/identity"
	"github.com/openziti/edge/rest_management_api_client"
	identity_edge "github.com/openziti/edge/rest_management_api_client/identity"
	rest_model_edge "github.com/openziti/edge/rest_model"
	sdk_config "github.com/openziti/sdk-golang/ziti/config"
	"github.com/openziti/sdk-golang/ziti/enroll"
	"github.com/sirupsen/logrus"
	"time"
)

func enableHandler(_ identity.EnableParams, principal *rest_model_zrok.Principal) middleware.Responder {
	// start transaction early; if it fails, don't bother creating ziti resources
	tx, err := str.Begin()
	if err != nil {
		logrus.Errorf("error starting transaction: %v", err)
		return identity.NewCreateAccountInternalServerError().WithPayload(rest_model_zrok.ErrorMessage(err.Error()))
	}

	client, err := edgeClient()
	if err != nil {
		logrus.Errorf("error getting edge client: %v", err)
		return identity.NewEnableInternalServerError().WithPayload(rest_model_zrok.ErrorMessage(err.Error()))
	}
	ident, err := createIdentity(principal.Username, client)
	if err != nil {
		logrus.Error(err)
		return identity.NewEnableInternalServerError().WithPayload(rest_model_zrok.ErrorMessage(err.Error()))
	}
	cfg, err := enrollIdentity(ident.Payload.Data.ID, client)
	if err != nil {
		logrus.Error(err)
		return identity.NewEnableInternalServerError().WithPayload(rest_model_zrok.ErrorMessage(err.Error()))
	}

	iid, err := str.CreateIdentity(int(principal.ID), &store.Identity{ZitiId: ident.Payload.Data.ID}, tx)
	if err != nil {
		logrus.Errorf("error storing created identity: %v", err)
		_ = tx.Rollback()
		return identity.NewCreateAccountInternalServerError().WithPayload(rest_model_zrok.ErrorMessage(err.Error()))
	}
	if err := tx.Commit(); err != nil {
		logrus.Errorf("error committing: %v", err)
		return identity.NewCreateAccountInternalServerError().WithPayload(rest_model_zrok.ErrorMessage(err.Error()))
	}
	logrus.Infof("recorded identity '%v' with id '%v' for '%v'", ident.Payload.Data.ID, iid, principal.Username)

	resp := identity.NewEnableCreated().WithPayload(&rest_model_zrok.EnableResponse{
		Identity: ident.Payload.Data.ID,
	})

	var out bytes.Buffer
	enc := json.NewEncoder(&out)
	enc.SetEscapeHTML(false)
	err = enc.Encode(&cfg)
	if err != nil {
		panic(err)
	}
	resp.Payload.Cfg = out.String()

	return resp
}

func createIdentity(username string, client *rest_management_api_client.ZitiEdgeManagement) (*identity_edge.CreateIdentityCreated, error) {
	iIsAdmin := false
	iId, err := randomId()
	if err != nil {
		return nil, err
	}
	name := fmt.Sprintf("%v-%v", username, iId)
	identityType := rest_model_edge.IdentityTypeUser
	i := &rest_model_edge.IdentityCreate{
		Enrollment:          &rest_model_edge.IdentityCreateEnrollment{Ott: true},
		IsAdmin:             &iIsAdmin,
		Name:                &name,
		RoleAttributes:      nil,
		ServiceHostingCosts: nil,
		Tags:                nil,
		Type:                &identityType,
	}
	req := identity_edge.NewCreateIdentityParams()
	req.Identity = i
	resp, err := client.Identity.CreateIdentity(req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func enrollIdentity(id string, client *rest_management_api_client.ZitiEdgeManagement) (*sdk_config.Config, error) {
	p := &identity_edge.DetailIdentityParams{
		Context: context.Background(),
		ID:      id,
	}
	p.SetTimeout(30 * time.Second)
	resp, err := client.Identity.DetailIdentity(p, nil)
	if err != nil {
		return nil, err
	}
	tkn, _, err := enroll.ParseToken(resp.GetPayload().Data.Enrollment.Ott.JWT)
	if err != nil {
		return nil, err
	}
	flags := enroll.EnrollmentFlags{
		Token:  tkn,
		KeyAlg: "RSA",
	}
	conf, err := enroll.Enroll(flags)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

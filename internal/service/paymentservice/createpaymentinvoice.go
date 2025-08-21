package paymentservice

import (
	"context"
	"errors"
	"github.com/cynx-io/cynx-core/src/logger"
	pbhermes "github.com/cynx-io/plutus-payment/api/proto/gen/hermes"
	pb "github.com/cynx-io/plutus-payment/api/proto/gen/plutus"
	"github.com/cynx-io/plutus-payment/internal/constant"
	helper2 "github.com/cynx-io/plutus-payment/internal/helper"
	"github.com/cynx-io/plutus-payment/internal/model/entity"
	"github.com/cynx-io/plutus-payment/internal/model/response"
	"github.com/xendit/xendit-go/v7/invoice"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

func (s *Service) CreatePaymentInvoice(ctx context.Context, req *pb.CreatePaymentInvoiceRequest, resp *pb.PaymentInvoiceResponse) error {

	customer, err := s.TblCustomer.GetCustomerByUserId(ctx, req.UserId)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			response.ErrorInternal(resp)
			return err
		}
	}

	if err != nil || customer == nil {
		userResp, err := s.HermesUserClient.GetUserById(ctx, &pbhermes.GetUserByIdRequest{
			Base: req.Base,
			Id:   req.UserId,
		})
		if err != nil {
			response.ErrorHermes(resp)
			return err
		}

		if userResp.Base.Code != response.CodeSuccess.String() {
			logger.Debug(ctx, "failed to get user from Hermes: ", userResp.Base.Code, ": ", userResp.Base.Desc)
			response.ErrorHermes(resp)
			return errors.New("failed to get user from Hermes")
		}

		user := userResp.User
		customer = &entity.TblCustomer{
			Email:    &user.Email,
			FullName: &user.Name,
			Provider: constant.ProviderXendit,
			UserId:   req.UserId,
		}

		customer, err = s.TblCustomer.CreateCustomer(ctx, *customer)
		if err != nil {
			response.ErrorDatabaseCustomer(resp)
			return err
		}
	}

	paymentInvoiceId := helper2.FormatXenditExternalId(customer.Id, req.Base.RequestId)

	createInvoiceReq := *invoice.NewCreateInvoiceRequest(req.Base.RequestId, float64(req.Amount))
	createInvoiceReq.SetCurrency(req.Currency)
	createInvoiceReq.SetInvoiceDuration(float32(req.DurationInSeconds))
	createInvoiceReq.SetExternalId(paymentInvoiceId)
	createInvoiceReq.SetDescription(req.Description)
	customerObj := invoice.CustomerObject{}

	if customer.FullName != nil && *customer.FullName != "" {
		customerObj.GivenNames = *invoice.NewNullableString(customer.FullName)
	}

	if customer.Email != nil && *customer.Email != "" {
		customerObj.Email = *invoice.NewNullableString(customer.Email)
	}

	if customer.PhoneNumber != nil && *customer.PhoneNumber != "" {
		customerObj.MobileNumber = *invoice.NewNullableString(customer.PhoneNumber)
	}

	createInvoiceReq.SetCustomer(customerObj)
	createInvoiceReq.SetSuccessRedirectUrl(req.SuccessReturnUrl)
	createInvoiceReq.SetFailureRedirectUrl(req.FailureReturnUrl)

	logger.Debug(ctx, "Creating Xendit invoice with payload:", map[string]interface{}{
		"external_id":      paymentInvoiceId,
		"amount":           float64(req.Amount),
		"currency":         "IDR",
		"description":      req.Description,
		"invoice_duration": float32(req.DurationInSeconds),
		"customer_email":   customer.Email,
		"customer_name":    customer.FullName,
		"customer_phone":   customer.PhoneNumber,
		"success_url":      req.SuccessReturnUrl,
		"failure_url":      req.FailureReturnUrl,
	})

	createInvoiceResp, httpResp, xenditErr := s.XenditClient.InvoiceApi.CreateInvoice(ctx).
		CreateInvoiceRequest(createInvoiceReq).
		Execute()

	if xenditErr != nil {
		logger.Error(ctx, "failed to create invoice: ", xenditErr.Error())
		response.ErrorXendit(resp)
		return xenditErr
	}

	if httpResp == nil {
		logger.Error(ctx, "failed to create invoice: http response is nil")
		response.ErrorXendit(resp)
		return errors.New("failed to create invoice: http response is nil")
	}

	if httpResp.StatusCode != http.StatusOK && httpResp.StatusCode != http.StatusCreated {
		logger.Error(ctx, "failed to create invoice: http response status is ", httpResp.StatusCode, " - ", httpResp.Status)
		response.ErrorXendit(resp)
		return errors.New("failed to create invoice: http response status is " + strconv.Itoa(httpResp.StatusCode) + " - " + httpResp.Status)
	}

	if createInvoiceResp == nil {
		logger.Error(ctx, "failed to create invoice: response is nil")
		response.ErrorXendit(resp)
		return errors.New("failed to create invoice: response is nil")
	}

	paymentInvoice := &entity.TblPaymentInvoice{
		ExpiresAt:      createInvoiceResp.ExpiryDate,
		Id:             paymentInvoiceId,
		ExternalId:     *createInvoiceResp.Id,
		RequestId:      req.Base.RequestId,
		Status:         strings.ToUpper(createInvoiceResp.Status.String()),
		PaymentLinkUrl: createInvoiceResp.InvoiceUrl,
		Currency:       string(createInvoiceResp.GetCurrency()),
		Description:    createInvoiceResp.GetDescription(),
		Provider:       constant.ProviderXendit,
		Amount:         int64(createInvoiceResp.Amount),
		CustomerId:     customer.Id,
	}
	if err := s.TblPaymentInvoice.CreatePaymentInvoice(ctx, paymentInvoice); err != nil {
		logger.Error(ctx, "failed to create payment invoice in database: ", err.Error())
		response.ErrorDatabaseInvoice(resp)
		return err
	}

	response.Success(resp)
	resp.Payment = paymentInvoice.Response()
	return nil
}

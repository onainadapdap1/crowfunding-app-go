package payment

import (
	"bwastartup/user"
	"os"
	"strconv"

	midtrans "github.com/veritrans/go-midtrans"
)

type service struct {
}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	midclient := midtrans.NewClient()
	midtransServerKey :=  os.Getenv("MIDTRANS_SERVERKEY")
	midtransClientKey := os.Getenv("MIDTRANS_CLIENTKEY")
	midclient.ServerKey = midtransServerKey
	midclient.ClientKey = midtransClientKey
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}
	
	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}

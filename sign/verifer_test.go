package sign

import (
	"testing"
)

func TestGoVerify_ParseQuery(t *testing.T) {
	requestUri := "/restful/api/numbers?from=chanliao&timestamp=1693929600&sign=12c93799e4e075a03713c766e99f3112"

	verifier := NewVerifier()

	if err := verifier.ParseQuery(requestUri); nil != err {
		t.Fatal(err)
	}

	signer := NewSigner()
	signer.SetBody(verifier.GetBodyWithoutSign())
	signer.SetSecret("d93047a4d6fe6111")


	sign := signer.GetSignature()
	if verifier.MustString("sign") != sign {
		t.Fatal("校验失败")
	}

}
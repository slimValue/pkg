package sign

import (
	"fmt"
	"testing"
)

func TestSignMd5(t *testing.T) {
	signer := NewSigner()
	signer.SetTimestamp(1693929600)
	signer.AddBody("from", "chanliao")
	signer.SetSecret("d93047a4d6fe6111")
	fmt.Println("生成签字字符串：" + signer.GetSignBodyString())
	fmt.Println("生成签名sign：" + signer.GetSignature())
	fmt.Println("输出URL字符串：" + signer.GetSignedQuery())
	if "from=chanliao&timestamp=1693929600&sign=12c93799e4e075a03713c766e99f3112" != signer.GetSignedQuery() {
		t.Fatal("Md5校验失败")
	}
}

func TestSign(t *testing.T) {
	signer := NewSigner()
	signer.SetTimestamp(1694143440)
	signer.SetOwner("test")
	signer.SetApp("test")
	signer.SetSecret("6fe2e3a88ee325762c6111e0cc3d4c7e")

	t.Logf(signer.GetSignature())
}

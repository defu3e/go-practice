package billing

import (
	"io/ioutil"
	"log"
	"strconv"
)

type BillingData struct { 
	CreateCustomer bool 
	Purchase bool 
	Payout bool 
	Recurring bool 
	FraudControl bool 
	CheckoutPage bool 
 } 

func GetBillingData (srcFilePath string) BillingData  {
	log.Println("\n=== Getting email data ===")
    str, err := ioutil.ReadFile(srcFilePath)
    if err != nil {
        log.Fatal(err)
	}

	mask := getMaskNumb(str)
	
	return BillingData{
		CreateCustomer: isActiveBit(mask, 1),
		Purchase: isActiveBit(mask,2),
		Payout: isActiveBit(mask, 3),
		Recurring: isActiveBit(mask, 4),
		FraudControl: isActiveBit(mask,5),
		CheckoutPage: isActiveBit(mask,6),
	}

}

func getMaskNumb (s []byte) (mask uint8) {
	for i, p := (len(s) - 1), 0; i >= 0; i, p = i-1, p+1 {
		v, _ := strconv.ParseUint(string(s[i]), 10, 64)
		mask += uint8(v) << p
	}
	return
}

// get bit at m postion from n number (rtl)
// from: https://tproger.ru/articles/awesome-bits/
func isActiveBit (n uint8, m uint8) bool {
	return (n >> (m-1)) & 1 == 1;
}


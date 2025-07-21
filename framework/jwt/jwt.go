package jwt

import (
	"time"
	"github.com/golang-jwt/jwt/v4"
)
type CustomClaims struct {
	ID string `json:"id"`
	Roles []string `json:"roles"`
	WalletAddress string `json:"walletAddress"`
	jwt.RegisteredClaims
}
//expired is second
func GenerateJWTToken(key_sign,id,wallet_address string,roles []string,expired int,args ...string) (string,error){
	
	signingKey := []byte(key_sign)
	// Create the claims
	claims := CustomClaims{
		id,
		roles,
		wallet_address,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expired) *  time.Second)),			
			Issuer:    "Rivalz",
		},
	}
	//fmt.Printf("%+v\r\n",claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	res, err := token.SignedString(signingKey)
	if err!=nil{
		return "",err
	}
	return res,nil
}
//return 
// - int: number of second token expired , 0 if not expired
/*
func TokenExpiredTime(key,token_string string) float64{
	var claims CustomClaims
	_, err := jwt.ParseWithClaims(token_string, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err==nil{
		return 0
	}
	v, _ := err.(*jwt.ValidationError)
	if v.Errors == jwt.ValidationErrorExpired{
		//tm := time.Unix(claims.ExpiresAt, 0)
		return time.Now().Sub(claims.ExpiresAt.Time).Seconds()
	}
	return 0
}
func VerifyJWTToken(key,token_string string) (*CustomClaims,*e.Error){
	if key==""{
		return nil,e.New("Key is empty for verify token","KEY_IS_EMPTY")
	}
	if token_string==""{
		return nil,e.New("token is empty","TOKEN_IS_EMPTY")
	}
	token, err := jwt.ParseWithClaims(token_string, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		return nil, e.NewErr(err, "PARSE_TOKEN")
	}
	claim, ok := token.Claims.(*CustomClaims); 
	//fmt.Printf("%+v\r\n",token.Claims)
	if ok && token.Valid {
		return claim,nil
	} 
	return nil,e.NewErr(err,"PARSE_TOKEN")
}
*/
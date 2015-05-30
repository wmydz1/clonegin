package clonegin
import (

    "encoding/base64"
    "crypto/subtle"
)
const (
    AuthUserKey = "user"
)
type (
Accounts map[string]string
authPair struct {
    Value string
    User string
}
authPairs []authPair
)
func (a authPairs) searchCredential(authValue string) (string, bool) {
    if len(authValue)==0 {
        return "", false
    }
    for _, pair := range a {
        if pair.Value==authValue {
            return pair.User, true
        }
    }
    return "", false
}
func BasicAuthForRealm(accounts Accounts, realm string) HandlerFunc {

}
func BasicAuth(accounts Accounts) HandlerFunc {
    return BasicAuthForRealm(accounts, "")
}
func processAccounts(accounts Accounts) authPairs {
    if len(accounts) ==0 {
        panic("Empty list of authorized credentials")
    }
    pairs := make(authPair, 0, len(accounts))
    for user, password := range accounts {
        if len(user)==0 {
            panic("User can not be empty")
        }
        value := authorizationHeader(user, password)
        pairs =append(pairs, authPair{
            Value:value,
            User:user,

        })
    }
    return pairs
}
func authorizationHeader(user, password string) string {
    base := user+":"+password
    return "Basic"+base64.StdEncoding.EncodeToString([]byte(base))
}
func secureCompare(given, acual string) bool {
    if subtle.ConstantTimeEq(int32(len(given)), int32(len(acual))) ==1 {
        return subtle.ConstantTimeCompare([]byte(given), []byte(acual))==1
    }else {
        return subtle.ConstantTimeCompare([]byte(acual), []byte(acual))==1&&false
    }
}
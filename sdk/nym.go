package sdk

import (
	"time"

	"github.com/domenetwork/dome-lib/pkg/cfg"
	"github.com/domenetwork/dome-lib/pkg/common"
	"github.com/domenetwork/dome-lib/pkg/log"
)

// User returns the user object for the logged in user.
func (sdk *SDK) User() (user *common.User, err error) {
	log.D("sdk", "nym", "user")
	var v interface{}
	url := cfg.ServiceURL("nym")
	if v, err = sdk.fetch.Get(url, "/user"); err != nil {
		return
	}

	err = common.MapToInterface(v.(map[string]interface{}), user)
	return
}

// Register adds the provided user information into the Nym service.
// If there are any conflicts an error will be returned.
func (sdk *SDK) Register(user *common.User) (err error) {
	log.D("sdk", "nym", "register", user)
	var v interface{}
	url := cfg.ServiceURL("nym")
	if v, err = sdk.fetch.Post(url, "/register", user); err != nil {
		return
	}

	err = common.MapToInterface(v.(map[string]interface{}), user)
	return
}

// Login with the provided user information and setup the authorization
// token on success or return an error on failure.
func (sdk *SDK) Login(user *common.User) (err error) {
	log.D("sdk", "nym", "login", user)
	var v interface{}
	url := cfg.ServiceURL("nym")
	if v, err = sdk.fetch.Post(url, "/login", user); err != nil {
		return
	}

	res := v.(map[string]interface{})
	log.D("sdk", "login", "response", res)
	sdk.fetch.Auth(res["token"].(string))

	err = common.MapToInterface(res["user"].(map[string]interface{}), user)
	return
}

// Lookup a user and get a hardened key for that user.  These keys are single
// use and should not be used more than once.
func (sdk *SDK) Lookup(user *common.User) (err error) {
	log.D("sdk", "nym", "lookup", user)
	var v interface{}
	url := cfg.ServiceURL("nym")
	if v, err = sdk.fetch.Post(url, "/lookup", user); err != nil {
		return
	}

	err = common.MapToInterface(v.(map[string]interface{}), user)
	return
}

// Key will allow for the uploading of new generated hardened keys from a
// user.  Keys are currently manually managed by a user but will be automated
// in some of the protocols in the future like secure messaging.
func (sdk *SDK) Key(key *common.Key) (err error) {
	log.D("sdk", "nym", "key", key)
	url := cfg.ServiceURL("nym")
	_, err = sdk.fetch.Post(url, "/key", key)
	return
}

// Keys will return all of the keys for the logged in user or return an error
// if there are no keys currently uploaded.
func (sdk *SDK) Keys() (keys []*common.Key, err error) {
	log.D("sdk", "nym", "keys")
	var v interface{}
	url := cfg.ServiceURL("nym")
	if v, err = sdk.fetch.Get(url, "/keys"); err != nil {
		return
	}

	keys = make([]*common.Key, 0)
	for _, o := range v.([]interface{}) {
		key := o.(map[string]interface{})
		createdAt, err := time.Parse("2006-01-02T15:04:05.999999999Z", key["createdAt"].(string))
		if err != nil {
			log.E("sdk", "keys", "createdAt", key)
			continue
		}

		keys = append(keys, &common.Key{
			CreatedAt: createdAt,
			Name:      key["name"].(string),
			PublicKey: key["publicKey"].(string),
			Used:      key["used"].(bool),
		})
	}
	return
}

// Update the logged in user with the provided user information or error.
func (sdk *SDK) Update(user *common.User) (err error) {
	log.D("sdk", "nym", "update", user)
	url := cfg.ServiceURL("nym")
	_, err = sdk.fetch.Post(url, "/update", user)
	return
}

func (sdk *SDK) Forgot(user *common.User) (otp string, err error) {
	log.D("sdk", "nym", "forgot", user)
	var v interface{}
	url := cfg.ServiceURL("nym")
	if v, err = sdk.fetch.Post(url, "/forgot", user); err != nil {
		return
	}

	otp = v.(string)
	return
}

// Reset a user's authentication information using the provided OTP.  The OTP
// is returned by making a Forgot request.
func (sdk *SDK) Reset(otp string, user *common.User) (err error) {
	log.D("sdk", "nym", "reset", otp, user)
	req := map[string]interface{}{
		"otp":  otp,
		"user": user,
	}
	url := cfg.ServiceURL("nym")
	_, err = sdk.fetch.Post(url, "/reset", &req)
	return
}

// Search with the provided term for all matching users and returned a summarized
// list of those users.  If a user key is desired use the specific information a
// a user and make a Lookup request.
func (sdk *SDK) Search(term string) (users []*common.User, err error) {
	log.D("sdk", "nym", "search", term)
	var v interface{}
	req := map[string]interface{}{
		"term": term,
	}
	url := cfg.ServiceURL("nym")
	if v, err = sdk.fetch.Post(url, "/search", &req); err != nil {
		return
	}

	users = make([]*common.User, 0)
	for _, o := range v.([]interface{}) {
		user := o.(map[string]interface{})
		users = append(users, &common.User{
			Domain:   user["domain"].(string),
			Username: user["username"].(string),
		})
	}
	return
}

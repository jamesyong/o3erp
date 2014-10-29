package sessions

import (
	"github.com/gorilla/sessions"
)

const PARTY_ID string = "partyId"
const USER_LOGIN_ID string = "userLoginId"
const WEB_SOCKET string = "webSocket"

var SessionStore = sessions.NewCookieStore([]byte("something-very-secret"))

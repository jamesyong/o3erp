package sessions

import (
	"github.com/gorilla/sessions"
)

const PARTY_ID string = "partyId"
const WEB_SOCKET string = "webSocket"

var SessionStore = sessions.NewCookieStore([]byte("something-very-secret"))

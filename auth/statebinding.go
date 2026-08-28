package auth

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"net/http"
	"strings"
)

// Binding the Google federation leg to the browser that started it.
//
// Without this, `state` is only a server-side lookup key that anyone can mint.
// An attacker registers a client (registration is open, as MCP requires), calls
// /authorize to get a state of their choosing, and lures a victim into
// completing Google consent with that state. Google returns the victim's code
// to our callback; we exchange it, see the *victim's* identity — which passes
// the domain allowlist, because the victim is legitimately allowed — and hand
// the resulting authorization code to the *attacker's* registered redirect_uri.
// The attacker then completes the token exchange with the PKCE verifier they
// chose, and holds an access token backed by the victim's Google credentials.
//
// At /authorize we set an opaque, HttpOnly cookie and store only its SHA-256
// next to the state row. At the callback we require a cookie whose hash
// matches. The victim's browser never visited the attacker's /authorize, so it
// carries no such cookie and the flow is refused.
//
// Residual risk, deliberately accepted: the defence assumes an attacker cannot
// plant a binding cookie they know into the victim's browser for our host.
// Cookies have no origin or scheme integrity, so an attacker holding a sibling
// subdomain, or able to intercept plain http to any host under the parent
// domain, could set a Domain-scoped cookie the victim's browser would send
// here. A `__Host-` prefix would close that, but it mandates `Path=/` and
// `Secure`, which conflicts with scoping the cookie to the callback and with
// running locally over http. Tracked separately rather than decided here; the
// bar is still far above the pre-fix state, which needed no cookie at all.
const (
	bindingCookieName   = "gtm_fed_binding"
	bindingCookieMaxAge = 600
)

func hashBinding(value string) string {
	sum := sha256.Sum256([]byte(value))
	return hex.EncodeToString(sum[:])
}

// bindingMatches compares in constant time and treats either value being
// absent as a mismatch, so a state row without a recorded binding can never be
// satisfied by a request without a cookie.
func bindingMatches(cookieValue, expectedHash string) bool {
	if cookieValue == "" || expectedHash == "" {
		return false
	}
	return subtle.ConstantTimeCompare([]byte(hashBinding(cookieValue)), []byte(expectedHash)) == 1
}

// setBindingCookie issues the binding to the browser starting the flow.
//
// SameSite=Lax is deliberate: Google's redirect back to us is a cross-site
// top-level GET navigation, which Lax permits and Strict would drop. The
// cookie is marked Secure whenever the issuer we resolved for this request is
// https, so a plain-http local run still works.
func setBindingCookie(w http.ResponseWriter, value, issuer string) {
	http.SetCookie(w, &http.Cookie{
		Name:     bindingCookieName,
		Value:    value,
		Path:     "/oauth/callback",
		MaxAge:   bindingCookieMaxAge,
		HttpOnly: true,
		Secure:   strings.HasPrefix(issuer, "https://"),
		SameSite: http.SameSiteLaxMode,
	})
}

func bindingFromRequest(r *http.Request) string {
	c, err := r.Cookie(bindingCookieName)
	if err != nil {
		return ""
	}
	return c.Value
}

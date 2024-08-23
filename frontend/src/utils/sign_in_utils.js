import store from "@/store"

/** Redirects user to the correct google sign in page */
export const signInGoogle = ({
  state = null,
  selectAccount = false,
  requestCalendarPermission = false,
  requestContactsPermission = false,
  loginHint = "",
}) => {
  const clientId =
    "523323684219-jfakov2bgsleeb6den4ktpohq4lcnae2.apps.googleusercontent.com"
  const redirectUri = `${window.location.origin}/auth`

  let scope = "openid email profile "
  if (requestCalendarPermission) {
    scope +=
      "https://www.googleapis.com/auth/calendar.calendarlist.readonly https://www.googleapis.com/auth/calendar.events.readonly "
  }
  if (requestContactsPermission) {
    scope +=
      "https://www.googleapis.com/auth/contacts.readonly https://www.googleapis.com/auth/directory.readonly "
  }
  scope = encodeURIComponent(scope)

  let stateString = ""
  if (state !== null) {
    state = encodeURIComponent(JSON.stringify(state))
    stateString = `&state=${state}`
  }

  let promptString = ""
  if (selectAccount) {
    promptString = "&prompt=select_account+consent"
  } else {
    promptString = "&prompt=consent"
    if (loginHint.length > 0) {
      promptString += `&login_hint=${loginHint}`
    } else if (store.state.authUser) {
      promptString += `&login_hint=${store.state.authUser.email}`
    }
  }

  const url = `https://accounts.google.com/o/oauth2/v2/auth?client_id=${clientId}&redirect_uri=${redirectUri}&response_type=code&scope=${scope}&access_type=offline${promptString}${stateString}&include_granted_scopes=true`
  window.location.href = url
}

export const signInOutlook = ({
  state = null,
  requestCalendarPermission = false,
}) => {
  const clientId = "d27c1c46-4be7-45c4-ad98-626b2fa3a527"
  const tenant = "common"
  const redirectUri = encodeURIComponent(`${window.location.origin}/auth`)

  let scope = "offline_access User.Read"
  if (requestCalendarPermission) {
    scope += " Calendars.Read.Shared"
  }
  scope = encodeURIComponent(scope)

  let stateString = ""
  if (state !== null) {
    state = encodeURIComponent(JSON.stringify(state))
    stateString = `&state=${state}`
  }

  const url = `https://login.microsoftonline.com/${tenant}/oauth2/v2.0/authorize?client_id=${clientId}&response_type=code&redirect_uri=${redirectUri}&response_mode=query&scope=${scope}${stateString}`
  window.location.href = url
}

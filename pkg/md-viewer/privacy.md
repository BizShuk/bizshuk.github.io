# MD Viewer Privacy Policy

Effective August 10, 2026  
Last updated August 10, 2026

MD Viewer is a Markdown reader and editor for GitHub repositories. This policy
explains what the App accesses, where information is stored, and the choices
available to you.

For the styled version of this policy, visit
[Privacy Policy](https://bizshuk.github.io/pkg/md-viewer/privacy.html).

## Privacy at a glance

- **No tracking.** The App contains no advertising, analytics, cross-app
  tracking, or data-broker integrations.
- **Direct connection.** GitHub requests travel directly between your device
  and GitHub. MD Viewer does not operate an intermediary backend.
- **Local control.** Credentials, settings, navigation history, and cached
  content are kept on your device and can be removed by you.

## 1. Scope

This Privacy Policy applies to the MD Viewer iPhone application (the "App"). It
covers the App's GitHub sign-in, repository browsing, Markdown viewing and
editing, local caching, reader preferences, and on-device document translation
features.

The App is designed as a client for content you choose to access. The developer
does not run a server that receives your GitHub credentials, repository content,
or App activity.

## 2. Information the App handles

MD Viewer handles only the information needed to provide the features you
request. Depending on how you use the App, that may include:

- **GitHub account information.** Your GitHub account identifier, username,
  installation list, and repository access metadata are requested from GitHub
  to establish the correct account scope and show repositories available to
  you.
- **Authentication information.** GitHub OAuth access and refresh tokens are
  obtained through GitHub Device Flow. They are held in the iOS Keychain and
  are not exposed to the App's web interface or sent to the developer. During
  sign-in, the temporary GitHub verification code is placed on the system
  clipboard so you can paste it on GitHub's verification page.
- **Repository content.** Repository names, branches, file trees, Markdown
  files, `README.todo`, linked assets, and related Git metadata are accessed
  when needed to display or cache your selected repository.
- **Content you change.** Markdown edits, TODO changes, and commit information
  you submit are sent directly to GitHub and become part of the repository you
  selected.
- **Device-local preferences.** Reader style, custom colors, interface
  language, repository selection, navigation history, cache settings, and
  cached repository content are stored locally to preserve your App experience
  and support faster or offline reading.
- **Translation content.** On supported devices, readable Markdown text is
  processed by Apple's on-device Translation framework. Translations remain in
  memory for the current App session and are not written back to your
  repository.

MD Viewer does not request your contacts, photos, camera, microphone, health
data, precise location, payment information, or advertising identifier.

## 3. How information is used

The App uses the information above only to:

- authenticate your GitHub session and maintain access;
- list repositories and verify their granted permissions;
- load, render, search, navigate, and cache selected content;
- save edits and commits you explicitly submit;
- restore your preferences and reading position; and
- translate eligible text on supported Apple devices.

MD Viewer does not use your information for advertising, marketing, user
profiling, analytics, or sale. It does not combine your information with data
from other apps or services for tracking.

## 4. Sharing and external services

### GitHub

The App communicates with GitHub to sign you in and perform the repository
actions you choose. GitHub processes those requests under its own
[Privacy Statement](https://docs.github.com/en/site-policy/privacy-policies/github-general-privacy-statement).
Access is limited by your GitHub permissions and the repositories authorized
for the MD Viewer GitHub App.

### Apple

iOS provides Keychain, clipboard, Dynamic Type, WebView storage, and on-device
Translation capabilities used by the App. Apple's handling of information
through its platform is governed by Apple's policies and your device settings.

MD Viewer does not share information with advertising networks, analytics
providers, or data brokers. If you follow a link to an external website, that
website's own privacy practices apply.

MD Viewer only uses external services needed to provide the features you
request. Any service provider engaged to process information for the App must
provide privacy protections consistent with this policy and applicable law.
The developer does not authorize those providers to use App information for MD
Viewer advertising or marketing.

## 5. Storage and retention

### Credentials

GitHub OAuth tokens remain in the iOS Keychain until you sign out, the
credentials expire, or you revoke the App's authorization in GitHub. Signing
out removes the local Keychain credentials.

### Preferences and navigation

Reader settings, interface language, selected repository, and navigation
coordinates stay on your device until you clear them, sign out where
applicable, or remove the App.

### Repository cache

Repository content may be cached on your device. With offline retention turned
off, repository caches expire after one day and are checked periodically. With
offline retention turned on, automatic expiry is disabled so content remains
available until you clear the cache, sign out, or remove the App. On-demand
file content is also subject to a size-bounded local cache.

### GitHub records

Content and commits you submit are stored by GitHub as part of the selected
repository. Their retention and deletion follow your GitHub repository settings
and GitHub's policies.

## 6. Your privacy choices

You control the App's access and local data. You may:

- choose which repositories the GitHub App can access;
- turn offline retention on or off;
- clear the local repository cache from App settings;
- sign out to remove credentials, history, and cached content;
- revoke MD Viewer from your
  [GitHub App installations](https://github.com/settings/installations); and
- remove the App to delete its remaining device-local data.

To remove content already committed to GitHub, use GitHub's repository controls
or contact the relevant repository owner.

Depending on where you live, you may also have rights to access, correct,
delete, restrict, or obtain a copy of personal information. The developer does
not maintain an MD Viewer account database or a server-side copy of your App
data. Use GitHub's privacy controls for information held by GitHub, or contact
the developer if you believe the developer holds personal information about
you.

## 7. Security

MD Viewer uses GitHub's HTTPS endpoints for network communication and stores
OAuth credentials using the iOS Keychain. Credentials are separated from
repository caches and are never written into repository content, App logs, or
this website.

No method of storage or transmission is completely secure. Keep your device
and GitHub account protected, review the repositories granted to the App, and
revoke access if you believe your account or device has been compromised.

## 8. Children's privacy

MD Viewer is a developer and productivity tool and is not directed to children
under 13. The developer does not knowingly collect personal information from
children. GitHub account eligibility is governed by GitHub's terms.

## 9. Changes to this policy

This policy may be updated when MD Viewer's features, data practices, or legal
obligations change. The effective date at the top of this page will be revised
when an update is published.

## 10. Contact

For privacy questions, support, or requests concerning MD Viewer, contact the
developer through the
[public support issue tracker](https://github.com/BizShuk/bizshuk.github.io/issues/new?title=MD%20Viewer%20Privacy%20Request).

Do not include access tokens, private repository content, or other secrets in a
public issue.

Copyright 2026 MD Viewer.

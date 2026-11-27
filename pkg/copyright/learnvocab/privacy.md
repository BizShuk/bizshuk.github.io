# LearnVocab Privacy Policy

Effective November 27, 2026
Last updated November 27, 2026

LearnVocab shows study cards for language, economy, appearance, and coding,
including on the Lock Screen. Spoken playback uses the on-device speech
synthesizer and is not uploaded. Card text and images are requested from an
stts server you configure. The default address is `https://stts.shuks.dev:443`.
The developer does not create a LearnVocab account and does not operate an
advertising or analytics service inside the App.

## 1. Scope

This policy covers the iPhone app, its Lock Screen widget, and the Mac app:
the server address and API key you enter, the study prompt built from the
subject you select, the cards and images that come back, on-device speech,
and the local card history.

## 2. Information the App handles

The App handles the stts address and API key you enter, the subject, level,
and categories you select, the prompt built from those settings, and the card
text and images returned by that server. The key is stored in the device
Keychain. Cards, a history of up to 1000 cards for each subject, and
downloaded images are stored in the App's shared container. Other settings
are stored in local app settings. The App does not request contacts, photos,
camera, microphone, location, health, payment, or an advertising identifier.
Economy cards are study material. The App does not ask for your financial
accounts.

## 3. How information is used

Information is used to request and show the next cards, speak a card on the
device, and keep a local review history. The developer does not use it for
analytics, advertising, profiling, or sale.

## 4. Sharing and external services

Requests go only to the stts address you set. Until you change it, that
address is `https://stts.shuks.dev:443`, a server operated in connection with
this developer. The App sends the study prompt with your API key to
`POST /v1/chat/completions` and, when a card image is requested,
`POST /v1/images/generations`. The request names the models `MiniMax-M3` and
`MiniMax-H3-Image-Highspeed`. The App does not bundle those models and does
not contact their provider itself.

This policy does not state how long `stts.shuks.dev` keeps prompts, logs, or
generated images. That retention is not established by the App. If you point
the App at a different address, that operator's practices apply instead, and
this policy does not describe them.

## 5. Storage and retention

The API key is stored in the device Keychain. Study settings are stored in
local app settings. Card text, the per-subject history, and images stay in
the App's shared container until you remove the App, which removes that local
data. This policy does not promise deletion of prompts or images held by the
server you configured.

## 6. Your privacy choices

You choose the server address, the API key, the subject, and the level. You
can clear the key and the address in Settings, or remove the App. LearnVocab
does not create an account, so there is no LearnVocab account to delete. The
App does not include a control that deletes prompts from the server. Do not
put an API key or private study content in a public support issue.

## 7. Security

Connections use TLS. For `stts.shuks.dev`, the App also trusts one certificate
authority bundled with the App, and it still uses the system trust evaluation
for the host name and the certificate period. It does not accept every
certificate. The API key stays in the Keychain. Do not post it in public.

## 8. Children's privacy

LearnVocab is a general-audience study app. It is not in the Kids Category
and it does not create an account for a child. A parent or guardian should
decide whether a child may use it and which server receives study prompts.

## 9. Changes to this policy

This policy may change when App features, data practices, or legal obligations
change. The effective date will be revised when a new version is published.

## 10. Contact

Use the [public support issue tracker](https://github.com/bizshuk/bizshuk.github.io/issues)
for privacy and support requests. Never post an API key, a private prompt, or
other credential in a public issue.

Copyright 2026 BizShuk.

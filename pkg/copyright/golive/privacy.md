# golive Privacy Policy

Effective November 27, 2026
Last updated November 27, 2026

golive is a voice tool for iPhone and Mac. You supply your own Gemini API key.
The developer does not operate a golive account, advertising system, or
analytics service. When you start dictation, translation, or a live
conversation, microphone audio and the session text needed for that request
are sent to Google Gemini Live.

## 1. Scope

This policy covers the iPhone app, its keyboard, and the Mac menu-bar app:
microphone capture, transcription, translation, live conversation, the local
text history, and placing resulting text at the cursor or on the clipboard.

## 2. Information the App handles

The App handles microphone audio while a session you start is running;
transcripts, translations, live replies, and error text; the Gemini API key
you paste; and settings such as language, model, and voice. On a Mac, the App
reads the current clipboard so it can insert text and then restore what was
there. On iPhone and Mac, you can copy a history record onto the clipboard.
The App does not request contacts, photos, location, health, payment, or an
advertising identifier. It does not write microphone audio into its history
file.

## 3. How information is used

Information is used to transcribe speech, translate it, hold a live
conversation you start, show a local text history, and place the resulting
text at the cursor or on the clipboard. The developer does not use it for
analytics, advertising, profiling, or sale.

## 4. Sharing and external services

The App sends audio and the session setup over the operating system's TLS to
Google at `generativelanguage.googleapis.com`, using the API key you supply.
The session setup can include the selected model, language hints, custom
vocabulary, the translation target, and any live instruction you type. Live
mode also asks Google for a resumption handle so that dropped session can
reconnect. The handle is kept for the running session.

Google's logging, abuse monitoring, training, and retention depend on the
terms that apply to your key and project. Paid and unpaid keys are not
described the same way in Google's own documents, and a project owner can
change logging. This policy does not promise a retention period for audio or
text held by Google. Read the current documents:

- [Gemini API terms](https://ai.google.dev/gemini-api/terms)
- [Gemini API logs policy](https://aistudio.google.com/docs/logs-policy)
- [Google Privacy Policy](https://policies.google.com/privacy)

golive does not send the audio or the history to a server operated for this
App. There is no golive account.

## 5. Storage and retention

The API key is stored in the device Keychain. Settings are stored in local
app settings. Transcript text is stored on the device, at most 500 records.
On iPhone that file is in the App's shared container. On a Mac it is
`transcript_history.json` under `.config/golive/data` in your home directory.
Removing the App removes this local data. A copied record can remain on the
clipboard. Records held by Google follow Google's terms for your key, which
this policy does not restate as a fixed period.

## 6. Your privacy choices

You choose whether to allow the microphone and keyboard access, whether to
enter a Gemini API key, and which mode to use. You can clear the key in
Settings, change the system permissions, or remove the App. golive does not
create an account, so there is no golive account to delete. To manage data
Google holds for your key, use the controls in your Google project. Do not
put an API key or a private transcript in a public support issue.

## 7. Security

The API key stays in the Keychain. The connection uses the operating system's
TLS. golive does not add its own end-to-end encryption. Treat the key and
private transcripts as sensitive, and do not post them in public.

## 8. Children's privacy

golive is not in the Kids Category and is not directed to children under 13.
It does not create an account for a child. A parent or guardian should decide
whether a child may use the microphone feature.

## 9. Changes to this policy

This policy may change when App features, data practices, or legal obligations
change. The effective date will be revised when a new version is published.

## 10. Contact

Use the [public support issue tracker](https://github.com/bizshuk/bizshuk.github.io/issues)
for privacy and support requests. Never post an API key, private transcript,
or other credential in a public issue.

Copyright 2026 BizShuk.

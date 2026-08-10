# Push Privacy Policy

Effective August 10, 2026
Last updated August 10, 2026

Push is an iOS notification workbench. The current App does not operate a
developer backend, account service, advertising system, or analytics service.

## 1. Scope

This policy covers notification authorization, local notification examples,
remote notification registration and callbacks, APNs device-token display and
copying, notification actions, user replies, and the local arrival journal.

## 2. Information the App handles

The App handles notification permission status, notification content and
metadata, action selections, reply text entered through a notification, the
APNs device token assigned to the installation, and a local journal of arrival
events. The full token appears only when you explicitly copy it; the interface
otherwise uses a fingerprint. The App does not request contacts, photos,
camera, microphone, location, health, payment, or advertising identifiers.

## 3. How information is used

Information is used only to schedule and present notifications, respond to
notification actions, display registration and delivery state, and maintain a
bounded local diagnostic journal. The developer does not use it for analytics,
advertising, profiling, or sale.

## 4. Sharing and external services

Apple Push Notification service assigns and transports the device token and
remote notification payloads according to Apple's platform operation. The
current App does not send the token or journal to a developer server. If you
copy the full token and give it to a sender, that sender receives it and may
send notification content through APNs under its own practices.

## 5. Storage and retention

The most recent APNs device token and the bounded arrival journal are stored in
device-local settings. Notification content may also remain in Notification
Center until removed by you or the system. Removing the App removes its local
App data, while records held by an independent sender follow that sender's
retention policy.

## 6. Your privacy choices

You choose whether to authorize notifications, trigger local examples, copy a
full token, reply to an action, and provide a token to a sender. You may change
notification permissions in iOS Settings, clear pending or delivered
notifications and the local journal where provided, or remove the App. Contact
an independent sender to manage records it holds.

## 7. Security

Treat an APNs device token as sensitive routing information. The App masks it
by default, but copying or sharing it moves it outside the App's control. Do not
post tokens, payloads, or private notification content in public support.

## 8. Children's privacy

Push is a developer and testing utility and is not directed to children under
13. The developer does not knowingly collect children's information through an
App-operated service.

## 9. Changes to this policy

This policy may change when App features, data practices, or legal obligations
change. The effective date will be revised when published.

## 10. Contact

Use the [public support issue tracker](https://github.com/bizshuk/bizshuk.github.io/issues)
for privacy and support requests. Never post device tokens, private payloads,
reply text, or credentials in a public issue.

Copyright 2026 BizShuk.

# iPhone Sync Privacy Policy

Effective August 10, 2026
Last updated August 10, 2026

iPhone Sync transfers selected Photo Library media directly to a paired
computer on the same local network. The developer operates no cloud relay or
account service for this transfer.

## 1. Scope

This policy applies to the iPhone and iPad source App and the paired macOS
receiver. It covers album selection, local-network pairing, media transfer,
transfer history, automatic sync settings, and optional deletion after sync.

## 2. Information the App handles

The App handles selected photo and video resources, album names, asset
identifiers, media metadata needed for faithful export, receiver identity,
pairing secrets, transfer status, destination bookmarks, and local operation
logs. It does not request contacts, camera, microphone, location, health,
payment, or advertising identifiers. Media available only through iCloud is
skipped because network download is disabled.

## 3. How information is used

Information is used only to discover and pair devices, read selected media,
transfer original resources, verify receiver acknowledgement, resume work, and
show local status. The developer does not use it for analytics, advertising,
profiling, or sale.

## 4. Sharing and external services

Media travels directly over TLS to the computer you pair on the same local
network. Pairing uses a pre-shared secret stored in system credential storage.
The developer does not receive the media or transfer records. Apple provides
PhotoKit, Bonjour, Keychain, background scheduling, and other platform
services. Destination files may be processed by software you choose on the
receiver computer.

## 5. Storage and retention

The source device stores selected albums, receiver configuration, automatic
sync settings, and pending deletion state locally. Pairing credentials are
stored in Keychain. The receiver stores a local transfer manifest, destination
bookmark, completed media, and incomplete transfer files in the destination
you choose. Local logs are bounded. Destination media remains until you delete
it; removing the source App does not delete receiver files.

## 6. Your privacy choices

You choose albums, destination, receiver, automatic sync, and whether Delete
After Sync is enabled. You may forget a pairing, reset the source, change the
destination, clear local logs, remove the App, or delete destination files.
Delete After Sync is off by default and requires system confirmation. Library
deletion may also affect devices synchronized through iCloud Photos.

## 7. Security

Pairing is explicit and transfer uses TLS 1.2 with a pre-shared secret. Keep
both devices and the local network protected. No storage or transmission
method is completely secure, and a receiver computer may expose copied files
according to its own permissions and backup settings.

## 8. Children's privacy

iPhone Sync is a personal backup utility and is not directed to children under
13. The developer does not knowingly collect children's personal information
through an App-operated service.

## 9. Changes to this policy

This policy may change when App features, data practices, or legal obligations
change. The effective date will be updated when a revision is published.

## 10. Contact

Use the [public support issue tracker](https://github.com/bizshuk/bizshuk.github.io/issues)
for privacy and support requests. Never post personal media, pairing secrets,
device identifiers, or private logs in a public issue.

Copyright 2026 BizShuk.

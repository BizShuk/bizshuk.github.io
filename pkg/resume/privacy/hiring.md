# Hiring Privacy Compliance / 招募隱私法遵

What a multinational employer must tell a job candidate, broken down by the law that
forces each disclosure. Worked example: the TSMC candidate privacy notice (read
2026-09-09), which is used throughout as the reference specimen because it stacks
six regimes into one document.

跨國雇主在招募流程中必須向應徵者揭露的事項, 依`法源`拆解. 範例取自 TSMC 應徵者隱私
通知 —— 它把六套法規疊在同一份文件裡, 是很好的教材.

**This is a reading aid, not legal advice.** Every citation below is traceable to a
named statute or article so it can be checked against the primary text.

---

## 1. Where the rules come from / 法源總覽

| # | Jurisdiction | Trigger / 觸發條件 | Legal source | Appears in the notice as |
|---|---|---|---|---|
| 1 | EU (EEA-wide) | Applying to an entity established in the EEA | GDPR (EU) 2016/679, Art. 13 (notice), Art. 6 (lawful basis), Art. 9 (special categories) | "Additional Notice in relation to the EU and the UK" |
| 2 | Netherlands | Controller is TSMC Europe B.V. (Amsterdam) | GDPR + Dutch UAVG; retention practice from NVP Sollicitatiecode | Named controller, DPO `privacy_tsmceu@tsmc.com`, Dutch AP as supervisory authority |
| 3 | Germany | Controller is ESMC GmbH (Dresden) | GDPR + BDSG §26(2) (consent in employment), §26(3) (special categories) | Named controller, DPO `privacy_esmc@esmc.eu`, Saxon DPA |
| 4 | UK | Applying to a UK establishment | UK GDPR + Data Protection Act 2018 | Named in the heading only — `no UK controller, no ICO reference` |
| 5 | Japan | Applying to a TSMC Japan entity | APPI (個人情報保護法) Art. 27(5)(iii) — joint use / 共同利用 | "Additional Notice in relation to Japan", the five-item block |
| 6 | US federal | Any US application | ADA, 42 U.S.C. §12112(b)(5) — accommodation in the application process | "Please inform TSMC North America HR if you require an ADA accommodation" |
| 7 | California | Applicant resides in CA | CCPA/CPRA — applicant exemption expired 2023-01-01; §1798.100(a) notice at collection | "TSMC does not sell nor share your personal information" |
| 8 | Massachusetts | Position located in MA | M.G.L. c.149 §19B(2) — mandatory lie-detector notice on every application | The lie-detector paragraph, quoted near-verbatim from the statute |
| 9 | Taiwan | Data lands in the TSMC HQ candidate database | 個資法 Art. 8 (告知義務), Art. 21 (跨境傳輸) — `inferred, not cited in the notice` | Only implied by "transferred to Taiwan Semiconductor Manufacturing Company Limited" |

Rule of thumb: `法規跟著法人實體與應徵者所在地走`, 不是跟著職缺標題. 同一份通知對
阿姆斯特丹與德勒斯登的應徵者, 適用的責任主體與保存期限就已經不同.

---

## 2. What each regime actually requires

### 2.1 EU — GDPR baseline

The mandatory content of a candidate-facing notice comes from Art. 13. Five things
must be present, and all five are traceable in the specimen:

| Requirement | GDPR | Specimen |
|---|---|---|
| Identity + contact of the controller | Art. 13(1)(a) | Full legal name and registered address per entity |
| DPO contact | Art. 13(1)(b) | One mailbox per controller |
| Purpose **and** lawful basis, per purpose | Art. 13(1)(c) | Stated separately for account creation, events, programs, applications, marketing |
| Recipients + third-country transfers with the safeguard named | Art. 13(1)(e)(f) | 25+ affiliates listed; SCCs named, copy on request |
| Retention period, rights, right to complain | Art. 13(2)(a)(b)(d) | 6 months DE / 4 weeks NL; AP and Saxon DPA named |

Lawful bases used, and why they differ:

- `Art. 6(1)(b)` performance of a contract — the core application processing. Answering
  "can we hire you" is a pre-contractual step, so consent is not needed and cannot be
  withdrawn to stop it.
- `Art. 6(1)(f)` legitimate interest — website accounts, event registration, campus
  programs. `Carries a right to object` under Art. 21.
- `Art. 6(1)(a)` consent — talent-network retention beyond the rejection date, and all
  direct marketing. `Withdrawable at any time` under Art. 7(3).
- `Art. 9(2)(b)` + BDSG §26(3) — special categories, in practice disability status
  volunteered by the candidate.

這三種法律依據`不能互換`. 依 contract 處理的部分你撤不掉 (撤了等於退出應徵);
依 consent 處理的部分你隨時撤得掉, 且撤回不影響應徵本身.

### 2.2 Netherlands — TSMC Europe B.V.

- Controller: TSMC Europe B.V., WTC Amsterdam Tower H7, Zuidplein 60, 1077 XV Amsterdam.
- Supervisory authority: Autoriteit Persoonsgegevens.
- Retention if not hired: `~4 weeks` after the procedure ends. This is the Dutch
  recruitment-code norm rather than a statutory deadline, and it is the shortest
  default in the whole document.

### 2.3 Germany — ESMC GmbH

- Controller: ESMC GmbH, Rosenstrasse 32, 01067 Dresden.
- Supervisory authority: Saxon Data Protection and Transparency Officer.
- Retention if not hired: `~6 months`. Derived from AGG §15(4) — a discrimination claim
  must be raised within 2 months and litigated within 3, so 6 months covers the
  defence window under Art. 17(3)(e).
- BDSG §26(2) adds a German-specific gate: consent in an employment context is only
  valid if `voluntariness` can be shown, given the power imbalance. §26(3) governs
  special-category data such as severe-disability status.

### 2.4 UK — the gap

The heading says "EU and the UK", but the body names `no UK controller`, `no UK
representative`, and `no ICO complaint route`, and the transfer section cites EU SCCs
without the `UK International Data Transfer Addendum`. If you apply to a UK-based role,
this is the first thing to ask about.

### 2.5 Japan — APPI joint use

Japan's distinctive mechanism is `共同利用` (joint use), APPI Art. 27(5)(iii): a group of
companies may share personal data without treating each other as third parties, provided
five items are disclosed in advance. The specimen discloses all five:

| # | Required item | Specimen |
|---|---|---|
| 1 | The fact of joint use | "those companies may jointly use your personal information" |
| 2 | Categories of data | Section (B) of the Privacy Policy |
| 3 | Scope of joint users | TSMC Ltd. + 4 named Japanese entities |
| 4 | Purpose of joint use | Section (F) of the Privacy Policy |
| 5 | Responsible entity, address, and representative's name | TSMC Ltd., Hsinchu; C.C. Wei |

Item 5 including the `representative's name` is the tell that the notice was updated for
the 2020 APPI amendment (in force 2022-04-01). An older notice would name the company only.

Two clauses in the Japan block are `contract terms, not privacy disclosures`, and deserve
separate attention: an open-ended auto-extending retention term, and a waiver of "all of
TSMC's liabilities" except gross negligence or intentional misconduct. A liability waiver
does not remove APPI's statutory obligations or the regulator's authority.

### 2.6 US federal — ADA

Not a privacy rule. 42 U.S.C. §12112(b)(5) requires a reasonable accommodation in the
`application process` itself, which is why the contact and phone number sit inline. The
at-will statement alongside it is employment law, included because a written application
can otherwise be argued to imply a term.

### 2.7 California — CCPA/CPRA

The applicant/employee exemption expired on 2023-01-01, so candidates now hold full
consumer rights. What is required:

- A `notice at collection` — §1798.100(a) — stating categories collected, purposes, and
  retention, at or before collection.
- A statement on `selling` and `sharing` (cross-context behavioural advertising) —
  §1798.120. The specimen's denial of both is the disclosure.
- `Sensitive personal information` handling — §1798.121, the right to limit use.
- Rights to know, delete, and correct — §1798.110, §1798.105, §1798.106.

### 2.8 Massachusetts — lie detector notice

M.G.L. c.149 §19B(2) requires this notice on `every` job application used in the state,
in substantially the statutory wording. Its presence indicates the form is a single
global template rather than a per-state variant.

### 2.9 Taiwan — inferred, not cited

The notice never cites Taiwan's 個人資料保護法, yet routes the candidate database, the
website, and backups to the Taiwan parent. Under 個資法 Art. 8 a collector must state the
purpose, categories, period, scope of use, and the candidate's rights; Art. 21 governs
outbound transfer. `Treat this as an open question`, not as an established position.

---

## 3. The transfer chain / 跨境傳輸鏈

```
Candidate (EEA)
  → controller: TSMC Europe B.V. (NL) or ESMC GmbH (DE)
  → processors: LinkedIn / StepStone / XING, ATS cloud vendor, RPO and headhunters,
                background-check and IT-support vendors
  → group transfer: TSMC Ltd. (TW) + 25 named affiliates across TW/US/JP/KR/CN/SG/CA/DE
```

- Taiwan has `no EU adequacy decision`, so the safeguard is the EU Standard Contractual
  Clauses (2021/914) under Art. 46(2)(c). The notice offers a copy on request — Art. 15(2)
  makes that an entitlement, not a courtesy.
- The 25-affiliate group transfer is `conditioned on consent`: it applies "if we have
  obtained your consent to store ... in the TSMC Talent Network". Declining that consent
  narrows the distribution list without affecting the application itself.
- Applying via a job platform means the `platform's own policy governs its side` of the
  processing. Two controllers, two notices, two sets of rights to exercise separately.

---

## 4. What "I Confirm" bundles

One button, four legally distinct acts. Worth separating before clicking:

| Act | Nature | Reversible? |
|---|---|---|
| Acknowledging the notice was read | Art. 13 information duty discharged | n/a |
| Certifying the application is true and complete | Contractual representation — grounds for withdrawal or dismissal if false | No |
| Consent to talent-network retention and future matching | GDPR Art. 6(1)(a) / BDSG §26(2) | `Yes` — Art. 7(3), effective for the future |
| Consent to direct marketing (job alerts, career events) | GDPR Art. 6(1)(a) | `Yes` — same |

Candidate-side actions worth taking: keep a dated copy of the notice as submitted;
request the SCCs in writing if the application is EEA-based; diary the retention deadline
(4 weeks NL / 6 months DE) and file an Art. 17 erasure request afterwards if no consent
was given; note that withdrawing talent-network consent does not withdraw the application.

---

## 5. Gaps in the specimen / 待釐清事項

1. `No UK controller, representative, ICO route, or IDTA/Addendum` despite the heading.
2. `No retention period for the US or Japan` — only the EU figures are quantified.
3. `Taiwan's own PDPA is never cited`, though Taiwan is the primary storage location.
4. `No statement on automated decision-making or AI screening` — GDPR Art. 22 requires
   disclosure where it exists, and several US jurisdictions now regulate it directly.
5. The Japan liability waiver and auto-extending retention term are `contract clauses
   embedded in a privacy notice`, which is not where a candidate expects to find them.

---

## 6. Not covered here / 其他區域主題

Regimes that would apply to a TSMC-group application but are absent from this notice.
Each is a candidate for its own file in this folder.

| Region | Instrument | Why it would matter |
|---|---|---|
| Korea | PIPA | TSMC Korea Limited is a named recipient; PIPA has its own consent formalities |
| China | PIPL | TSMC China / Nanjing are named; PIPL Arts. 38-39 govern outbound transfer |
| Singapore | PDPA | SSMC Pte. Ltd. is a named recipient |
| Canada | PIPEDA + Quebec Law 25 | TSMC Design Technology Canada Inc. is a named recipient |
| Illinois (US) | BIPA; AI Video Interview Act | Biometrics and AI-scored video interviews |
| New York City | Local Law 144 | Bias audit + notice for automated employment decision tools |
| Washington (US) | My Health My Data Act | Broad definition of health data, private right of action |

---

## Sources

- TSMC candidate privacy notice, read 2026-09-09 (Solutions Architect posting).
- Regulation (EU) 2016/679 (GDPR), Arts. 6, 7, 9, 13, 15, 17, 21, 22, 46.
- Bundesdatenschutzgesetz (BDSG) §26; Allgemeines Gleichbehandlungsgesetz (AGG) §15(4).
- 個人情報の保護に関する法律 (APPI) Art. 27(5).
- California Civil Code §§1798.100, 1798.105, 1798.106, 1798.110, 1798.120, 1798.121.
- M.G.L. c.149 §19B(2); 42 U.S.C. §12112(b)(5).
- 個人資料保護法 Arts. 8, 21 (Taiwan) — applied by inference, not cited by the notice.

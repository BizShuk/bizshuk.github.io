# Scope / 適用範圍

## What this dataset supports

- Within-market change across periods, same instrument
- Within-market ranking across sectors, same instrument, same period
- Direction of travel in the experienced-hire proxy indicators
- Locating the primary source behind any figure

可靠的是**同一市場、同一工具的時間序列與內部排序**。

## What it does not support

**Cross-market comparison of headline rates.** This dataset makes no attempt to
harmonise definitions across markets and provides no cross-walk. The figures are
recorded as each market publishes them.

本資料集**不做跨市場比較**，也不提供口徑換算。數字按各市場公布的原樣記錄。

The reason is that the headline graduate employment rates published by these seven
markets are not the same quantity, and harmonising them would require assumptions
this dataset cannot document honestly. Rather than encode a false equivalence, the
dataset states each market's definition in its own `.md` and stops there.

Illustrative, all describing broadly healthy markets:

| Market | Headline | Window | Base | Counts as employed |
|---|---|---|---|---|
| JP | ~98% | at/before graduation | job seekers only | any employment |
| SG | ~89% | 6 months | in labour force | any employment |
| US | ~86% | 6 months | known outcomes only | employment or further study |
| AU | ~75% | 4–6 months | in labour force | full-time permanent |

Japan appears 23 points ahead of Australia. Most of that gap is definitional:
Japan's base excludes roughly a quarter of the cohort who never declared an intent
to seek work, and its observation point sits at or before graduation. Australia
counts only full-time work. Neither difference says anything about which market is
easier to enter.

生活化比喻：三家餐廳都說「滿意度 95%」，但一家只問吃完沒抱怨的人，一家問所有進門的人，
一家只算願意再訪的人。數字相同，意義完全不同。

## Where the definitions live

In the `.md` half of each pair, under `## Metric definitions`. That section is
mandatory and machine-enforced: `validate.py` fails a pair if any populated field
in the `.json` goes unexplained.

The `.json` retains `window_months`, `denominator` and `numerator_includes` as
typed enums. Their purpose is **grouping and filtering**, and supporting the
validator's invariant that a rate without a stated base is rejected. They are not a
comparison mechanism, and no code in this repository ranks markets against each
other.

`.json` 裡保留這三個欄位，用途是**分組、過濾、以及讓 validator 擋掉「有數值卻沒分母」
的情況**，不是為了比較。

## If you need a comparison anyway

Do it downstream, in your own analysis, and write down the assumptions there. Do
not add a derived comparable value back into a `.json` — the structured layer holds
what the sources published, not what someone computed from it.

If a derivation is genuinely useful, state it in the `.md` prose and mark it as
derived.

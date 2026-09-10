# Autobiography — Liu, Teng-Yuan (Shuk Liu)

> Working draft, 2026-09-09. Voice: first person, plain, no résumé adjectives.
> Source of facts: `pkg/resume/Resume.md`. Update both together.

## The one-paragraph version

I am a systems person who has spent twelve years learning that the tools matter far
less than the patterns behind them. I have built poker games, backend-as-a-service
platforms, healthcare pipelines, an AI platform for a semiconductor giant, and data
centre foundations that had to survive regional compliance law. Between those jobs I
have cleaned motel rooms in Utah, argued with immigration officers in Ireland, and
sat in the wrong spot on a reef in Bali until I learned to read the swell. The same
instinct runs through all of it: find the real structure of a thing, then work with
it instead of against it.

## Act I — Hualien, Jongli, and the habit of building

I grew up in Taipei and went east to Hualien for a computer science degree at
National Dong-Hua, then north to Jongli for a master's at National Central. The most
formative thing I did in school was not a course. The university library's e-book
search was expensive and slow, so I built one that bypassed it. It saved north of
twenty thousand dollars a year in licensing and maintenance. Nobody assigned it. I
learned there that the fastest way to be useful is to look at the thing everyone
tolerates and ask why it has to be that way.

## Act II — Moab, and the year I was not an engineer

After graduating I took a J1 visa to Moab, Utah, and worked as a housekeeper at a
Motel 6. My colleagues were American, Mexican, Indian, Russian, Chinese and Turkish.
I stripped beds for five months in a desert town.

People sometimes read that line as a gap. I read it as the year I learned that a team
is a set of humans with different defaults, and that most coordination failure is not
a skill problem. Every organisation I have designed since has been shaped more by
that motel than by any architecture book.

## Act III — Games, data, and learning the whole loop

My first commercial job was at Gamesofa, a two-hundred-person studio in Taipei making
its own Mahjong, poker, FPS and RPG titles. I shipped a Windows 8 client app, then got
pulled into the backend team, then spent a year inside the marketing department doing
data modelling and A/B analysis. Three seats, one company. I optimised the time
complexity of an Omaha hand evaluator and I also sat with analysts asking what a user
group actually was.

That was the moment my career stopped being about a language and started being about
a loop: someone has a question, code answers it, the answer changes the product, the
product changes the question.

## Act IV — Infrastructure, and Git as the single source of truth

At Droi Tech I built the CI/CD spine of a backend-as-a-service platform — Jenkins,
Docker, Kubernetes, Terraform, Ansible. The decision I still defend is making Git the
single source of truth for environment state, so that every change to the platform
landed through review rather than through somebody's shell history.

Then I left. I took a working holiday visa to Ireland, discovered that the policy did
not actually recognise a permanent engineering role no matter how clearly the
immigration office explained it, spent eighty days there and thirty more crossing
Europe, and came home to study algorithms.

I do not present that as a detour. I present it as a data point about how I make
decisions: I would rather test an assumption in the real world at personal cost than
argue about it from a desk.

## Act V — Remote, regulated, and responsible for other people

Four and a half years at Change Healthcare, remote from Taipei into a US company. I
led a team of five, sat in more than thirty interview loops, and wrote the shared
frameworks other teams built on — an EventBus, a Spring Secret Manager plugin, event
annotation, Lambda scaffolding. I also went looking for the security problems nobody
had filed: JWT-signed unsubscribe links, KMS in the Jenkins pipeline, the Fortify
backlog.

Healthcare taught me that "it works" and "it is allowed to work" are two different
engineering problems, and the second one is usually harder.

## Act VI — TSMC, and leaving well

At TSMC I architected the AI platform and owned MLOps for the inference tier: a
scheduler that packs queued inference jobs onto a fixed GPU pool and admits the next
job the instant one finishes, so the accelerators never idle and the queue keeps
draining. I ran delivery as GitOps and held the team to functions under thirty lines
and coverage above ninety percent.

I was there five months. The culture and I did not fit, and I said so and left with
the support of both my managers and the department head. I include this deliberately.
I would rather be legible about a mismatch than quietly stay somewhere I am wrong for.

## Act VII — Singapore, and scale that has a legal shape

Three and a half years at ByteDance as a tech lead in Singapore, where I am now a
permanent resident. The work started by aligning internal and external services onto a
new data centre, then owning the wallet data architecture roadmap on top of it: a
deadline-driven EU migration, a data SDLC that made residency and retention a lifecycle
property rather than a per-project checklist, a flexible channel between data centres
that satisfied both compliance and stability, and global redundancy that cut
implementation and maintenance effort by two job grades. I designed a high-level task
orchestration framework and pushed a regional compliance migration framework into
adoption across departments.

At that scale, architecture is organisation design. You cannot draw a boundary
between two systems without also drawing one between two teams, and pretending
otherwise is how programmes die.

## Act VIII — Now: building alone, on purpose

Since mid-2026 I have been working for myself as an AI engineer. I work across
Gemini, Claude, GPT, Grok, MiniMax, DeepSeek, Qwen and ElevenLabs, and I built an
Agent SDK that hides all of them behind one interface, because betting a product on a
single vendor is a bet I have watched other people lose. I have shipped iOS tool
apps, a product service, a surfing tracker and two voice agent prototypes, and I run
my own infrastructure end to end — accounts, payments, cloud.

This is the first time in my career that the architecture, the code, the operations
and the consequences are all mine. It is the clearest feedback loop I have ever had.

## The part that is not a job

I surf. Out on the reef there is no roadmap and no backlog, and none of my seniority
paddles for me. The ocean keeps its own changelog — swell period, tide push, reef
shadow — and reading it honestly is the whole game. Getting the read wrong costs skin.

That is also, I think, the honest summary of how I work. Read the system as it is,
not as the documentation claims. Sit in the right spot. Commit early on the drop.

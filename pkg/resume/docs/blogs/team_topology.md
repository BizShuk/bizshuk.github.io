# Team Topologies — Reading Notes & Key Concepts

> An organization's architecture and team structure are intrinsically linked. Optimizing for fast flow, psychological safety, and sustainable cognitive load requires treating teams as the fundamental building block of software delivery.

---

## 📖 Executive Summary & Core Foundations

### The 4 Fundamental Team Topologies

| Team Type | Primary Mission & Characteristics |
| :--- | :--- |
| **Stream-Aligned Team** | Dedicated to continuous delivery of value along a single stream of work; fully responsive to feedback cycles. |
| **Complicated-Subsystem Team** | Responsible for building and maintaining components requiring deep, specialist domain expertise. |
| **Platform Team** | Provides underlying services and APIs to reduce cognitive load on stream-aligned teams. |
| **Enabling Team** | Cross-cuts domains to research, pilot, and transfer new capabilities/tools to accelerate stream-aligned teams. |

### 3 Types of Organizational Structures
1. **Formal Structure**: The official reporting lines and hierarchical org charts.
2. **Informal Structure**: The emergent social network and influence paths between individuals.
3. **Value Creation Structure**: How value actually flows across teams to customers.

### Key Drivers of High-Performing Teams
- **Motivation Triad**: Autonomy, Mastery, and Purpose (*Drive* by Daniel Pink).
- **Global vs. Local Optimization**: Local optimizations improve individual teams but do not necessarily improve end-to-end value delivery if systemic bottlenecks exist upstream or downstream.

---

## 🏛️ Part I: Teams as the Means of Delivery

### Conway's Law & Reverse Conway Maneuver

> *"Organizations which design systems are constrained to produce designs which are copies of the communication structures of these organizations."* — Melvin Conway

- **Organizational Architecture vs. System Architecture**: If the architecture of the system and the structure of the organization conflict, **the organizational structure always wins**.
- **The Pitfall of Functional Silos**: Segregating teams into functional silos (e.g., dedicated QA, DBA, or Security silos) impedes end-to-end flow and leads to brittle, fractured systems.
- **Reverse Conway Maneuver (Inverse Conway Maneuver)**:
  - Intentionally structure teams and communication paths to match the desired software architecture (e.g., align independent cross-functional teams around loosely-coupled microservices).

### Communication: Less is Often More

- **Restrict Unnecessary Communication**: Many organizations assume more communication is always better, but unrestricted communication creates tight coupling and paralysis.
- **Key Architectural Tests**:
  - Does the team structure minimize necessary cross-team communication paths?
  - Does the structure encourage communication between teams that genuinely need to collaborate?
- **Software Architecture Best Practices**:
  - Loose coupling & High cohesion.
  - Clear and well-defined version compatibility.
  - Structured, contract-driven cross-team integration testing.

---

## 🧠 Part II: Team-First Architecture & Cognitive Load

### Small, Long-Lived Teams as the Standard

> *"Best-performing teams accomplish remarkable feats not simply because of the individual qualifications of their members, but because those members coalesce into a single organism."* — *Team of Teams*

- **Optimal Team Size**: Stable groups of **5 to 9 people** (Dunbar's number limits deep trust to ~5 close collaborators, ~15 close trust partners).
- **Assign Work to Teams, Not Individuals**: Foster shared ownership and eliminate single points of failure.
- **Long-Lived Teams Enable Flow**:
  - Takes 2 weeks to 3 months to form a cohesive unit (forming → storming → norming → performing).
  - High churn / reassigning people every few months disrupts momentum (*Brooks' Law*, *The Mythical Man-Month*).
- **Team-First Mindset & Collective Ownership**:
  - Embrace diversity in perspectives and problem-solving styles (*Peopleware*).
  - Reward the entire team, not individual contributors, to avoid misaligned incentives and internal friction.

### Managing Cognitive Load

Cognitive load determines the upper limit of a team's effectiveness. Software subsystem boundaries must be sized according to team cognitive capacity, not arbitrary technical slicing.

#### Domain Complexity Guidelines per Team
1. **1 team : 1 complex/complicated domain**
2. **1 team : 2–3 simple domains**
3. **Never assign 2+ complex domains to a single team** (leads to context-switching overhead and disengagement).

#### Practical Strategies to Reduce Cognitive Load
- **Clear Boundaries**: Align service boundaries with domain boundaries and team cognitive load.
- **Team APIs**: Treat team interactions like published APIs — clear contracts, documentation, SLAs, and dedicated communication channels.
- **Developer Experience (DevEx)**: Improve internal platform usability, self-service portals, and SDK documentation.
- **Management Style**: Practice *"Eyes On, Hands Off"* (*Team of Teams*) — communicate clear outcomes and intent rather than micromanaging the "how".
- **Shared Mental Models**: Cultivate a shared understanding of architecture to reduce defects and streamline decision-making.

---

## 🔄 Part III: Team Interaction & Flow of Change

### Anti-Patterns in Team Design

- **Ad-Hoc Specialized Teams**: Standing up temporary, disconnected operational teams (e.g., ad-hoc DBA or middleware units) without end-to-end integration creates handoff delays.
- **Constant Member Shuffling**: Treats engineers as interchangeable cogs; overlooks context switching, loss of institutional knowledge, and differing conventions.
- **Feature Teams without High Engineering Maturity**:
  - Editing shared codebases without strict ownership or boy-scout disciplines results in codebase degradation and eroded cross-team trust.

### Team Topologies Evolution & Models

- **The Spotify Model Framework**:
  - **Squad**: Small, cross-functional autonomous team (Stream-Aligned).
  - **Tribe**: Collection of squads working in a related domain.
  - **Chapter**: Functional specialization area across squads for craft quality and mentoring.
  - **Guild**: Loose community of interest across the entire organization.
- **Flow & Sensing Mechanisms**: Actively monitor cross-team dependencies and communication health to dynamically evolve topologies as business context shifts.

---

## 📚 References & Recommended Reading

- *Team Topologies* — Matthew Skelton & Manuel Pais
- *Team of Teams* — General Stanley McChrystal
- *The Mythical Man-Month* — Frederick P. Brooks Jr.
- *Peopleware: Productive Projects and Teams* — Tom DeMarco & Timothy Lister
- *Lean Enterprise* — Jez Humble, Joanne Molesky, Barry O'Reilly
- *The Principles of Product Development Flow* — Donald G. Reinertsen
- *The Five Dysfunctions of a Team* — Patrick Lencioni
- *Improving Performance: How to Manage the White Space on the Organization Chart* — Geary A. Rummler & Alan P. Brache
- *A Leader's Framework for Decision Making* (Cynefin framework) — David J. Snowden & Mary E. Boone

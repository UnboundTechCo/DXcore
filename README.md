# DXcore (Proprietary Core)

## Overview

DXcore, a proprietary and closed core designed to protect users from censorship and network restrictions, powers advanced VPN capabilities in [DefyxVPN](https://github.com/UnboundTechCo/defyxVPN), a fully open-source client developed with Flutter.

The decision to keep DXcore closed-source is intentional: it implements advanced and customized censorship-circumvention protocols. Public disclosure of these implementations would make it significantly easier for adversarial systems (e.g., GFW) to analyze, fingerprint, and block them, directly endangering users. This is not secrecy for its own sake — it is an operational security measure in an ongoing struggle to ensure safe, reliable, and unrestricted connectivity.

![Cover](.github/images/cover.png)

---

## What is DXcore

DXcore is a set of engines and protocol adaptations developed to withstand:

- **Traffic analysis**
- **Protocol fingerprinting**
- **Active probing**

These components are the result of targeted engineering for use in hostile network environments. Publicly exposing implementation details would rapidly weaken their effectiveness and reduce user safety.

---

## Why DXcore is Closed

- **Operational security over exposure.** Public code would allow hostile actors to reverse-engineer and design precise countermeasures.
- **Transparency in design, not in sensitive implementation.** We keep designs, API formats, and architectural documentation open while withholding exploitable implementation details.
- **User protection as the priority.** Closed sourcing DXcore is a risk-mitigation step to protect users in repressive network environments.

---

## Our Commitment to Transparency & Independent Review

We balance openness with operational safety through controlled evaluation:

1. **Design transparency.** High-level architecture, APIs, and configuration formats are public.
2. **Independent audits.** DXcore source is made available to accredited security teams under strict legal and technical controls (NDA, secure review environments).
3. **Public summaries.** Post-audit, summary reports and remediation statements are published without leaking sensitive implementation details.
4. **Reproducible builds & signed binaries.** Wherever possible, reproducible builds and signed releases are provided, so integrity can be verified without exposing internals.
5. **Responsible disclosure.** We maintain a formal vulnerability reporting process with coordinated timelines to protect users.

---

## Audit Process (Overview)

When accredited teams request DXcore access for audit:

1. **Legal & procedural agreements.** NDA or audit contracts define scope, reporting format, and non-disclosure of exploitation details.
2. **Controlled access.** Access is provided via private repositories or secure environments, with time limits and monitoring.
3. **Scope definition.** Focus on cryptographic correctness, protocol implementation, dependency hygiene, and network behavior.
4. **Reporting & remediation.** Vulnerabilities are reported privately; fixes are prioritized and coordinated for disclosure.
5. **Public summary.** Once resolved, a high-level summary of findings and fixes is shared publicly.

---

## Responsible Disclosure

If you discover a potential vulnerability:

- **Email:** security@defyxvpn.com
- **Report contents (recommended):** short summary, impact assessment, reproduction steps or limited PoC (please avoid full exploit code), proposed mitigations, and your contact info.
- **Coordination:** For critical issues, we coordinate disclosure timelines with the reporter to ensure user safety.

---

## FAQ

**Q: Does closed-source mean users cannot trust DXcore?**  
A: No. Closing the code is not our first choice for transparency, but in censorship-circumvention contexts, open implementations can directly harm users. We offset this by providing open designs, controlled audits, public summaries, and verifiable binaries.

**Q: Can I review DXcore’s real source?**  
A: Accredited security organizations and teams can request access under NDA and secure review conditions. Public audit summaries are published afterward.

**Q: How does this protect users?**  
A: By withholding sensitive implementation details, adversaries cannot easily adapt to block or fingerprint the system. This increases the durability of user connections in hostile environments.

---

## DXcore Mock / Stub

To support developers building with Flutter, we provide a **DXcore Mock** (stub implementation) in this repository.

**Download from the [DXcore Releases page](https://github.com/UnboundTechCo/DXcore/releases).**

The mock is designed **only** to allow the app to compile and run in development and to enable UI/integration testing without needing the proprietary core.

- The mock calls the same expected DXcore methods.
- It returns predefined/sample responses.
- It **does not** contain any real censorship-circumvention logic, traffic-resistance, or sensitive protocol code.
- It must **not** be used for security or network-behavior testing.

For proper auditing, testing, or analysis, the real DXcore (available only to accredited auditors under NDA) must be used.

---

## Short Summary

DXcore is closed-source not for secrecy but for **user protection** against censorship systems. Designs and APIs remain open; sensitive code is withheld but made available for independent audits under NDA. Mock/stub implementations are available for developers: [DXcore Releases](https://github.com/UnboundTechCo/DXcore/releases).

Contact: **security@defyxvpn.com**

---

## Contact

For audit requests, security reports, or inquiries about controlled DXcore access:

**security@defyxvpn.com**

---

## Third-Party Licenses

This project uses the following third-party components:

- [Xray](https://github.com/XTLS/Xray-core): Licensed under terms specified by its authors. Please refer to the official repository for full license details.
- [Vwarp](https://github.com/voidr3aper-anon/Vwarp): Licensed under terms specified by its authors. Please refer to the official repository for full license details.
- [Psiphon](https://github.com/UnboundTechCo/psiphon-tunnel-core): Licensed under terms specified by its authors. Please refer to the official repository for full license details.
- [Outline](https://github.com/Jigsaw-Code/outline-sdk): Licensed under terms specified by its authors. Please refer to the official repository for full license details.

If you are a copyright holder and believe your license is not properly attributed, please contact us at **info@defyxvpn.com**.

---

## Legal Notice

DXcore is proprietary. It is withheld from public release to protect users against censorship adversaries. Access for audits is provided under NDA to accredited institutions, with public audit summaries published afterward. The DXcore Mock is safe for development/testing only and not a substitute for the real implementation.

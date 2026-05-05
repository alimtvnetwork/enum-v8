# 04 — No Email-Based CI Notifications

## Symptom

N/A — this is a permanent constraint, not a bug.

## Root Cause

User explicitly receives too many emails (e.g. Dependabot bump notifications) and rejects all email-based notification flows.

## Fix / Workaround

The AI must NEVER:
- Configure Dependabot recipients.
- Add SMTP / email-sending code.
- Set up CI alerts that email anyone.
- Suggest "we'll notify you by email" as a feature.

When configuring Dependabot or any CI workflow that has notification options, leave email recipients empty / unconfigured. Use repo-level GitHub UI notifications only (which the user can mute themselves).

## Status

🚫 Permanent constraint.

## Related

- `mem://constraints/no-email-notifications.md`
- `.lovable/strictly-avoid.md` § "Communication & delivery"

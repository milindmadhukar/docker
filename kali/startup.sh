#!/bin/bash

mkdir -p /root/.vnc
echo "root" | vncpasswd -f > /root/.vnc/passwd
chmod 600 /root/.vnc/passwd

# Extract rockyou.txt if it exists as .gz
if [ -f /usr/share/wordlists/rockyou.txt.gz ]; then
    gunzip /usr/share/wordlists/rockyou.txt.gz
fi

# Start VNC server
vncserver :1 -geometry 1920x1080 -depth 24

# Keep the container running
tail -f /dev/null

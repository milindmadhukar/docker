#!/bin/bash

mkdir -p /root/.vnc
echo "root" | vncpasswd -f > /root/.vnc/passwd
chmod 600 /root/.vnc/passwd

# Start VNC server
vncserver :1 -geometry 1280x800 -depth 24

# Keep the container running
tail -f /dev/null

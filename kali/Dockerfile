FROM kalilinux/kali-rolling

ENV DEBIAN_FRONTEND=noninteractive
ENV USER=root

RUN apt update && apt upgrade -y && apt install -y \
    metasploit-framework \
    set \
    xfce4 \
    tightvncserver \
    dbus-x11 \
    x11-xserver-utils \
    net-tools \
    iputils-ping \
    curl \
    git \
    hydra \
    openssh-client \
    nmap \
    seclists \
    wordlists \
    && apt clean

COPY startup.sh /root/startup.sh
RUN chmod +x /root/startup.sh

CMD ["/root/startup.sh"]

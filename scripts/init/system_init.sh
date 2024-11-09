#!/bin/bash

# 系统依赖安装
install_dependencies() {
    apt-get update
    apt-get install -y \
        docker.io \
        docker-compose \
        postgresql-client \
        redis-tools \
        consul \
        prometheus \
        grafana
}

# 创建系统用户和组
setup_users() {
    useradd -r -s /bin/false gfc
    usermod -aG docker gfc
}

# 配置系统限制
configure_system_limits() {
    cat >> /etc/security/limits.conf <<EOF
gfc soft nofile 65536
gfc hard nofile 65536
gfc soft nproc 32768
gfc hard nproc 32768
EOF
}

# 配置系统参数
configure_sysctl() {
    cat >> /etc/sysctl.conf <<EOF
net.core.somaxconn = 65535
net.ipv4.tcp_max_syn_backlog = 65535
net.ipv4.tcp_fin_timeout = 30
net.ipv4.tcp_keepalive_time = 300
net.ipv4.tcp_keepalive_probes = 5
net.ipv4.tcp_keepalive_intvl = 15
EOF
    sysctl -p
}

# 主函数
main() {
    install_dependencies
    setup_users
    configure_system_limits
    configure_sysctl
}

main 
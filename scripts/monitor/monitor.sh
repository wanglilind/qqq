#!/bin/bash

# 监控配置
PROMETHEUS_URL="http://prometheus:9090"
ALERT_THRESHOLD=90
CHECK_INTERVAL=60  # 秒
LOG_FILE="/var/log/gfc/monitor.log"

# 颜色输出
RED='\033[0;31m'
YELLOW='\033[1;33m'
GREEN='\033[0;32m'
NC='\033[0m'

# 日志函数
log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1" | tee -a $LOG_FILE
}

# 检查服务健康状态
check_service_health() {
    local service=$1
    local endpoint="http://${service}:8080/health"
    
    response=$(curl -s -o /dev/null -w "%{http_code}" $endpoint)
    if [ "$response" == "200" ]; then
        log "${GREEN}Service ${service} is healthy${NC}"
        return 0
    else
        log "${RED}Service ${service} is unhealthy (HTTP $response)${NC}"
        return 1
    fi
}

# 检查系统资源使用情况
check_system_resources() {
    # CPU使用率
    cpu_usage=$(top -bn1 | grep "Cpu(s)" | awk '{print $2}' | cut -d. -f1)
    if [ "$cpu_usage" -gt "$ALERT_THRESHOLD" ]; then
        log "${RED}High CPU usage: ${cpu_usage}%${NC}"
        send_alert "HIGH_CPU" "CPU usage is ${cpu_usage}%"
    fi

    # 内存使用率
    memory_usage=$(free | grep Mem | awk '{print ($3/$2 * 100)}' | cut -d. -f1)
    if [ "$memory_usage" -gt "$ALERT_THRESHOLD" ]; then
        log "${RED}High memory usage: ${memory_usage}%${NC}"
        send_alert "HIGH_MEMORY" "Memory usage is ${memory_usage}%"
    fi

    # 磁盘使用率
    disk_usage=$(df -h | grep '/dev/sda1' | awk '{print $5}' | cut -d% -f1)
    if [ "$disk_usage" -gt "$ALERT_THRESHOLD" ]; then
        log "${RED}High disk usage: ${disk_usage}%${NC}"
        send_alert "HIGH_DISK" "Disk usage is ${disk_usage}%"
    fi
}

# 检查服务性能指标
check_service_metrics() {
    local service=$1
    
    # 查询Prometheus获取服务指标
    response=$(curl -s "${PROMETHEUS_URL}/api/v1/query" \
        --data-urlencode "query=rate(${service}_request_duration_seconds_count[5m])")
    
    # 解析响应并检查性能指标
    # 这里需要根据实际的Prometheus查询和指标进行调整
}

# 发送告警
send_alert() {
    local alert_type=$1
    local message=$2
    
    # 记录告警
    log "${RED}ALERT: ${alert_type} - ${message}${NC}"
    
    # 发送告警到告警系统
    # 这里可以集成具体的告警系统，如Slack、邮件等
}

# 主监控循环
main() {
    log "Starting monitoring service..."
    
    while true; do
        # 检查核心服务��康状态
        for service in "identity" "transaction" "consensus" "monitor"; do
            check_service_health $service
        done
        
        # 检查系统资源
        check_system_resources
        
        # 检查服务性能指标
        for service in "identity" "transaction" "consensus" "monitor"; do
            check_service_metrics $service
        done
        
        # 等待下一次检查
        sleep $CHECK_INTERVAL
    done
}

# 启动监控
main 
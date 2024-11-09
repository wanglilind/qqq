#!/bin/bash

# 安全审计配置
AUDIT_DIR="security/audit"
SCAN_TARGETS=("identity-service" "transaction-service" "consensus-service")
REPORT_DIR="security/reports"

# 创建目录
mkdir -p $AUDIT_DIR $REPORT_DIR

# 日志函数
log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1" | tee -a "$REPORT_DIR/audit.log"
}

# 依赖检查
check_dependencies() {
    local deps=("nmap" "nikto" "sonarqube" "trivy")
    for dep in "${deps[@]}"; do
        if ! command -v $dep &> /dev/null; then
            log "Error: $dep is required but not installed."
            exit 1
        fi
    done
}

# 漏洞扫描
vulnerability_scan() {
    local target=$1
    log "Starting vulnerability scan for $target..."
    
    # 使用Trivy扫描容器镜像
    trivy image "gfc/$target:latest" \
        --format json \
        --output "$REPORT_DIR/${target}_vulnerabilities.json"
    
    # 使用Nmap进行端口扫描
    nmap -sV -sC -oX "$REPORT_DIR/${target}_ports.xml" "$target"
}

# 代码安全分析
code_security_analysis() {
    log "Starting code security analysis..."
    
    # 运行SonarQube分析
    sonar-scanner \
        -Dsonar.projectKey=gfc \
        -Dsonar.sources=. \
        -Dsonar.host.url=http://localhost:9000 \
        -Dsonar.login=$SONAR_TOKEN
}

# 配置审计
config_audit() {
    log "Auditing configuration files..."
    
    # 检查配置文件权限
    find ./configs -type f -name "*.yaml" -exec ls -l {} \;
    
    # 检查敏感信息
    grep -r "password\|secret\|key" ./configs
}

# 权限审计
permission_audit() {
    log "Auditing permissions..."
    
    # 检查文件权限
    find . -type f -perm /o=w -ls
    
    # 检查特权容器
    kubectl get pods -o json | jq '.items[] | select(.spec.containers[].securityContext.privileged == true)'
}

# 生成审计报告
generate_report() {
    log "Generating audit report..."
    
    # 合并所有扫描结果
    cat > "$REPORT_DIR/audit_report.html" <<EOF
    <!DOCTYPE html>
    <html>
    <head>
        <title>Security Audit Report</title>
        <style>
            body { font-family: Arial, sans-serif; }
            .vulnerability { color: red; }
            .warning { color: orange; }
            .info { color: blue; }
        </style>
    </head>
    <body>
        <h1>Security Audit Report</h1>
        <div id="summary">
            <!-- 摘要信息将通过脚本动态插入 -->
        </div>
        <div id="details">
            <!-- 详细信息将通过脚本动态插入 -->
        </div>
    </body>
    </html>
EOF

    # 处理扫描结果并更新报告
    python3 scripts/security/process_audit_results.py \
        --input-dir "$REPORT_DIR" \
        --output "$REPORT_DIR/audit_report.html"
}

# 主审计流程
main() {
    log "Starting security audit..."
    
    # 检查依赖
    check_dependencies
    
    # 执行各项审计
    for target in "${SCAN_TARGETS[@]}"; do
        vulnerability_scan $target
    done
    
    code_security_analysis
    config_audit
    permission_audit
    
    # 生成报告
    generate_report
    
    log "Security audit completed"
}

main 
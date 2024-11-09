#!/bin/bash

# 性能测试配置
TEST_DURATION=300  # 测试持续时间（秒）
CONCURRENT_USERS=1000  # 并发用户数
RAMP_UP_TIME=60  # 爬坡时间（秒）
TEST_HOST="localhost"
REPORT_DIR="test/reports/performance"

# 创建报告目录
mkdir -p $REPORT_DIR

# 日志函数
log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1" | tee -a "$REPORT_DIR/performance_test.log"
}

# 身份服务性能测试
test_identity_service() {
    log "Starting identity service performance test..."
    
    # 使用k6进行负载测试
    k6 run - <<EOF
    import http from 'k6/http';
    import { check, sleep } from 'k6';

    export let options = {
        stages: [
            { duration: '${RAMP_UP_TIME}s', target: ${CONCURRENT_USERS} },
            { duration: '${TEST_DURATION}s', target: ${CONCURRENT_USERS} },
            { duration: '30s', target: 0 }
        ],
    };

    export default function() {
        // 身份验证测试
        let res = http.post('http://${TEST_HOST}:8080/identity/verify', {
            'user_id': 'test-user',
            'biometric_data': 'test-data'
        });
        check(res, {
            'status is 200': (r) => r.status === 200,
            'response time < 500ms': (r) => r.timings.duration < 500
        });
        sleep(1);
    }
EOF
}

# 交易服务性能测试
test_transaction_service() {
    log "Starting transaction service performance test..."
    
    k6 run - <<EOF
    import http from 'k6/http';
    import { check, sleep } from 'k6';

    export let options = {
        stages: [
            { duration: '${RAMP_UP_TIME}s', target: ${CONCURRENT_USERS} },
            { duration: '${TEST_DURATION}s', target: ${CONCURRENT_USERS} },
            { duration: '30s', target: 0 }
        ],
    };

    export default function() {
        // 交易创建测试
        let res = http.post('http://${TEST_HOST}:8080/transaction/create', {
            'sender': 'user1',
            'recipient': 'user2',
            'amount': '100'
        });
        check(res, {
            'status is 200': (r) => r.status === 200,
            'response time < 1s': (r) => r.timings.duration < 1000
        });
        sleep(1);
    }
EOF
}

# 生成性能报告
generate_report() {
    log "Generating performance report..."
    
    # 收集测试结果
    cat > "$REPORT_DIR/report.html" <<EOF
    <!DOCTYPE html>
    <html>
    <head>
        <title>Performance Test Report</title>
    </head>
    <body>
        <h1>Performance Test Results</h1>
        <div id="results">
            <!-- 测试结果将通过脚本动态插入 -->
        </div>
    </body>
    </html>
EOF

    # 使用jq处理k6输出的JSON结果
    # 这里需要根据实际的k6输出格式调整
}

# 主测试流程
main() {
    log "Starting performance tests..."
    
    # 检查依赖
    command -v k6 >/dev/null 2>&1 || { log "k6 is required but not installed."; exit 1; }
    
    # 运行测试
    test_identity_service
    test_transaction_service
    
    # 生成报告
    generate_report
    
    log "Performance tests completed"
}

main 
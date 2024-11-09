#!/bin/bash

# 安装 iconv
yum install -y iconv

# 修复文件编码
fix_encoding() {
    local file=$1
    echo "Fixing encoding for $file"
    
    # 创建临时文件
    temp_file="${file}.tmp"
    
    # 转换编码
    iconv -f ISO-8859-1 -t UTF-8 "$file" > "$temp_file"
    
    # 检查转换是否成功
    if [ $? -eq 0 ]; then
        mv "$temp_file" "$file"
        echo "Successfully fixed encoding for $file"
    else
        rm -f "$temp_file"
        echo "Failed to fix encoding for $file"
    fi
}

# 需要修复的文件列表
files=(
    "pkg/consensus/pbft/pbft.go"
    "pkg/contract/interop/interface.go"
    "pkg/contract/migrate/migrator.go"
    "pkg/contract/upgrade/proxy.go"
    "pkg/contract/validator/validator.go"
    "pkg/contract/deploy/deployer.go"
    "pkg/contract/debug/debugger.go"
    "pkg/contract/event/emitter.go"
    "pkg/contract/test/suite.go"
    "pkg/contract/docs/generator.go"
    "pkg/contract/profiler/profiler.go"
    "pkg/contract/monitor/watcher.go"
    "pkg/contract/dependency/manager.go"
    "pkg/contract/benchmark/runner.go"
    "pkg/contract/engine.go"
    "pkg/contract/backup/manager.go"
    "pkg/contract/version/manager.go"
    "pkg/contract/security/checker.go"
    "pkg/contract/stdlib/standard.go"
    "pkg/contract/optimize/cache.go"
    "pkg/contract/audit/analyzer.go"
    "pkg/blockchain/block.go"
    "pkg/crypto/hash.go"
    "pkg/crypto/signature.go"
    "pkg/discovery/consul.go"
    "pkg/config/loader.go"
    "pkg/database/store.go"
    "pkg/database/postgres.go"
    "pkg/database/dao/dao.go"
    "pkg/state/manager.go"
    "pkg/logger/logger.go"
    "pkg/network/p2p/node.go"
    "pkg/vm/vm.go"
    "cmd/monitor-service/main.go"
    "cmd/transaction-service/main.go"
    "cmd/consensus-service/main.go"
    "cmd/identity-service/main.go"
    "internal/consensus/service/service.go"
    "internal/monitor/alert/manager.go"
    "internal/monitor/security/auditor.go"
    "internal/monitor/service/service.go"
)

# 修复所有文件
for file in "${files[@]}"; do
    if [ -f "$file" ]; then
        fix_encoding "$file"
    else
        echo "File not found: $file"
    fi
done

# 设置 Git 配置
git config --global core.autocrlf false
git config --global core.filemode true
git config --global core.safecrlf true

echo "Encoding fix completed" 
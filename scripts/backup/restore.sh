#!/bin/bash

# 恢复配置
RESTORE_DIR="/var/restore/gfc"
DB_HOST="localhost"
DB_PORT="5432"
DB_USER="gfc_user"
DB_NAME="gfc_db"
S3_BUCKET="gfc-backups"

# 日志函数
log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1" >> "$RESTORE_DIR/restore.log"
}

# 从云存储下载备份
download_from_cloud() {
    local backup_date=$1
    log "Downloading backup files from cloud storage..."
    
    mkdir -p $RESTORE_DIR
    
    # 下载指定日期的备份
    aws s3 cp "s3://${S3_BUCKET}/database/gfc_backup_${backup_date}.dump.gpg" $RESTORE_DIR/
    aws s3 cp "s3://${S3_BUCKET}/configs/gfc_backup_${backup_date}_configs.tar.gz.gpg" $RESTORE_DIR/
    
    if [ $? -eq 0 ]; then
        log "Download completed successfully"
    else
        log "Download failed"
        exit 1
    fi
}

# 解密备份文件
decrypt_backup() {
    log "Decrypting backup files..."
    
    # 使用GPG解密
    gpg --decrypt "$RESTORE_DIR/*.dump.gpg" > "$RESTORE_DIR/database.dump"
    gpg --decrypt "$RESTORE_DIR/*.tar.gz.gpg" > "$RESTORE_DIR/configs.tar.gz"
    
    if [ $? -eq 0 ]; then
        log "Decryption completed successfully"
    else
        log "Decryption failed"
        exit 1
    fi
}

# 恢复数据库
restore_database() {
    log "Starting database restore..."
    
    # 停止相关服务
    systemctl stop gfc-services
    
    # 删除现有数据库
    dropdb -h $DB_HOST -p $DB_PORT -U $DB_USER $DB_NAME
    createdb -h $DB_HOST -p $DB_PORT -U $DB_USER $DB_NAME
    
    # 恢复数据库
    pg_restore -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -F c "$RESTORE_DIR/database.dump"
    
    if [ $? -eq 0 ]; then
        log "Database restore completed successfully"
    else
        log "Database restore failed"
        exit 1
    fi
}

# 恢复配置文件
restore_configs() {
    log "Restoring configuration files..."
    
    # 解压配置文件
    tar -xzf "$RESTORE_DIR/configs.tar.gz" -C /etc/gfc/
    
    if [ $? -eq 0 ]; then
        log "Configuration restore completed successfully"
    else
        log "Configuration restore failed"
        exit 1
    fi
}

# 验证恢复
verify_restore() {
    log "Verifying restore..."
    
    # 检查数据库连接
    psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "\dt"
    
    # 检查配置文件
    if [ -f "/etc/gfc/configs/config.yaml" ]; then
        log "Configuration files verified"
    else
        log "Configuration files verification failed"
        exit 1
    fi
}

# 启动服务
start_services() {
    log "Starting services..."
    systemctl start gfc-services
    
    # 检查服务状态
    if systemctl is-active --quiet gfc-services; then
        log "Services started successfully"
    else
        log "Services failed to start"
        exit 1
    fi
}

# 主恢复流程
main() {
    if [ -z "$1" ]; then
        echo "Usage: $0 <backup_date>"
        echo "Example: $0 20240101_120000"
        exit 1
    fi
    
    local backup_date=$1
    log "Starting restore process for backup date: $backup_date"
    
    # 执行恢复步骤
    download_from_cloud $backup_date
    decrypt_backup
    restore_database
    restore_configs
    verify_restore
    start_services
    
    log "Restore process completed successfully"
}

main "$@" 
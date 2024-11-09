#!/bin/bash

# 备份配置
BACKUP_DIR="/var/backups/gfc"
DB_HOST="localhost"
DB_PORT="5432"
DB_USER="gfc_user"
DB_NAME="gfc_db"
RETENTION_DAYS=30
S3_BUCKET="gfc-backups"

# 创建备份目录
mkdir -p $BACKUP_DIR

# 生成备份文件名
TIMESTAMP=$(date +%Y%m%d_%H%M%S)
BACKUP_FILE="$BACKUP_DIR/gfc_backup_$TIMESTAMP"

# 日志函数
log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1" >> "$BACKUP_DIR/backup.log"
}

# 数据库备份
backup_database() {
    log "Starting database backup..."
    
    # 导出数据库
    pg_dump -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -F c -f "${BACKUP_FILE}.dump"
    
    if [ $? -eq 0 ]; then
        log "Database backup completed successfully"
    else
        log "Database backup failed"
        exit 1
    fi
}

# 配置文件备份
backup_configs() {
    log "Starting configuration backup..."
    
    # 打包配置文件
    tar -czf "${BACKUP_FILE}_configs.tar.gz" /etc/gfc/configs/
    
    if [ $? -eq 0 ]; then
        log "Configuration backup completed successfully"
    else
        log "Configuration backup failed"
        exit 1
    fi
}

# 加密备份文件
encrypt_backup() {
    log "Encrypting backup files..."
    
    # 使用GPG加密
    gpg --encrypt --recipient gfc-backup "${BACKUP_FILE}.dump"
    gpg --encrypt --recipient gfc-backup "${BACKUP_FILE}_configs.tar.gz"
    
    # 删除未加密的文件
    rm "${BACKUP_FILE}.dump" "${BACKUP_FILE}_configs.tar.gz"
}

# 上传到云存储
upload_to_cloud() {
    log "Uploading backup files to cloud storage..."
    
    # 上传到S3
    aws s3 cp "${BACKUP_FILE}.dump.gpg" "s3://${S3_BUCKET}/database/"
    aws s3 cp "${BACKUP_FILE}_configs.tar.gz.gpg" "s3://${S3_BUCKET}/configs/"
    
    if [ $? -eq 0 ]; then
        log "Cloud upload completed successfully"
    else
        log "Cloud upload failed"
        exit 1
    fi
}

# 清理旧备份
cleanup_old_backups() {
    log "Cleaning up old backups..."
    
    # 删除本地旧备份
    find $BACKUP_DIR -type f -mtime +$RETENTION_DAYS -delete
    
    # 删除云存储中的旧备份
    aws s3 ls "s3://${S3_BUCKET}" | while read -r line;
    do
        createDate=`echo $line|awk {'print $1" "$2'}`
        createDate=`date -d"$createDate" +%s`
        olderThan=`date -d"-$RETENTION_DAYS days" +%s`
        if [[ $createDate -lt $olderThan ]]
        then
            fileName=`echo $line|awk {'print $4'}`
            if [[ $fileName != "" ]]
            then
                aws s3 rm "s3://${S3_BUCKET}/$fileName"
            fi
        fi
    done
}

# 主备份流程
main() {
    log "Starting backup process..."
    
    # 检查依赖
    command -v pg_dump >/dev/null 2>&1 || { log "pg_dump is required but not installed."; exit 1; }
    command -v gpg >/dev/null 2>&1 || { log "gpg is required but not installed."; exit 1; }
    command -v aws >/dev/null 2>&1 || { log "aws cli is required but not installed."; exit 1; }
    
    # 执行备份步骤
    backup_database
    backup_configs
    encrypt_backup
    upload_to_cloud
    cleanup_old_backups
    
    log "Backup process completed successfully"
}

main 
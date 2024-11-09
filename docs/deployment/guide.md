# 全球公平数字货币系统部署指南

## 1. 系统要求

### 1.1 硬件要求
- CPU: 至少8核心
- 内存: 至少16GB RAM
- 存储: 至少500GB SSD
- 网络: 千兆网络带宽

### 1.2 软件要求
- Ubuntu Server 20.04 LTS
- Docker 20.10+
- Kubernetes 1.22+
- 结构化数据使用 CockroachDB
- 区块数据使用 LevelDB
- 热点数据使用 Redis
- Go 1.20+

## 2. 前置准备

### 2.1 安装基础软件

更新系统
```bash
sudo apt update && sudo apt upgrade -y
```
安装基础工具
```bash
sudo apt install -y git make curl wget
```
安装 Docker
curl -fsSL https://get.docker.com | sh
sudo usermod -aG docker $USER
安装 kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
安装 Helm
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
### 2.2 配置数据库
bash
安装 PostgreSQL
sudo apt install -y postgresql postgresql-contrib
创建数据库和用户
sudo -u postgres psql <<EOF
CREATE DATABASE gfc_db;
CREATE USER gfc_user WITH ENCRYPTED PASSWORD 'your_secure_password';
GRANT ALL PRIVILEGES ON DATABASE gfc_db TO gfc_user;
EOF

## 3. 部署步骤

### 3.1 克隆代码仓库
bash
git clone https://github.com/your-org/global-fair-currency.git
cd global-fair-currency

### 3.2 配置环境变量
bash
创建环境变量文件
cat > .env <<EOF
ENVIRONMENT=production
DB_HOST=localhost
DB_PORT=5432
DB_USER=gfc_user
DB_PASSWORD=your_secure_password
DB_NAME=gfc_db
REDIS_HOST=localhost
REDIS_PORT=6379
JWT_SECRET=your_jwt_secret
EOF

### 3.3 初始化系统
bash
运行初始化脚本
./scripts/init/system_init.sh
运行数据库迁移
make migrate-up

### 3.4 部署核心服务

#### 3.4.1 身份服务
bash
kubectl apply -f deployments/kubernetes/identity-service.yaml

#### 3.4.2 交易服务

bash
kubectl apply -f deployments/kubernetes/transaction-service.yaml

#### 3.4.3 共识服务
bash
kubectl apply -f deployments/kubernetes/consensus-service.yaml

#### 3.4.4 监控服务
bash
kubectl apply -f deployments/kubernetes/monitor-service.yaml

### 3.5 配置负载均衡
bash
安装 Ingress Controller
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm install ingress-nginx ingress-nginx/ingress-nginx
应用 Ingress 配置
kubectl apply -f deployments/kubernetes/ingress.yaml

## 4. 监控与维护

### 4.1 部署监控工具
bash
部署 Prometheus
kubectl apply -f deployments/kubernetes/prometheus.yaml
部署 Grafana
kubectl apply -f deployments/kubernetes/grafana.yaml

### 4.2 配置日志收集
bash
部署 EFK 栈
kubectl apply -f deployments/kubernetes/efk/

### 4.3 设置告警
bash
配置 AlertManager
kubectl apply -f deployments/kubernetes/alertmanager.yaml

## 5. 安全加固

### 5.1 网络安全
bash
应用网络策略
kubectl apply -f deployments/kubernetes/network-policies.yaml
配置防火墙规则
sudo ufw allow 80,443/tcp
sudo ufw allow 6443/tcp # Kubernetes API

### 5.2 证书配置
bash
安装 cert-manager
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.8.0/cert-manager.yaml
配置 TLS 证书
kubectl apply -f deployments/kubernetes/certificates.yaml

## 6. 备份策略

### 6.1 配置自动备份
bash
设置定时备份任务
kubectl apply -f deployments/kubernetes/backup-cronjob.yaml

### 6.2 验证备份恢复
bash
测试备份恢复
./scripts/backup/test-restore.sh

## 7. 故障排除

### 7.1 检查服务状态
bash
查看所有服务状态
kubectl get pods -n gfc
查看服务日志
kubectl logs -f deployment/identity-service -n gfc

### 7.2 常见问题解决
- 数据库连接问题
- 网络连接问题
- 性能问题
- 内存泄漏问题

## 8. 扩容指南

### 8.1 水平扩展
bash
扩展服务实例
kubectl scale deployment identity-service --replicas=5 -n gfc

### 8.2 垂直扩展
bash
修改资源限制
kubectl edit deployment identity-service -n gfc

## 9. 更新流程

### 9.1 滚动更新
bash
更新服务镜像
kubectl set image deployment/identity-service identity=gfc/identity-service:new-version -n gfc

### 9.2 回滚操作
bash
回滚到上一版本
kubectl rollout undo deployment/identity-service -n gfc

## 10. 性能优化

### 10.1 系统调优
bash
应用系统优化配置
sudo sysctl -p /etc/sysctl.d/99-gfc.conf

### 10.2 资源监控
- 配置资源使用告警
- 设置自动扩缩容策略
- 优化缓存策略

## 附录：常用命令
```bash
查看服务状态
kubectl get all -n gfc
查看日志
kubectl logs -f <pod-name> -n gfc
进入容器
kubectl exec -it <pod-name> -n gfc -- /bin/sh
查看资源使用
kubectl top pods -n gfc
```
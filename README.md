# Kubernetes Monitor in Go

This is a **hobby project** written in Go to monitor a Kubernetes cluster. It uses the official `client-go` library to interact with the Kubernetes API and provides basic observability into:

- Current Pods and their statuses  
- Log summaries from Pods whose names match a specific prefix  
- Live events for Pod additions, deletions, and modifications

> ⚙️ This tool is designed for CLI-based monitoring, especially useful when working directly with clusters in dev/test environments.

---

## 🧠 Features

- 📦 **List all Pods** in all namespaces, showing name and phase  
- 📄 **Fetch logs** from all containers in Pods that start with a specific prefix (e.g., `p-`)  
- 📺 **Watch for live Pod events**: added, modified, or deleted

---

## 🔧 Requirements

- Go 1.24 or newer  
- Access to a valid Kubernetes cluster (via kubeconfig)

Install dependencies:

```bash
go mod tidy
```

Run the app:

```bash
go run main.go
```

## 💡 Example Output
▶️ Listing current Pods:
```sql
- beaconator-minio-tenant-default/beaconator-pool-0-0 (Phase: Running)
- cert-manager/cert-manager-c756b4fd7-zs8bq (Phase: Running)
- cert-manager/cert-manager-cainjector-694df7dc87-lpb7f (Phase: Running)
- cert-manager/cert-manager-webhook-596dbf79cd-m5zgb (Phase: Running)
- default/my-release-postgresql-ha-client (Phase: Failed)
- default/my-release-postgresql-ha-pgpool-55494f984b-6tl7f (Phase: Running)
- default/my-release-postgresql-ha-postgresql-0 (Phase: Running)
- default/my-release-postgresql-ha-postgresql-1 (Phase: Running)
- default/my-release-postgresql-ha-postgresql-2 (Phase: Running)
- edb-migration-copilot/edb-migration-copilot-f845d4bb6-cgwk2 (Phase: Running)
- edb-migration-portal/cluster-migration-portal-1 (Phase: Running)
- edb-migration-portal/k8s-mp-67f4c794bf-zkx2q (Phase: Running)
- edb-migration-portal/mp-epas-17-1 (Phase: Running)
```

📄 Fetching logs for pods starting with 'p-':

```css
  --> Getting logs for p-35v0yfm35f/p-35v0yfm35f-1 [container: postgres]
    [SUMMARY] 3534 lines, 3533 error(s)
  --> Getting logs for p-35v0yfm35f/p-35v0yfm35f-1 [container: beacon-agent-2ded2d08-7c30-413b-a4e8-c0b4b1750bfc]
    [SUMMARY] 380 lines, 28 error(s)
  --> Getting logs for p-35v0yfm35f/p-35v0yfm35f-2 [container: postgres]
    [SUMMARY] 406 lines, 405 error(s)
  --> Getting logs for p-35v0yfm35f/p-35v0yfm35f-2 [container: beacon-agent-adec2249-4f97-478d-a9df-c145aaa14d48]
    [SUMMARY] 54 lines, 0 error(s)
```

📡 Watching for Pod events:
```less
[ADDED] Pod: beaconator-minio-tenant-default/beaconator-pool-0-0
[ADDED] Pod: cert-manager/cert-manager-c756b4fd7-zs8bq
...

[ADDED] Pod: upm-replicator/kubernetes-replicator-7b54695bc-dztqz
[ADDED] Pod: upm-ui/upm-ui-5bf978b894-xqknt
[MODIFIED] Pod: p-35v0yfm35f/p-35v0yfm35f-2
[MODIFIED] Pod: p-35v0yfm35f/p-35v0yfm35f-2
[MODIFIED] Pod: p-35v0yfm35f/p-35v0yfm35f-2
[DELETED] Pod: p-35v0yfm35f/p-35v0yfm35f-2
[ADDED] Pod: p-35v0yfm35f/p-35v0yfm35f-2
[MODIFIED] Pod: p-35v0yfm35f/p-35v0yfm35f-2
```
## 🛠️ TODO (Planned Features)
- Add Prometheus integration for metrics
- Expose as a REST API or TUI (Terminal UI)
- Filter by label selectors or resource limits
- Log to file or external system

## 📁 Project Structure
```bash
.
├── main.go         # Entry point
├── go.mod / sum    # Dependency declarations
└── README.md       # You're here!
```

### 🧑‍💻 Author:
This project is maintained by me as a hobby project to sharpen my Go skills and better understand how Kubernetes works under the hood.

### 📜 License
MIT License – free to use, modify, or build upon.
```vbnet
Let me know if you'd like me to include a badge or add a contribution section!
```
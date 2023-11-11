# TermK

Easy show log pods using kubeconf over multiples cluster

![show_logs_one_pod.png](images%2Fshow_logs_one_pod.png)


### use 

copy data file `config` kubernetes in path root termk, with filename `kubeconfig.conf`, example:

```bash
cp ~/.kube/config kubeconf.conf
```

### keys

Common keys on screens 

| key                                | description                                |  
|------------------------------------|--------------------------------------------|
| <enter>                            | enter each item                            | 
| <backspace> <ctrl-backspace> <esc> | return item before                         | 
| <ctrl-q>                           | delete pod (use carefully)                 |
| r                                  | update the list of namespaces, pods, logs. |
| g                                  | top                                        |
| G                                  | Buttom                                     |
| J                                  | Down                                       |
| k                                  | up                                         |
| <C-d>                              | Page Down                                  |
| <C-u>                              | Page Up                                    |
| q                                  | exit program                               |

### Logs
Get last 500 record logs POD
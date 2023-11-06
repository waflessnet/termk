# TermK

Easy show log pods using kubeconf over multiples cluster

![show_logs_one_pod.png](images%2Fshow_logs_one_pod.png)


### use 

copy data file `config` kubernetes in path root termk, with filename `kubeconfig.conf`, example:

```bash
cp ~/.kube/config kubeconf.conf
```

### keys


| key                            | description                |  
|--------------------------------|----------------------------|
| enter                          | enter each item            | 
| backspace or control-backspace | return item before         | 
| d                              | delete pod (use carefully) |
| m                              | over logs, get new logs    |
| q                              | exit                       |


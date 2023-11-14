# TermK

Easy show log pods using kubeconf over multiples cluster k8

![show_logs_one_pod.png](images%2Fshow_logs_one_pod.png)


### Install 

download last release from repo https://github.com/waflessnet/termk/releases 
 or can download with script:

Linux Family 
```bash
wget https://github.com/waflessnet/termk/releases/download/v0.2.2-alpha/termk-linux-amd64 -o /tmp/termk && sudo mv termk-linux-amd64 /usr/local/bin/termk && sudo chmod +x /usr/local/bin/termk
```



### Use

App work over terminal S.O, call directly binary `termk`. 
Termk first search “kubeconfig.conf” in current path. You can indicate the path of a specific conf with kconfig

```bash
termk --kconfig ~/.kube/config
```
#### *Optional 
can create `kubeconfig.conf` copy data file `config` kubernetes in path o root termk, with filename `kubeconfig.conf`, example:

```bash
cp ~/.kube/config kubeconf.conf
```

then open terminal in path created kubeconf.conf and run `termk`

### Keys

Common keys on screens 

| key                                     | description                                |  
|-----------------------------------------|--------------------------------------------|
| \<enter\>                               | enter each item                            | 
| \<backspace\> <ctrl-backspace\> \<esc\> | return item before                         | 
| \<ctrl-q\>                              | delete pod (use carefully)                 |
| r                                       | update the list of namespaces, pods, logs. |
| g                                       | top                                        |
| G                                       | Buttom                                     |
| J                                       | Down                                       |
| k                                       | up                                         |
| \<C-d\>                                 | Page Down                                  |
| \<C-u\>                                 | Page Up                                    |
| q                                       | exit program                               |

### Logs
Get last 500 record logs POD
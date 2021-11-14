# rpc-account

## 运维

```
ops --variable .cfg/dev.yaml -a dep
ops --variable .cfg/dev.yaml -a run --env dev --task config --cmd=put
ops --variable .cfg/dev.yaml -a run --env dev --task image
ops --variable .cfg/dev.yaml -a run --env dev --task mysql
ops --variable .cfg/dev.yaml -a run --env dev --task helm --cmd=upgrade
```
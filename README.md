# goflatten

# install

```
go install github.com/nozo-moto/goflatten
```

# how to use

pipe jsondata to goflatten

```
$ cat example/test1.json| goflatten | jq
{
  "123_456": 789,
  "alpha_beta_gamma_delta_epsilon_zeta": "eta",
  "alpha_beta_gamma_theta": "lota"
}
```

